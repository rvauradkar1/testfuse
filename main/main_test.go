package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/rvauradkar1/fuse/mock"
)

func TestMain(m *testing.M) {
	//f, errors := funcName()
	//fmt.Println(f)
	//fmt.Println(errors)
	fmt.Println("Begin test main....")
	os.Exit(m.Run())
	fmt.Println("End test main....")
}

func Test_s(t *testing.T) {
	fmt.Println("tttt")
}

/*
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
	m.Generate()
}

*/

func Test_reg(t *testing.T) {
	m := mock.New("main")
	entries := Entries()
	errors := m.Register(entries)
	fmt.Println(errors)

	m.Generate()

}
