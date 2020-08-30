

package db
import (

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


// Begin of mock for DBSvc and its methods
type MockDBSvc struct{
	
}





type AddCart func(s1 string) (error)
var MockDBSvc_AddCart AddCart
func (p *MockDBSvc) AddCart(s1 string) (error) {
	capture("MockDBSvc_AddCart", []interface{}{s1 })
	return MockDBSvc_AddCart( s1)
}




type AddOrder func(s1 string) (error)
var MockDBSvc_AddOrder AddOrder
func (p *MockDBSvc) AddOrder(s1 string) (error) {
	capture("MockDBSvc_AddOrder", []interface{}{s1 })
	return MockDBSvc_AddOrder( s1)
}

// End of mock for DBSvc and its methods

