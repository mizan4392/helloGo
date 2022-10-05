package main

import "fmt"

func addsTwoNumber(num1 int,num2 int) int{
	return num1 + num2
}

func main(){
	myCountryIs("Bangladesh")

	var result = addsTwoNumber(1,2)

	fmt.Println("Result:",result)
}


func myCountryIs(country string){
	fmt.Println("My country is ",country)
}