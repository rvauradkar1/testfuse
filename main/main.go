package main

import (
	"fmt"

	"github.com/rvauradkar1/testfuse/main/ctrl/ord"

	"github.com/rvauradkar1/testfuse/main/ctrl/ord/db"

	"github.com/rvauradkar1/testfuse/main/ctrl/cache"

	"github.com/rvauradkar1/testfuse/main/ctrl/auth"

	"github.com/rvauradkar1/testfuse/main/ctrl"
	"github.com/rvauradkar1/testfuse/main/ctrl/cart"
	"github.com/rvauradkar1/testfuse/main/setup"

	"github.com/rvauradkar1/fuse/fuse"
)

func main() {
	fmt.Println("Hello testfuse")
	cs := make([]fuse.Entry, 0)
	cs = append(cs, fuse.Entry{Name: "OrdCtrl", Stateless: true, Instance: &ctrl.OrderController{}})
	cs = append(cs, fuse.Entry{Name: "CartSvc", Stateless: true, Instance: &cart.Service{}})
	cs = append(cs, fuse.Entry{Name: "AuthSvc", Stateless: true, Instance: &auth.Service{}})
	cs = append(cs, fuse.Entry{Name: "CacheSvc", Stateless: true, Instance: &cache.Service{}})
	cs = append(cs, fuse.Entry{Name: "DBSvc", Stateless: true, Instance: &db.Service{}})
	cs = append(cs, fuse.Entry{Name: "OrderSvc", Stateless: true, Instance: &ord.Service{}})

	f := fuse.New()
	setup.Find = f.Find
	errors := f.Register(cs)
	fmt.Println(errors)
	i := f.Find("OrdCtrl")
	ii := i.(*ctrl.OrderController)
	err := ii.Order("raj", "order123")
	fmt.Println("Return from 1 ", err)
}
