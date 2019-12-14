package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	// v.Scale(2)'i method olduğu için golang "(&v).Scale(2)" olarak yorumlar
	// pointer receiver otomatik olarak çağrılır, işaret edilir.

	// bir fonksiyon olduğu için ScaleFunc()'da ise; ya biz değişken yaratırken
	// pointer'e işaret ederiz ya da fonksiyona parametreyi gönderirken işaret
	// ettiririz

	v := Vertex{3, 4}
	v.Scale(2)        // 6,8
	ScaleFunc(&v, 10) // 60,80

	p := &Vertex{4, 3}
	p.Scale(3)      // 12,9
	ScaleFunc(p, 8) // 96,72

	fmt.Println(v, p)
}
