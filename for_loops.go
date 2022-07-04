package main

import "fmt"

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println("Nilai counter = ", counter)
		counter++ //->ini contoh post statement, dieksekusi setelah perulangan
	}
	//for dengan statement, init statement dieksekusi sebelum perulangan -> i++
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//For Range
	slice := []string{"Maratus", "Sholihah", "Zayyan", "Erdiansyah"}
	//umumnya for untuk array,slice, dan map seperti berikut
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	//dengan for range menjadi berikut
	for key, value := range slice {
		fmt.Println("index ", key, " = ", value)
	}
	//jika hanya butuh data value saja tanpa indexnya maka gunakan _
	for _, value := range slice {
		fmt.Println("Data = ", value)
	}

}
