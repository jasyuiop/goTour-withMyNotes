/*
pointer receiver kullanmanın 2 nedeni vardır;
1- method'un receiver'inin pointer ile gösterdiği değeri değiştirebilmesidir.
2-her method çağrıldığında value'nun ram'de kopyalanmasını önlemektir.

Genel olarak, belirli type'lardaki tüm method'lar value ve pointer
receivers'lerine sahip olmalı, ancak her ikisinin de karışımı
olmamalıdır.


*/

// Abs methodunun receiver'inin değerini değiştirmesi gerekmese de
// ikisi de burada *Vertex olarak yazılmış.

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
