package main

import "fmt"

func main() {
	//deklarasi variable dengan tipe data
	var name string
	name = "Maratus Sholihah"
	fmt.Println(name)

	name = "Zayyan Erdiansyah"
	fmt.Println(name)

	//deklarasi variable tanpa tipe data
	var age = 17
	fmt.Println(age)

	//deklarasi variable tanpa kata kunci var tapi pakai ':='
	country := "Indonesia"
	fmt.Println(country)

	//deklarasi multi variable
	var (
		firstname = "Maratus"
		lastname  = "Sholihah"
	)
	fmt.Println(firstname)
	fmt.Println(lastname)

	//deklarasi multi variable
	firstword, lastword := "Zayyan", "Erdiansyah"
	fmt.Println(firstword)
	fmt.Println(lastword)
}
