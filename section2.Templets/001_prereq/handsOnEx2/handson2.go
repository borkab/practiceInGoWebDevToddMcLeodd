package main

type Person struct {
	fName string
	lName string
}

type SecretAgent struct {
	Person
	licenceToKill bool
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
