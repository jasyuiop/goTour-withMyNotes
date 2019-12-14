package main

import "fmt"

// interface'ler başka paketlere implement edilebilir
type I interface {
	M()
}

type T struct {
	S string
}

// bu method'ta T type'ının I interface'ine implement(dahil etme) yaptık.
// implement etmek için kullanılan bir keyword yok, açıkça belirtemiyoruz yani.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
