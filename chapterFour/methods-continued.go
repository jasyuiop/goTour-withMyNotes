package main

import (
	"fmt"
	"math"
)

/* method tanımlarken struct bir type sahip alıcı gerekli değildir
type olarak tanımlanmış bir değişken olması da yeterlidir.
*/
/*
bir method tanımlarken verilen alıcının, method ile aynı pakette
olması gereklidir.
type başka bir pakette tanımlanmış ise o method'a alıcı olarak o
type verilmez.
*/
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
