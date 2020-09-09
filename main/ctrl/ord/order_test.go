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
	fmt.Println("Number of calls to MockDBSvc_AddOrder = ", NumCalls("MockDBSvc_AddOrder"))
	// Print params passed to MockDBSvc_AddOrder
	for i, c := range CallParams("MockDBSvc_AddOrder") {
		fmt.Println(" Call Number ", i, " Params` = ", c)
	}
	// Ensure call to MockCacheSvc_AddOrd was made
	fmt.Println("Number of calls to MockCacheSvc_AddOrd = ", NumCalls("MockCacheSvc_AddOrd"))
	// Print params passed to MockCacheSvc_AddOrd
	for i, c := range CallParams("MockCacheSvc_AddOrd") {
		fmt.Println(" Call Number ", i, " Params` = ", c)
	}
}
