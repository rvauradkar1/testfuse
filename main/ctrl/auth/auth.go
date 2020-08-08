package auth

import "fmt"

type IService interface {
	Auth(user string) error
}

type AuthSvc struct {
}

func (a *AuthSvc) Auth(user string) error {
	fmt.Printf("Auth for user [%s]\n", user)
	return nil
}
