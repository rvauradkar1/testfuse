package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/rvauradkar1/fuse/mock"
)

func Test_Generate(t *testing.T) {
	fmt.Println("Start mock generation....")
	// Instantiaze mock, "main" is the name of the directory where your application resides
	m := mock.New("main")
	// Reuse the 'Entries' method from 'main.go'
	entries := Entries()
	// Register the entries, mocks are generated based upon the component structures
	errors := m.Register(entries)
	if len(errors) != 0 {
		log.Fatal(errors)
	}
	// Generate the mocks
	errors = m.Generate()
	if len(errors) != 0 {
		log.Fatal(errors)
	}
	fmt.Println("End mock generation....")
}
