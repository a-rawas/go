package main

import "fmt"

func main(){

	x,y := 5,6.9

	fmt.Println(add(x,y))

	fmt.Println("Factorial of ",x,"is" ,factorial(x))
}

func add (num1 int, num2 float64) float64{
	f := float64(num1)
	return f+num2
}

func factorial(num int) int{

	if num==0{
		return 1
	}
	return num*factorial(num-1)
}