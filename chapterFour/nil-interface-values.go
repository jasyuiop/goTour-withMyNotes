package main

import "fmt"

/*
nil interface value'su değer ve somut bir tür barındırmaz.
"nil" interface üzerinden metod çağırmak runtime hatasına neden oluyor

interface'i temsil eden değer,tür çiftinde tür bilgisi olmadığı için
hangi method çağrılacak bilinemez.
*/
type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)

}
