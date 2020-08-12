
package db
import (

)

type MockDBSvc struct{
	
}




type AddCart func(s1 string) (error)
var AddCartFunc AddCart
func (p *MockDBSvc) AddCart(s1 string) (error) {
	return AddCartFunc( s1)
}


type AddOrder func(s1 string) (error)
var AddOrderFunc AddOrder
func (p *MockDBSvc) AddOrder(s1 string) (error) {
	return AddOrderFunc( s1)
}


