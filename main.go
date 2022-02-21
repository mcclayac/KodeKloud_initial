package main

import (
	"fmt"
	"reflect"
	"strconv"
)

const PI = 3.14

var (
	cube = func(i int) string {
		c := i * i * i
		return strconv.Itoa(c)
	}
)

func main() {

	//x := cube(8)
	//fmt.Printf("%T %v", x, x)

	// fmt.Println(quote.Go())

	// Chapter 1
	// dataTypesAndVariables()

	//Chapter2
	// operatorsAndControlFlow()

	// Chapter 3
	//arraysAndSlices()

	// Chapter 4
	// functions()

	// Chapter 5
	//pointers()

	// Chapter 6
	structsMethodsInterfaces()

}

func structsMethodsInterfaces() {
	fmt.Println("--------------------------------")
	fmt.Println("Chapter 6")
	fmt.Println("structsMethodsInterfaces\n\n")

	declaringAndInitializingStruct()
	accessingFieldsStructs()
	passingStructsToFunctions()
	comparingStructs()

}

func comparingStructs() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("Comparring Structs \n")

	//c1 := s1{5}
	//c2 := s2{5}

	// compile-type error
	//if c1 == c2 {
	//	fmt.Println("yes")
	//} else {
	//	fmt.Println("no")
	//}

	c1 := s1{5}
	c2 := s1{6}
	c3 := s1{5}

	if c1 != c2 {
		fmt.Println("c1 and c2 have different values")
	}

	if c1 == c3 {
		fmt.Println("c is same as c2")
	}

}

type s1 struct {
	x int
}

type s2 struct {
	x int
}

func passingStructsToFunctions() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("Passng Structs to Function \n")

	d := Circle2{5, 5, 5, 0, 0}
	fmt.Printf("%+v\n", d)
	calcAreaStruct(&d)
	fmt.Printf("%+v\n", d)

}

func calcAreaStruct(c *Circle2) {
	const PI float64 = 3.14
	var area float64
	area = (PI * c.radius * c.radius)
	(*c).area = area
	// or
	c.area2 = area + 10
}

func accessingFieldsStructs() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("Accessing a struct field \n")

	var c Circle
	c.x = 5
	c.y = 5
	c.radius = 5

	fmt.Printf("%+v\n", c)

}

type Circle struct {
	x      float64
	y      float64
	radius float64
}

type Circle2 struct {
	x      int
	y      int
	radius float64
	area   float64
	area2  float64
}

type Student struct {
	name   string
	rollNo int
	marks  []int
	grades map[string]int
}

type Student2 struct {
	name   string
	rollNo int
}

func declaringAndInitializingStruct() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("Declaring and Initoa==ilaizing a Struct \n")

	// Initial ver 1
	var s Student
	fmt.Printf("%+v\n", s)

	// Initial ver 2
	t := new(Student)
	fmt.Printf("%+v\n", t)

	// Initial ver 3
	u := Student{
		name:   "Joe",
		rollNo: 12,
	}
	fmt.Printf("%+v\n", u)

	// Initial ver 3
	v := Student2{"Tony", 32}
	fmt.Printf("%+v\n", v)
}

func pointers() {
	fmt.Println("--------------------------------")
	fmt.Println("Chapter 5")
	fmt.Println("pointers\n\n")

	addressAndDereferenceOperator()
	declaringAPointer()
	initializingAPointer()
	dereferencingAPointer()
	passingByValue()
	passByReference()

}

func passByReference() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("passByReference \n")
	a := "hello"
	fmt.Println(a)
	modifyRef(&a)
	fmt.Println(a)

	slice := []int{10, 20, 30}
	fmt.Println(slice)
	modifySlice(slice)
	fmt.Println(slice)

	ascii_code := make(map[string]int)
	ascii_code["A"] = 65
	ascii_code["F"] = 70
	fmt.Println(ascii_code)
	modifyMap(ascii_code)
	fmt.Println(ascii_code)
}

func modifyMap(code map[string]int) {

	code["K"] = 75
}

func modifySlice(slice []int) {

	slice[0] = 100

}

