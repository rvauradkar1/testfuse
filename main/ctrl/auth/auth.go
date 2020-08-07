package auth

import "fmt"

type IService interface {
	Auth(user string) error
}

type Service struct {
}

func (a *Service) Auth(user string) error {
	fmt.Printf("Auth for user [%s]\n", user)
	return nil
}
