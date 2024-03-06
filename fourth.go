package main

import "fmt"

func main(){

	var EvenNum[5] int

	EvenNum[0]=0
	EvenNum[1]=2
	EvenNum[2]=4
	EvenNum[3]=6
	EvenNum[4]=8

	for i:=0;i<len(EvenNum); i++{
		fmt.Println("EvenNum[",i,"] -",EvenNum[i])
	}

	OddNum := [5]int{1,3,5,7,9}

	// You can avoid index if not useful inside the loop
	for _, value:=range OddNum{
		fmt.Println("OddNum[] -",value)
	}

	for i, value:=range OddNum {
		fmt.Println("OddNum[",i,"] -",value)
	}

	numSlice:= []int{5,4,3,2,1}
	sliced := numSlice[3:5]
	fmt.Println("numSlice", numSlice)
	fmt.Println("numSlice[3:5]", sliced)

	slice1 := numSlice[0:]
	fmt.Println("numSlice[0:]", slice1)

	slice2 := numSlice[0:2]
	fmt.Println("numSlice[0:2]", slice2)
    // Using make() to create a slice with a specified length and capacity
	newSlice := make([]int,5,6)
	fmt.Println("newSlice[]", newSlice)

	copy(newSlice,numSlice)
	fmt.Println("newSlice[]", newSlice)
	newSlice = append(newSlice, 11,12,13,14,15,16,17,18)
	fmt.Println("newSlice[]", newSlice)

}