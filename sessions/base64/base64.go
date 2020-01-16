package base64

import (
	"encoding/base64"
)

func Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Decode(data string) ([]byte, error) {
	buff, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return buff, nil
}
