package main

import "fmt"

const spanish = "Spanish"
const French = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloprefix = "Bonjour, "

// seperating domain from the side-effect
func Hello (name string, language string) string{
	if(name == ""){
		return "Hello, world"
	}

	if language == spanish{
		return spanishHelloPrefix + name
	}

	if language == "French"{
		return frenchHelloprefix + name
	}
	return englishHelloPrefix + name
}

func main(){
	fmt.Println(Hello("prateek", "English"))
}