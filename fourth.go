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
}