package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/pownthep/electioner/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleHH"))
	fmt.Println(cmp.Diff("Hello World", "Hello GoGGGGGGGGGG"))
	array := [...]float64{7.0, 8.5, 9.1, 4}
	x := sum(&array) // Note the explicit address-of operator
	fmt.Printf("%f", x)
}

func sum(a *[4]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}
