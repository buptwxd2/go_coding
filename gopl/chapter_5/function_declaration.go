package main

import "fmt"

func add(x, y int) int{return x + y}

func main()  {
	fmt.Printf("%T\n", add)
}
