package main

import "net/http"
import "fmt"

//! Creating an INTERFACE
// import "log"

// type Animal interface {
// 	Says() string
// 	NumberOfLegs() int
// }

// type Dog struct {
// 	Name string
// 	Breed string
// }

// type Gorilla struct {
// 	Name string
// 	Color string
// 	NumberOfTeeth int
// }

// func main() {
// 	dog := Dog {
// 		Name: "Samson",
// 		Breed: "German Shepherd",
// 	}

// 	PrintInfo(dog)

// 	gorilla := Gorilla {
// 		Name: "King Kong",
// 		Color: "Black",
// 		NumberOfTeeth: 32,
// 	}

// 	PrintInfo(gorilla)
// }

// func (d Dog) Says() string {
// 	return "woof"
// }

// func (d Dog) NumberOfLegs() int{
// 	return 4
// }
// func (d Gorilla) Says() string {
// 	return "grunt"
// }

// func (d Gorilla) NumberOfLegs() int{
// 	return 2
// }
// //! I had to import "log" for log to work
// func PrintInfo(a Animal){
// 	log.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
// }

//! HELLO WORLD.....
//? Take 1
// func main(){
// 	fmt.Println("Hello, World of GO!")
// }

//? Take 2
func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		n, err := fmt.Fprintf(w, "Hello, World!")
		// if there IS an error
		if err != nil { 
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	})

	_ = http.ListenAndServe(":8080", nil)
}