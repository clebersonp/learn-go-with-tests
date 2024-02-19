package v3

type Dictionary map[string]string

func (d Dictionary) Search(key string) (value string, err error) {
	return d[key], nil
}
