package main

import "fmt"

type bot interface{
	getGreeting() string
}

type EnglishBot struct {}
type BanglaBot struct {}

func main(){
	eb:=EnglishBot{}
	bb:=BanglaBot{}

	printGreeting(eb)
	printGreeting(bb)
}

func printGreeting(b bot){
	fmt.Println(b.getGreeting())
}

func (EnglishBot) getGreeting()string{
	//custom logic for returning english
	return "English"
}

func (BanglaBot) getGreeting()string{
	//custom logic for returning bangla
	return "Bangla"  
}