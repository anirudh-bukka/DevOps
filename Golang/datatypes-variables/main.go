package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a string
	var b int
	fmt.Print("Enter string and number: ")

	/*
		count: counts no. of arguments the function writes to
		erri: any error thrown during the execution of the function

		THESE ARE NOT KEYWORDS
	*/
	count, erri := fmt.Scanf("%s %d", &a, &b)

	fmt.Println("Count: ", count)
	fmt.Println("Error: ", erri)
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

	var grade int = 99
	var message string

	fmt.Print("Enter a message: ")
	fmt.Scanf("%s", &message)

	fmt.Printf("Variable grade = %v is of the type: %T", grade, grade)
	fmt.Printf("Variable message = %v is of the type: %v", message, reflect.TypeOf(message))
}

// func datatypes() {
// 	var grade int = 99
// 	var message string

// 	fmt.Scanf("Enter a message: ", &message)

// 	fmt.Printf("Variable grade = %v is of the type: %T", grade, grade)
// 	fmt.Printf("Variable message = %v is of the type: %v", grade, reflect.TypeOf(message))
// }
