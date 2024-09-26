package main

import "fmt"

const LoginToken="dwibfdwjwjovnfevnf" // public 
// if anything which declared with the first letter as capital will be public

func main() {
	fmt.Println("Variables")
fmt.Println("==========================","\n","\n")
	var username string="Darshan"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool= true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8= 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64= 255.1234567823456789 // in float 32 we get 5 digits after decimal this is mainly like 32 bit and 64 bit 
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// aliases and default values 

	var anotherVariable int
	fmt.Println(anotherVariable)
		fmt.Printf("Variable is of type: %T \n", anotherVariable)

		var anotherBool bool
		fmt.Println(anotherBool)
		fmt.Printf("Variable is of type: %T \n", anotherBool)

		var anotherString string
		fmt.Println(anotherString)
		fmt.Printf("Variable is of type: %T \n", anotherString)

		var anotherFloat float64
		fmt.Println(anotherFloat)
		fmt.Printf("Variable is of type: %T \n", anotherFloat)

		var anotherunint uint
		fmt.Println(anotherunint)
		fmt.Printf("Variable is of type: %T \n", anotherunint)


		// implicit way of declaring a variable 
		var variable="Darshan"
		fmt.Println(variable)


		// no var type of variable declaration

		Name := "Darshan" // this method is only allowed inside a function or method 
		// if we declare this outside then it will give error we have to follow implicit way or the long procedure way 
		fmt.Println(Name)

		accessingTheLoginToken:=LoginToken
		fmt.Println(accessingTheLoginToken)
		fmt.Printf("the type of the variable is %T", accessingTheLoginToken)

}