func modifyRef(s *string) {

	*s = "World!"

}

func passingByValue() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("passingByValue \n")

	a := "hello"
	fmt.Println(a)
	modify(a)
	fmt.Println(a)

}

func modify(s string) {

	s = "world!"
}

func dereferencingAPointer() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("dereferencingAPointer \n")

	s := "Hello"
	fmt.Printf("%T %v \n", s, s)
	ps := &s
	*ps = "world"
	fmt.Printf("%T %v \n", s, s)

}

func initializingAPointer() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("initializingAPointer \n")

	s := "Hello"
	var b *string = &s
	fmt.Println(b)
	var a *string = &s
	fmt.Println(a)
	var c *string = &s
	fmt.Println(c)

	fmt.Println(*b)

}

func declaringAPointer() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("declaringAPointer \n")

	var i *int
	var p *string
	fmt.Println(i)
	fmt.Println(p)

}

func addressAndDereferenceOperator() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("addressAndDereferenceOperator \n")

	x := 77
	fmt.Printf("%T, %v \n", &x, &x)
	fmt.Printf("%T, %v \n", *(&x), *(&x))

}

func functions() {
	fmt.Println("--------------------------------")
	fmt.Println("Chapter 4")
	fmt.Println("Functions\n\n")

	basicFunction()
	multiplereturnTypes()
	variadicFunction()
	recursiveFunctions()
	anonymousFunctions()
	// highOrderFunctions()
	DeferStatement()

}

func DeferStatement() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("DeferStatement \n")

	printName("Joe")
	defer printRollNo(123)
	printAddress("16315 Lawndale")

}

func printName(str string) {
	fmt.Println(str)
}

func printRollNo(rno int) {
	fmt.Println(rno)
}

func printAddress(adr string) {
	fmt.Println(adr)
}

func highOrderFunctions() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("highOrderFunctions \n")

	var query int
	var radius float64
	fmt.Println("High Order Function case 1")
	fmt.Println("Enter the raidus of the circle:")
	fmt.Scanf("%f", &radius)
	fmt.Printf("Enter \n 1 - area \n 2 - perimeter \n 3 - diameter \n enter:")
	fmt.Scanf("%d", &query)

	if query == 1 {
		fmt.Println("Result:", calcArea(radius))
	} else if query == 2 {
		fmt.Println("Result:", calcPerimeter(radius))
	} else if query == 3 {
		fmt.Println("Result:", calcDiameter(radius))
	} else {
		fmt.Println("Error")
	}

	fmt.Println("High Order Function case 2")
	fmt.Println("High Order Function case 2")
	fmt.Println("Enter the raidus of the circle:")
	fmt.Scanf("%f", &radius)
	fmt.Printf("Enter \n 1 - area \n 2 - perimeter \n 3 - diameter \n enter:")
	fmt.Scanf("%d", &query)

	printResult(radius, getFunction(query))

}

func printResult(radius float64, calcFunction func(r float64) float64) {
	result := calcFunction(radius)
	fmt.Println("result :", result)
	fmt.Println("Thank you!")
}

func getFunction(query int) func(r float64) float64 {

	query_to_func := map[int]func(r float64) float64{
		1: calcArea,
		2: calcPerimeter,
		3: calcDiameter,
	}
	return query_to_func[query]
}

func calcArea(r float64) float64 {
	return 3.14 * r * r
}
func calcPerimeter(r float64) float64 {
	return 2 * 3.14 * r
}
func calcDiameter(r float64) float64 {
	return 2 * r
}

func anonymousFunctions() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("anonymousFunctions \n")

	x := func(l int, b int) int {
		return l * b
	}
	fmt.Printf("%T \n", x)
	fmt.Println(x(20, 10))
	fmt.Printf("\n\n")

	y := func(l int, b int) int {
		return l * b
	}(20, 30)

	fmt.Printf("%T \n", y)
	fmt.Println(y)

}

func recursiveFunctions() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("variadicFunction \n")

	fmt.Println("factorial(5) = ", factorial(5))

}

