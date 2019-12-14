package main

/*
error türü fmt.stringer interface'ine benzer şekilde, fmt paketine
yerleşik bir interface'dir.
type error interface {
    Error() string
}
Fonksiyonlar sıklıkla bir error değeri döndürür ve
fonksiyonu çağıran kod, hata değerinin nil değerine eşit olup olmadığını
kontrol ederek hataları yakalamalıdır.
*/
import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}
func main() {
	/*
			run() fonksiyonunu err değişkenine attığımda
		    MyError struct'unun içerisindeki değerleri doldurup
			Error() fonksiyonunu kendisi otomatik çağırıyor
			çünkü bu error interface'i içerisinde çağrılan bir fonksiyon.

			err değişkenimin içi değerlerle doldu,
			err nil değilse ekrana yazdır dediğim zaman Error()
			method'undan dönen değeri yazdırıyor.

	*/
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
