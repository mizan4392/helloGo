package main

import "fmt"


type ContactInfo struct{
	email string
	zipCode int
}

type Person struct{
	firstName string
	lastName string
	contact ContactInfo
}

func main(){
	mizan := Person{
		firstName: "Mizan",
		lastName: "Sobuz",
		contact:ContactInfo{
			email: "mizan@gmail.com",
			zipCode: 1000,
		},
	}
	mizan.print()

	mizan.updateFirstName("Sozib")
	mizan.print()

}

func (refToPerson *Person) updateFirstName(newFirstName string){
	(*refToPerson).firstName = newFirstName
}

func (p Person) print(){
	fmt.Println(p)
}