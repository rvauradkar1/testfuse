package ctrl

import (
	"fmt"

	"github.com/rvauradkar1/testfuse/main/ctrl/auth"
	"github.com/rvauradkar1/testfuse/main/ctrl/cart"
)

type OrderController struct {
	CartSvc *cart.Service `_fuse:"CartSvc"`
	AuthSvc auth.IService `_fuse:"AuthSvc"`
}

func (o *OrderController) Order(user string, order string) error {
	fmt.Println("Calling Auth")
	err := o.AuthSvc.Auth(user)
	if err != nil {
		return err
	}
	fmt.Println("Calling Add Order")
	err = o.CartSvc.Add(order)
	return err
}