func factorial(n int) int {

	if n == 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func variadicFunction() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("variadicFunction \n")

	var sum int = 0
	sum = addNumbers(1, 2, 3, 4)
	fmt.Printf("%d , %d, %d, %d = %d\n", 1, 2, 3, 4, sum)

	fmt.Println("addNumbers() = ", addNumbers())
	fmt.Println("addNumbers(10) = ", addNumbers(10))
	fmt.Println("addNumbers(10,20) = ", addNumbers(10, 20))
	fmt.Println("addNumbers(10,20,30) = ", addNumbers(10, 20, 30))

}

func addNumbers(numbers ...int) int {

	var sum int = 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func multiplereturnTypes() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("multiplereturnTypes \n")

	var val1, val2 int = 8, 9
	addition, substraction := operatios(val1, val2)
	fmt.Printf("%d + %d = %d\n", val1, val2, addition)
	fmt.Printf("%d - %d = %d\n\n", val1, val2, substraction)

	val1, val2 = 10, 20
	addition, substraction = operatios2(val1, val2)
	fmt.Printf("%d + %d = %d\n", val1, val2, addition)
	fmt.Printf("%d - %d = %d\n\n", val1, val2, substraction)

}

func operatios2(val1 int, val2 int) (add int, sub int) {

	// return values with named return-type arguement
	add = val1 + val2
	sub = val1 - val2
	return add, sub
}

func operatios(val1 int, val2 int) (int, int) {

	var add, sub int
	add = val1 + val2
	sub = val1 - val2
	return add, sub
}

func basicFunction() {

	var val1, val2 int = 6, 9
	sum := addTwo(val1, val2)
	fmt.Printf("%d + %d = %d\n\n", val1, val2, sum)

}

func addTwo(value1 int, value2 int) int {

	var sum int
	sum = value1 + value2
	return sum
}

func arraysAndSlices() {
	fmt.Println("--------------------------------")
	fmt.Println("Chapter 3")
	fmt.Println("\n\n Arrays and Slices")

	arrays()
	slices()
	maps()

}

func maps() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("maps {\n\n")

	//fmt.Println("This is an error")
	//var code map[string]string
	//code["en"] = "English"
	//fmt.Println(code)

	codes2 := map[string]string{"en": "English", "fr": "French"}
	fmt.Println("codes2 := map[string]string{\"en\":\"English\",\"fr\":\"French\"}")
	fmt.Println("codes2 = ", codes2, "\n")

	fmt.Println("Map with make function")
	codes3 := make(map[string]int)
	fmt.Println("codes3 := make(map[string]int) =", codes3)
	fmt.Println("empty map")

	codes4 := map[string]string{"en": "English", "fr": "French", "hi": "Hindi"}
	fmt.Println("codes4 := map[string]string{\"en\": \"English\", \"fr\": \"French\",\"hi\":\"Hindi\"}")
	fmt.Println("codes4 =", codes4)
	fmt.Println("len(codes4) = ", len(codes4))

	fmt.Println("Accessing the keys of an map - codes4")
	fmt.Println("codes4[\"en\"] =", codes4["en"])
	fmt.Println("codes4[\"fr\"] =", codes4["fr"])
	fmt.Println("codes4[\"hi\"] =", codes4["hi"])

	fmt.Println("\n ----- getting a value from a map - value, found")
	code5 := map[string]int{"en": 1, "fr": 2, "hi": 3}
	fmt.Println("code5 := map[string]int{\"en\": 1, \"fr\": 2, \"hi\": 3}")
	value, found := code5["en"]
	fmt.Println("value, found := code5[\"en\"]")
	fmt.Println("value, found = ", value, ",", found)
	value, found = code5["hh"]
	fmt.Println("value, found = code5[\"hh\"]")
	fmt.Println("value, found = ", value, ",", found)

	fmt.Println("\n---- Adding a value to a map")
	codes6 := map[string]string{"en": "English", "fr": "French", "hi": "Hindi"}
	fmt.Println("codes6 := map[string]string{\"en\": \"English\", \"fr\": \"French\", \"hi\": \"Hindi\"}")
	codes6["it"] = "Italian"
	fmt.Println("codes6[\"it\"] = \"Italian\"")
	fmt.Println("codes6 =", codes6)

	fmt.Println("\n------ Update a key")
	fmt.Println("codes6 =", codes6)
	codes6["en"] = "English language"
	fmt.Println("codes6[\"en\"] = \"English language\"")
	fmt.Println("codes6 =", codes6)

	fmt.Println("\n------ Delete a key")
	fmt.Println("codes6 =", codes6)
	delete(codes6, "en")
	fmt.Println("delete(codes6, \"en\")")
	fmt.Println("codes6 =", codes6)

	fmt.Println("\n------ Interate over a map")
	fmt.Println("codes6 =", codes6)
	for key, value := range codes6 {
		fmt.Println(key, "=>", value)
	}

	fmt.Println("\n Truncate a map ")
	fmt.Println("codes6 =", codes6)
	for key, _ := range codes6 {
		delete(codes6, key)
	}
	fmt.Println("codes6 =", codes6)

	fmt.Println("\n Truncate option #2")
	codes7 := map[string]string{"en": "English", "fr": "French", "hi": "Hindi"}
	fmt.Println("codes7 =", codes7)
	codes7 = make(map[string]string)
	fmt.Println("codes7 = make(map[string]string)")
	fmt.Println("codes7 =", codes7)

}

