package v4

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the key you were looking for")

func (d Dictionary) Search(key string) (value string, err error) {
	// map can return 2 values. the first is the value and the second is a boolean which indicates if the key
	// was found successfully
	value, ok := d[key]
	if !ok {
		return value, ErrNotFound
	}

	return d[key], nil
}

// Add - We can modify a map without passing as an address pointer (&dictionary)
// So when you pass a map to a function/method, you are indeed copying it, but just the pointer part
// Map can be nil value. A nil map behaves like an empty map when READING.
// But attempts to write a nil map will cause a runtime panic
// to void panic with nil map, you should never initialize an empty map variable like this: var m map[string]string
// Instead, you can initialize an empty map like this: var dictionary = map[string]string{}
// OR
// var dictionary = make(map[string]string)
// make is a built-in function in Go to create things
// Both approaches create an empty hash map and point dictionary at it. Which ensures that you will never get a runtime
// panic
func (d Dictionary) Add(key, value string) {
	d[key] = value
}
