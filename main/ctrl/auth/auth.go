package auth

import (
	"fmt"
	"time"
)

type IService interface {
	Auth(user string) error
}

type AuthSvc struct {
	t time.Duration
}

func (a *AuthSvc) Auth(user string) error {
	fmt.Printf("Auth for user [%s]\n", user)
	return nil
}
