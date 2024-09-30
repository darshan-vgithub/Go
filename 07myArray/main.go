package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in golang")

	var fruitlist [4]string

	fruitlist[0] = "Apple"
	fruitlist[2] = "Mango"
	fruitlist[3] = "Peach"

	fmt.Println("Fruit list is : ", fruitlist)
	fmt.Println("Fruit list is : ", len(fruitlist))

	fmt.Println("Welcome to slices in golang")

	var fruitlist1 []string

	fruitlist1 = append(fruitlist1, "Apple")
	fruitlist1 = append(fruitlist1, "Mango")
	fruitlist1 = append(fruitlist1, "Peach")
	fruitlist1 = append(fruitlist1, "Grapes")

	fmt.Println("Fruit list is : ", fruitlist1)
	fmt.Println("Length of fruit list is : ", len(fruitlist1))

}
