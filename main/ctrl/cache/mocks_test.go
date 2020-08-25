

package cache
import (

)

// Begin of mock for CacheSvc and its methods
type MockCacheSvc struct{
	
}





type AddCrt func(s1 string) (error)
var MockCacheSvc_AddCrt AddCrt
func (p *MockCacheSvc) AddCrt(s1 string) (error) {
	return MockCacheSvc_AddCrt( s1)
}




type AddOrd func(s1 string) (error)
var MockCacheSvc_AddOrd AddOrd
func (p *MockCacheSvc) AddOrd(s1 string) (error) {
	return MockCacheSvc_AddOrd( s1)
}

// End of mock for CacheSvc and its methods

