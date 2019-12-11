package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// map'ler key ve value'dan oluşur.
var m map[string]Vertex

func main() {
	// make fonksiyonu verilen tipte bir map dönderiyor
	// burada string tipli key ve vertex tipli value
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68, -74.39,
	}
	fmt.Println(m["Bell Labs"])
}
