package ctrl

import (
	"github.com/rvauradkar1/testfuse/main/ctrl/cache"
	"github.com/rvauradkar1/testfuse/main/ctrl/ord/db"
)

type MockAuthSvc struct {
}

type Auth func(s1 string) error

var AuthFunc Auth

func (p *MockAuthSvc) Auth(s1 string) error {
	return AuthFunc(s1)
}

type MockCartSvc struct {
	CacheSvc cache.IService
	DBSvc    db.IService
}

type Add func(s1 string) error

var AddFunc Add

func (p *MockCartSvc) Add(s1 string) error {
	return AddFunc(s1)
}
