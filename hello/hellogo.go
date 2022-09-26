package main

import (
	"fmt"

	"rsc.io/quote"
)

var pl = fmt.Println

func main() {
	pl(quote.Go())
}
