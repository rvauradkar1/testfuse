package main

import (
	"fmt"

	"github.com/rvauradkar1/fuse/fuse"
)

func main() {
	fmt.Println("Hello testfuse")

	fuse := fuse.New()
	fmt.Println(fuse)
}
