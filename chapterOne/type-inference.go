package main

import "fmt"

func main() {
	// := ve = kullanırken değişken türü açıkca belirlenmezse
	// değişkenin türü sağ taraftaki aldığı değere göre belirlenir.
	v := 3 // change me!
	fmt.Printf("v is of type %T\n", v)
	s := 3.142 // change me!
	fmt.Printf("s is of type %T\n", s)
	g := 0.867 + 0.5i // change me!
	fmt.Printf("g is of type %T\n", g)
}
