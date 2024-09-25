package main

import (
	"fmt"
	"strings"
)

func moreOnStrings() {
	str:="Hello World!!"
	length:=len(str)
	fmt.Println(str)
	fmt.Println(length)


	str1:="Go programming"
	str2:="go programming"
	fmt.Println(strings.EqualFold(str1,str2))

	str3:="Hello World!!"
	fmt.Println((strings.Index(str3,"W")))


	substr:=str3[0:3] //slicing (starting index will be inculded and ending index will be eliminated)
	fmt.Println(substr)

	str4:="Hello Go!!"
	fmt.Println(strings.Replace(str4,"Go","World!",1))
}