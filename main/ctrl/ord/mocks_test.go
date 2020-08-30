package ord

import (
	"github.com/rvauradkar1/testfuse/main/ctrl/cache"
	"github.com/rvauradkar1/testfuse/main/ctrl/ord/db"
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

// Begin of mock for CacheSvc and its methods
type MockCacheSvc struct {
}

type AddCrt func(s1 string) error

var MockCacheSvc_AddCrt AddCrt

func (p *MockCacheSvc) AddCrt(s1 string) error {
	capture("MockCacheSvc_AddCrt", []interface{}{s1})
	return MockCacheSvc_AddCrt(s1)
}

type AddOrd func(s1 string, s2 string) error

var MockCacheSvc_AddOrd AddOrd

func (p *MockCacheSvc) AddOrd(s1 string, s2 string) error {
	capture("MockCacheSvc_AddOrd", []interface{}{s1, s2})
	return MockCacheSvc_AddOrd(s1, s2)
}

// End of mock for CacheSvc and its methods

// Begin of mock for DBSvc and its methods
type MockDBSvc struct {
}

type AddCart func(s1 string) error

var MockDBSvc_AddCart AddCart

func (p *MockDBSvc) AddCart(s1 string) error {
	capture("MockDBSvc_AddCart", []interface{}{s1})
	return MockDBSvc_AddCart(s1)
}

type AddOrder func(s1 string) error

var MockDBSvc_AddOrder AddOrder

func (p *MockDBSvc) AddOrder(s1 string) error {
	capture("MockDBSvc_AddOrder", []interface{}{s1})
	return MockDBSvc_AddOrder(s1)
}

// End of mock for DBSvc and its methods

// Begin of mock for OrderSvc and its methods
type MockOrderSvc struct {
	CacheSvc cache.IService
	DBSvc    db.IService
}

type SaveOrder func(s1 string, s2 string) error

var MockOrderSvc_SaveOrder SaveOrder

func (p *MockOrderSvc) SaveOrder(s1 string, s2 string) error {
	capture("MockOrderSvc_SaveOrder", []interface{}{s1, s2})
	return MockOrderSvc_SaveOrder(s1, s2)
}

// End of mock for OrderSvc and its methods
