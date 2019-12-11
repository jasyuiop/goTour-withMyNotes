package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {

	// range slice ve map'lerde yenileme yapabilen bir keyworddür
	// bir slice üzerinde range kullanılırken her döngü için iki değer döndürülür
	// birincisi index, ikincisi ise bu arraydeki elemanın bir kopyası
	// burada her döngüde "i" bana index'i "v" ise slice'ımda o indisteki elemanı döndürüyor.
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
