package main

import (
	"fmt"
	"math"
)

/*
Go dilinde classlar yoktur ancak type'ler üzerine methods'lar tanımlanabilir.
bir method aslında özel bir alıcıya sahip fonksiyondur.
*/
type Vertex struct {
	X, Y float64
}

/*
fonksiyonu yazarken alttaki gibi özel bir alıcı belirttiğimizde;
bu tipe artık (burada "Vertex") bu fonksiyon tanımlanıyor
*/
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.Y + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
