package dictionary

import "errors"

type Dictionary map[string]string

var ErrorWordNotFound = errors.New("word not found.")

func (d Dictionary) Search(word string) (string, error) {
    definition, found := d[word]

    if !found {
        return "", ErrorWordNotFound 
    }

    return definition, nil
}

func (d Dictionary) Add(word, definition string) {
    d[word] = definition
}
