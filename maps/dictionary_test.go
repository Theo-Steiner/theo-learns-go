package dictionary

import (
	"testing"
)

const (
	word       = "whale"
	definition = "big mammal living in the ocean"
)

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
	retrieved, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, retrieved, definition)
}

func assertWordNotInDictionary(t *testing.T, dictionary Dictionary, word string) {
	t.Helper()
	_, err := dictionary.Search(word)
	if err == nil {
		t.Fatal("should not find non-existing word:", err)
	}
	assertError(t, err, ErrNotFound)
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{word: definition}

	t.Run("known word", func(t *testing.T) {
		got, err := dictionary.Search(word)
		want := definition

		assertError(t, err, nil)

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		val, err := dictionary.Search("unknown")

		assertError(t, err, ErrNotFound)

		assertStrings(t, val, "")
	})
}

func TestAdd(t *testing.T) {
	t.Run("add new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add(word, definition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("try add existing word", func(t *testing.T) {
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "overriding word")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing", func(t *testing.T) {
		dictionary := Dictionary{word: "old definition"}
		err := dictionary.Update(word, definition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("update non-existing", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)
		assertWordNotInDictionary(t, dictionary, word)
	})
}

func TestDelete(t *testing.T) {
	t.Run("update existing", func(t *testing.T) {
		dictionary := Dictionary{word: definition}
		dictionary.Delete(word)
		assertWordNotInDictionary(t, dictionary, word)
	})
}
