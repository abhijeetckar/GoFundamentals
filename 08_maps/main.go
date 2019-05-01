package main

import(
	"fmt"
)

func main(){
	//Define map
	// emails := make(map[string]string)

	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"
	// emails["Mike"] = "mike@gmail.com"	
	
	// fmt.Println(emails["Bob"])

	//Declare map and add key values

	emails := map[string]string{"Bob":"bob@gmail.com","Sharon":"sharon@gmail.com"}

	emails["Mike"] = "mike@gmail.com"
	fmt.Println(emails)
	fmt.Println(len(emails))


	//Delete from map
	delete(emails,"Bob")
	fmt.Println(emails)
}