package main

import "fmt"

func main() {
	var nilai32 int32 = 32767
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	nama := "Muhammad Rizky Perdana"
	charNama := nama[0]
	stringNama := string(charNama)

	fmt.Println(nama)
	fmt.Println(charNama)
	fmt.Println(stringNama)
}
