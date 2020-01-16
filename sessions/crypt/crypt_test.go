package crypt

import (
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	expected := "Ma Ma is not the law...I am the law!"
	encData, err := Encrypt([]byte(expected))
	if err != nil {
		t.Error(err.Error())
	}
	observed, err := Decrypt(encData)
	if err != nil {
		t.Error(err.Error())
	}
	if string(observed) != expected {
		t.Error("Expected", expected, "got", observed)
	}

}
func TestCreateHash(t *testing.T) {
	expected := "62082b6d0e06f4307d3489d32556700e"
	observed := createHash("test-secret-key")
	if observed != expected {
		t.Error("Expected", expected, "got", observed)
	}
}
