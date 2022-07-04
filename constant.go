package main

import "fmt"

func main() {
	//deklarasi dan langsung diberi value-nya
	//boleh pakai tipe data atau tidak pakai tipe data
	const firstname string = "Maratus"
	const phi = 3.14

	//bisa multi deklarasi const
	const nama, bahasa, usia = "Maratus", "Golang", 30

	//variable const tidak harus digunakan meski telah dideklarasikan
	//berbeda dengan jenis variable biasa yang harus digunakan setelah dideklarasikan
	//ex: firstname tidak menampilkan error meski tidak digunakan di code berikutnya
	fmt.Println(phi)

}
