package main

import (
	"fmt"
	"time"

	"github.com/freundallein/token-session/sessions"
)

func main() {
	expiration := 60 * time.Second
	secretKey := "my-secret-key"
	sessions.Init(secretKey, expiration)
	data := map[string]string{
		"username": "Test",
		"role":     "user",
	}
	session := sessions.Create(data)
	fmt.Println(session)
	fmt.Println(session.Data())
	token, err := session.Token()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(token)
	session, err = sessions.Get(token)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(session)
}
