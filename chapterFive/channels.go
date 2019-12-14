/*
Kanallar ("channels"), <- kanal operatörü vasıtası ile değer
gönderip alabildiğiniz bir veri hattıdır.

ch <- v    // v'yi ch kanalına yolla.
v := <-ch  // ch'den al ve değeri v'ye ata.

(Veri akışı ok yönünde gerçekleşir.)

tıpkı map ve slice'lar gibi chaneller'de kullanılmadan önce
oluşturulmalıdır.
ch := make(chan int)

Varsayılan olarak, veri gönderimleri ve alımları diğer taraf
hazır olana kadar bloklanır. Bu,
"goroutine"lerin açık kilitler ("explicit locks") veya
koşul değişkenleri ("conditional variables") kullanmadan eşzamanlanmasını sağlar.

Örnek kod, slice'daki sayıların toplamını hesap yükünü iki goroutine'e
paylaştırarak hesaplar. Her iki goroutine de işlerini tamamladığında
nihai sonuç hesaplanır.
*/

package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s { // index değerini almadık
		sum += v
	}
	c <- sum // sum'u c channel'ine gönderdik
}

func main() {
	s := []int{3, 5, 7, 9, -4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c) // s array'inin 0'ından uzunluğunun yarısına kadar olan
	// elemanları ve oluşturduğum c channel'ini eşzamanlı çalışacak şekilde sum
	// fonksiyonuna parametre olarak gönderdim.

	go sum(s[len(s)/2:], c) // diğer yarısını da gönerdim

	// iki eşzamanlı işlem bitene kadar değer atamaları bloklanacak!
	x, y := <-c, <-c // c channelinden ilk birinci eşzamanlı işlem bitince x'e
	// ikinci eşzamanlı işlem bitince y'ye atama gerçekleşecek.

	fmt.Println(x, y, x+y)
}
