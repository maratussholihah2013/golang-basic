package main

import "fmt"

func main() {
	name := "Zayyan"

	switch name {
	case "Maratus":
		fmt.Println("Hai Maratus")
	case "Sholihah":
		fmt.Println("Hai SHolihah")
	case "Erdiansyah":
		fmt.Println("Hai Erdiansyah")
	default:
		fmt.Println("Siapa namamu?")
	}

	//menambahkan short statement dalam if
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Namamu cukup bagus")
	}

	//switch tanpa kondisi, kondisi di casenya
	length_1 := len(name)
	switch {
	case length_1 > 5:
		fmt.Println("Nama terlalu panjang")
	case length_1 <= 5:
		fmt.Println("Namamu cukup bagus")
	}
}
