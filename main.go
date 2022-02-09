package main

import (
	"fmt"
	"reflect"
	"strconv"
)

const PI = 3.14

func main() {

	// Chapter 1
	// dataTypesAndVariables()

	//Chapter2
	operatorsAndControlFlow()

}

func operatorsAndControlFlow() {
	fmt.Println("Chapter 2")
	fmt.Println("\n\n Operators and Control Flow")

	// Comparison Operators
	comparisonOperator()
	arithmeticOperators()
	logicalOperators()
	assignmentOperators()

}

func assignmentOperators() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("assignmentOperators\n")

	var x int = 10
	var y int
	y = x
	var multiline string
	multiline = `var x int = 10
var y int
y = x`
	fmt.Println(multiline)
	fmt.Printf("y = x, y = %v\n", y)
	multiline = `
// add and assign
	x = 10
	y = 20
	x += y // x = x + y `
	x = 10
	y = 20
	x += y // x = x + y
	fmt.Println(multiline)
	fmt.Printf("x += y = %d", x, x)

	x = 10
	y = 20
	x -= y // x = x - y
	multiline = `
// subtract and assign
    x = 10
	y = 20
	x -= y // x = x - y
`
	fmt.Println(multiline)
	fmt.Printf("x -= y // x = x - y = %d\n\n", x)

	x = 10
	y = 20
	x *= y
	multiline = `
// multiply and assign
	x = 10
	y = 20
	x *= y 
`
	fmt.Println(multiline)
	fmt.Printf("x *= y = %d\n\n", x)

	x = 200
	y = 20
	x /= y
	multiline = `
\\ devide and assign
	x = 200
	y = 20
	x /= y
`
	fmt.Println(multiline)
	fmt.Printf("x /= y = %d\n\n", x)

	x, y = 210, 20
	x %= y
	multiline = `
// modulus and assign
	x , y = 210, 20
	x %= y
`
	fmt.Println(multiline)
	fmt.Println("x " + " % " + "= y = " + strconv.Itoa(x))

}

func logicalOperators() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("logicalOperators\n")

	fmt.Println("And && Operator\n")
	var x int = 10
	fmt.Println("var x int = 10")
	fmt.Printf("(x < 100) && (x < 200)=%v\n", (x < 100) && (x < 200))
	fmt.Printf("(x < 300) && (x < 0)=%v\n\n", (x < 300) && (x < 0))

	fmt.Println("Or || Operator\n")
	var y int = 10
	fmt.Println("var y int = 10")
	fmt.Printf("(y < 0) || (y < 200)=%v\n", (y < 0) || (y < 200))
	fmt.Printf("(y < 0) || (y > 200)=%v\n\n", (y < 0) || (y > 200))

	fmt.Println("NOT ! Operator\n")
	var a, b int = 10, 20
	fmt.Println("var a, b int = 10, 20")
	fmt.Printf("!(a  > b) = %v\n", !(a > b))
	fmt.Printf("!(true) = %v\n", !(true))
	fmt.Printf("!(false) = %v\n\n", !(false))

}

func arithmeticOperators() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("arithmeticOperators\n")

	var a, b string = "foo", "bar"
	fmt.Printf("a + b = %s\n", a+b)

	var c, d float64 = 79.02, 75.66
	fmt.Printf("c - d = %.2f\n", c-d)

	var e, f int = 12, 2
	fmt.Printf("e * f = %d\n", e*f)

	var g, h int = 24, 2
	fmt.Printf("g / h = %d\n", g/h)

	var i, j int = 24, 7
	fmt.Println("i % j =", (i % j))
	fmt.Println(i % j)

	var k int = 1
	k++
	fmt.Printf("k++ = %d\n", k)

	var l int = 1
	l--
	fmt.Printf("l++ = %d\n", l)

}

func comparisonOperator() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("comparisonOperator\n")

	var (
		city_1 string = "Harvey"
		city_2 string = "Markham"
	)
	var a, b int = 5, 10
	var c, d int = 10, 10
	var e, f int = 20, 10
	var g, h int = 20, 20

	fmt.Println("city_1 == city_2 = ", city_1 == city_2)

	fmt.Println("city_1 != city_2 = ", city_1 != city_2)

	fmt.Printf("a(%d) < b(%d) = %t \n", a, b, a < b)

	fmt.Printf("c(%d) <= d(%d) = %t \n", c, d, c <= d)

	fmt.Printf("e(%d) > f(%d) = %t \n", e, f, e > f)

	fmt.Printf("g(%d) >= h(%d) = %t \n", g, h, g >= h)

}

func dataTypesAndVariables() {

	fmt.Println("Chapter 1")
	fmt.Println("\n\n Data Types and Variables")

	dataTypes()
	printVariables()
	printFormatted()
	declaringVariables()
	scopeVariable()
	initialVariableValues()
	scanfFunction()
	findTheVariableType()
	convertingDataTypes()
	constantVariables()

}

