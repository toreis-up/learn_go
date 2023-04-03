package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader *rot13Reader) Read(b []byte) (int, error) {
	n, err := reader.r.Read(b)
	if err == nil {
		for i := range b {
			switch true {
			case b[i] >= 97 && b[i] <= 109:
				b[i] += 13
			case b[i] > 109 && b[i] <= 122:
				b[i] -= 13
			case b[i] >= 65 && b[i] <= 77:
				b[i] += 13
			case b[i] > 77 && b[i] <= 90:
				b[i] -= 13
			}
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
