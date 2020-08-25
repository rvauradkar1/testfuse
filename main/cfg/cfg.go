package cfg

import (
	"github.com/rvauradkar1/fuse/fuse"
)

var Find func(name string) interface{}

/*
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

*/

func Fuse(entries []fuse.Entry) []error {
	f := fuse.New()
	Find = f.Find
	errors := f.Register(entries)
	if len(errors) != 0 {
		return errors
	}
	errors = f.Wire()
	return errors
}

/*
func Mock() {
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
