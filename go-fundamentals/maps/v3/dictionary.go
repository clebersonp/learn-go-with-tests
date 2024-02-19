package v3

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(key string) (value string, err error) {
	// map can return 2 values. the first is the value and the second is a boolean which indicates if the key
	// was found successfully
	value, ok := d[key]
	if !ok {
		return value, errors.New("could not find the key you were looking for")
	}

	return d[key], nil
}
