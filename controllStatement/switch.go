package main

import "fmt"

func main(){
	var number int
	fmt.Println("Enter a Number")
	fmt.Scan(&number)

	switch number{
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")

	default:
		fmt.Println("Out of scope")
	}
}