package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "This is a test"}
	t.Run("Known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "This is a test"
		assertStrings(t, got, want)
	})
	t.Run("unkown word", func(t *testing.T) {
		_, err := dictionary.Search("abcd")
		assertError(t, err, ErrnotFound)
	})
}
func TestAdd(t *testing.T) {
	t.Run("New Word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "This is a test"
		err := dictionary.Add(word, definition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "This is a test"}
		word := "test"
		definition := "This is a test"
		err := dictionary.Add(word, "new test")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})

}
func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %s want %s", got, want)
	}
}
func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, _ := dictionary.Search(word)
	if got != definition {
		t.Errorf("got %s want %s", got, definition)
	}
}
