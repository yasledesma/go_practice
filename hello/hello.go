package main

const englishGreeting = "Hello, "
const frenchGreeting = "Bonjour, "
const spanishGreeting = "Hola, "
const russianGreeting = "Привет, "

const spanish = "Spanish"
const french = "French"
const russian = "Russian"

func Hello(name string, language string) string {
    if name == "" {
        name = "world"
    } 

    return getGreetingPrefix(language) + name
}

func getGreetingPrefix(language string) (prefix string) {
    switch language {
        case french:
            prefix = frenchGreeting
        case spanish:
            prefix = spanishGreeting
        case russian:
            prefix = russianGreeting
        default:
            prefix = englishGreeting
    }
    return
}
