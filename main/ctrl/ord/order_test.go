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

	for i, c := range Calls("MockDBSvc_AddOrder") {
		fmt.Println(" i ", i, " c = ", c)
	}

	for i, c := range Calls("MockCacheSvc_AddOrd") {
		fmt.Println(" i ", i, " c = ", c)
	}

}
