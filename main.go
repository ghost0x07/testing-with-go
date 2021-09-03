package main

import (
	"hello/di"
	"os"
)

func main() {
	di.Greet(os.Stdout, "Faiz")
}
