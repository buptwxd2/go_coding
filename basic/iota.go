package main

import (
	"fmt"
)

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
)

func main()  {
	fmt.Println(Sunday, Monday, Tuesday)


	type Flags uint
	const(
		FlagUp Flags = 1 << iota
		FlagBroadcast
		FlagLoopback
		FlagPointToPoint
		FlagMulticast
	)

	fmt.Print(FlagUp, FlagBroadcast, FlagLoopback, FlagPointToPoint, FlagMulticast)

}
