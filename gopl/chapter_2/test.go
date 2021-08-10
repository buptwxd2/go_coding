package main

import "fmt"

func main()  {

	ascii := 'a'
	unicode := '国'
	newline := '\n'

	fmt.Printf("Type is %T\n", ascii)

	fmt.Printf("%x, %o, %b, %d\n", unicode, unicode, unicode, unicode)

	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]q\n", newline)


	fmt.Println("============")

	s := "中"

	fmt.Printf("Length is %d\n", len(s))

	fmt.Println(s[0], s[1], s[2])
	fmt.Printf("Type is %T\n", s)
	fmt.Printf("s[0] type is %T\n", s[0])





}
