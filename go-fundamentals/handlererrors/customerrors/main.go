package main

import (
	"errors"
	"fmt"
	"log"
)

type norgateMathError struct {
	lat  string
	long string
	err  error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Printf("%T\n", err)
		log.Printf("error is of type norgateMathError? %v", errors.As(err, &norgateMathError{}))
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
		return 0, norgateMathError{
			lat:  "50.2289 N",
			long: "99.4656 W",
			err:  nme,
		}
	}
	return 42, nil
}
