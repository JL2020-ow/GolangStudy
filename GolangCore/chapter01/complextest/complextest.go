package main

import (
	"fmt"
)

func main() {
	var v = complex(2.1, 3)
	a := real(v)
	b := imag(v)
	fmt.Println(a)
	fmt.Println(b)
}
