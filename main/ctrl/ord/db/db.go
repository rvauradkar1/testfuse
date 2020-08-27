package db

import "fmt"

var cartCounts map[string]*int
var orderCounts map[string]*int

func init() {
	cartCounts = make(map[string]*int, 0)
	orderCounts = make(map[string]*int, 0)
}

type IService interface {
	AddOrder(product string) error
	AddCart(cart string) error
}

type DBSvc struct {
}

func (db *DBSvc) AddCart(cart string) error {
	fmt.Println("Inside AddCart of DBSvc")
	if count, ok := cartCounts[cart]; ok {
		*count++
		return nil
	}
	count := 1
	cartCounts[cart] = &count
	return nil
}

func (db *DBSvc) AddOrder(order string) error {
	if count, ok := orderCounts[order]; ok {
		*count++
		return nil
	}
	count := 1
	orderCounts[order] = &count
	return nil
}
