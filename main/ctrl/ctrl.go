package ctrl

import (
	"github.com/rvauradkar1/testfuse/main/ctrl/cart"
)

type OrderController struct {
	s       string
	CartSvc *cart.CartService `_fuse:"CartSvc"`
}

func (o *OrderController) Order(order string) error {
	o.CartSvc.SaveOrder(order)
	return nil
}
