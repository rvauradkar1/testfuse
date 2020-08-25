package cfg

import (
	"github.com/rvauradkar1/fuse/fuse"
)

// Find is used by application packages as a Resource Locator
var Find func(name string) interface{}

// Fuse is used by application main package to provide a list of compoenets to register and fuse
func Fuse(entries []fuse.Entry) []error {
	f := fuse.New()
	Find = f.Find
	errors := f.Register(entries)
	if len(errors) != 0 {
		return errors
	}
	return f.Wire()
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
