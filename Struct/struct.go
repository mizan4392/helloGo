package main

import "fmt"


type Student struct{
	id int
	name string
	age int
}


func displayStudent(s Student){
	fmt.Printf("ID: %v ; name: %v ; age:%v \n",s.id,s.name,s.age)
}

func (x *Student) increaseAge(age int){
	x.age = x.age + age
}

func main(){
	 rahim := Student{101,"Rahim",20}
	 karim := Student{102,"Rahim",18}

	 displayStudent(rahim)
	 displayStudent(karim)

	 rahim.increaseAge(5)
	 displayStudent(rahim)
}