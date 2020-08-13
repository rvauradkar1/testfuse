

package cache
import (

)

// Begin of mock for CacheSvc and its methods
type MockCacheSvc struct{
	
}



type AddCrt func(s1 string) (error)
var MockAddCrt AddCrt
func (p *MockCacheSvc) AddCrt(s1 string) (error) {
	return MockAddCrt( s1)
}


type AddOrd func(s1 string) (error)
var MockAddOrd AddOrd
func (p *MockCacheSvc) AddOrd(s1 string) (error) {
	return MockAddOrd( s1)
}

// End of mock for CacheSvc and its methods

