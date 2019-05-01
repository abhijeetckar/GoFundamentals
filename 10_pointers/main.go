package main

import(
	"fmt"
)

func main(){
	a := 5
	b := &a

	fmt.Println(a,b)
	fmt.Printf("%T,%T\n",a,b)
	fmt.Println(a,*b)
	fmt.Println(a,*&a)


	//Change value with pointer

	*b = 10
	fmt.Println(a)

	//Everything in go is pass by value
}