package services

import (
	"context"
	"time"

	"github.com/MatheusMikio/config"
	"github.com/MatheusMikio/dto/auth"
	"github.com/MatheusMikio/dto/professional"
	"github.com/MatheusMikio/models"
	"github.com/MatheusMikio/repository"
	"github.com/MatheusMikio/schemas"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	googleOAuth2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type IAuthProfessionalService interface {
	RequestMagicLink(email string) *models.ErrorMessage
	VerifyMagicLink(token string) (*auth.LoginResponse[professional.ProfessionalResponse], *models.ErrorMessage)
	GetGoogleAuthUrl(state string) string
	HandleGoogleCallBack(code string) (*auth.LoginResponse[professional.ProfessionalResponse], *models.ErrorMessage)
}

type AuthProfessionalService struct {
	ProfessionalRepository repository.IProfessionalRepository
	MagicLinkRepository    repository.IMagicLinkRepository
	GoogleConfig           *oauth2.Config
}

func NewAuthProfessionalService(
	professionalRepo repository.IProfessionalRepository,
	magicLinkRepo repository.IMagicLinkRepository,
) IAuthProfessionalService {
	googleConfig := &oauth2.Config{
		ClientID:     config.GetGoogleClientId(),
		ClientSecret: config.GetGoogleClientSecret(),
		RedirectURL:  config.GetGoogleRedirectURL("professional"),
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}

	return &AuthProfessionalService{
		ProfessionalRepository: professionalRepo,
		MagicLinkRepository:    magicLinkRepo,
		GoogleConfig:           googleConfig,
	}
}

func (a *AuthProfessionalService) GetGoogleAuthUrl(state string) string {
	return a.GoogleConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

func (a *AuthProfessionalService) HandleGoogleCallBack(code string) (*auth.LoginResponse[professional.ProfessionalResponse], *models.ErrorMessage) {
	logger := config.GetLogger("AuthProfessional:GoogleCallback")
	ctx := context.Background()

	token, err := a.GoogleConfig.Exchange(ctx, code)
	if err != nil {
		logger.Errorf("Failed to exchange code: %v", err)
		return nil, models.CreateErrorMessage("Google", "Falha ao autenticar com Google")
	}

	oauth2Service, err := googleOAuth2.NewService(ctx, option.WithTokenSource(a.GoogleConfig.TokenSource(ctx, token)))
	if err != nil {
		logger.Errorf("Failed to create oauth2 service: %v", err)
		return nil, models.CreateErrorMessage("Google", "Erro ao obter informações do Google")
	}

	userInfo, err := oauth2Service.Userinfo.Get().Do()
	if err != nil {
		logger.Errorf("Failed to get user info: %v", err)
		return nil, models.CreateErrorMessage("Google", "Erro ao obter informações do usuário")
	}

	var professionalData *schemas.Professional
	professionalData, err = a.ProfessionalRepository.GetByGoogleId(userInfo.Id)
	if err != nil {
		professionalData, err = a.ProfessionalRepository.GetByEmail(userInfo.Email)
		if err != nil {
			logger.Errorf("Professional not found: %v", err)
			return nil, models.CreateErrorMessage("Professional", "Usuário não cadastrado no sistema")
		}
		professionalData.GoogleId = &userInfo.Id
	}

	professionalData.EmailVerified = true
	if err := a.ProfessionalRepository.Update(professionalData); err != nil {
		logger.Errorf("Failed to update professional: %v", err)
	}

	jwtToken, _, err := generateJWTProfessional(
		professionalData.ID,
		professionalData.Email,
		professionalData.Role,
	)
	if err != nil {
		logger.Errorf("Failed to generate JWT: %v", err)
		return nil, models.CreateErrorMessage("System", "Erro ao gerar token de autenticação")
	}

	response := &auth.LoginResponse[professional.ProfessionalResponse]{
		User: professional.ProfessionalResponse{
			ID:        professionalData.ID,
			FirstName: professionalData.FirstName,
			LastName:  professionalData.LastName,
			Email:     professionalData.Email,
			Phone:     professionalData.Phone,
			CompanyID: professionalData.CompanyID,
			Role:      professionalData.Role,
		},
		Token: jwtToken,
	}

	return response, nil
}

func (a *AuthProfessionalService) RequestMagicLink(email string) *models.ErrorMessage {
	logger := config.GetLogger("AuthProfessional:RequestMagicLink")

	professionalData, err := a.ProfessionalRepository.GetByEmail(email)
	if err != nil {
		logger.Errorf("Professional not found: %v", err)
		return models.CreateErrorMessage("Professional", "Not found!")
	}

	token, err := generateToken()
	if err != nil {
		logger.Errorf("Failed to generate token: %v", err)
		return models.CreateErrorMessage("Token", "Failed to generate!")
	}

	magicLink := &schemas.MagicLink{
		Email:      professionalData.Email,
		Token:      token,
		EntityType: "professional",
		ExpiresAt:  time.Now().Add(15 * time.Minute),
	}

	if err := a.MagicLinkRepository.Create(magicLink); err != nil {
		logger.Errorf("Failed to save magic link: %v", err)
		return models.CreateErrorMessage("Magic link", "Failed to save!")
	}

	magicLinkURL := "http://localhost:8080/api/v1/auth/professional/verify?token=" + token
	logger.Infof("Magic Link: %s", magicLinkURL)

	return nil
}

func (a *AuthProfessionalService) VerifyMagicLink(token string) (*auth.LoginResponse[professional.ProfessionalResponse], *models.ErrorMessage) {
	logger := config.GetLogger("AuthProfessional:VerifyMagicLink")

	magicLink, err := a.MagicLinkRepository.GetByToken(token)
	if err != nil {
		logger.Errorf("Invalid or expired token: %v", err)
		return nil, models.CreateErrorMessage("Token", "Token inválido ou expirado")
	}

	if magicLink.EntityType != "professional" {
		logger.Error("Token is not for professional entity")
		return nil, models.CreateErrorMessage("Token", "Token inválido para este tipo de usuário")
	}

	professionalData, err := a.ProfessionalRepository.GetByEmail(magicLink.Email)
	if err != nil {
		logger.Errorf("Professional not found: %v", err)
		return nil, models.CreateErrorMessage("Professional", "Profissional não encontrado")
	}

	if err := a.MagicLinkRepository.MarkUsed(magicLink.ID); err != nil {
		logger.Errorf("Failed to mark as used: %v", err)
	}

	professionalData.EmailVerified = true
	if err := a.ProfessionalRepository.Update(professionalData); err != nil {
		logger.Errorf("Failed to update professional: %v", err)
	}

	jwtToken, _, err := generateJWTProfessional(
		professionalData.ID,
		professionalData.Email,
		professionalData.Role,
	)
	if err != nil {
		logger.Errorf("Failed to generate JWT: %v", err)
		return nil, models.CreateErrorMessage("System", "Erro ao gerar token de autenticação")
	}

	response := &auth.LoginResponse[professional.ProfessionalResponse]{
		User: professional.ProfessionalResponse{
			ID:        professionalData.ID,
			FirstName: professionalData.FirstName,
			LastName:  professionalData.LastName,
			Email:     professionalData.Email,
			Phone:     professionalData.Phone,
			CompanyID: professionalData.CompanyID,
			Role:      professionalData.Role,
		},
		Token: jwtToken,
	}

	return response, nil
}
