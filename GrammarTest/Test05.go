package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Now you have %g problems.",
		math.Nextafter(2, 3))
}
