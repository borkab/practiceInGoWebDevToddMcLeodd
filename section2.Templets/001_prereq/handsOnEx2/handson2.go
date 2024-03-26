package main

type Person struct {
	fName string
	lName string
}

type SecretAgent struct {
	Person
	liceneToKill bool
}

func (p Person) Speak() string {
	return "I am " +p.fName + " " + p.lName 
}

func