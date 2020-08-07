package cache

type ICacheService interface {
	AddProduct(product string) error
	AddCart(cart string) error
}

type CacheService struct {
	ProductCounts map[string]*int
	OrderCounts   map[string]*int
}

func (c *CacheService) init() {
	c.ProductCounts = make(map[string]*int, 0)
	c.OrderCounts = make(map[string]*int, 0)
}

func (a *CacheService) auth(user string) bool {
	return true
}

func (c *CacheService) AddProduct(product string) error {
	if count, ok := c.ProductCounts[product]; ok {
		*count++
		return nil
	}
	count := 1
	c.ProductCounts[product] = &count
	return nil
}

func (c *CacheService) AddCart(product string) error {
	if count, ok := c.OrderCounts[product]; ok {
		*count++
		return nil
	}
	count := 1
	c.OrderCounts[product] = &count
	return nil
}
