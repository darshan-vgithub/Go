package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza application")
	fmt.Println("Please rate our pizza between 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}



	//practice
	scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("Enter your name: ")
    scanner.Scan()
    name := scanner.Text()

    fmt.Print("Enter your age: ")
    scanner.Scan()
    ageStr := scanner.Text()
    age, err := strconv.Atoi(ageStr)
    if err != nil {
        fmt.Println("Invalid age. Please enter a number.")
        return
    }

    fmt.Print("Enter your designation: ")
    scanner.Scan()
    designation := scanner.Text()

    fmt.Printf("Hello, my name is %s, I am %d years old, and I am a %s.\n", name, age, designation)
}
