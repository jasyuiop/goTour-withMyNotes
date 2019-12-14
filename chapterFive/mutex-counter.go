package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Peki ya haberleşmeye ihtiyacımız yoksa? Ya tek istediğimiz bir değişkene,
çakışmalardan sakınmak için, aynı anda sadece bir gorutinin erişmesini sağlamak ise?

Bu mekanizmaya _karşılıklı dışlama_ ("mutual exclusion") denilmekte ve mekanizmayı
sağlayan veri türüne geleneksel olarak mutex adı verilmektedir.

 Standart Go kitaplığında karşılıklı dışlama sync.Mutex veri türü ve ona ait
 iki metodla sağlanır:
    Lock
    Unlock

Karşılıklı dışlamayla çalıştırılacak bir kod öbeğini,
Inc metodunda görüldüğü gibi, kodu Lock ve Unlock çağrılarıyla çevreleyerek
tanımlayabiliriz.

Mutex kilidinin kaldırıldığından emin olmak için,
alue metodunda görüldüğü gibi, defer yerleşik metodunu da kullanabiliriz.
*/
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock() // kilitledik, böylecek aynı anda sadece tek goroutine map'e erişebilir.
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 100; i++ {
		go c.Inc("SomeKey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("SomeKey"))
}
