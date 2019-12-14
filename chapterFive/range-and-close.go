package main

import "fmt"

/*
Gönderici, bir kanaldan daha fazla değer gönderilmeyeceğini belirtmek adına
close fonksiyonu ile bu kanalı kapatabilir.

alıcı ifadeye, ikinci bir parametre ataması yaparak kanalın
kapalı olup olmadığını öğrenebilir:
v, ok := <-ch

Eğer alınacak başka değer kalmamış ve
kanal kapatılmışsa ok değişkeni false değerini alacaktır.

for i := range c döngüsü kanal kapanana kadar kanaldan tekrar tekrar değer alır.

Yalnızca gönderen kanalı kapatmalıdır, alıcı asla kapatmamalıdır.
Kapalı kanala bir şeyler göndermek "panic"'e yol açacaktır.

gönderen burada fibonacci fonksiyonumuz, o yüzden burada işler bitince kanalı
kapatıyoruz.

 Kanallar dosyalar gibi değildir; kapatmanız genellikle gerekmez.
 Kapatma işlemi yalnızca alıcıya daha başka değer gelmeyeceğini bildirmek
 maksadıyla yapılır, tıpkı bir range döngüsünü sonlandırmak gibi.
*/
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // c kanalını kapattım
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}
}
