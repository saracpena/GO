package main

import "log"

type myStruct struct {
	FirstName string
}
// ! receiver ties function to myStruct line 5
func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "Sara"

// assigning value of
	myVar2 := myStruct{
		FirstName: "Mary",
	}
	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())
}