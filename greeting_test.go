package greeting

import "testing"

func TestGreeting(t *testing.T) {
	given := "Pallat"
	want := "Hello, Pallat"

	result := greeting(given)
	if result != want {
		t.Errorf("greeting(%q) = %q, want %q", given, result, want)
	}
}

func TestGreetingWithAge(t *testing.T) {
	givenName := "Pallat"
	givenAge := 30
	want := "Hello, Pallat. You are 30 years old."

	result := greetingWithAge(givenName, givenAge)
	if result != want {
		t.Errorf("greetingWithAge(%q, %d) = %q, want %q", givenName, givenAge, result, want)
	}
}

func TestGreetingWithAgeAndDrink(t *testing.T) {
	givenName := "Pallat"
	givenAge := 30
	givenDrink := "Cola"
	want := "Hello, Pallat. You are 30 years old and your favorite drink is Cola."

	result := greetingWithAgeAndDrink(givenName, givenAge, givenDrink)
	if result != want {
		t.Errorf("greetingWithAgeAndDrink(%q, %d, %q) = %q, want %q", givenName, givenAge, givenDrink, result, want)
	}
}
