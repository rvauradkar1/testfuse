package ord

import (
	"fmt"

	"github.com/rvauradkar1/testfuse/main/ctrl/cache"
	"github.com/rvauradkar1/testfuse/main/ctrl/ord/db"
)

type IService interface {
	SaveOrder(order string) error
}

type OrderSvc struct {
	CacheSvc cache.IService `_fuse:"CacheSvc"`
	DBSvc    db.IService    `_fuse:"DBSvc"`
}

func (o *OrderSvc) SaveOrder(order string) error {
	fmt.Println("Begin of AddOrder  on OrderSvc")
	err := o.DBSvc.AddOrder(order)
	if err != nil {
		return err
	}
	o.CacheSvc.AddOrder(order)
	return nil
}