func slices() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("slices {\n\n")

	fmt.Println("a Slice")
	aSlice := []int{10, 20, 30}
	fmt.Println("aSlice = ", aSlice)

	fmt.Println("declairing and initializing a slice")
	aArrays2 := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("aArrays2 = ", aArrays2)
	aSlice2 := aArrays2[1:8]
	fmt.Println(aSlice2)
	aSubSlice2 := aSlice2[0:3]
	fmt.Println("aSubSlice2  = aSlice2[0:3] = ", aSubSlice2)

	aSlice3 := make([]int, 5, 8)
	fmt.Println("aSlice3 := make([]int, 5 , 8) = ", aSlice3)
	fmt.Println("len(aslice3) =", len(aSlice3))
	fmt.Println("cap(aslice3) =", cap(aSlice3))

	aArray4 := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	aSlice4 := aArray4[1:8]
	fmt.Println("aArray4 = ", aArray4)
	fmt.Println("aSlice4 := aArray4[1:8]")
	fmt.Println("aSlice4 = ", aSlice4)
	fmt.Println("cap(aArray4) = ", cap(aArray4))
	fmt.Println("cap(aSlice4) = ", cap(aSlice4))

	aArray5 := [5]int{10, 20, 30, 40, 50}
	aSlice5 := aArray5[:3]
	fmt.Println("array5 =", aArray5)
	fmt.Println("aSlice5 := aArray5[:3]")
	fmt.Println("aSlice5 =", aSlice5)
	aSlice5[1] = 9000
	fmt.Println("aSlice5[1] = 9000")
	fmt.Println("After modification")
	fmt.Println("array5 =", aArray5)
	fmt.Println("aSlice5 =", aSlice5)

	fmt.Println("\n\nAppending a slice")
	aArray6 := [4]int{10, 20, 30, 40}
	aSlice6 := aArray6[1:3]
	fmt.Println("aArray6 := [4]int {10,20,30,40}")
	fmt.Println(" aSlice6 := aArray6[1:3] = ", aSlice6)
	fmt.Println("len(aSlice6) = ", len(aSlice6))
	fmt.Println("cap(aSlice6) = ", cap(aSlice6))

	aSlice6 = append(aSlice6, 900, -90, 50)
	fmt.Println("aSlice6 = append(aSlice6, 900, -90, 50)")
	fmt.Println(" aSlice6 = ", aSlice6)
	fmt.Println("len(aSlice6) = ", len(aSlice6))
	fmt.Println("cap(aSlice6) = ", cap(aSlice6))

	aArray7 := [5]int{10, 20, 30, 40, 50}
	aSlice7 := aArray7[:2]
	aArray2a := [5]int{5, 15, 25, 35, 45}
	ASlice7a := aArray2a[:2]

	aMessage := `
aArray7 := [5]int{10, 20, 30, 40, 50}
	aSlice7 := aArray7[:2]
	aArray2a := [5]int{5, 15, 25, 35, 45}
	ASlice7a := aArray2a[:2]`

	fmt.Println(aMessage)
	aNewSlice := append(aSlice7, ASlice7a...)
	fmt.Println("aNewSlice := append(aSlice7, ASlice7a...)")
	fmt.Println("aNewSlice", aNewSlice)

	fmt.Println("\n --  Deleting a slice")
	aArray8 := [5]int{10, 20, 30, 40, 50}
	fmt.Println("aArray8 := [5]int{10, 20, 30, 40, 50}")
	i8 := 2
	fmt.Println("aArray8 = ", aArray8)
	aSlice8_1 := aArray8[:i8]
	aSlice8_2 := aArray8[i8+1:]
	aNewSlice8 := append(aSlice8_1, aSlice8_2...)
	fmt.Println("aNewSlice8 := append(aSlice8_1, aSlice8_2...)")
	fmt.Println("aNewSlice8 = ", aNewSlice8)

	fmt.Println("\n---   Copying a slice")
	src_slice9 := []int{10, 20, 30, 40, 50}
	dest_slice9 := make([]int, 3)
	num := copy(dest_slice9, src_slice9)
	fmt.Println("num := copy(dest_slice9, src_slice9)")
	fmt.Println("dest_slice9 = ", dest_slice9)
	fmt.Println("number of elements copied", num)

	fmt.Println("\n ----- Looping through a slice")
	for index, value := range src_slice9 {
		fmt.Println(index, "=>", value)
	}

	fmt.Println("\n ----- Looping through a slice, index not needed")
	for _, value := range src_slice9 {
		fmt.Println(value)
	}

}

