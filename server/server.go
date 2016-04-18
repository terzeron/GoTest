package main

import "github.com/go-martini/martini"

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello, world!"	  
	})
	m.Get("/error", func() (int, string) {
		return 301, "I'm a kettle"
	})
	m.Run()
}
