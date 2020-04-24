package main

package main

import (  
    "fmt"
    "time"
)

func hello() {  
    fmt.Println("Hello world goroutine")
}
func main() {  
    go hello()
	fmt.Println("main function")
	time.Sleep(5 * time.Second)
}