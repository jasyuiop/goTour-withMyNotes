package main

import "fmt"

func main() {
	i, j := 42, 2701
	// "&" operatörünü önüne eklediğin değişkene işaret eder
	// "*" operatörünü önüne eklediğin değişkenin değerini verir
	p := &i         // i'e işaret ettim
	fmt.Println(*p) // işaretci vasıtasıyla i'nin değerini okudum
	*p = 21         // işaretci vasıtasıyla i'nin değerini değiştirdim
	fmt.Println(i)  // değiştirdiğim değeri direkt i'den okudum

	p = &j         // j'nin bellekteki adresine işaret ettim
	*p = *p / 37   // işaretci vasıtasıyla j'nin değerini 37 ile bölüp kendisine atadım.
	fmt.Println(j) // değiştirdiğim değeri direkt j'den okudum

}
