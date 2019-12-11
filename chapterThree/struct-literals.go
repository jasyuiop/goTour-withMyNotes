package main

import "fmt"

type Insan struct {
	isim, soyisim string
}

var (
	i1 = Insan{"emre", "nefesli"}
	i2 = Insan{isim: "levent"}       // soyisim boş
	i3 = Insan{}                     // isim ve soyisim boş
	p  = &Insan{"burak", "polatkan"} // Insan struct'ına pointer ile işaret ettim.
)

func main() {
	fmt.Println(i1, i2, i3, p)

}
