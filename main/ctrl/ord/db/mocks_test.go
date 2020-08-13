

package db
import (

)

// Begin of mock for DBSvc and its methods
type MockDBSvc struct{
	
}



type AddCart func(s1 string) (error)
var MockAddCart AddCart
func (p *MockDBSvc) AddCart(s1 string) (error) {
	return MockAddCart( s1)
}


type AddOrder func(s1 string) (error)
var MockAddOrder AddOrder
func (p *MockDBSvc) AddOrder(s1 string) (error) {
	return MockAddOrder( s1)
}

// End of mock for DBSvc and its methods

