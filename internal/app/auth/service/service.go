package auth_service

type AuthService interface {
	Login(username, password string) (string, error)
}