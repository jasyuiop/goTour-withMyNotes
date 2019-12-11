package main

import "fmt"

func main() {

	// struct literalleri aynı array literalleri gibidir
	// sadece slice literallerinde dizinin uzunluğu bildirilmez.

	// [3]bool{true, true, false} => bu array literalidir

	// []bool{true, true, false} => bu ise slice literalidir
	// bu aslında yukarıda verilen array literaliyle aynı işi yapar
	// bir dizi oluşturur, fakat aynı zamanda bu arrayı referans alan bir
	// slice ile birlikte!
	q := []int{2, 3, 4, 5, 6}
	fmt.Println(q)

	r := []bool{true, false, false}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{4, true},
	}
	fmt.Println(s)
}
