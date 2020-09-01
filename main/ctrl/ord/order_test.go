package ord

import (
	"fmt"
	"testing"
)

func Test_SaveOrder(t *testing.T) {
	// Make instance and initialize mocked dependencies.
	svc := OrderSvc{DBSvc: &MockDBSvc{}, CacheSvc: &MockCacheSvc{}}
	// Mock dependent methods.
	MockDBSvc_AddOrder = func(order string) error { return nil }
	MockCacheSvc_AddOrd = func(cart string, status string) error { return nil }
	// Execute logic
	err := svc.SaveOrder("new order", "stats")
	// Error handling omitted
	fmt.Println(err)
	// Ensure call to MockDBSvc_AddOrder was made
	for i, c := range Calls("MockDBSvc_AddOrder") {
		fmt.Println(" i ", i, " c = ", c)
	}
	// Ensure call to MockCacheSvc_AddOrd was made
	for i, c := range Calls("MockCacheSvc_AddOrd") {
		fmt.Println(" i ", i, " c = ", c)
	}
}
