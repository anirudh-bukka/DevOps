package main

import (
	"fmt"
	"reflect"
	"strconv"
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

	// --------------------------------------------------------
	// Finding type of variables
	var grade int = 99
	var message string

	fmt.Print("\nEnter a message: ")
	fmt.Scanf("%s", &message)

	fmt.Printf("Variable grade = %v is of the type: %T\n", grade, grade)
	fmt.Printf("Variable message = %v is of the type: %v\n\n", message, reflect.TypeOf(message))

	// --------------------------------------------------------
	// CONVERTING BETWEEN TYPES
	// -- Float to integer
	var floatA float64 = 45.89
	var integerA int = int(floatA)
	fmt.Printf("%v\n", integerA)
	// -- Integer to float
	var integerB int = 50
	var floatB float64 = float64(integerB)
	fmt.Printf("%.2f\n", floatB)

	// -- String to Integer and vice-versa
	// -- Itoa(): Converts integer to string | returns 1 value
	var integerC int = 23
	var stringC string = strconv.Itoa(integerC)
	fmt.Printf("\n%v, is of type %T\n", stringC, stringC)

	// -- Atoi(): Converts string to integer | returns corresponding integer and error
	var stringD string = "200"
	integerD, errorMessageD := strconv.Atoi(stringD)
	fmt.Printf("%v, is of the type %T\n\tErrors: %v (and type): %T", integerD, integerD, errorMessageD, errorMessageD)

	var stringE string = "abc"
	integerE, errorMessageE := strconv.Atoi(stringE)
	fmt.Printf("\n%v, is of the type %v\n\tErrors: %v (and type): %T", integerE, reflect.TypeOf(integerE), errorMessageE, errorMessageE)

	// Untyped Constant
	const untyped_age = 12
	const untyped_name, untyped_roll = "anirudh", 22

	// Typed Constant
	// // // cannot declare multiple constants in line
	const typed_name string = "anirudh"
	const typed_age int = 22
}
