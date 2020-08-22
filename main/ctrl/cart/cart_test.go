package cart

import (
	"fmt"
	"testing"

	"github.com/rvauradkar1/testfuse/main/ctrl/cache"

	"github.com/rvauradkar1/fuse/fuse"
	"github.com/rvauradkar1/testfuse/main/ctrl/ord/db"
	"github.com/rvauradkar1/testfuse/main/find"
)

var mockCache cache.IService = &MockCacheSvc{}
var mockDB db.IService = &MockDBSvc{}

func init() {
	fmt.Println("pppppp")
	f := fuse.New()
	f.RegisterMock("OrderSvc", &MockOrderSvc{})
	find.Find = f.Find
}

func Test_Add(t *testing.T) {
	c := CartSvc{CacheSvc: mockCache, DBSvc: mockDB}
	MockSaveOrder = func(cart string) error {
		return nil
	}
	MockAddCrt = func(cart string) error {
		return nil
	}
	MockAddCart = func(cart string) error {
		return nil
	}

	err := c.Add("new cart")
	if err != nil {
		t.Errorf("there should have been no error but got %s\n", err)
	}
	fmt.Println(c)
}
