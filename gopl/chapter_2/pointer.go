package main

import "fmt"

func main()  {
	x := foo()

	fmt.Printf("Address is %p", x)

	fmt.Printf("Value is %v", *x)

	*x = 10
	fmt.Printf("Value is %v", *x)

}

func foo() *int {
	var foo int = 5
	p := &foo

	make()

	return p
}
