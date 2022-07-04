package main

import "fmt"

func main() {
	//operasi perbandingan: >,<,>=,<=,==,!=
	name1 := "Maratus"
	name2 := "Mara tus"
	result := name1 == name2
	val1, val2 := 300, 500
	fmt.Println(result)
	fmt.Println(val1 > val2)
	fmt.Println(val1 < val2)
	fmt.Println(val1 != val2)
	fmt.Println(val1 == val2)

	//operasi boolean: &&, ||, !
	var nilaiakhir = 90
	absensi := 80

	lulusUjian := nilaiakhir > 80
	lulusAbsensi := absensi > 80

	lulus := lulusAbsensi && lulusUjian

	fmt.Println(lulus)
}
