package main

/*
import (
		"fmt"
)
*/

import "fmt"

func main() {
	var firstName string = "Srijon"
	var lastName string = "Paul"
	initials := "S.P"
	fmt.Println("His first name is ", firstName, ", last name is ", lastName, " and initials are ", initials)
	b, c := 1, 2
	d := -1 == (b - c)
	fmt.Println(d)
}
