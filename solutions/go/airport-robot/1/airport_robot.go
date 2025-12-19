package airportrobot

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

type Italian struct {}

func (Italian) LanguageName() string {
	return "Italian"
}

func (Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

type Portuguese struct {}

func (Portuguese) LanguageName() string {
    return "Portuguese"
}

func (Portuguese) Greet(name string) string{
    return fmt.Sprintf("Olá %s!", name)
}

func SayHello(name string, greeter Greeter) string {
	return fmt.Sprintf(
		"I can speak %s: %s",
		greeter.LanguageName(),
		greeter.Greet(name),
	)
}