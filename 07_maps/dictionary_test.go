package dictionary

import "testing"

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		expected := "this is just a test"

		err := dictionary.Add(key, expected)

		assertError(t, err, nil)
		assertSearch(t, dictionary, "test", expected)
	})

	t.Run("existing word", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, value)

		assertError(t, err, ErrWordExists)
		assertSearch(t, dictionary, key, value)
	})
}

func assertSearch(t testing.TB, dictionary Dictionary, key string, expected string) {
	got, err := dictionary.Search("test")

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, expected)
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("word exists", func(t *testing.T) {
		got, err := dictionary.Search("test")
		expected := "this is just a test"

		if err != nil {
			t.Fatal("expected to not get an error")
		}

		assertStrings(t, got, expected)
	})

	t.Run("word does not exist", func(t *testing.T) {
		_, err := dictionary.Search("doesnotexist")

		assertError(t, err, ErrNotFound)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		initialValue := "this is just a test"
		expected := "updated"

		dictionary := Dictionary{word: initialValue}

		dictionary.Update(word, expected)

		assertSearch(t, dictionary, word, expected)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		expected := "updated"

		dictionary := Dictionary{}

		err := dictionary.Update(word, expected)

		assertError(t, err, ErrWordNotExists)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	value := "some value"

	dictionary := Dictionary{word: value}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)

	assertError(t, err, ErrNotFound)
}

func assertStrings(t testing.TB, got, expected string) {
	t.Helper()

	if got != expected {
		t.Errorf("expected %q but got %q instead", expected, got)
	}
}

func assertError(t testing.TB, got, expected error) {
	t.Helper()

	if got != expected {
		t.Errorf("expected %q but got %q instead", expected, got)
	}
}
