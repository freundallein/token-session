package utils

import (
	"testing"
)

func TestUtils(t *testing.T) {
	expected := map[string]string{
		"Judge Dredd": "Ma Ma is not the law...I am the law!",
	}
	observed := ExtractBytes(Bytes(expected))
	_, ok := observed["Judge Dredd"]
	if !ok {
		t.Error("test key lost")
	}
	if observed["Judge Dredd"] != expected["Judge Dredd"] {
		t.Error("Expected", expected, "got", observed)
	}
}
