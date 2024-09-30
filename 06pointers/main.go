package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers in golang")

	// var ptr *int
	// fmt.Println("value of pointer is : ", ptr) this is how we create a pointer

	myNumber:=22
	var ptr=&myNumber
	fmt.Println("value of actual pointer is : ", ptr)
	fmt.Println("value of actual pointer is : ", *ptr)

	*ptr= *ptr+2
	fmt.Println("new value is : ", myNumber)
}
