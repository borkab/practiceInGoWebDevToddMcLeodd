package main

import "testing"

func TestPSpeak(t *testing.T) {
	t.Run("attach a person", func(t *testing.T) {
		pIda := Person{fName: "Ida",
			lName: "Krohnenberg"}

		got := pIda.pSpeak()
		want := "I am Ida Krohnenberg"

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("attach a secret agent", func(t *testing.T) {
		pMor := SecretAgent{Person: Person{fName: "Mortimer", lName: "Morrison"}, licenceToKill: true}

		got := pMor.pSpeak()
		want := "I am Mortimer Morrison"

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSaSpeak(t *testing.T) {

	t.Run("has a licence to kill", func(t *testing.T) {
		pMor := SecretAgent{Person: Person{fName: "Mortimer", lName: "Morrison"}, licenceToKill: true}

		got := pMor.saSpeak()
		want := "My name is Mortimer Morrison and I am a secret agent"

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}

	})

	t.Run("does not have a licence to kill", func(t *testing.T) {
		pWil := SecretAgent{Person: Person{fName: "Willi", lName: "Wondraschek"}, licenceToKill: false}

		got := pWil.saSpeak()
		want := "My name is Willi Wondraschek and I am NOT a secret agent"

		if got != want {
			t.Errorf("got %v want %v", got, want)
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
		want := "I am Mortimer Morrison, agent"

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}

	})
}
