package maps

const ErrWordNotFound = DictionaryErr("could not find the word specified")
const ErrWordExists = DictionaryErr("word already exists in the dictionary")

type DictionaryErr string

func (err DictionaryErr) Error() string {
	return string(err)
}

type Dictionary map[string]string

func (dict Dictionary) Search(key string) (string, error) {
	if word, ok := dict[key]; ok {
		return word, nil
	}
	return "", ErrWordNotFound
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		return ErrWordNotFound
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
