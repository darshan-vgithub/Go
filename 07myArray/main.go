package main

import "fmt"

func main() {
	// these arrays are not much used in the go language but still there are some use cases for these also
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


	var veitableList[4]string
	
	veitableList[0] = "Carrot"
	veitableList[1] = "Beans"
	veitableList[2] = "Potato"
	veitableList[3] = "Onion"

	fmt.Println("Vegetable list is : ", veitableList)
	fmt.Println("Vegetable list is : ", len(veitableList))

	var vegList[]string

	vegList = append(vegList, "Carrot")
	vegList = append(vegList, "Beans")
	vegList = append(vegList, "Potato")

	fmt.Println("Vegetable list is : ", vegList)
	fmt.Println("Vegetable list is : ", len(vegList))

}
