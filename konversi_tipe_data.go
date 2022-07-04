package main

import "fmt"

func main() {
	var nilai32 int32 = 100000
	//nilai kecil ke lebih besar
	var nilai64 int64 = int64(nilai32)
	//konversi dari besar ke kecil mengakibatkan perubahan nilai
	var nilai8 int8 = int8(nilai32)

	//lihat perbedaan hasilnya melalui code berikut
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	//bandingkan dengan berikut
	var nilai32_2 int32 = 30
	var nilai64_2 int64 = int64(nilai32_2)
	var nilai8_2 int8 = int8(nilai32_2)

	fmt.Println(nilai32_2)
	fmt.Println(nilai64_2)
	fmt.Println(nilai8_2)

	//ubah byte ke string
	var name = "Maratus Sholihah"
	var e byte = name[4]
	var eString = string(e)
	var estring_2 = string(name[4])
	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
	fmt.Println(estring_2)
}
