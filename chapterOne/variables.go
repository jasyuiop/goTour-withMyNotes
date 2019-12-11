package main

import "fmt"

var isim, soyisim, hisLive bool

// sonunda yazdığım bool değeri değişkenler için initializer işlemi yapıyor
// yani onları oluşturuyor, var ediyor,başlatıyor. değişkenler initalizerin değerini alıyor.

func main() {
	var i int
	fmt.Println(i, isim, soyisim, hisLive)
	// := kısatlması var bildiriminin yerini alabiliyor. consturct ile kullanılmıyor
	// çünkü fonksiyonların dışında her statment bir keyword ile başlamak zorunda.
	isim := "emre"
	soyisim := "nefesli"
	fmt.Println(isim, soyisim)
}
