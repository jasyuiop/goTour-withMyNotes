package main

import "fmt"

type Kopekler struct {
	renk string
	cins string
	yas  int
}

func main() {
	rex := Kopekler{"beyaz", "dogo argentino", 2}
	fmt.Println(rex)

	// "." ile struct'umun içersindeki fields'lara erişebiliyorum.
	rex.cins = "labrador"
	fmt.Println(rex)
}
