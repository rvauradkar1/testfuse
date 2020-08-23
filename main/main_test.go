package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	f, errors := funcName()
	fmt.Println(f)
	fmt.Println(errors)
	os.Exit(m.Run())
}

func Test_s(t *testing.T) {
	fmt.Println("tttt")
}
