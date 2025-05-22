package main

import "fmt"

const (
	A int = 1
	B     = 3.14
	C     = "Hi!"
)

func main() {
	fmt.Println("hello rishabh sharma")
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	var i string = "Hello"
	var j int = 15

	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("j has value: %v and type: %T", j, j)
	var x uint = 500
	var y uint = 4500
	fmt.Printf("Type: %T, value: %v", x, x)
	fmt.Printf("Type: %T, value: %v", y, y)
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)
}

//%v is used to print the value of the arguments
// %T is used to print the type of the arguments
