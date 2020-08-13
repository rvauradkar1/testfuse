package main

import (
	"fmt"
	"runtime"

	"github.com/rvauradkar1/fuse/mock"

	"github.com/rvauradkar1/testfuse/main/ctrl/ord"

	"github.com/rvauradkar1/testfuse/main/ctrl/ord/db"

	"github.com/rvauradkar1/testfuse/main/ctrl/cache"

	"github.com/rvauradkar1/testfuse/main/ctrl/auth"

	"github.com/rvauradkar1/testfuse/main/ctrl"
	"github.com/rvauradkar1/testfuse/main/ctrl/cart"
	"github.com/rvauradkar1/testfuse/main/find"

	"github.com/rvauradkar1/fuse/fuse"
)

func main() {
	call(100)
	fmt.Println("Hello testfuse")
	cs := make([]fuse.Entry, 0)
	cs = append(cs, fuse.Entry{Name: "OrdCtrl", Stateless: true, Instance: &ctrl.OrderController{}})
	cs = append(cs, fuse.Entry{Name: "CartSvc", Stateless: true, Instance: &cart.CartSvc{}})
	cs = append(cs, fuse.Entry{Name: "AuthSvc", Stateless: true, Instance: &auth.AuthSvc{}})
	cs = append(cs, fuse.Entry{Name: "CacheSvc", Stateless: true, Instance: &cache.CacheSvc{}})
	cs = append(cs, fuse.Entry{Name: "DBSvc", Stateless: true, Instance: &db.DBSvc{}})
	cs = append(cs, fuse.Entry{Name: "OrderSvc", Stateless: false, Instance: &ord.OrderSvc{}})

	f := fuse.New()
	find.Find = f.Find
	errors := f.Register(cs)
	fmt.Println(errors)
	i := f.Find("OrdCtrl")
	ii := i.(*ctrl.OrderController)
	err := ii.Order("raj", "order123")
	fmt.Println("Return from 1 ", err)
	genMocks()
}

func call(i int) {
	fmt.Println(runtime.Caller(1))

}

func genMocks() {
	m := mock.MockGen{}
	comps := make([]mock.Component, 0)
	path := "/Users/rvauradkar/go_code/src/github.com/rvauradkar1/testfuse/main"
	fmt.Println("path ====== " + path)
	comps = append(comps, mock.Component{PtrToComp: &ctrl.OrderController{}, Basepath: path + "/ctrl", Name: "OrdCtrl"})
	comps = append(comps, mock.Component{PtrToComp: &cart.CartSvc{}, Basepath: path + "/ctrl/cart", Name: "CartSvc"})
	comps = append(comps, mock.Component{PtrToComp: &auth.AuthSvc{}, Basepath: path + "/ctrl/auth", Name: "AuthSvc"})
	comps = append(comps, mock.Component{PtrToComp: &cache.CacheSvc{}, Basepath: path + "/ctrl/cache", Name: "CacheSvc"})
	comps = append(comps, mock.Component{PtrToComp: &db.DBSvc{}, Basepath: path + "/ctrl/ord/db", Name: "DBSvc"})
	comps = append(comps, mock.Component{PtrToComp: &ord.OrderSvc{}, Basepath: path + "/ctrl/ord", Name: "OrderSvc"})
	m.Comps = comps
	m.Gen()
}
