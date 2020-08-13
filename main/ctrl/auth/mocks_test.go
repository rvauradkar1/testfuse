

package auth
import (

)

// Begin of mock for AuthSvc and its methods
type MockAuthSvc struct{
	
}



type Auth func(s1 string) (error)
var MockAuth Auth
func (p *MockAuthSvc) Auth(s1 string) (error) {
	return MockAuth( s1)
}

// End of mock for AuthSvc and its methods

