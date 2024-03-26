package main

type Person struct {
	fName string
	lName string
	SecretAgent
}

type SecretAgent struct {
	liceneToKill bool
}

func (p Person) Speak {}