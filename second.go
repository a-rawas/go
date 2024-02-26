package main

import(
	"fmt"
	"reflect"
)


func main(){

	// printf formating

	const pi float64 = 3.1415

	x:= 33

	isbool := true

	fmt.Printf("Value of pi in float is %f \n", pi)

	fmt.Printf("Type of boolean variable is %T \n", isbool)

	fmt.Printf("%t \n", pi)

	fmt.Printf("base10 formatting for integer %d \n", x)

	fmt.Printf("Type of integer variable is %t \n", x)

	fmt.Printf("Binary representation of integer is %b \n", x)

	fmt.Printf("Char representation of integer is %c \n", x)

	fmt.Printf("Hex representation of integer is %x \n", x)

	fmt.Printf("Type for integer using reflect.TypeOf is %v \n", reflect.TypeOf(x))



	var name string = "World is"

	fmt.Println("String length is ", len(name) ,"characters long")

	fmt.Println(name + " Great!")

}
