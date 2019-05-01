package main

import(
	"fmt"
)

func main(){
	//Arrays
	// var fruitArray [2] string

	// fruitArray[0] = "Apple"
	// fruitArray[1] = "Orange"
	// fmt.Println(fruitArray);

	// fruitArray := [2]string{"Apple","Orange"}
	// fmt.Println(fruitArray);

	fruitSlice := []string{"Apple","Orange","Grapes","Banana"}
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])

}
