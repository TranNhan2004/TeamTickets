package auth

type AuthService interface {
	Register(username, password string) (string, error)
	VerifyEmail(token string) (string, error)
	Login(username, password string) (string, error)
	RefreshToken(refreshToken string) (string, error)
}
