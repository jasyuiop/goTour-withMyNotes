package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Stringer fmt paketinde tanımlı bir interface'tir
// kendini string olarak tanımlayan bir türdür.
// fmt paketi ve bir çok paket değerleri yazdırmak için bu interface'e bakar.
/* type Stringer interface {
    String() string
}*/
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Emre", 20}
	z := Person{"Levo", 23}
	fmt.Println(a, z)
}
