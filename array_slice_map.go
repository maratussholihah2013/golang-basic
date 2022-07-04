package main

import "fmt"

func main() {
	///////   ARRAY   ////////
	//deklarasi satu-satu
	var names [3]string
	names[0] = "Maratus"
	names[1] = "Sholihah"
	names[2] = "Erdiansyah"
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	//deklarasi langsung
	var values = [3]int{10, 30, 60}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	//mendapatkan panjang array
	fmt.Println(len(values))

	///////  SLICE   /////////
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"Nopember",
		"Desember",
	}
	//membuat slice low 4 high 7
	var slice1 = months[9:10]
	fmt.Println(slice1)
	//mendapatkan length dari slice
	fmt.Println(len(slice1))
	//mendapatkan capacity dari slice
	fmt.Println(cap(slice1))

	//jika array diubah maka slicenya pun akan berubah begitu juga sebaliknya slice berubah maka array juga berubah
	months[9] = "AnotherOktober"
	fmt.Println(slice1[0])

	//menambah data baru ke slice jadi slice baru,
	//jika masih ada capacity maka akan merubah array months dan slice sebelumnya
	//jika tidak ada capacity maka akan membuat array baru
	//bandingkan array, slice1,slice2,slice3
	slice2 := append(slice1, "Bulan baru")
	fmt.Println(slice2)
	fmt.Println(slice1)
	fmt.Println(months)
	//slice 3 tidak mengubah array months karena melebihi capacity maka membuat array baru lagi
	slice3 := append(slice2, "Tahun baru", "Tahun super baru")
	fmt.Println(slice3)

	//make membuat slice baru
	newSlice := make([]string, 2, 5)
	newSlice[0] = "TK"
	newSlice[1] = "SD"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//copy slice, pastikan length dan capacity sama agar data tidak terpotong
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	/////////   MAP    ///////
	person := map[string]string{
		"name":    "Maratus Sholihah",
		"address": "Bandung",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Golang Language"
	book["author"] = "Maratus Sholihah"
	book["edition"] = "Ketiga"

	fmt.Println(len(book))
	fmt.Println(book)

	delete(book, "edition")

	fmt.Println(len(book))
	fmt.Println(book)
}
