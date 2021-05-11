package main

import "fmt"

type rectange struct {
	width, height int
}

func (r *rectange) area() int {
	return r.width * r.height
}

func (r rectange) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rectange{width: 10, height: 5}
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
