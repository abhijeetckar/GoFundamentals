package main

import(
	"fmt"
)

func adder()func(int)int{
	sum := 0
	//fmt.Println(sum)
	return func(x int)int{
		sum+= x
		fmt.Println(sum)
		return sum
	}
}

func main(){
	sum := adder()

	for i := 0; i < 10; i++{
		fmt.Println(sum(i))
	}
}