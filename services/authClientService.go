package services

import (
	"context"
	"time"

	"github.com/MatheusMikio/config"
	"github.com/MatheusMikio/dto/auth"
	"github.com/MatheusMikio/dto/client"
	"github.com/MatheusMikio/models"
	"github.com/MatheusMikio/repository"
	"github.com/MatheusMikio/schemas"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	googleOAuth2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type IAuthClientService interface {
	RequestMagicLink(email string) *models.ErrorMessage
	VerifyMagicLink(token string) (*auth.LoginResponse[client.ClientResponse], *models.ErrorMessage)
	GetGoogleAuthUrl(state string) string
	HandleGoogleCallBack(code string) (*auth.LoginResponse[client.ClientResponse], *models.ErrorMessage)
}

type AuthClientService struct {
	ClientRepository   repository.IClientRepository
	MagiLinkRepository repository.IMagicLinkRepository
	GoogleConfig       *oauth2.Config
}

func NewAuthClientService(clientRepo repository.IClientRepository, magicLinkRepo repository.IMagicLinkRepository) IAuthClientService {
	googleConfig := &oauth2.Config{
		ClientID:     config.GetGoogleClientId(),
		ClientSecret: config.GetGoogleClientSecret(),
		RedirectURL:  config.GetGoogleRedirectURL("client"),
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}
	return &AuthClientService{
		ClientRepository:   clientRepo,
		MagiLinkRepository: magicLinkRepo,
		GoogleConfig:       googleConfig,
	}
}

func (a *AuthClientService) GetGoogleAuthUrl(state string) string {
	return a.GoogleConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

func (a *AuthClientService) HandleGoogleCallBack(code string) (*auth.LoginResponse[client.ClientResponse], *models.ErrorMessage) {
	logger := config.GetLogger("AuthClient:GoogleCallback")
	ctx := context.Background()

	token, err := a.GoogleConfig.Exchange(ctx, code)
	if err != nil {
		logger.Errorf("Failed to exchange code: %v", err)
		return nil, models.CreateErrorMessage("Google", "Failed to authenticate with Google.")
	}

	oauth2Service, err := googleOAuth2.NewService(ctx, option.WithTokenSource(a.GoogleConfig.TokenSource(ctx, token)))
	if err != nil {
		logger.Errorf("Failed to create oauth2 service: %v", err)
		return nil, models.CreateErrorMessage("Google", "Error retrieving information from Google.")
	}

	userInfo, err := oauth2Service.Userinfo.Get().Do()
	if err != nil {
		logger.Errorf("Failed to get user info: %v", err)
		return nil, models.CreateErrorMessage("Google", "Error retrieving information from Google.")
	}

	clientData := &schemas.Client{}

	clientData, err = a.ClientRepository.GetByGoogleId(userInfo.Id)
	if err != nil {
		clientData, err := a.ClientRepository.GetByEmail(userInfo.Email)
		if err != nil {
			logger.Errorf("Client not found: %v", err)
			return nil, models.CreateErrorMessage("Client", "Not found")
		}
		clientData.GoogleId = &userInfo.Id
	}

	clientData.EmailVerified = true
	if err := a.ClientRepository.Update(clientData); err != nil {
		logger.Errorf("Failed to update client: %v", err)
		return nil, models.CreateErrorMessage("Client", "Failed update client")
	}

	jwtToken, _, err := generateJWT(clientData.ID, clientData.Email, "client", "CLIENT")
	if err != nil {
		logger.Errorf("Failed to generate JWT: %v", err)
		return nil, models.CreateErrorMessage("System", "Erro ao gerar token de autenticação")
	}

	appointmentsSummary := make([]client.AppointmentSummary, 0, len(clientData.Appointments))
	for _, scheduling := range clientData.Appointments {
		appointmentsSummary = append(appointmentsSummary, client.AppointmentSummary{
			ID:   scheduling.ID,
			Date: scheduling.Date,
		})
	}

	response := &auth.LoginResponse[client.ClientResponse]{
		User: client.ClientResponse{
			ID:                  clientData.ID,
			FirstName:           clientData.FirstName,
			LastName:            clientData.LastName,
			Cpf:                 clientData.Cpf,
			Email:               clientData.Email,
			Phone:               clientData.Phone,
			CreatedAt:           clientData.CreatedAt,
			UpdatedAt:           clientData.UpdatedAt,
			AppointmentsSummary: appointmentsSummary,
		},
		Token: jwtToken,
	}

	return response, nil
}

func (a *AuthClientService) RequestMagicLink(email string) *models.ErrorMessage {
	logger := config.GetLogger("AuthClient: RequestMagicLink")
	client, err := a.ClientRepository.GetByEmail(email)
	if err != nil {
		logger.Errorf("Client not found: %v", err)
		return models.CreateErrorMessage("Email", "Email not found!")
	}
	token, err := generateToken()
	if err != nil {
		logger.Errorf("Failed to generate token: %v", err)
		return models.CreateErrorMessage("Token", "Error generate token")
	}

	magicLink := &schemas.MagicLink{
		Email:      client.Email,
		Token:      token,
		EntityType: "client",
		ExpiresAt:  time.Now().Add(15 * time.Minute),
	}

	if err := a.MagiLinkRepository.Create(magicLink); err != nil {
		logger.Errorf("Failed to save magic link: %v", err)
		return models.CreateErrorMessage("Database", "Failed to save magic link!")
	}

	magicLinkURL := "http://localhost:8080/api/v1/auth/client/verify?token=" + token
	logger.Infof("Magic Link: %s", magicLinkURL)

	return nil
}

func (a *AuthClientService) VerifyMagicLink(token string) (*auth.LoginResponse[client.ClientResponse], *models.ErrorMessage) {
	logger := config.GetLogger("AuthClient:VerifyMagicLink")
	magicLink, err := a.MagiLinkRepository.GetByToken(token)
	if err != nil {
		logger.Errorf("Invalid or expired token: %v", err)
	}

	if magicLink.EntityType != "client" {
		logger.Error("Invalid token")
		return nil, models.CreateErrorMessage("Token", "Invalid")
	}

	clientData, err := a.ClientRepository.GetByEmail(magicLink.Email)

	if err != nil {
		logger.Errorf("Client not found: %v", err)
		return nil, models.CreateErrorMessage("Client", "Not found!")
	}

	if err := a.MagiLinkRepository.MarkUsed(magicLink.ID); err != nil {
		logger.Errorf("Failed to mark as used: %v", err)
		return nil, models.CreateErrorMessage("MagicLink", "Fail mark used")
	}

	clientData.EmailVerified = true

	if err := a.ClientRepository.Update(clientData); err != nil {
		logger.Errorf("Failed to update client: %v", err)
		return nil, models.CreateErrorMessage("Client", "Failed to update client")
	}

	jwtToken, _, err := generateJWT(clientData.ID, clientData.Email, "client", "CLIENTS")
	if err != nil {
		logger.Errorf("Failed to generate JWT: %v", err)
		return nil, models.CreateErrorMessage("System", "Erro ao gerar token de autenticação")
	}

	appointmentSummary := make([]client.AppointmentSummary, 0, len(clientData.Appointments))
	for _, scheduling := range clientData.Appointments {
		appointmentSummary = append(appointmentSummary, client.AppointmentSummary{
			ID:   scheduling.ID,
			Date: scheduling.Date,
		})
	}

	response := &auth.LoginResponse[client.ClientResponse]{
		User: client.ClientResponse{
			ID:                  clientData.ID,
			FirstName:           clientData.FirstName,
			LastName:            clientData.LastName,
			Cpf:                 clientData.Cpf,
			Email:               clientData.Email,
			Phone:               clientData.Phone,
			AppointmentsSummary: appointmentSummary,
			CreatedAt:           clientData.CreatedAt,
			UpdatedAt:           clientData.UpdatedAt,
		},
		Token: jwtToken,
	}

	return response, nil
}
