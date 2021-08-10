package main

import "fmt"

func main()  {

	type Student struct {
		name string
		age	int
	}

	var s = Student{
		name: "alex",
		age: 10,
	}

	var b = struct {
		name string
		age  int
	}{
		name: "alex",
		age: 10,
	}


	fmt.Printf("Type is %T", s)
	fmt.Println()
	fmt.Printf("Type is %T", b)

}
