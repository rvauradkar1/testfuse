package cart

import (
	"fmt"
	"testing"

	"github.com/rvauradkar1/testfuse/main/cfg"

	"github.com/rvauradkar1/testfuse/main/ctrl/cache"

	"github.com/rvauradkar1/testfuse/main/ctrl/ord/db"
)

var mockCache cache.IService = &MockCacheSvc{}
var mockDB db.IService = &MockDBSvc{}

func init() {
	fmt.Println("Calling init...")
	cfg.Find = func(name string) interface{} {
		switch name {
		case "OrderSvc":
			return &MockOrderSvc{}
		}
		return nil
	}
}

func Test_Add(t *testing.T) {
	c := CartSvc{CacheSvc: mockCache, DBSvc: mockDB}
	MockSaveOrder = func(cart string) error {
		fmt.Println("Inside MockSaveOrder... ")
		return nil
	}
	MockAddCrt = func(cart string) error {
		fmt.Println("Inside MockAddCrt... ")
		return nil
	}
	MockAddCart = func(cart string) error {
		fmt.Println("Inside MockAddCart... ")
		return nil
	}

	err := c.Add("new cart")
	if err != nil {
		t.Errorf("there should have been no error but got %s\n", err)
	}
	fmt.Println(c)
}
