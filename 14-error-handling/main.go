package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	text, err := readFile("test.txtt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(text)
	}
}

func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("unable to open %s", filename)
	}

	var buf [1024]byte
	for {
		_, err := file.Read(buf[:])
		if err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}
	}

	return string(buf[:]), nil
}
