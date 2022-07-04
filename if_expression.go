package main

import "fmt"

func main() {
	name := "Maratus"
	if name == "Maratus" {
		fmt.Println("Hello Maratus")
	} else if name == "Sholihah" { //else harus sebaris dengan penutup kurawal "}" dari if
		fmt.Println("Hai Sholihah")
	} else {
		fmt.Println("Hei, siapa kamu?")
	}

	//menambahkan short statement dalam if
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	}
}
