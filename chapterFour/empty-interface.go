package main

import "fmt"

/*herhangi bir method belirtilmemişse bu interface empty olarak adlandırılıyo
empty interface herhangi bir türde değer barındırabilir.
empty interface'lerin bilinmeyen türde değerleri yakalaması gerektiğinde
kullanılması makbuldür.
*/
func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)
	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
