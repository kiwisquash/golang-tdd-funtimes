package main

type Dictionary map[string]string

const (
	ErrNotFound = DictionaryErr("key not found")
	ErrWordAlreadyExists = DictionaryErr("word already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update non-existent word")
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
		return ErrWordAlreadyExists
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