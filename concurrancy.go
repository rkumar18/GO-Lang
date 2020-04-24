package main


import (
	"fmt";
	"time"
)
func h1(s string){
	fmt.Printf("hello %s",s)
}
func h2(s string){
		fmt.Printf("hello %s",s)
	
}
func main(){
	go h1("rohit")
	h2("mohit")
	fmt.Printf("hello world")
	time.Sleep(10 * time.Second)
}