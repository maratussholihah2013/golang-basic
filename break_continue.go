package main

import "fmt"

func main() {
	//break untuk menghentikan perulangan
	//continue untuk skip dan lanjut ke perulangan berikutnya
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Sudah waktunya break")
			break
		}
		if i == 3 {
			continue
		}
		fmt.Println("Perulangan ke ", i)
	}
}
