package v1

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, name string) {
	fmt.Printf("Hello, %s\n", name)
}
