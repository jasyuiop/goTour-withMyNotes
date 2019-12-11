package main

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func main() {
	// make kullanarak 5 kapasiteli, elemanlarını 0 ile ilk değer atanması yapılmış
	// bir array yaratıyorum
	a := make([]int, 5) // len 5 cap 5
	printSlice("a", a)

	// 5 kapasiteli fakat içerisine ilk değer atanması yapılmamış boş elemanlı bir
	// dizi yaratıyorum.
	b := make([]int, 0, 5) // len 0 cap 5
	printSlice("b", b)

	// bir slice ile b dizisini referans aldığımda ise
	// 2'nci elemanına kadar talep ediyorum, b dizisinin kapasitesine sahip
	// bir slice oluyor, ve aldığım o 2 elemana ilk değer ataması yapılıyor.
	c := b[:2] // len 2 cap 5
	printSlice("c", c)

	// ilk baştaki 2 elemanı atınca geriye bana 3 kapasiteli bir dizi kalıyor
	// ve 2:5 arası aldığım için buradan da 3 elemana ilk değer ataması yapılıyor.
	d := c[2:5]
	printSlice("d", d)

	/* birkaç örnek daha ...
	b = b[:cap(b)] // len(b)=5, cap(b)=5
	b = b[1:]      // len(b)=4, cap(b)=4
	*/
}
