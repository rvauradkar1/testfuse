package main

import (
	"fmt"
	"runtime"

	"github.com/rvauradkar1/testfuse/main/ctrl"
	"github.com/rvauradkar1/testfuse/main/ctrl/auth"
	"github.com/rvauradkar1/testfuse/main/ctrl/cache"
	"github.com/rvauradkar1/testfuse/main/ctrl/cart"
	"github.com/rvauradkar1/testfuse/main/ctrl/ord"
	"github.com/rvauradkar1/testfuse/main/ctrl/ord/db"
	"github.com/rvauradkar1/testfuse/main/find"

	"github.com/rvauradkar1/fuse/fuse"
)

func main() {
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

func funcName() (fuse.Fuse, []error) {
	cs := make([]fuse.Entry, 0)
	cs = append(cs, fuse.Entry{Name: "OrdCtrl", State: false, Instance: &ctrl.OrderController{}})
	cs = append(cs, fuse.Entry{Name: "CartSvc", State: false, Instance: &cart.CartSvc{}})
	cs = append(cs, fuse.Entry{Name: "AuthSvc", State: false, Instance: &auth.AuthSvc{}})
	cs = append(cs, fuse.Entry{Name: "CacheSvc", State: false, Instance: &cache.CacheSvc{}})
	cs = append(cs, fuse.Entry{Name: "DBSvc", State: false, Instance: &db.DBSvc{}})
	cs = append(cs, fuse.Entry{Name: "OrderSvc", State: true, Instance: &ord.OrderSvc{}})

	f := fuse.New()
	find.Find = f.Find
	errors := f.Register(cs)
	return f, errors
}

func call(i int) {
	fmt.Println(runtime.Caller(1))

}

/*
func genMocks() {
	m := mock.MockGen{}
	comps := make([]mock.Component, 0)
	path := "/Users/rvauradkar/go_code/src/github.com/rvauradkar1/testfuse/main"
	fmt.Println("path ====== " + path)
	comps = append(comps, mock.Component{Instance: &ctrl.OrderController{}, Basepath: path + "/ctrl", Name: "OrdCtrl"})
	comps = append(comps, mock.Component{Instance: &cart.CartSvc{}, Basepath: path + "/ctrl/cart", Name: "CartSvc"})
	comps = append(comps, mock.Component{Instance: &auth.AuthSvc{}, Basepath: path + "/ctrl/auth", Name: "AuthSvc"})
	comps = append(comps, mock.Component{Instance: &cache.CacheSvc{}, Basepath: path + "/ctrl/cache", Name: "CacheSvc"})
	comps = append(comps, mock.Component{Instance: &db.DBSvc{}, Basepath: path + "/ctrl/ord/db", Name: "DBSvc"})
	comps = append(comps, mock.Component{Instance: &ord.OrderSvc{}, Basepath: path + "/ctrl/ord", Name: "OrderSvc"})
	m.Comps = comps
	m.Generate()
}
*/
