package sessions

import (
	"testing"
	"time"

	"github.com/freundallein/token-session/sessions/crypt"
)

func TestInit(t *testing.T) {
	opts := &Options{Secret: "test", Expiration: 1 * time.Second}
	Init(opts)
	if crypt.SecretKey != opts.Secret {
		t.Error("Expected", opts.Secret, "got", crypt.SecretKey)
	}
	if expirationTime != opts.Expiration {
		t.Error("Expected", opts.Expiration, "got", expirationTime)
	}
}

func TestCreate(t *testing.T) {
	expected := map[string]string{
		"testKey": "value",
	}
	observed := Create(expected)
	_, ok := observed.data["lastSeen"]
	if !ok {
		t.Error("Expected lastSeen key")
	}
	val, ok := observed.data["testKey"]
	if !ok || val != "value" {
		t.Error("Expected testKey: value, got", val)
	}
}

func TestData(t *testing.T) {
	expected := map[string]string{
		"testKey": "value",
	}
	observed := Create(expected).Data()
	for key, value := range expected {
		val, ok := observed[key]
		if !ok {
			t.Error("Expected", key)
		}
		if val != value {
			t.Error("Expected", value, "got", val)
		}
	}
}
func TestToken(t *testing.T) {
	expected := map[string]string{
		"testKey": "value",
	}
	observed, err := Create(expected).Token()
	if err != nil {
		t.Error(err.Error())
	}
	if observed == "" {
		t.Error("Expected", observed, "got nothing")
	}
}

func TestGet(t *testing.T) {
	data := map[string]string{
		"testKey": "value",
	}
	expected := Create(data)
	token, err := expected.Token()
	if err != nil {
		t.Error(err.Error())
	}
	observed, err := Get(token)
	if err != nil {
		t.Error(err.Error())
	}
	for key, value := range data {
		val, ok := observed.Data()[key]
		if !ok {
			t.Error("Expected", key)
		}
		if val != value {
			t.Error("Expected", value, "got", val)
		}
	}

}
