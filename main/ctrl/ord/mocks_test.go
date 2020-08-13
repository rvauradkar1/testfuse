

package ord
import (
"github.com/rvauradkar1/testfuse/main/ctrl/cache"
"github.com/rvauradkar1/testfuse/main/ctrl/ord/db"

)

// Begin of mock for CacheSvc and its methods
type MockCacheSvc struct{
	
}



type AddCrt func(s1 string) (error)
var MockAddCrt AddCrt
func (p *MockCacheSvc) AddCrt(s1 string) (error) {
	return MockAddCrt( s1)
}


type AddOrd func(s1 string) (error)
var MockAddOrd AddOrd
func (p *MockCacheSvc) AddOrd(s1 string) (error) {
	return MockAddOrd( s1)
}

// End of mock for CacheSvc and its methods

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

// Begin of mock for OrderSvc and its methods
type MockOrderSvc struct{
	CacheSvc cache.IService
DBSvc db.IService

}



type SaveOrder func(s1 string) (error)
var MockSaveOrder SaveOrder
func (p *MockOrderSvc) SaveOrder(s1 string) (error) {
	return MockSaveOrder( s1)
}

// End of mock for OrderSvc and its methods

