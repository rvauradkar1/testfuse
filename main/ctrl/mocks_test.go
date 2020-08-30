

package ctrl
import (
"github.com/rvauradkar1/testfuse/main/ctrl/cart"
"github.com/rvauradkar1/testfuse/main/ctrl/auth"
"github.com/rvauradkar1/testfuse/main/ctrl/cache"
"github.com/rvauradkar1/testfuse/main/ctrl/ord/db"
"time"

)

// Start of method calls and parameter capture
var stats = make(map[string]*FuncCalls, 0)

type FuncCalls struct {
	Count  int
	Params [][]interface{}
}

func (f FuncCalls) First() []interface{} {
	for _, p := range f.Params {
		return p
	}
	return []interface{}{}
}

func (f FuncCalls) All() [][]interface{} {
	return f.Params
}

func (f FuncCalls) NumCalls() int {
	return f.Count
}

func capture(key string, params []interface{}) {
	val, ok := stats[key]
	if !ok {
		val = &FuncCalls{}
		val.Count = 0
		val.Params = make([][]interface{}, 0)
		stats[key] = val
	}
	val.Count++
	val.Params = append(val.Params, params)
}

func calls(key string) FuncCalls {
	if val, ok := stats[key]; ok {
		return *val
	}
	return FuncCalls{}
}
// End of method calls and parameter capture


// Begin of mock for AuthSvc and its methods
type MockAuthSvc struct{
	t time.Duration

}





type Auth func(s1 string) (error)
var MockAuthSvc_Auth Auth
func (p *MockAuthSvc) Auth(s1 string) (error) {
	capture("MockAuthSvc_Auth", []interface{}{s1 })
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
	capture("MockOrderController_Order", []interface{}{s1 ,s2 })
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
	capture("MockCartSvc_Add", []interface{}{s1 })
	return MockCartSvc_Add( s1)
}

// End of mock for CartSvc and its methods

