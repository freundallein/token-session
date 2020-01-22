# token-session
[![Build Status](https://travis-ci.org/freundallein/token-session.svg?branch=master)](https://travis-ci.org/freundallein/token-session)
[![Coverage Status](https://coveralls.io/repos/github/freundallein/token-session/badge.svg?branch=master)](https://coveralls.io/github/freundallein/token-session?branch=master)

Client-side session lib.

Allows to encrypt session data in plain token   
`5tMF4QqEUiWRsIIjrlDmefSa+Nis58xfbsWwveNHSpba+h8SzZpC/qBa9Em8S7LKlQee316eGwyEMTzngFBtb99X7ikiRS+wiPa5DzxpOEZK`  
You can set it as cookie, or send to client side in other way.

Expiration time is 30 seconds by default.
If you don't call `session.Token()` and don't renew it on client side, session will expire.

## Installation
```
$> go get gihub.com/freundallein/token-session
```


## Usage

```
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
```