func arrays() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("arrays\n")

	var grades [5]int
	fmt.Println(grades)
	var fruits [3]string

	fmt.Println(fruits)

	fmt.Println("intialized arrays")
	var grades2 [3]int = [3]int{10, 20, 30}
	fmt.Println(grades2)
	var fruits2 [3]string = [3]string{"apples", "oranges", "grapes"}
	fmt.Println(fruits2)

	fmt.Println("intialized arrays - shorthand")
	grades3 := [3]int{10, 20, 30}
	fmt.Println(grades3)

	fmt.Println("intialized arrays - shorthand ..Ellipses")
	grades4 := [...]int{10, 20, 30}
	fmt.Println(grades4)

	fmt.Println("intialized arrays - shorthand ..Ellipses -  stings")
	fruits3 := [...]string{"apples", "oranges", "grapes"}
	fmt.Println(fruits3)

	fmt.Println("intialized arrays - shorthand ..Ellipses -  int")
	marks := [...]int{10, 20, 30}
	fmt.Println(marks)

	fmt.Println("len function")
	fmt.Println("len(fruits2) = " + strconv.Itoa(len(fruits2)))

	fmt.Println("Array Indexes")
	fmt.Println("grades2[1] = " + strconv.Itoa(grades2[1]))

	fmt.Println("Array Indexes #2 strings")
	fruits5 := [...]string{"apples", "oranges", "grapes", "mango", "papaya"}
	fmt.Printf("fruits5[2] = %s\n", fruits5[2])

	fmt.Println("Array Indexes #2 int = change value")
	grades5 := [...]int{90, 80, 70, 80, 97}
	fmt.Println(grades5)
	grades5[1] = 100
	fmt.Println(grades5)

	fmt.Println("loop through array")
	for i := 0; i < len(grades5); i++ {
		fmt.Println(grades5[i])
	}

	fmt.Println("loop through array .. with range")
	for index, element := range grades5 {
		fmt.Println(index, "=>", element)
	}

	fmt.Println("Multi-dementional array")
	arr1 := [3][2]int{{2, 4}, {4, 16}, {8, 64}}
	arr2 := [3][2]int{{2, 4}, {4, 16}, {8, 64}}
	fmt.Println(arr1)
	fmt.Println("arr1[2][1] = ", arr1[2][1])
	fmt.Println(arr2)
	fmt.Println("arr2[2][1] = ", arr2[2][1])

}

