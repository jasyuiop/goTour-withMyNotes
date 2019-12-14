package main

import "fmt"

/*
select keyword'ü bir "goroutine"nin birden
fazla iletişim işleminde beklemesine izin verir.

select, case yollarından biri çalışabilir hale gelinceye dek bloke eder
Eğer birden fazla çalıştırılabilir case varsa aralarından rastgele birini seçer.
*/

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}() // asekron olarak bu fonksiyon çalışacak.
	fibonacci(c, quit)
}
