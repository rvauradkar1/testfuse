# testfuse

Simple project to demonstrate usage of the fuse and mock packages.  
These packages are available at the repo - [https://github.com/rvauradkar1/fuse](https://github.com/rvauradkar1/fuse)

Please review the package structure - folders and sub-folders are provided to
demonstrate wiring of components and generation of mocks to aid unit-testing.

A graph of the components is provided below.

1. The vertices represent componets.
2. The edges represent dependencies.
3. The labels represent state.

![Component Graph](graph.png "Component Graph")

## Usage of package fuse

Steps to follow:

1. Defining components and configuring dependencies
2. Defining component entries
3. Creating a config package to avoid cyclic dependencies
4. Fusing the components
5. Providing a finder to find stateful components

### Defining components and configuring dependencies

Dependencies are configured using `struct` tags.

Example: The code below defines a `struct` that is registered as a component. This component
depends on 2 other components (look at the component graph above)

1. CartSvc (pointer variable) - tagged with key _fuse and the name of the component.
2. AuthSvc (interface variable) - tagged with key _fuse and the name of the component.
```
type OrderController struct {
	CartSvc *cart.CartSvc `_fuse:"CartSvc"`
	AuthSvc auth.IService `_fuse:"AuthSvc"`
}
```




```
func Entries() []fuse.Entry {
	fmt.Println("Hello testfuse")
	entries := make([]fuse.Entry, 0)
	entries = append(entries, fuse.Entry{Name: "OrdCtrl", State: false, Instance: &ctrl.OrderController{}})
	entries = append(entries, fuse.Entry{Name: "CartSvc", State: false, Instance: &cart.CartSvc{}})
	entries = append(entries, fuse.Entry{Name: "AuthSvc", State: false, Instance: &auth.AuthSvc{}})
	entries = append(entries, fuse.Entry{Name: "CacheSvc", State: false, Instance: &cache.CacheSvc{}})
	entries = append(entries, fuse.Entry{Name: "DBSvc", State: false, Instance: &db.DBSvc{}})
	entries = append(entries, fuse.Entry{Name: "OrderSvc", State: true, Instance: &ord.OrderSvc{}})

	return entries
}
```

