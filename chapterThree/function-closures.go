package main

import "fmt"

/*
bir fonksiyon kendi kapsamı dışındaki bir değişkeni kullanıyorsa
ve bu fonksiyon başka bir kapsamdan çalıştırılsa bile
o değişkene hala erişimi vardır.*/

/*Closure bir fonksiyonun,
başka bir lexical scope tarafından çağırılsa bile kendi lexical
scope’unu hatırlamasıdır*/
func adder() func(int) int {
	// main içerisindeki for döngüsünde her turda bu sum değişkeninin içindeki değer
	// hatırlanıyor, bir nevi fonksiyon ile değişken arasında bağ oluşuyor.
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}
