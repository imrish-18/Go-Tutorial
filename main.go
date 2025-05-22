package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, Go!")
	fmt.Println("hello rishabh sharma...")
	//display()
	//print()
	ratingPizza()
}

func display() {
	var student1 string = "John" //type is string
	var student3 = "Jane"        //type is inferred
	var student2 = "sharma"
	x := 2 //type is inferred
	y := 3
	fmt.Print(y)
	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Print(student3)
	fmt.Println(x)

	fmt.Println("Hello, Go!")
	fmt.Println("hello rishabh sharma...")

	// Variable Declaration Without Initial Value
	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Value Assignment After Declaration
	var student4 string
	student1 = "John"
	fmt.Println(student4)
}

var a int
var b int = 2
var c = 3
var d = 4

func print() {
	a = 1
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Print(d)
	var a, b, c, d int = 1, 3, 5, 7

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func ratingPizza() {
	fmt.Println("welcome to user input ...")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating for the pizzza")

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for the rating ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("added 1 rating to your rating ", numRating+1)
	}

}
