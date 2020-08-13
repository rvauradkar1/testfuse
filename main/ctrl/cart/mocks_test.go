package cart

import (
	"github.com/rvauradkar1/testfuse/main/ctrl/cache"
	"github.com/rvauradkar1/testfuse/main/ctrl/ord/db"
)

type MockCacheSvc struct {
}

type AddCrt func(s1 string) error

var AddCrtFunc AddCrt

func (p *MockCacheSvc) AddCrt(s1 string) error {
	return AddCrtFunc(s1)
}

type AddOrd func(s1 string) error

var AddOrdFunc AddOrd

func (p *MockCacheSvc) AddOrd(s1 string) error {
	return AddOrdFunc(s1)
}

type MockDBSvc struct {
}

type AddCart func(s1 string) error

var AddCartFunc AddCart

func (p *MockDBSvc) AddCart(s1 string) error {
	return AddCartFunc(s1)
}

type AddOrder func(s1 string) error

var AddOrderFunc AddOrder

func (p *MockDBSvc) AddOrder(s1 string) error {
	return AddOrderFunc(s1)
}

type MockOrderSvc struct {
	CacheSvc cache.IService
	DBSvc    db.IService
}

type SaveOrder func(s1 string) error

var SaveOrderFunc SaveOrder

func (p *MockOrderSvc) SaveOrder(s1 string) error {
	return SaveOrderFunc(s1)
}

type MockCartSvc struct {
	CacheSvc cache.IService
	DBSvc    db.IService
	DEPS_    interface{}
}

type Add func(s1 string) error

var AddFunc Add

func (p *MockCartSvc) Add(s1 string) error {
	return AddFunc(s1)
}
