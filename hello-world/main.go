package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloprefix = "Bonjour, "

// seperating domain from the side-effect
func Hello (name string, language string) string{
	if(name == ""){
		name = "world"
	}

	prefix:= englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloprefix
	case spanish:
		prefix = spanishHelloPrefix
	}

	return prefix + name
}

func main(){
	fmt.Println(Hello("prateek", "English"))
}