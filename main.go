package main

import "fmt"

func main() {

	fmt.Println("These Nutzs")

	dataTypes()
	printVariables()
	printFormatted()

}

func printFormatted() {

	var city string = "Kolkata"
	var user string = "macDady"

	fmt.Printf("\n\ncity: %v usrname: %v \n", city, user)
}

func printVariables() {
	var city string = "Kolkata"
	var user string = "macDady"
	fmt.Println("Welcome to", city, ", ", user)
	fmt.Println("----------------------------")
	fmt.Println("Print")
	fmt.Print(city, "\n")
	fmt.Print(user, "\n")
	fmt.Println("----------------------------")
	fmt.Println("Println")
	fmt.Println(city)
	fmt.Println(user)

}

func dataTypes() {

	var s string = "Mother"
	var i int = 100
	var b bool = false
	var f float64 = 3.12
	fmt.Println(s)
	fmt.Println(i)
	fmt.Println(b)
	fmt.Println(f)
}
