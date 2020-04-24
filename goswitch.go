package main

import (
	"fmt"
)

func main(){
	a :=3
	switch a{
	case 1:
	fmt.Print("case one")
	case 2:
	fmt.Print("case two")
	default:
	fmt.Print("default")
	}
}