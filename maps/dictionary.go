package maps

import "errors"

var ErrNotFound = errors.New("could not find the word you were looking for")

func Search(dictionary map[string]string, key string) string {
	return dictionary[key]
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {

	//  It can return 2 values. The second value is a boolean which indicates if the key was found successfully.
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}
