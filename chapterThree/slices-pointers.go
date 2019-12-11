package main

import "fmt"

func main() {
	// slicelar arrayları referans alır
	// bu demektir ki slice'lar dataları barındırmaz
	// sadece referans aldığı array'in belirli bir bölümünü ifade eder

	// slice'ın elemanlarını değiştirmek referans alınan dizinin ilgili
	// elemanlarında değişikliğe yol açar

	// bir array'i birden fazla slice kullanıyor ise
	// bu değişiklik onlara da yansır.

	names := [4]string{
		"emre",
		"levent",
		"burak",
		"ethem",
	}

	fmt.Println(names)

	a := names[0:2] // dizinin 0 ve 1'ncı elemanını alıyor.
	b := names[2:4] // dizinin 2 ve 3'ncü elemanını alıyor
	fmt.Println(a, b)

	b[0] = "yıkık" // b'nın 0'ı yani "burak"'ı değiştirdim
	fmt.Println(a, b)
	fmt.Println(names) // bu arraya da yansıyor

}
