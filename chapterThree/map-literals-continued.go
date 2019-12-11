package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// map'in value'suna verilen tip eğer tip adıysa onu literalden çıkartabiliriz.
var m = map[string]Vertex{
	"Bell Labs":/* yani buradan*/ {40.68433, -74.39967},
	"Google": {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
