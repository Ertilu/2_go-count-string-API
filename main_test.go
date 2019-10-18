package main

import (
	c "countWebService/count"
	"testing"
)

func TestCount(t *testing.T) {
	// test is the input is nil
	t.Run("Nil input", func(t *testing.T) {
		input := ""
		consonants, vowels := c.Count(input)
		expectedConsonant := "Huruf mati: 0,"
		expectedVowels := "Huruf hidup: 0"
		if consonants != expectedConsonant && vowels != expectedVowels {
			t.Fatalf("return must be %s %s instead of Huruf mati: %s Huruf hidup: %s", expectedConsonant, expectedVowels, consonants, vowels)
		}
	})

	// test is the input is no vowels
	t.Run("No vowels input", func(t *testing.T) {
		input := "thsnptsnvwls"
		consonants, vowels := c.Count(input)
		expectedConsonant := "Huruf mati: 12,"
		expectedVowels := "Huruf hidup: 0"
		if consonants != expectedConsonant && vowels != expectedVowels {
			t.Fatalf("return must be %s %s instead of Huruf mati: %s Huruf hidup: %s", expectedConsonant, expectedVowels, consonants, vowels)
		}
	})

	// test is the input is no vowels
	t.Run("No vowels input", func(t *testing.T) {
		input := "thsnptsnvwls"
		consonants, vowels := c.Count(input)
		expectedConsonant := "Huruf mati: 12,"
		expectedVowels := "Huruf hidup: 0"
		if consonants != expectedConsonant && vowels != expectedVowels {
			t.Fatalf("return must be %s %s instead of Huruf mati: %s Huruf hidup: %s", expectedConsonant, expectedVowels, consonants, vowels)
		}
	})

	// test is the input is no consonants
	t.Run("No vowels input", func(t *testing.T) {
		input := "aiueo"
		consonants, vowels := c.Count(input)
		expectedConsonant := "Huruf mati: 0,"
		expectedVowels := "Huruf hidup: 5"
		if consonants != expectedConsonant && vowels != expectedVowels {
			t.Fatalf("return must be %s %s instead of Huruf mati: %s Huruf hidup: %s", expectedConsonant, expectedVowels, consonants, vowels)
		}
	})

	main()
}
