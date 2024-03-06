package main

import "fmt"

func main(){

	fmt.Println(addemup(10,20,30,40,50))
	fmt.Println(concat("Hello ", "World", "!"))
}

func addemup(args...int) int{

	sum := 0
	//prints in array format
	fmt.Println(args)
	//access like elements of array
	fmt.Println(args[1])
	//_i is for index, value is for values in array
	for _i, value := range args{
		fmt.Println("i:",_i)
		fmt.Println("val:",value)

		sum += value
	}
	return sum
}

func concat(args ...string) string{
	s:=""
	for _, v := range args{
		//cool string operation
		s += v
	}
	return s
}