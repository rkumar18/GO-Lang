package main

import (
	"fmt"
)

func main() {
	hello()
	sum :=add(2,3)
	fmt.Print(sum)
}

func hello(){
	fmt.Println("hello runs")
}

func add(a int, b int) int{
	return a+b
}