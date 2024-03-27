package main

type Person struct {
	fName string
	lName string
}

type SecretAgent struct {
	Person
	licenceToKill bool
}

type Human interface {
	pSpeak()
	saSpeak()
}

func (p Person) pSpeak() string {
	return "I am " + p.fName + " " + p.lName
}

func (sa SecretAgent) saSpeak() string {

	if sa.licenceToKill { //if sa.licenceToKill == true   -> you can simplify it
		return "My name is " + sa.fName + " " + sa.lName + " and I am a secret agent"
	} else {
		return "My name is " + sa.fName + " " + sa.lName + " and I am NOT a secret agent"
	}
}

func Vomit(h Human) string {

	switch v := h.(type) {
	case Person:
		return "I am " + v.fName + " just a person"
	case SecretAgent:
		return "I am " + v.fName + ", agent"
	default:
		return "I am not a human"
	}
	//impossible type switch case: Person
	//h (variable of type Human) cannot have dynamic type Person (wrong type for method pSpeak)

	//impossible type switch case: SecretAgent
	//h (variable of type Human) cannot have dynamic type SecretAgent (wrong type for method pSpeak)
}
