package main

import (
	"bytes"
	"fmt"
	"os"
)

func FileType(FileName string) (string, error) {
	file, err := os.Open(FileName)

	if err != nil {
		return "", err
	}

	defer file.Close()

	data := make([]byte, 64)

	if _, err := file.Read(data); err != nil {
		return "", err
	}

	if bytes.Equal(data[:4], []byte{0x00, 0x00, 0x01, 0x00}) {
		return "ico", nil
	}

	if bytes.Equal(data[:6], []byte{0x47, 0x49, 0x46, 0x38, 0x37, 0x61}) {
		return "gif", nil
	}

	if bytes.Equal(data[:4], []byte{0xFF, 0xD8, 0xFF, 0xF0}) {
		return "jpg", nil
	}

	if bytes.Equal(data[:8], []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}) {
		return "png", nil
	}

	return "", fmt.Errorf("%q - unknown file type", FileName)
}
