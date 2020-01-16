package utils

import (
	"bytes"
	"fmt"
	"strings"
)

func Bytes(data map[string]string) []byte {
	buffer := new(bytes.Buffer)
	for key, value := range data {
		fmt.Fprintf(buffer, "%s=%s&", key, value)
	}
	return buffer.Bytes()
}

func ExtractBytes(sequence []byte) map[string]string {
	extracted := map[string]string{}
	pairs := strings.Split(string(sequence), "&")
	for _, keyval := range pairs {
		pair := strings.Split(keyval, "=")
		if len(pair) != 2 {
			continue
		}
		extracted[pair[0]] = pair[1]
	}
	return extracted
}
