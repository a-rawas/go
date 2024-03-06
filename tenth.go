package main

import "fmt"

func main(){

	rect1 := Rectangle{height:30,width:40}
	fmt.Println("Height:",rect1.height)

	rect2 := Rectangle{30,40}
	fmt.Println("Width:",rect2.width)

	fmt.Println("Rect1 Area:", rect1.area())

}

type Rectangle struct{
	height float32
	width float32
}

func (rect *Rectangle) area() float32{
	return rect.height*rect.width
}