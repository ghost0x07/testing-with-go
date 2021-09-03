package maps

import "testing"

func TestDictionary(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "This is just a test"}
		got, _ := dictionary.Search("test")
		want := "This is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "This is just a test"}
		_, err := dictionary.Search("test1")

		if err == nil {
			t.Fatal("Expected to get an error.")
		}

		want := ErrWordNotFound.Error()

		assertStrings(t, err.Error(), want)

	})
}

func TestAdd(t *testing.T) {

	t.Run("New Word", func(t *testing.T) {
		dict := Dictionary{}
		meaning := "A part of a Sentence"
		dict.Add("word", meaning)

		assertDefinition(t, dict, "word", meaning)
	})

	t.Run("Word Already Exists", func(t *testing.T) {
		def := "Just a test"
		dict := Dictionary{"test": def}

		err := dict.Add("test", "New meaning")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, "test", def)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("Word Exists", func(t *testing.T) {
		dict := Dictionary{"test": "Old Meaning"}

		err := dict.Update("test", "New Meaning")

		assertNoError(t, err)
		assertDefinition(t, dict, "test", "New Meaning")
	})
	t.Run("Word Doesn't Exist", func(t *testing.T) {
		dict := Dictionary{}

		err := dict.Update("test", "New Meaning")

		assertError(t, err, ErrWordNotFound)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrWordNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("Encountered an error %v", err)
	}
}

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

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}