func constantVariables() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("Constant variables\n")

	const name = "Harry Potter"
	const is_muggle = false
	const age = 14

	fmt.Printf("name %v: %T\n", name, name)
	fmt.Printf("is a muggle %v: %T\n", is_muggle, is_muggle)
	fmt.Printf("age: %v: %T\n\n", age, age)

	var radius float64 = 5.0
	var area float64
	area = PI * radius * radius
	fmt.Printf("area = %.2f\n\n", area)

}

func convertingDataTypes() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("Converting Data Types\n")

	fmt.Println("Interger to float")
	var i = 90
	var f float64 = float64(i)
	fmt.Printf("%.2f\n", f)

	fmt.Println("Float  to Integer")
	var f2 float64 = 54.89
	var i2 int = int(f2)
	fmt.Printf("%v\n", i2)

	fmt.Println("strconv Package")
	var i3 int = 42
	var s string = strconv.Itoa(i3) // convert integer to string
	fmt.Printf("%q", s)

	fmt.Println("\n------------\n")
	fmt.Println(" strconv string to integer")

	var s4 string = "200"
	i, err := strconv.Atoi(s4)
	fmt.Printf(" i %v, %T \n", i, i)
	fmt.Printf("err : %v, %T\n\n", err, err)

	s4 = "200abc"
	i, err = strconv.Atoi(s4)
	fmt.Printf(" i %v, %T \n", i, i)
	fmt.Printf("err : %v, %T", err, err)

}

func findTheVariableType() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("Finf variable Type\n")
	var (
		grades  int     = 42
		message string  = "hello world"
		isCheck bool    = true
		amount  float32 = 5466.54
	)
	fmt.Printf("variable grades = %v is of type %T\n", grades, grades)
	fmt.Printf("variable message = %v is of type %T\n", message, message)
	fmt.Printf("variable isCheck = %v is of type %T\n", isCheck, isCheck)
	fmt.Printf("variable amount = %v is of type %T\n\n", amount, amount)

	fmt.Printf("Type: %v \n", reflect.TypeOf(1000))
	fmt.Printf("Type: %v \n", reflect.TypeOf("priyanka"))
	fmt.Printf("Type: %v \n", reflect.TypeOf(46.0))
	fmt.Printf("Type: %v \n\n", reflect.TypeOf(true))

	fmt.Printf("variable grades = %v is of type %v\n", grades, reflect.TypeOf(grades))
	fmt.Printf("variable message = %v is of type %v\n", message, reflect.TypeOf(message))
	fmt.Printf("variable isCheck = %v is of type %v\n", isCheck, reflect.TypeOf(isCheck))
	fmt.Printf("variable amount = %v is of type %v\n\n", amount, reflect.TypeOf(amount))

}

func scanfFunction() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("scanf function\n")

	//var name string
	//fmt.Print("Enter your name: ")
	//fmt.Scanf("%s", &name)
	//fmt.Printf("\nHello %s\n", name)
	//
	//var (
	//	name2      string
	//	is_muggle2 bool
	//)
	//fmt.Printf("\nEnter your name & are you a muggle : ")
	//fmt.Scanf("%s %t", &name2, &is_muggle2)
	//fmt.Printf("Hello %s ... are you a muggle: %t\n", name2, is_muggle2)

	//var (
	//	a string
	//	b int
	//)
	//fmt.Print("\nEnter a sting and a number :")
	//count, err := fmt.Scanf("%s %d", &a, &b)
	//fmt.Println("count : ", count)
	//fmt.Println("err : ", err)
	//fmt.Printf("a: %s b: %d", a, b)

}

func initialVariableValues() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("Initial Variable Values\n")

	var i int
	fmt.Printf("Inital vaiable value : %d", i)

}

func scopeVariable() {
	fmt.Println("\n\n--------------------------------------------")
	fmt.Println("Vaiable scope\n")
	// blocks

	city := "Markham"
	{ // outter Block
		country := "USA"

		{ // inner block
			contenent := "North America"
			fmt.Printf("\ncity: %s, country: %s, contenent: %s\n", city, country, contenent)
		}
	}
	// fmt.Printf("\ncity: %s, country: %s, contenent: %s\n", city, country, contenent)
}

func declaringVariables() {

	fmt.Println("\n\nDeclaring variables")
	fmt.Println("------------------------------")

	var name string
	name = "Tony"

	fmt.Printf("\n name %s\n", name)

	// shorthand #1
	var s, t string = "foo", "bar"
	fmt.Printf("\nPrinting foo,bar = %s, %s\n", s, t)

	// shorthand #2
	var (
		s2 string = "Tony"
		i2 int    = 52
	)
	fmt.Printf("\nName: %s Age: %d", s2, i2)

	// short variable declaration
	name3 := "Maria"
	age3 := 40
	fmt.Printf("\nWoman: %s age: %d", name3, age3)

	name3 = "lisa"
	age3 = 34
	fmt.Printf("\nWoman: %s age: %d", name3, age3)

}

func printFormatted() {

	var city string = "Kolkata"
	var user string = "macDady"
	var grade int = 32

	fmt.Printf("\n\ncity: %v usrname: %v \n", city, user)

	fmt.Printf("\n-------------------\nprintf()\n\ncity: %v username: %v  grade %d", city, user, grade)
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
