package main

import (
	"fmt"

	"github.com/rvauradkar1/testfuse/main/cfg"

	"github.com/rvauradkar1/testfuse/main/ctrl"
	"github.com/rvauradkar1/testfuse/main/ctrl/auth"
	"github.com/rvauradkar1/testfuse/main/ctrl/cache"
	"github.com/rvauradkar1/testfuse/main/ctrl/cart"
	"github.com/rvauradkar1/testfuse/main/ctrl/ord"
	"github.com/rvauradkar1/testfuse/main/ctrl/ord/db"

	"github.com/rvauradkar1/fuse/fuse"
)

/*
func main1() {
	fmt.Println("Hello testfuse")
	f, errors := funcName()
	errors = f.Wire()
	fmt.Println(errors)
	comp := f.Find("OrdCtrl")
	ctrl := comp.(*ctrl.OrderController)
	err := ctrl.Order("raj", "order123")
	fmt.Println("Return from 1 ", err)
	//genMocks()
}
*/

func main() {
	fmt.Println("Hello testfuse")
	entries := Entries()
	errors := cfg.Fuse(entries)
	fmt.Println(errors)

	comp := cfg.Find("OrdCtrl")
	ctrl := comp.(*ctrl.OrderController)
	err := ctrl.Order("raj", "order123")
	fmt.Println("Return from 1 ", err)
}

func Entries() []fuse.Entry {
	fmt.Println("Hello testfuse")
	entries := make([]fuse.Entry, 0)
	entries = append(entries, fuse.Entry{Name: "OrdCtrl", State: false, Instance: &ctrl.OrderController{}})
	entries = append(entries, fuse.Entry{Name: "CartSvc", State: false, Instance: &cart.CartSvc{}})
	entries = append(entries, fuse.Entry{Name: "AuthSvc", State: false, Instance: &auth.AuthSvc{}})
	entries = append(entries, fuse.Entry{Name: "CacheSvc", State: false, Instance: &cache.CacheSvc{}})
	entries = append(entries, fuse.Entry{Name: "DBSvc", State: false, Instance: &db.DBSvc{}})
	entries = append(entries, fuse.Entry{Name: "OrderSvc", State: true, Instance: &ord.OrderSvc{}})

	return entries
}
