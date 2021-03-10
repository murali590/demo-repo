package main

import (
	"fmt"
)

func main() {
	var car vech.vechile = vech.vechile{modelname: "fx10", price: 1250000.50}

	fmt.Printf("%v\n", car)
}
