package main

import "fmt"

func main() {
	//string
	var agency string="Fast Track"
	fmt.Println("Welcome to",agency)
	
	// practice
	var name="Darshan"
	fmt.Println(name)
	
	name1:="Darshan" // variable name and : shorthand property to declare a variable 
	fmt.Println(name1)
	fmt.Println("Hello,",name1) 

	totalCars:=50
	fmt.Println("Our fleet consists of ",totalCars)

	StartingPrice:=29.99

	fmt.Printf("Our cars starting price is %v \n",StartingPrice)
	fmt.Println("take your pick")

	includesInsurance:=false
	fmt.Println("the insurance is",includesInsurance)


	var name2 string
	fmt.Println("name2",name2)

	var number int
	fmt.Println("number",number)

	var price float64
	fmt.Println("price",price)

	var isActive bool
	fmt.Println("isActive",isActive)

	// more on strings
	moreOnStrings()
}