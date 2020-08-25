

package ctrl
import (
"github.com/rvauradkar1/testfuse/main/ctrl/cart"
"github.com/rvauradkar1/testfuse/main/ctrl/auth"
"github.com/rvauradkar1/testfuse/main/ctrl/cache"
"github.com/rvauradkar1/testfuse/main/ctrl/ord/db"
"time"

)

// Begin of mock for AuthSvc and its methods
type MockAuthSvc struct{
	t time.Duration

}





type Auth func(s1 string) (error)
var MockAuthSvc_Auth Auth
func (p *MockAuthSvc) Auth(s1 string) (error) {
	return MockAuthSvc_Auth( s1)
}

// End of mock for AuthSvc and its methods

// Begin of mock for OrderController and its methods
type MockOrderController struct{
	CartSvc *cart.CartSvc
AuthSvc auth.IService

}





type Order func(s1 string,s2 string) (error)
var MockOrderController_Order Order
func (p *MockOrderController) Order(s1 string,s2 string) (error) {
	return MockOrderController_Order( s1, s2)
}

// End of mock for OrderController and its methods

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

