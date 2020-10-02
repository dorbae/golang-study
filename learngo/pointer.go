package learngo

import "fmt"

func pointerExample() {
	//
	// Pointer
	//
	a := 2
	b := a
	a = 10
	fmt.Println(a, b)
	fmt.Println(&a, &b)

	c := 2
	cAddress := &c
	fmt.Println(c, cAddress, *cAddress)
}
