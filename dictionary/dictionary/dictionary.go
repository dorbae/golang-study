package dictionary

import "errors"

// #2.4 Ditionary part One
// https://nomadcoders.co/go-for-beginners/lectures/1516

// Dictionary type
// Structure X, Alias O
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// #2.5 Add method
// https://nomadcoders.co/go-for-beginners/lectures/1517

var errWordExists = errors.New("Word already exists")

// Add word and definition
func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)
	// if err == errNotFound {
	// 	d[word] = definition
	// } else if err == nil {
	// 	return errWordExists
	// }
	// return nil
	switch err {
	case errNotFound:
		d[word] = definition
	case nil:
		return errWordExists
	}
	return nil
}

// #2.6 Update Delete
// https://nomadcoders.co/go-for-beginners/lectures/1518

var errCantUpdate = errors.New("Can't update nont-existing word")

// Update the word
func (d Dictionary) Update(word, deifinition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = deifinition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete the word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
