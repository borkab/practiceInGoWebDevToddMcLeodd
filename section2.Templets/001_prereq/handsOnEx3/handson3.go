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

func Vomit(h Human) (string, bool) {
	return h.Communicate()
}
