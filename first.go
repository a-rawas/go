package main

import "fmt"

func main(){
	fmt.Println("Hello World! This is my first GO Programme!")

// A constant is a variable with a value that can't be changed
	const pi float64 = 3.14159265359

//You can declare multiple variables like this

	var(
		varA = 2
		varB = 3
	)
// Another way to declare variables
	x,y := 14, 15

	fmt.Println(varA, varB)
	fmt.Println(pi)
	fmt.Println(x, y)

// Arithmatic operations

	fmt.Println("x+y",x+y)
	fmt.Println("x-y",x-y)
	fmt.Println("x*y",x*y)
	fmt.Println("x/y",x/y)
	fmt.Println("x mod y",x%y)

// Boolean operations
	isbool := true
	hate := false

	fmt.Println("isbool :", isbool)
	fmt.Println("isbool && hate :", isbool && hate)
	fmt.Println("isbool || hate :", isbool||hate)
	fmt.Println("!isbool :", !isbool)


// Pointers

	fmt.Println("Where is x ?", &x)
	fmt.Println("Where is y ?", &y)

// Functions for pointers - pass by reference
	changeval(&x)
	fmt.Println("New value of x is ", x)
	fmt.Println("Where is x now ?", &x)

	//Strings are a series of characters between " or '

	var Name string = "Anand Rawas"

	// Get string length

	fmt.Println("Length of ", Name ,"is", len(Name))

// Functions

}

func changeval(addr *int){
	*addr = 22;
}
