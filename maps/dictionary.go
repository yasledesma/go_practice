package dictionary

type Dictionary map[string]string
type DictionaryError string

const (
    ErrorWordNotFound = DictionaryError("word not found.")
    ErrorWordAlreadyExists = DictionaryError("a definition for this word already exists and can't be modified.")
    ErrorCanNotUpdateWord = DictionaryError("definition can't be updated because the word doesn't exist.")
) 

func (e DictionaryError) Error() string {
    return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
    definition, found := d[word]

    if !found {
        return "", ErrorWordNotFound 
    }

    return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
    _, err := d.Search(word)

    switch err {
    case ErrorWordNotFound:
        d[word] = definition
    case nil:
        return ErrorWordAlreadyExists
    default:
        return err
    }

    return nil
}

func (d Dictionary) Update(word, definition string) error {
    _, err := d.Search(word)

    switch err {
    case ErrorWordNotFound:
        return ErrorCanNotUpdateWord
    case nil:
        d[word] = definition
    default:
        return err
    }

    return nil
}
