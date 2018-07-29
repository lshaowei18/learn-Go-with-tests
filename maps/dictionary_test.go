package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "This is a test"}

	t.Run("Existing Word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		assertStrings(t, got, "This is a test")
	})

	t.Run("Nonexisting Word", func(t *testing.T) {
		_, got := dictionary.Search("wrong")
		assertErrors(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("Add Word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("add", "This is to add")
		assertErrors(t, err, nil)
		assertDefinition(t, dictionary, "add", "This is to add")
	})

	t.Run("Add existing word", func(t *testing.T) {
		dictionary := Dictionary{"add": "Added"}
		err := dictionary.Add("add", "Added")
		assertErrors(t, err, WordExistsError)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update existing word", func(t *testing.T) {
		word := "add"
		definition := "Added"
		dictionary := Dictionary{word: definition}
		newDefinition := "Updated"
		dictionary.Update(word, newDefinition)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("Update nonexisting word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update("update", "Updated")
		assertErrors(t, err, WordDoesNotExistsError)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete existing word", func(t *testing.T) {
		word := "add"
		dictionary := Dictionary{"add": "To be deleted"}
		dictionary.Delete(word)
		_, err := dictionary.Search(word)
		if err != ErrNotFound {
			t.Error("Expected error; given none")
		}
	})

	t.Run("Delete nonexisting word", func(t *testing.T) {
		word := "add"
		dictionary := Dictionary{}
		err := dictionary.Delete(word)
		if err == nil {
			t.Error("Expected error; given none")
		}
	})
}

func TestDictError(t *testing.T) {
	dictErr := DictionaryErr("An Error")
	got := dictErr.Error()
	want := "An Error"
	assertStrings(t, got, want)
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got : %s; want %s", got, want)
	}
}

func assertDefinition(t *testing.T, dict Dictionary, word, definition string) {
	t.Helper()
	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("Should find added word.", err)
	}
	if definition != got {
		t.Errorf("got '%s'; want '%s'", got, definition)
	}
}

func assertErrors(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error '%s'; want '%s'", got, want)
	}
}
