package cache

import "fmt"

type IService interface {
	AddOrd(product string, status string) error
	AddCrt(cart string) error
}

var cartCounts map[string]*int
var orderCounts map[string]*int

type CacheSvc struct {
}

func init() {
	cartCounts = make(map[string]*int, 0)
	orderCounts = make(map[string]*int, 0)
}

func (c *CacheSvc) AddCrt(cart string) error {
	fmt.Println("Inside AddCrt of CacheSvc")
	if count, ok := cartCounts[cart]; ok {
		*count++
		return nil
	}
	count := 1
	cartCounts[cart] = &count
	return nil
}

func (c *CacheSvc) AddOrd(order string, status string) error {
	fmt.Println("Inside AddOrd of CacheSvc")
	if count, ok := orderCounts[order]; ok {
		*count++
		return nil
	}
	count := 1
	orderCounts[order] = &count
	return nil
}
