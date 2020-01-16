package base64

import "testing"

func TestEncode(t *testing.T) {
	expected := "TWEgTWEgaXMgbm90IHRoZSBsYXcuLi5JIGFtIHRoZSBsYXch"
	observed := Encode([]byte("Ma Ma is not the law...I am the law!"))
	if observed != expected {
		t.Error("Expected", expected, "got", observed)
	}
}
func TestDecode(t *testing.T) {
	observed, err := Decode("TWEgTWEgaXMgbm90IHRoZSBsYXcuLi5JIGFtIHRoZSBsYXch")
	if err != nil {
		t.Error(err.Error())
	}
	expected := "Ma Ma is not the law...I am the law!"
	if string(observed) != expected {
		t.Error("Expected", expected, "got", observed)
	}
}

func TestDecodeErr(t *testing.T) {
	observed, err := Decode("!!!")
	if err == nil {
		t.Error("should be error")
	}
	if observed != nil {
		t.Error("should be nil")
	}
}
