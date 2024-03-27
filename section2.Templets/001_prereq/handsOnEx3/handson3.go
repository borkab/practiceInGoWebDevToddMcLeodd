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
	Communicate() (string, bool)
}

func (p Person) Communicate() (string, bool) {
	return "I am " + p.fName + " " + p.lName, false
}

func (sa SecretAgent) Communicate() (string, bool) {
	if sa.licenceToKill {
		return "I am " + sa.fName + " " + sa.lName, true
	} else {
		return "I am " + sa.fName + " " + sa.lName, false
	}
}

func Vomit(h Human) string {

	switch v := h.(type) {
	case Person:
		return "I am " + v.fName + ", just a person"
	case SecretAgent:
		return "I am " + v.fName + ", agent"
	default:
		return "I am not a human"
	}

}
