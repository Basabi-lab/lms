package main

import (
	"github.com/Basabi-lab/lms/loader"
)

func main() {
	defPath := "./tests/song"
	host := "http://localhost:8080"

	scanner := loader.NewScanner(defPath)
	accessor := loader.NewAccessor(host)

	l := loader.NewLoader(scanner, accessor)
	l.Load()
}
