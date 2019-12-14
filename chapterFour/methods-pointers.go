package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
method tanımlarken alıcıyı pointer ile gösterebiliriz.
pointer ile alıcıyı(receivers) tanımladığımızda, receivers'in
işaret ettiği value'ları değiştirebiliriz (aşağıda Scale bunu yapıyor).

method'lar genelde receivers'lerini değiştirmesi gerektiğinden (napacağına göre...)
pointers receivers kullanımı normalinden daha yaygındır.

Scale method'undaki alıcının pointeri kaldırıldığında içerisindeki x ve y değerleri
değiştirilemiyor.
değiştirmesi için pointer'a sahip olması gerekiyor.
*/

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	v.Scale(10)
	fmt.Println(v.Abs())
}
