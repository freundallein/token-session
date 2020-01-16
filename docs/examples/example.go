package main

import (
	"fmt"

	"github.com/freundallein/token-session/sessions"
)

func main() {
	data := map[string]string{
		"username": "Test",
		"role":     "user",
	}
	sessions.Init("SECRETKEY")
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
