package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (reader MyReader) Read(b []byte) (int, error) {
	for i := range b { // bのキャパに合わせる
		b[i] = 'A'
	}
	return len(b), nil // 入ってるサイズを返す
}

func main() {
	reader.Validate(MyReader{})
}
