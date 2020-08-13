package cart

import (
	"fmt"
	"testing"

	"github.com/rvauradkar1/testfuse/main/find"

	"github.com/rvauradkar1/fuse/fuse"

	"github.com/rvauradkar1/testfuse/main/ctrl/cache"
	"github.com/rvauradkar1/testfuse/main/ctrl/ord/db"
)

var mockCache cache.IService = &MockCacheSvc{}
var mockDB db.IService = &MockDBSvc{}

func init() {
	fmt.Println("pppppp")
	fuse := fuse.New()
	fuse.RegisterMock("OrderSvc", &MockOrderSvc{})
	find.Find = fuse.Find
}

func Test_Add(t *testing.T) {
	c := CartSvc{CacheSvc: mockCache, DBSvc: mockDB}
	SaveOrderFunc = func(cart string) error {
		return nil
	}
	AddCartFunc = func(cart string) error {
		return nil
	}
	AddCrtFunc = func(cart string) error {
		return nil
	}

	c.Add("new cart")
	fmt.Println(c)
}
