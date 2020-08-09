package ctrl

import (
	"github.com/rvauradkar1/testfuse/main/ctrl/auth"
	"github.com/rvauradkar1/testfuse/main/ctrl/cache"
	"github.com/rvauradkar1/testfuse/main/ctrl/cart"
	"github.com/rvauradkar1/testfuse/main/ctrl/ord/db"
)

type MockAuthSvc struct {
}

type Auth func(s1 string) error

var AuthFunc Auth

func (p *MockAuthSvc) Auth(s1 string) error {
	return AuthFunc(s1)
}

type MockOrderController struct {
	CartSvc *cart.CartSvc
	AuthSvc auth.IService
}

type Order func(s1 string, s2 string) error

var OrderFunc Order

func (p *MockOrderController) Order(s1 string, s2 string) error {
	return OrderFunc(s1, s2)
}

type MockCartSvc struct {
	CacheSvc cache.IService
	DBSvc    db.IService
	_GEN_    interface{}
}

type Add func(s1 string) error

var AddFunc Add

func (p *MockCartSvc) Add(s1 string) error {
	return AddFunc(s1)
}
