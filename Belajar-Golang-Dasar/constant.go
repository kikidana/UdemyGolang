package main

import "fmt"

func main() {
	const firstName string = "Rizky"
	const lastName = "Perdana"

	// error jika diubah nilainya
	//firstName = "Rizku"
	//lastName = "Perdanu"

	const (
		firstName1 string = "Rizky"
		lastName1         = "Perdana"
	)

	fmt.Println(firstName, lastName)
	fmt.Println(firstName1, lastName1)
}
