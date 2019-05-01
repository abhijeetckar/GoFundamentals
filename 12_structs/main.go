package main

import(
	"fmt"
	"strconv"
)

//Define Person Struct

type Person struct{
	firstName, 
	lastName, 
	city, 
	gender string
	age int
}

// Greeting method (value reciever)

func (p Person) greet() string{
	return "Hello, my name is "+ p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}
//What if I call this func() for a different struct???


// hasBirthDay method (pointer reciever)

func (p *Person) hasBirthDay(){
	p.age++
}

// getMarried (pointer Reciever)
func (p *Person)getMarried(spouseLastname string){
	if p.gender == "M"{
		return
	}else{
		p.lastName = spouseLastname
	}
}

func main(){

	//Init person struct
	// person1 := Person{
	// 	firstName: "Samantha",
	// 	lastName:"Smith",
	// 	city:"Boston",
	// 	gender:"F",
	// 	age:25}

	//Alternative
	person1 := Person{"Samantha","Smith","Boston","F",25}
	person2 := Person{"Bob","Jordan","New York","M",30}
	
	// fmt.Println(person1)
	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1)	
	
	person1.hasBirthDay()
	person1.getMarried("Williams")
	person2.getMarried("Thompson")
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())	
}
