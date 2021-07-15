package main

import "fmt"

func main()  {

	var x uint8 = 1 << 1 | 1 << 5

	fmt.Printf("%08b\n", x)

	f := 1e2
	fmt.Print(f)
	fmt.Printf("%d", f)

	y := int(f)
	fmt.Println(y)

}
