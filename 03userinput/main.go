package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	greetings:="Welcome to userinput"
	fmt.Println(greetings)


	raeder:=bufio.NewReader(os.Stdin)

	fmt.Println("Enter your rating for pizza")
//comma ok syntax or comma error syntax
	input,_:=raeder.ReadString('\n')  // input is like a try and _ is like a catch instead of _ we can slo write err or we can also write like _,_
	fmt.Println("Thanks for rating, ",input)
	fmt.Printf("the type of this variable is : %T",input,"\n")



	 scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("Enter your name: ")
    scanner.Scan()
    name := scanner.Text()
    fmt.Println("Hello, ", name)
}