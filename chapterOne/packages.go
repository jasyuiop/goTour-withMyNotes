package main

import (
	"fmt"
	"math/rand"
	// tüm kütüphaneyi import etmektense
	// math sınıfından rand ile başlayan paketleri import ettim
)

func main() {
	fmt.Println("en sevdiğim sayı", rand.Intn(5))
}
