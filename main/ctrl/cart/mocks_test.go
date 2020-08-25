

package cart
import (
"github.com/rvauradkar1/testfuse/main/ctrl/cache"
"github.com/rvauradkar1/testfuse/main/ctrl/ord/db"

)

// Begin of mock for CacheSvc and its methods
type MockCacheSvc struct{
	
}





type AddCrt func(s1 string) (error)
var MockCacheSvc_AddCrt AddCrt
func (p *MockCacheSvc) AddCrt(s1 string) (error) {
	return MockCacheSvc_AddCrt( s1)
}




type AddOrd func(s1 string) (error)
var MockCacheSvc_AddOrd AddOrd
func (p *MockCacheSvc) AddOrd(s1 string) (error) {
	return MockCacheSvc_AddOrd( s1)
}

// End of mock for CacheSvc and its methods

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

// Begin of mock for OrderSvc and its methods
type MockOrderSvc struct{
	CacheSvc cache.IService
DBSvc db.IService

}





type SaveOrder func(s1 string) (error)
var MockOrderSvc_SaveOrder SaveOrder
func (p *MockOrderSvc) SaveOrder(s1 string) (error) {
	return MockOrderSvc_SaveOrder( s1)
}

// End of mock for OrderSvc and its methods

// Begin of mock for CartSvc and its methods
type MockCartSvc struct{
	CacheSvc cache.IService
DBSvc db.IService
DEPS_ interface {}

}





type Add func(s1 string) (error)
var MockCartSvc_Add Add
func (p *MockCartSvc) Add(s1 string) (error) {
	return MockCartSvc_Add( s1)
}

// End of mock for CartSvc and its methods

