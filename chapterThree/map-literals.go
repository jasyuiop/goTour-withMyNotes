package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// map literali'de struct literaline benzer fakat oluştururken key gereklidir.
var m = map[string]Vertex{
	"Bell Labs": Vertex{40.68, -74.39},
	"Google":    Vertex{37.42, -122.08},
}

func main() {
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
}
