package main

import "fmt"


func main(){
	cl := map[string]string{
		"red":"#ff0000",
		"green":"#4bf745",
	}

	printMap(cl)

}


func printMap(c map[string]string){
	for color,hex :=range c{
		fmt.Println("Hex code for ",color," is ",hex)
	}
}