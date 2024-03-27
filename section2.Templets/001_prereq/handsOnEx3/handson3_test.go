package main

import "testing"

func TestCommunicate(t *testing.T) {
	t.Run("person", func(t *testing.T) {
		p := Person{fName: "Ida", lName: "Krohnenberg"}

		got, _ := p.Communicate()
		want := "I am Ida Krohnenberg"

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("agent with a licence", func(t *testing.T) {
		ag := SecretAgent{Person: Person{fName: "Mortimer", lName: "Morrison"}, licenceToKill: true}

		got, killer := ag.Communicate()
		want, _ := "I am Mortimer Morrison", true

		if got != want || !killer { //ha a method altal adott nev nem egyezik meg a vart nevvel, vagy nincs engedelye, akkor hiba
			t.Errorf("got %v want %v killer %v", got, want, killer)
		}
	})
}

func TestVomit(t *testing.T) {

	t.Run("person", func(t *testing.T) {
		pIda := Person{fName: "Ida", lName: "Krohnenberg"}

		got := Vomit(pIda) //cannot use pIda (variable of type Person) as Human value in argument to Vomit: Person does not implement Human (wrong type for method pSpeak)
		//have pSpeak() string
		//want pSpeak()compiler
		want := "I am Ida, just a person"

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("agent", func(t *testing.T) {
		pMor := SecretAgent{Person: Person{fName: "Mortimer", lName: "Morrison"}, licenceToKill: true}

		got := Vomit(pMor) //cannot use pMor (variable of type SecretAgent) as Human value in argument to Vomit: SecretAgent does not implement Human (wrong type for method pSpeak)
		//have pSpeak() string
		//want pSpeak()compiler
		want := "I am Mortimer, agent"

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}

	})
}
