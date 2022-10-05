package main

import "fmt"

func main(){
	var number int
	fmt.Println("Enter a Number")
	fmt.Scan(&number)

	if number % 2 == 0 {
		fmt.Println("Number is even!")
	}else{
		fmt.Println("Number is Odd!")
	}

}