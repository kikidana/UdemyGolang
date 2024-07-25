package main

import "fmt"

func main() {
	type NoKTP string

	var ktpRizky NoKTP = "3214321410103214"

	var contoh string = "1234123412341234"
	var contohKTP = NoKTP(contoh)

	fmt.Println(ktpRizky)
	fmt.Println(contohKTP)
}
