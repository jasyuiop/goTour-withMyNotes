package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}

	s = s[1:4] //2,3,4
	fmt.Println(s)

	// eğer ":"'dan önce birşey yazmaz isek dizinin 0'ıncı indisi demiş gibi oluruz
	s = s[:2] // 2,3
	fmt.Println(s)

	// eğer ":"'dan sonra birşey yazmaz isek dizinin son indisine kadar demiş oluruz.
	s = s[1:] // 3
	fmt.Println(s)

	/* örneğin;
	var a [10]int
	olsun...

	bu alttaki slice'ların hepsi birbirine eşittir.
	a[0:10]
	a[:10]
	a[0:]
	a[:]
	*/
}
