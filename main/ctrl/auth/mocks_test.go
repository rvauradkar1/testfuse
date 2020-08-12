
package auth
import (

)

type MockAuthSvc struct{
	
}




type Auth func(s1 string) (error)
var AuthFunc Auth
func (p *MockAuthSvc) Auth(s1 string) (error) {
	return AuthFunc( s1)
}


