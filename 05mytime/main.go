package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    raeder := bufio.NewScanner(os.Stdin)

    fmt.Println("Enter your name:")
    raeder.Scan()
    name := raeder.Text()
    fmt.Println("Your name is", name)

    fmt.Println("Enter your age:")
    raeder.Scan()
    agestr := raeder.Text()

    age, err := strconv.Atoi(agestr)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Your age is", age) // Print the age variable instead of agestr

    fmt.Println("Enter your designation:")
    raeder.Scan()
    designation := raeder.Text()
    fmt.Println("Your designation is", designation)


	fmt.Printf("my name is %v and my age is %v and my designation is %v", name, age, designation)
}