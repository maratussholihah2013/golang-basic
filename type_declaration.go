package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	//membuat var dr type declarations
	var noKTPnasabah NoKTP = "354829473920"
	var statusNikahNasabah Married

	fmt.Println(noKTPnasabah)
	fmt.Println(statusNikahNasabah)
}
