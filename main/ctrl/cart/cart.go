package cart

import (
	"fmt"

	"github.com/rvauradkar1/testfuse/main/ctrl/cache"
	"github.com/rvauradkar1/testfuse/main/ctrl/ord"
	"github.com/rvauradkar1/testfuse/main/ctrl/ord/db"
	"github.com/rvauradkar1/testfuse/main/setup"
)

type ICartService interface {
	Add(cart string) error
}

type Service struct {
	CacheSvc cache.IService `_fuse:"CacheSvc"`
	DBSvc    db.IService    `_fuse:"DBSvc"`
}

func (c *Service) Add(cart string) error {
	fmt.Println("Calling find on OrderSvc")
	o := setup.Find("OrderSvc")
	ordSvc := o.(ord.IService)
	err := ordSvc.SaveOrder(cart)
	if err != nil {
		return err
	}
	err = c.DBSvc.AddCart(cart)
	if err != nil {
		return err
	}
	c.CacheSvc.AddCart(cart)
	return nil
}
