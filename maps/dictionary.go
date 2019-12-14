package maps

import "errors"

// Dictionary ...
type Dictionary map[string]string

// ErrnotFound ...
var (
	ErrnotFound         = errors.New("could not find the word you are looking for")
	ErrWordExists       = errors.New("Cannot add word because it already exists")
	ErrWordDoesNotExist = errors.New("cannot update word because it does not exist")
)

// Search ...
func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", ErrnotFound
	}
	return definition, nil
}

// Add ...
func (d Dictionary) Add(key, definition string) error {
	_, err := d.Search(key)
	switch err {
	case ErrnotFound:
		d[key] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

// Update ...
func (d Dictionary) Update(key, definition string) error {
	_, err := d.Search(key)
	switch err {
	case ErrnotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = definition
	default:
		return err
	}
	return nil
}
