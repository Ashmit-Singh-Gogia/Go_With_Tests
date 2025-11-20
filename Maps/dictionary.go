package main

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	value, exists := d[key] // exists is a boolean which tells whether the value is present or not
	if !exists {
		return "", errors.New("could not find the word you were looking for")
	}
	return value, nil
}
