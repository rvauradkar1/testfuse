

package ctrl
import (
"github.com/rvauradkar1/testfuse/main/ctrl/cart"
"github.com/rvauradkar1/testfuse/main/ctrl/auth"
"github.com/rvauradkar1/testfuse/main/ctrl/cache"
"github.com/rvauradkar1/testfuse/main/ctrl/ord/db"

)

// Begin of mock for AuthSvc and its methods
type MockAuthSvc struct{
	
}



type Auth func(s1 string) (error)
var MockAuth Auth
func (p *MockAuthSvc) Auth(s1 string) (error) {
	return MockAuth( s1)
}

// End of mock for AuthSvc and its methods

// Begin of mock for OrderController and its methods
type MockOrderController struct{
	CartSvc *cart.CartSvc
AuthSvc auth.IService

}



type Order func(s1 string,s2 string) (error)
var MockOrder Order
func (p *MockOrderController) Order(s1 string,s2 string) (error) {
	return MockOrder( s1, s2)
}

// End of mock for OrderController and its methods

// Begin of mock for CartSvc and its methods
type MockCartSvc struct{
	CacheSvc cache.IService
DBSvc db.IService
DEPS_ interface {}

}



type Add func(s1 string) (error)
var MockAdd Add
func (p *MockCartSvc) Add(s1 string) (error) {
	return MockAdd( s1)
}

// End of mock for CartSvc and its methods

