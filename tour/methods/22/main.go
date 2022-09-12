package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (m MyReader) Read(p []byte) (n int, err error) {
	for i := range p {
		p[i] = 'A'
	}
	return len(p), nil
}
func main() {
	reader.Validate(MyReader{})
}
