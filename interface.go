package main

import "fmt"

func AnyThink(any interface{}){
	fmt.Println(any)
}
func main(){
	fmt.Println("vin")
	AnyThink(2)
	AnyThink(2.01)
	AnyThink(struct {}{})
}