package maps

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

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

func (d Dictionary) Add(key, value string) error {

	_, err := d.Search(key)

	// We can loop the switch on the same type -> Error

	switch err {
	case ErrNotFound:
		d[key] = value
		return nil
	case nil:
		return ErrWordExists
	default:
		return err
	}
}

func (d Dictionary) Update(key, value string) error {

	_, err := d.Search(key)

	if err != nil {
		return ErrWordDoesNotExist
	}

	d[key] = value

	return nil

}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
