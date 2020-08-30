package cart

import (
	"fmt"

	"github.com/rvauradkar1/testfuse/main/cfg"

	"github.com/rvauradkar1/testfuse/main/ctrl/cache"
	"github.com/rvauradkar1/testfuse/main/ctrl/ord"
	"github.com/rvauradkar1/testfuse/main/ctrl/ord/db"
)

type ICartService interface {
	Add(cart string) error
}

type CartSvc struct {
	CacheSvc cache.IService `_fuse:"CacheSvc"`
	DBSvc    db.IService    `_fuse:"DBSvc"`
	DEPS_    interface{}    `_deps:"OrderSvc"`
}

func (c *CartSvc) Add(cart string) error {
	fmt.Println("Calling find on OrderSvc111")
	o := cfg.Find("OrderSvc")
	ordSvc := o.(ord.IService)
	err := ordSvc.SaveOrder(cart, "stat")
	if err != nil {
		return err
	}
	err = c.DBSvc.AddCart(cart)
	if err != nil {
		return err
	}
	c.CacheSvc.AddCrt(cart)
	return nil
}
