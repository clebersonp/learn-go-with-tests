package v2

type Dictionary map[string]string

func (d Dictionary) Search(key string) (value string) {
	return d[key]
}
