package main 

import "log"

//! a Pointer points to a specific location in memory and gives you a means in getting that
// ! particular location in memory, pass a reference "&"

//! When I need to get a reference to a var use "&"
//! When I need to refer to an actual pointer use "*"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)
	changeUsingPointer(&myString)
	log.Println("after func call myString is set to", myString)

}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}