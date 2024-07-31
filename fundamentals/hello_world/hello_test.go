package main

import "testing"

func TestHelloWorld(t *testing.T) {
	t.Run(`Testing "HelloWorld"`, func(t *testing.T) {
		want := "Hello World!"
		got := HelloWorld()
		assertCorrectMessage(t, got, want)
	})
}

func TestHelloYou(t *testing.T) {
	t.Run(`testing "HelloYou" with user`, func(t *testing.T) {
		want := "Hello, Gabriel!"
		got := HelloYou("Gabriel")
		assertCorrectMessage(t, got, want)
	})
	t.Run(`testing "HelloYou" without user`, func(t *testing.T) {
		want := "Hello, World!"
		got := HelloYou("")
		assertCorrectMessage(t, got, want)
	})
}

func TestHelloYouLang(t *testing.T) {
	t.Run(`testing "HelloYouLang" with user and language "es"`, func(t *testing.T) {
		want := "Hola, Gabriel!"
		got := HelloYouLang("Gabriel", "es")
		assertCorrectMessage(t, got, want)
	})
	t.Run(`testing "HelloYouLang" with user and language "fr"`, func(t *testing.T) {
		want := "Bonjour, Gabriel!"
		got := HelloYouLang("Gabriel", "fr")
		assertCorrectMessage(t, got, want)
	})
	t.Run(`testing "HelloYouLang" with user and language "jp"`, func(t *testing.T) {
		want := "こんにちは, ガブリエル!"
		got := HelloYouLang("ガブリエル", "jp")
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
