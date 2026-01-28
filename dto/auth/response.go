package auth

type AuthResponse struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}

type LoginResponse[T any] struct {
	User  T      `json:"user"`
	Token string `json:"token"`
}

type GoogleAuthUrlResponse struct {
	Url string `json:"url"`
}
