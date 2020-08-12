
package cache
import (

)

type MockCacheSvc struct{
	
}




type AddCrt func(s1 string) (error)
var AddCrtFunc AddCrt
func (p *MockCacheSvc) AddCrt(s1 string) (error) {
	return AddCrtFunc( s1)
}


type AddOrd func(s1 string) (error)
var AddOrdFunc AddOrd
func (p *MockCacheSvc) AddOrd(s1 string) (error) {
	return AddOrdFunc( s1)
}


