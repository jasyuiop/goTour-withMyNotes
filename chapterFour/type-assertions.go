package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// i interface değerinin string türünü barındırdığını söylüyoruz
	// ve arka plandaki(burada "hello") s değişkenine atıyoruz.
	s := i.(string)
	fmt.Println(s)

	// interface değerinin belirtilen türü barındırıp barındırmadığını anlamak
	// için tür onayı(type-assertions) kullanırsak;
	// arkadaki değeri ve onaylamanın başarılı olup olmadığını anlamamız için
	// mantıksal değer dönderir

	// eğer i belirtilen türde bir somut değer barındırıyorsa ok değişkenine true
	// değeri dönecek.
	s, ok := i.(string)
	fmt.Println(s, ok)

	// eğer i belirtilen türde somut değer barındırmıyorsa o halde değere 0,
	// mantıksal değere ise false değeri dönecek.
	// böylelike panic gerçekleşmeyecek.
	f, ok := i.(float64)
	fmt.Println(f, ok)

	// type-assertions kullanmazsak burada float64 tipinde değer barındırmadığı için
	// panic hatası alacağız.
	f = i.(float64)
	fmt.Println(f)
}
