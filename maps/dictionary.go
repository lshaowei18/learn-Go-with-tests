package maps

type Dictionary map[string]string

const (
	ErrNotFound                = DictionaryErr("Could not find the word you are looking for.")
	WordExistsError            = DictionaryErr("Key Value pair already exists.")
	WordDoesNotExistsError     = DictionaryErr("Update failed cause could not find the word you are looking for.")
	DeleteWordDoesNotExistsErr = DictionaryErr("Delete failed cause could not find the word.")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	val, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return val, nil
}

func (d Dictionary) Add(key, val string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = val
	case nil:
		return WordExistsError
	}
	return nil
}

func (d Dictionary) Update(key, newVal string) error {
	_, err := d.Search(key)
	if err != nil {
		return WordDoesNotExistsError
	}
	d[key] = newVal
	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	if err != nil {
		return DeleteWordDoesNotExistsErr
	}
	delete(d, key)
	return nil
}
