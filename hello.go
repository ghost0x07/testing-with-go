package main

const englishHelloPrefix = "Hello, "

func hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name + "!"
}
