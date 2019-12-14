package main

import "fmt"

/* Interface'in içindeki somut değer nil ise method nil olan bir receiver ile
çağrılacaktır.
Bazı dillerde bu durum boş işaretçi istinası ("null pointer exception") üretir
Go'da ise receiver'in nil olması durumunu kontrol ederek metodlar yazmak sıkça rastlanan bir durumdur.
*/
type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I
	// somut değer burada nil, fakat interface value'sunun nil olmamasını sağladık.
	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
