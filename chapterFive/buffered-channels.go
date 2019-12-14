package main

import "fmt"

func main() {
	/*
		buffered channels(tanponlanmış kanallar);
		buffered channel'i ilklendirmek(initialize) için make fonksiyonunun
		ikinci parametresine tampon("buffer") uzunluğunu giriyoruz
		ch := make(chan int, 100)

		Tamponlanmış bir kanala veri gönderimleri,
		sadece tampon tamamen dolu olduğunda durur.
		Tamponlanmış bir kanaldan veri alımları ise tampon tamamen boş olduğunda durur.

		tampon bölge bir nevi o channel'e kaç ayrı işlem gelip hafızada tutulacak
		onu belirtiyoruz gibi birşey.
	*/
	ch := make(chan int, 2) // 2 tamponlu channel oluşturduk.
	ch <- 1
	ch <- 2
	fmt.Println(<-ch) // channel'den veriyi alıp fonksiyonun içine atıyoruz
	fmt.Println(<-ch)

}
