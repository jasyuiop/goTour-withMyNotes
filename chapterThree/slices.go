package main

import "fmt"

func main() {
	// slicelar dinamik olarak boyutlandırılabilir.
	sayilar := [6]int{2, 3, 4, 5, 6, 7} // bu bir dizi

	//a[low : high]
	// slicelar iki indis arasında belirtilen elemanları alır.
	// verilen ilk indis dahil edilip, son verilen ise alınmaz
	// yani bu örnekte 1'nci indisten başlayıp 3'ncü indis dahil alacak.
	var s []int = sayilar[1:4] // bu bir slice
	fmt.Println(s)
}
