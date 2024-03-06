package main

import "fmt"

func main(){
// Control structures
// defer, recover, panic

	defer first()
	second()

	fmt.Println(div(3,0))
	fmt.Println(div(5,3))
	demPanic()

}

//Note the use of recover()
func div(num1, num2 int) int{
	defer func(){
		fmt.Println("a:::",recover())
	}()
	solution := num1/num2
	return solution
}

// Note the use of panic()
func demPanic(){
	defer func(){
		fmt.Println("b:::",recover())
	}()

	panic("Panic")
}

func first(){
	fmt.Println("First!")
}

func second(){
	fmt.Println("Second!")
}