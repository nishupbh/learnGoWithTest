package main

import "errors"

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

type Dictionary map[string]string

// Search will return output value of dictionary from key
func (d Dictionary) Search(key string) (string, error) {

	definition, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

//Add method for adding key value pair in Dictionary
func (d Dictionary) Add(key string, text string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = text
	case nil:
		return ErrWordExists
	default:
		return err
	}
}
