package main

import (
	"fmt"
	"github.com/learn-with-test/prateek/testingImport"
	"github.com/learn-with-test/prateek/helloworld"
)

func main(){
	fmt.Println(testingImport.ImportTest())
	fmt.Println(helloworld.Hello("prateek", ""))
}