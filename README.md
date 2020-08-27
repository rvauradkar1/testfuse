# testfuse

Simple project to demonstrate usage of the fuse and mock packages.  
These packages are available at the repo - [https://github.com/rvauradkar1/fuse](https://github.com/rvauradkar1/fuse)

The following component graph is used to demonstrate the usage of fuse/mock.  
Please review the package structure of this project - folders and sub-folders are provided to
demonstrate wiring of components and generation of mocks to aid unit-testing.

A graph of the components is provided below.

1. The vertices represent componets.
2. The edges represent dependencies.
3. The labels represent state.

<img src="graph.png" alt="Component Graph" title="Component Graph" class="absent" />


## Usage of package fuse

**API for fuse:**

```
// Fuse is used by clients to configure dependency injection (DI)
type Fuse interface {
	// Register a slice of components
	Register(entries []Entry) []error
	// Wire injects dependencies into components.
	Wire() []error
	// Find is needed only for stateful components. Can also be used for 
	// stateless in case dependencies are not defined at instance level
	Find(name string) interface{}
}
```

Steps to follow:

<ol>
<li><a href="#step1">Declare components and configure dependencies</a></li>
<li><a href="#step2">Create a slice of component entries</a></li>
<li><a href="#step3">Create a config package to avoid cyclic dependencies</a></li>
<li><a href="#step4">Register and fuse the components (Dependency Injection pattern)</a></li>
<li><a href="#step5">Provide a finder to find stateful components (Resource Locator pattern)</a></li>
</ol>

1. Declare components and configure dependencies
2. Create a slice of component entries
3. Create a config package to avoid cyclic dependencies
4. Register and fuse the components (Dependency Injection pattern)
5. Provide a finder to find stateful components (Resource Locator pattern)


<div id="step1"><h3>1. Declare components and configure dependencies</h3></div>

Dependencies are configured using `struct` tags.

**Example**: The code below defines a `struct` which is later registered as a component. This component
depends on 2 other components (look at the component graph above)

1. CartSvc (pointer variable) - tagged with key _fuse and the name of the component.
2. AuthSvc (interface variable) - tagged with key _fuse and the name of the component.
```
type OrderController struct {
	CartSvc *cart.CartSvc `_fuse:"CartSvc"`
	AuthSvc auth.IService `_fuse:"AuthSvc"`
}
```
<div id="step1"><h3>2. Create a config package to avoid cyclic dependencies</h3></div>

**Note:** Config package is needed for several reasons:

1. Application packages cannot refer to 'main'.
2. Go does not allow cyclical references.
3. Application packges do not need to import 'fuse', they will import 'cfg' and
use just the 'find' method. Eases refactoring if projects want to discontinue usage of 'fuse'.


Following diagrams shows a high-level dependency:
1. Application package 'main' depends on 'cfg' and 'fuse'. 'main' provides 'cfg' with the list of component entries.
2. 'cfg' uses 'fuse', register and fuses the components together. The 'find' method
from the 'fuse' package is assigned to the 'find' method from the 'cfg' package. Other packages use this method
without importing 'fuse'
3. 'other packages' are all other application packages. They depend on 'cfg' and use only the 'find' method.

<img src="cfg.png" alt="Config Graph" title="Config Graph" class="absent" />


<div id="step1"><h3>3. Create a slice of component entries</h3></div>

The component slice is created in the application 'main' package. Helps in:
1. Isolating dependencies on the 'fuse' package.
2. One place to define the graph(s).
3. Multiple slices can be creatd in case of different component graphs.
4. 'mock' package uses these slices to create mock for unit-testing.

**Example:**

The code below is defined in the application main file. It can be named anything. This same method will be used by the mock generator.
Each component entry is recorded as a fuse.Entry with properties:

1. Name : Unique identifier for the component. This same identifier is used when defining the '_fuse' dependency.
2. State : Stateful or stateless. If stateless, the component instance provided is always used, no new instances are created. If stateful,
a new instance is created for each use. If this component has dependencies on other stateless components, they are recursively populated as well.
3. Instance : A pointer to the `struct`, if it is not a pointer variable, an erros is generated during registry process.


```
func Entries() []fuse.Entry {
	entries := make([]fuse.Entry, 0)
	entries = append(entries, fuse.Entry{Name: "OrdCtrl", State: false, Instance: &ctrl.OrderController{}})
	entries = append(entries, fuse.Entry{Name: "CartSvc", State: false, Instance: &cart.CartSvc{}})
	entries = append(entries, fuse.Entry{Name: "AuthSvc", State: false, Instance: &auth.AuthSvc{}})
	entries = append(entries, fuse.Entry{Name: "CacheSvc", State: false, Instance: &cache.CacheSvc{}})
	entries = append(entries, fuse.Entry{Name: "DBSvc", State: false, Instance: &db.DBSvc{}})
	entries = append(entries, fuse.Entry{Name: "OrderSvc", State: true, Instance: &ord.OrderSvc{}})

	return entries
}

func main() {
	fmt.Println("Hello testfuse")
	entries := Entries()
	errors := cfg.Fuse(entries)
	..........
}
```
**main():** Uses 'Entries()' and calls 'cfg.Fuse()'. 'cfg' registers and fuses.
```
entries := Entries()
errors := cfg.Fuse(entries)
```


<div id="step4"><h3>4. Register and fuse the components (Dependency Injection pattern)</h3></div>
### 5. Register and fuse the components (Dependency Injection pattern)

```
package cfg

import (
	"github.com/rvauradkar1/fuse/fuse"
)

// Find is used by application packages as a Resource Locator
var Find func(name string) interface{}

// Fuse is used by application main package to provide a list of compoenets to register and fuse
func Fuse(entries []fuse.Entry) []error {
    // Step 1. Instance of Fuse
	f := fuse.New()
	// Step 2. 'cfg.Find' now points to the 'fuse.Find'
	Find = f.Find
	// Step 3. Register entries
	errors := f.Register(entries)
	if len(errors) != 0 {
		return errors
	}
	// Step 4. Wire dependencies
	return f.Wire()
}
```


<div id="step5"><h3>5. Provide a finder to find stateful components (Resource Locator pattern)</h3></div>

Code snippet from above provides this functionality:

```
// Find is used by application packages as a Resource Locator
var Find func(name string) interface{}
```

```
// 'cfg.Find' now points to the 'fuse.Find'
Find = f.Find
```


## Usage of package mock

**API for mock:**

```
type Mock interface {
	// Register a slice of components
	Register(entries []fuse.Entry) []error
	// Find is needed primarily for stateful components.
	Find(name string) interface{}
	// Generates mocks
	Generate() []error
}
```

The graph below shows how the mock code generation is laid out.

1. main_test.go - Test file for the application main. It depends on the 'mock' package.
It uses the 'Entries' method of 'main.go' to register the components.
2. main.go - The application main file that supplies the component entries.
3. mock package - depends on the 'fuse' package.

<img src="mock.png" alt="Mock" title="Mock" class="absent" />

### Mock code generation pattern

All mocks are generated based on a fixed pattern.

'AuthSvc' is a registered component. Clients of this package need to mock this dependency during unit-testing.


```
package auth

import (
	"fmt"
	"time"
)

type IService interface {
	Auth(user string) error
}

type AuthSvc struct {
	t time.Duration
}

func (a *AuthSvc) Auth(user string) error {
	fmt.Printf("Auth for user [%s]\n", user)
	return nil
}
```

```
package auth

import (
	"time"
)

// Begin of mock for AuthSvc and its methods
type MockAuthSvc struct {
	t time.Duration
}

type Auth func(s1 string) error

var MockAuthSvc_Auth Auth

func (p *MockAuthSvc) Auth(s1 string) error {
	return MockAuthSvc_Auth(s1)
}

// End of mock for AuthSvc and its methods
```

**** Struct mock

```
// Original
type AuthSvc struct {
	t time.Duration
}
// Mock of Original
type MockAuthSvc struct {
	t time.Duration
}
```

#### Method mock

```
// Method from interface
type IService interface {
	Auth(user string) error
}
// Method mock
type Auth func(s1 string) error

// Client provide a mock implementation
var MockAuthSvc_Auth Auth

func (p *MockAuthSvc) Auth(s1 string) error {
    // Mocked method is used instead of real code
	return MockAuthSvc_Auth(s1)
} 
```

