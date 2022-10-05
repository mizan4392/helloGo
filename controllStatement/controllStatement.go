package main

import (
	"fmt"
)

//statement

func main(){
	var number int
	fmt.Println("Enter a Number")
	fmt.Scan(&number)

	if number > 0 {
		fmt.Println("Number is Positive")
	}else if number < 0 {
		fmt.Println("Number is Negative")
	}else{
		fmt.Println("Number is Zero")
	}
}