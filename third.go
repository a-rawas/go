package main

import "fmt"

func main(){
	

	for i:=1; i<=10; i++{
		fmt.Println(i)

	}

	j:=1

	for j<=10 {

		fmt.Println (j)
		j++
	}

	for i:=1; i<5; i++{
		for j:=1; j<=i; j++{
			fmt.Printf("*")
		}
		fmt.Println()
	}

	age := 18

	if age >=18 {
		fmt.Println("You are",age,". You can vote")
	} else {
			fmt.Println("You are",age,". You cannot vote")

	}

	era := 2023

	switch era{
	case 1947: fmt.Println("India won freedom !")
	case 1970: fmt.Println("India under emergency")
	case 2023: fmt.Println("Moon landing!")
	}

}