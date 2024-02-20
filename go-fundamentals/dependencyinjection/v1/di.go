package v1

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, name string) (bytes int, err error) {
	return fmt.Fprintf(writer, "Hello, %s", name)
}
