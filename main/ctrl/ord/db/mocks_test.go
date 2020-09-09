

package db
import (

)

// Start of method calls and parameter capture
var stats = make(map[string]*funcCalls, 0)

type funcCalls struct {
	Count  int
	Params [][]interface{}
}

type CallInfo struct {
	Ok     bool
	Name   string
	Params []interface{}
}

type Params []interface{}

func NumCalls(name string) int {
	call := forCall(name)
	return call.Count
}

func CallParams(name string) []Params {
	call := forCall(name)
	if call.Count > 0 {
		calls := make([]Params, 0)
		for i := 0; i < call.Count; i++ {
			calls = append(calls, call.Params[i])
		}
		return calls
	}
	return []Params{}
}

func capture(key string, params []interface{}) {
	val, ok := stats[key]
	if !ok {
		val = &funcCalls{}
		val.Count = 0
		val.Params = make([][]interface{}, 0)
		stats[key] = val
	}
	val.Count++
	val.Params = append(val.Params, params)

}

func forCall(key string) funcCalls {
	if val, ok := stats[key]; ok {
		return *val
	}
	return funcCalls{}
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

