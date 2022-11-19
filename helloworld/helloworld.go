package helloworld

const spanish = "Spanish"
const french = "French"
const hindi = "Hindi"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloprefix = "Bonjour, "
const hindiHelloprefix = "Namaste, "

// seperating domain from the side-effect
func Hello (name string, language string) string{

	// seperated the default cases
	if(name == ""){
		name = "world"
	}

	// checking the languages
	return greetingMessageBuilder(language) + name

}

func greetingMessageBuilder(language string)(prefix string){
	switch language {
	case french:
		prefix = frenchHelloprefix
	case spanish:
		prefix = spanishHelloPrefix
	case hindi:
		prefix = hindiHelloprefix
	default :
		prefix = englishHelloPrefix
	}
	return
}
