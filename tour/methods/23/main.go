package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (read *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = read.r.Read(p)

	for i, value := range p {
		switch {
		case 'A' <= value && value <= 'M':
			value += 13
		case 'M' < value && value <= 'Z':
			value -= 13
		case 'a' <= value && value <= 'm':
			value += 13
		case 'm' < value && value <= 'z':
			value -= 13
		}
		p[i] = value
	}

	return
}
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
