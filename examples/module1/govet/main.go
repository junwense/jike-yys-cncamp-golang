package main

import (
	"fmt"
)

func main() {
	name := "testing"
	s := "map be use these"
	fmt.Printf("%d\n", name)
	fmt.Printf("%s\n", name, name)

	fmt.Printf("%s\n", s)
	fmt.Printf("%s\n", name)
	fmt.Printf("%s%s\n", name, name)
}