func operatorsAndControlFlow() {
	fmt.Println("Chapter 2")
	fmt.Println("\n\n Operators and Control Flow")

	// Comparison Operators
	comparisonOperator()
	arithmeticOperators()
	logicalOperators()
	assignmentOperators()
	bitwiseOperators()

	// Control Flow
	if_elseOperator()
	loopControlOperator()

}

func loopControlOperator() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("Loop Control functions\n")

	fmt.Println("-------------")
	fmt.Println("For loop ")
	for i := 1; i <= 3; i++ {
		fmt.Println("Hello World")
	}

	fmt.Println("-------------")
	fmt.Println("For loop ")
	for i := 1; i <= 5; i++ {
		fmt.Println(i * i)
	}

	fmt.Println("-------------")
	fmt.Println("For loop no initialization or post increment")
	var i int = 1
	for i <= 5 {
		fmt.Println(i * i)
		i += 1
	}

	fmt.Println("-------------")
	fmt.Println("If with break .. ends with 3")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		if i == 3 {
			break
		}
	}

	fmt.Println("-------------")
	fmt.Println("If with continue .. skips 3")
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

}

func if_elseOperator() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("if_elseOperator\n")

	var a string = "happy"
	if a == "happy" {
		fmt.Println(a)
	}

	fmt.Printf("\n")

	var fruit string = "grapes"
	if fruit == "apples" {
		fmt.Println("fruit is apple")
	} else {
		fmt.Println("Fruit is not apple")
	}

	fmt.Printf("\n")

	fruit = "grapes"
	if fruit == "apples" {
		fmt.Println("I love apples")
	} else if fruit == "oranges" {
		fmt.Println("I love oranges")
	} else {
		fmt.Println("No appetite")
	}

	var i int = 800
	switch i {
	case 10:
		fmt.Println("i is 10")
	case 100, 200:
		fmt.Println("i is either 100, 200")
	default:
		fmt.Println("i is neither 0,100 or 200")
	}

	i = 10
	switch i {
	case -5:
		fmt.Println("-5")
	case 10:
		fmt.Println("10")
		fallthrough
	case 20:
		fmt.Println("20")
		fallthrough
	default:
		fmt.Println("default")
	}

	var c, d int = 10, 20
	switch {
	case c+d == 30:
		fmt.Println("equal to 30")
	case c+d <= 30:
		fmt.Println("less than or equal to 30")
	default:
		fmt.Println("greater than 30")
	}

}

func bitwiseOperators() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("bitwiseOperators\n")

	fmt.Println("bitwise And (&) ")
	var x, y int = 12, 25
	z := x & y
	fmt.Printf("z := x & y = %d\n\n", z)

	fmt.Println("bitwise Or (|) ")
	x, y = 12, 25
	z = x | y
	fmt.Printf("z = x | y = %d\n\n", z)

	fmt.Println("bitwise XOR (^) ")
	x, y = 12, 25
	z = x ^ y
	fmt.Printf("z = x ^ y = %d\n\n", z)

	fmt.Println("Left Shift (<<) ")
	x = 212
	z = x << 1
	fmt.Printf("z = x << 1 = %d\n\n", z)

	fmt.Println("Right Shift (>>) ")
	x = 212
	z = x >> 2
	fmt.Printf("z = x >> 2 = %d\n\n", z)

	x, y = 100, 90
	z = (x + y) >> 2
	fmt.Printf("(x+y) >> 2)  %d\n\n", z)

	x, y = 100, 90
	b := !(((x + y) >> 2) == 47)
	fmt.Printf("!(((x+y) >> 2 ) == 47  ", b)

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
	x = 10
	y = 20
	x += y // x = x + y `
	x = 10
	y = 20
	x += y // x = x + y
	fmt.Println(multiline)
	fmt.Printf("x += y = %d", x)

	x = 10
	y = 20
	x -= y // x = x - y

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
