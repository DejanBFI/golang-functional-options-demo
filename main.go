package main

import (
	"fmt"
	"playground/optstruct"
)

// default value:
// - string: ""
// - float64: 0.0
// - integers (int, uint): 0
// - bool: false
// - pointer: nil
// - slice: nil

func main() {
	simpleMail := optstruct.NewSimpleMail(
		optstruct.WithMessage("Hello world"),
		optstruct.WithSender("Dejan"),
	)
	fmt.Printf("%#v\n", simpleMail)
}
