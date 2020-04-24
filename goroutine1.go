package main

import (  
    "fmt"
)

func hello(i int) int {  
    fmt.Println("Hello world goroutine")
	
	
}
func main() {  
	go hello(5)
    fmt.Println("main function")
}