package main

import (
	"fmt"
	"time"
)

/*
Eğer hazırda çalışacak başka case yoksa select
içerisindeki default çalıştırılır.
Bloklanmadan veri alma ya da gönderme yapmak için default kullanın.
*/
func main() {
	tick := time.Tick(1000 * time.Millisecond)
	boom := time.After(3000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("boom")
			return
		default:
			fmt.Println(".........")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
