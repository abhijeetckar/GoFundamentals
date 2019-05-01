package main

import(
	"fmt"
)

func main(){
	// if else
	
	x := 15
	y := 10

	if x < y{
		fmt.Printf("%d is less than %d\n",x,y);
	}else{
		fmt.Printf("%d is less than %d\n",y,x);
	}


	//if elseif
	color := "blue"


	if color == "red"{
		fmt.Println("Color is red")
	}else if color == "blue"{
		fmt.Println("Color is blue")
	}else{
		fmt.Println("Color is NOT blue or red")
	}

	//switch case

	switch color{
		case "red":
			fmt.Println("Color is Red")
			// break
		case "blue":
			fmt.Println("Color is Blue")
			// break		
		default:
			fmt.Println("Color is not BLUE or RED")
	}
}