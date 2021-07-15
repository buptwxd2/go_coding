package main

import "fmt"
import "strconv"

func main(){
	test()
	//x := '世'
	//fmt.Printf("Type is %T", x)
	////fmt.Println("Len is", len(x))
	//fmt.Printf("%x", x)
	//
	//y := "世"
	//fmt.Println()
	//fmt.Printf("Type is %T", y)
	//fmt.Println("Len is", len(y))
	//fmt.Printf("%x", y)
}

func test()  {
	x, err := strconv.Atoi("65")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("Type is %T. And the value is %v\n", x, x)
	fmt.Printf("The binary format is %b", x)


}


