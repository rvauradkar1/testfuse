package ord

import (
	"fmt"
	"testing"
)

func Test_SaveOrder(t *testing.T) {
	svc := OrderSvc{DBSvc: &MockDBSvc{}, CacheSvc: &MockCacheSvc{}}
	MockDBSvc_AddOrder = func(order string) error { return nil }
	MockCacheSvc_AddOrd = func(cart string, status string) error { return nil }

	fmt.Println(svc.SaveOrder("new order", "stats"))

	fmt.Println("Num called = ", calls("MockDBSvc_AddOrder").NumCalls())

	for i, p := range calls("MockDBSvc_AddOrder").First() {
		fmt.Println(fmt.Sprintf("No : [%d], val = [%v]", i, p))
	}

	fmt.Println("Num called = ", calls("MockCacheSvc_AddOrd").NumCalls())

	for i, p := range calls("MockCacheSvc_AddOrd").First() {
		fmt.Println(fmt.Sprintf("No : [%d], val = [%v]", i, p))
	}

}
