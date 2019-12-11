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
	// "&" işaretcisi ile rex'in var olan yapısına erişip içerisindeki field'ini değiştirdim.
	// rex'in yapısına eriştikten sonra fieldlerine de erişebileceğim için "."
	// koyarak değerini almadan("*") erişebiliyorum.
	rexPointers := &rex
	rexPointers.cins = "pitbull"
	fmt.Println(rexPointers)
}
