package base64

import (
	"encoding/base64"
)

// Encode return base64 encoded string
func Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Decode decode string and return it bytes representation
func Decode(data string) ([]byte, error) {
	buff, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return buff, nil
}
