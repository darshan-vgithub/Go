package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello time study go language")

	PresentTime := time.Now()
	fmt.Println("present time is : ", PresentTime)

	fmt.Println(PresentTime.Format("Mon Jan 2 15:04:05 MST 2006"))
loc := time.Location{}
CreatedTime := time.Date(2022, time.December, 22, 0, 0, 0, 0, &loc)
	fmt.Println("created time is : ", CreatedTime)
}
