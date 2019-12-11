package main

import "fmt"

func main() {
	// dizilerin uzunlukları çalışma zamanında yeniden boyutlandırılamaz bunun için slicelar var
	var a [2]string // 2 elemanlı string bir dizi oluşturdum
	a[0] = "hello"  // 0 elemanına değer atadım
	a[1] = "world"
	fmt.Println(a[0], a[1]) // 0 ve 1'nci elemanlarını yazdırdım
	fmt.Println(a)          // direkt diziyi yazdırdım.

	sayilar := [6]int{1, 2, 3, 4, 5, 6}
	// 6 elemanlı int bir dizi oluşturdum ve değer atamalarını oluştururken yaptım.
	fmt.Println(sayilar)
}
