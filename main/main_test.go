package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/rvauradkar1/fuse/mock"
)

func TestMain(m *testing.M) {
	fmt.Println("Begin test main....")
	os.Exit(m.Run())
	fmt.Println("End test main....")
}

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

/*
func Test_reg(t *testing.T) {
	m := mock.New("main")
	entries := Entries()
	errors := m.Register(entries)
	fmt.Println(errors)

	m.Generate()

}
*/
