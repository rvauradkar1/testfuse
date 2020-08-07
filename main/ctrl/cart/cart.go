package cart

import "github.com/rvauradkar1/testfuse/main/ctrl/cache"

type ICartService interface {
	Add(cart string) error
}

type CartService struct {
	CacheSvc cache.ICacheService `_fuse:"CacheSvc"`
}

func (c *CartService) Add(cart string) error {
	err := c.CacheSvc.AddCart(cart)
	return err
}
