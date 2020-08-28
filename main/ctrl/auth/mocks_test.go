

package auth
import (
"time"

)

// Begin of mock for AuthSvc and its methods
type MockAuthSvc struct{
	t time.Duration

}





type Auth func(s1 string) (error)
var MockAuthSvc_Auth Auth
func (p *MockAuthSvc) Auth(s1 string) (error) {
	return MockAuthSvc_Auth( s1)
}

// End of mock for AuthSvc and its methods

