package main

import (
	"fmt"
	"math"
)

type I interface {
	// interface başlığının altındaki değerler
	// bir değer ("value") ve somut bir türden ("concrete type") oluşan bir çift
	// şeklinde düşünülebilir.
	//(value, type)

	// interface'de bir method çağırma; değer burada  M()
	// bir interface value'su üzerinden çağrılan method arkadaki türe ait
	// aynı isimli methodu çalıştırır.
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"} // burada (t *T) alıcısına sahip methodu
	describe(i)
	i.M()

	i = F(math.Pi) // burada (f F) alıcısına sahip methodu
	describe(i)
	i.M()
}
func describe(i I) {
	fmt.Println("(%v, %T)\n", i, i)
}
