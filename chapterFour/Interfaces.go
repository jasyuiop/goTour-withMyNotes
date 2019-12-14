/*
Interface sayesinde; function'dan dönen değerleri ve
değerlerin tiplerini bir başka yerde kullanmak için şekillendirilebilir.
Interface'ler nesneler arasındaki iletişimi sağlamak için kullanılır.
*/

package main

import (
	"fmt"
	"math"
)

// Abser adında bir interface oluşturdum.
type Abser interface {
	Abs() float64 // Abs fonksiyonunu çağıracak ve Interface'in geri dönüş tipi float64 olacak
}

func main() {
	var a Abser               // Abser interface'i tipinden bir değişken oluşturdum
	f := MyFloat(-math.Sqrt2) // -1.4xxx
	v := Vertex{3, 4}

	a = f
	fmt.Println(a.Abs()) // MyFloat alıcısına sahip fonksiyon çağrıldı

	a = &v // çağıracağım fonksiyon pointere sahip olduğu için & işareti ile
	// bellekteki adresine işaret ediyoruz.

	fmt.Println(a.Abs()) // Vertex alıcısına sahip fonksiyon çağrıldı
}

type MyFloat float64

type Vertex struct {
	X, Y float64
}

// iki fonksiyonun isimleri aynı fakat alıcıları farklı, overload ettik bir nevi.

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
