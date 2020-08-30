

package auth
import (
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

