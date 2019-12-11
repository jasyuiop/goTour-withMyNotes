package main

import "fmt"

func main() {
	m := make(map[string]int)

	// eleman eklemek
	m["Answer"] = 42
	fmt.Println("the value : ", m["Answer"])

	// elemanı güncellemek
	m["Answer"] = 48
	fmt.Println("the value : ", m["Answer"])

	// elemanı silmek
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// elemanın varlığını test etmek
	// eğer eleman map'te yok ise geri döndüreceği değer 0'dır
	// key map'in içince var ise burada ok'a döndüreceği değer true'dur.
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
