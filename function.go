package main

import "fmt"

type Filter func(string) string
type Blacklist func(string) bool

func main() {
	sayHello()
	sayHelloTo("Maratus", "Sholihah")
	sayHelloTo("Zayyan", "Erdiansyah")
	fmt.Println(getHello("Maratus Sholihah"))

	namaPertama, namaKedua := getFullName()
	fmt.Println(namaPertama, namaKedua)

	a, b, c := getCompleteID()
	fmt.Println(a, b, c)

	total := sumAll("Total variadic function", 10, 30, 50)
	fmt.Println(total)

	slice := []int{10, 30, 50}
	totalSlice := sumAll("Total slice", slice...)
	fmt.Println(totalSlice)

	sayHelloWithFilter("Maratus", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
	sayHelloWithFilter("Sholihah", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)

	//anonymous function
	blacklist := func(name string) bool {
		return name == "admin"
	}
	registerUser("admin", blacklist)
	registerUser("Maratus", blacklist)

	//factorialLoop
	loop := factorialLoop(5)
	fmt.Println(loop)
	//recursive
	recursive := factorialRecursive(5)
	fmt.Println(recursive)
}

func sayHello() {
	fmt.Println("Hello")
}

//function dengan parameter
func sayHelloTo(firstname string, lastname string) {
	fmt.Println("Hello ", firstname, lastname)
}

//function dengan return value
func getHello(name string) string {
	result := "Hai " + name
	return result
}

//returning multi value
func getFullName() (string, string) {
	return "Zaky", "Erdiansyah"
}

//named return values
func getCompleteID() (nik string, name string, age int) {
	nik = "73842927"
	name = "Maratus Sholihah"
	age = 30

	return
}

//variadic function
func sumAll(note string, numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}

//function as parameter
func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello ", nameFiltered)
}
func spamFilter(name string) string {
	if name == "Anjing" {
		return "...."
	} else {
		return name
	}
}

//function as parameter dengan type declarations, see import type
func sayHelloWithFilter2(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello ", nameFiltered)
}

//anonymous function
func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Blocked ", name)
	} else {
		fmt.Println("Welcome ", name)
	}
}

//factorial dengan loop
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

//factorial menggunakan recursive function
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
