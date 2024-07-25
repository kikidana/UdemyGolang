package main

import "fmt"

func main() {
	var name string
	name = "Rizky"
	fmt.Println(name)

	name2 := "Rizky2"
	name2 = "Perdana"
	fmt.Println(name2)

	var (
		firstName  = "Muhammad"
		middleName = "Rizky"
		lastName   = "Perdana"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
