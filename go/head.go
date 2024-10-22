package main

import (
	"fmt"
	"os"
)

func fileHead(fileName string, size int) ([]byte, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	buffeer := make([]byte, size)
	n, err := file.Read(buffeer)

	if err != nil {
		return nil, err
	}

	if n != size {
		return nil, fmt.Errorf("%q too small", fileName)
	}

	return buffeer, nil
}
