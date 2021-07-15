package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	l := len(os.Args)
	if l == 0 || l == 1{
		return
	}

	s := ""
	for _, arg := range(os.Args[1:]){
		s += arg + " "
	}
	s = strings.TrimSpace(s)
	fmt.Println(s)


	//var s, sep string
	//for i := 1; i < len(os.Args); i++{
	//	s += sep + os.Args[i]
	//	sep = " "
	//}
	//fmt.Println(s)
}
