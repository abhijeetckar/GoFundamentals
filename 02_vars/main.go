package main

import(
	"fmt"
)

//var name = "Abhijeet"


func main(){
	//Main Types
	//string
	//bool
	//int
	//int int8 int16 int32 int64
	//uint uint8 uint16 uint32 uint64
	//byte -- alias for int8
	//rune -- alias for int32
	//float32 float64
	//complex64 complex128

	//Using var

//	var name string = "Abhijeet"
//	var name = "Abhijeet"
//	var age int = 26
	var age int32 = 26
	const isCool = true
	//isCool = false

//	Shorthand
	//name := "Abhijeet"
	name, email := "abhijeet","abhijeet@gmail.com"
	height := 1.72
	fmt.Printf("%T\n",height)
	fmt.Println(name,age,isCool,email)
}