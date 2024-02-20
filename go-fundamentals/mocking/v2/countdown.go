package v2

import (
	"fmt"
	"io"
)

func Countdown(out io.Writer) {
	fmt.Fprintf(out, "3")
}
