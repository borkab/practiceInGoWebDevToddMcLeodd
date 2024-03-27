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
