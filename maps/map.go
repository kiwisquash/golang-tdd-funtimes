package main

const (
	// ErrNotFound - couldn't look up the word because it's not there
	ErrNotFound = DictionaryErr("word not found")

	// ErrWordExists - trying to add a word that's already in the dictionary 
	ErrWordExists = DictionaryErr("word already exists")

	// ErrWordDoesNotExist - 
	ErrWordDoesNotExist = DictionaryErr("cannot update non-existent word")
)


type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

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
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, val string) error {
	
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = val
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}