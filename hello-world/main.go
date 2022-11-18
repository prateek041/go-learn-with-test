package main

import "fmt"

const englishHelloPrefix = "Hello, "

// seperating domain from the side-effect
func Hello (name string) string{
	if(name == ""){
		return "Hello world"
	}
	return englishHelloPrefix + name
}

func main(){
	fmt.Println(Hello("prateek"))
}