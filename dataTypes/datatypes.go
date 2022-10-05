/* Data types: boolean, string, numeric - integer, floating,
derived types - pointer, array, structure, slice, map, interface etc.*/

package main

import "fmt"

func main(){
	//declaration
	var name string
	var age int
	var country string
	var  gpa float32	

	//initialization
	name = "Mizan"
	age = 26
	country = "Bangladesh"
	gpa = 3.30

	fmt.Println("Name:",name)
	fmt.Println("Age:",age)
	fmt.Println("Country",country,"and Gpa:",gpa)

	//declaration and initialization

	var _name = "Mizan"
	 _age := 26
	 _country := "Bangladesh"
	 _gpa := 3.30
	fmt.Println("Name:",_name)
	fmt.Println("Age:",_age)
	fmt.Printf("Country %v and Gpa:%v",_country,_gpa)


	fmt.Println("\nEnter your custom input")
	var __name string

	fmt.Scan(&__name)

	fmt.Println(__name)


}
