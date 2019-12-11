package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	// koşulsuz switch
	switch {
	case t.Hour() < 12:
		fmt.Println("günaydın")
	case t.Hour() < 22:
		fmt.Println("iyi akşamlar")
	default:
		fmt.Println("iyi geceler")
	}
}
