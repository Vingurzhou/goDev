package main

import (
	"errors"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func Read([]byte) (int, error) {

	return 0, errors.New("")
}
func main() {
	reader.Validate(MyReader{})
}
