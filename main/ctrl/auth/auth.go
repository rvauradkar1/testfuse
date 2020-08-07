package auth

type IAuthService interface {
	auth() bool
}

type AuthService struct {
}

func (a *AuthService) auth(user string) bool {
	return true
}

type OrderDB struct {
}
