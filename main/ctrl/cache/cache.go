package cache

type IService interface {
	AddOrder(product string) error
	AddCart(cart string) error
}

var cartCounts map[string]*int
var orderCounts map[string]*int

type Service struct {
}

func init() {
	cartCounts = make(map[string]*int, 0)
	orderCounts = make(map[string]*int, 0)
}

func (c *Service) AddCart(cart string) error {
	if count, ok := cartCounts[cart]; ok {
		*count++
		return nil
	}
	count := 1
	cartCounts[cart] = &count
	return nil
}

func (c *Service) AddOrder(order string) error {
	if count, ok := orderCounts[order]; ok {
		*count++
		return nil
	}
	count := 1
	orderCounts[order] = &count
	return nil
}
