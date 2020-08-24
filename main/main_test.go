package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/rvauradkar1/fuse/fuse"
	"github.com/rvauradkar1/fuse/mock"
	"github.com/rvauradkar1/testfuse/main/ctrl"
	"github.com/rvauradkar1/testfuse/main/ctrl/auth"
	"github.com/rvauradkar1/testfuse/main/ctrl/cache"
	"github.com/rvauradkar1/testfuse/main/ctrl/cart"
	"github.com/rvauradkar1/testfuse/main/ctrl/ord"
	"github.com/rvauradkar1/testfuse/main/ctrl/ord/db"
)

func TestMain(m *testing.M) {
	f, errors := funcName()
	fmt.Println(f)
	fmt.Println(errors)
	os.Exit(m.Run())
}

func Test_s(t *testing.T) {
	fmt.Println("tttt")
}

func Test_register(t *testing.T) {
	m := mock.New("mock")
	entries := make([]fuse.Entry, 0)
	entries = append(entries, fuse.Entry{Name: "OrdCtrl", Instance: &ctrl.OrderController{}})
	entries = append(entries, fuse.Entry{Name: "CartSvc", Instance: &cart.CartSvc{}})
	entries = append(entries, fuse.Entry{Name: "AuthSvc", Instance: &auth.AuthSvc{}})
	entries = append(entries, fuse.Entry{Name: "CacheSvc", Instance: &cache.CacheSvc{}})
	entries = append(entries, fuse.Entry{Name: "DBSvc", Instance: &db.DBSvc{}})
	entries = append(entries, fuse.Entry{Name: "OrderSvc", Instance: &ord.OrderSvc{}})
	errors := m.Register(entries)
	fmt.Println("errors = ", errors)
	m.Generate2()
}
