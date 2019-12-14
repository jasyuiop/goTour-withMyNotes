package main

import "fmt"

/*type switch seri halde çok sayıda type-assertions(tür onayı) yapılmasına
olanak sağlayan bir yapıdır.

type switch normal switch gibi çalışır, fakat her durum ilgili türü belirtir
değerleri değil. veya bu türler interface değerleriyle ilişkili türle karşılaştırılır.

type switch'de bildirim type-assertions'taki i.(T) sözdizimiyle aynı.
sadece buraya özel tür yazmak yerine type keyword'ünü yazıyoruz.
*/
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
