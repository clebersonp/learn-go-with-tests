package maps

const (
	ErrNotFound          = DictionaryErr("could not find the word you were looking for")
	ErrWordExists        = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExists = DictionaryErr("cannot update word because it does not exist")
)

// DictionaryErr my own error type which implements the error interface
type DictionaryErr string

// implementing error interface
func (e DictionaryErr) Error() string {
	return string(e)
}

// Dictionary type wrapper around a map. It's an underlying of map[string]string
// With the custom type we can create the Search method
// Avoid create map with syntax: var m map[string]string -> m will be == nil
// Use this instead: var m = map[string]string{} or var m = make(map[string]string) -> m will be empty map
type Dictionary map[string]string

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil:
		d[word] = newDefinition
	default:
		return err
	}
	
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
