

package db
import (

)

// Begin of mock for DBSvc and its methods
type MockDBSvc struct{
	
}





type AddCart func(s1 string) (error)
var MockDBSvc_AddCart AddCart
func (p *MockDBSvc) AddCart(s1 string) (error) {
	return MockDBSvc_AddCart( s1)
}




type AddOrder func(s1 string) (error)
var MockDBSvc_AddOrder AddOrder
func (p *MockDBSvc) AddOrder(s1 string) (error) {
	return MockDBSvc_AddOrder( s1)
}

// End of mock for DBSvc and its methods

