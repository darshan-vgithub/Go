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
}
