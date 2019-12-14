package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

/*Bir goroutine Go tarafından çalışma zamanında
yönetilen hafif bir iş parçacığıdır ("lightweight thread").

Goroutine'de bir işlem gerçekleştirmek istiyorsak başına sadece go keywordü
yazmamız yeterlidir.

Goroutine aslında fonksiyon gibi çalışır, eş zamanlı(asekron) fonksiyonları
çağırmak için kullanılır.

Gorutinleri aynı adres uzayında çalışır,
bu yüzden paylaşılmış bellek erişimi mutlaka eşzamanlanmalıdır.
sync paketi yararlı bazı eşzamanlama ilkelleri sağlar;
bununla birlikte Go'da bulunan diğer eşzamanlama ilkelleri
sayesinde buna çok ihtiyaç duymayacaksınız.
*/

func main() {
	go say("world")
	say("hello")
}
