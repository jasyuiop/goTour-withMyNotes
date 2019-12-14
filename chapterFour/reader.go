package main

import (
	"fmt"
	"io"
	"strings"
)

/*
io paketi veri akışının sonunu okuyan io.Reader arayüzünü kapsar.

 io.Reader arayüzü Read metoduna sahiptir:
 func (T) Read(b []byte) (n int, err error)

  Read metodu, kendisine verilen byte dilimi ile veri üretir ve üretilen
   verinin kaç byte olduğu ile hata değerini döner.
   Eğer veri akışı sonlanmışsa io.EOF hatası döner.

Örnek kod bir strings.Reader oluşturuyor ve tek seferde 8 baytlık çıktı işliyor.
*/
func main() {
	r := strings.NewReader("Hello Reader")

	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
