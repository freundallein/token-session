package sessions

import (
	"errors"
	"fmt"
	"github.com/freundallein/token-session/sessions/base64"
	"github.com/freundallein/token-session/sessions/crypt"
	"github.com/freundallein/token-session/sessions/utils"
	"time"
)

const (
	timeFormat = "02.01.2006 15:04:05"
)

var (
	expirationTime     = 30 * time.Second
	ErrSessionExpired  = errors.New("session expired")
	ErrInvalidToken    = errors.New("invalid token")
	ErrEncryptionError = errors.New("error during encryption")
)

func Init(key string, timeout time.Duration) {
	crypt.SecretKey = key
	expirationTime = timeout
}

func Create(data map[string]string) *Session {
	lastSeen := time.Now().UTC()
	data["lastSeen"] = lastSeen.Format(timeFormat)
	return &Session{data: data, lastSeen: lastSeen}
}

func Get(token string) (*Session, error) {
	session, err := FromToken(token)
	if err != nil {
		return nil, ErrInvalidToken
	}
	if ok := session.validate(); !ok {
		return nil, ErrSessionExpired
	}
	return session, nil
}

type Session struct {
	data     map[string]string
	lastSeen time.Time
}

func FromToken(token string) (*Session, error) {
	data, err := decrypt(token)
	if err != nil {
		return nil, err
	}
	lastSeenVal, ok := data["lastSeen"]
	if !ok {
		return nil, err
	}
	lastSeen, err := time.Parse(timeFormat, lastSeenVal)
	if err != nil {
		return nil, err
	}
	return &Session{data: data, lastSeen: lastSeen}, nil
}

func (s *Session) Data() map[string]string {
	return s.data
}

func (s *Session) String() string {
	return fmt.Sprintf("%v last: %v", s.data, s.lastSeen.Format(timeFormat))
}

func (s *Session) Token() (string, error) {
	s.data["lastSeen"] = time.Now().UTC().Format(timeFormat)
	encData, err := encrypt(s.data)
	if err != nil {
		return "", ErrEncryptionError
	}
	return encData, nil
}

func (s *Session) validate() bool {
	return time.Since(s.lastSeen) <= expirationTime
}

func encrypt(data map[string]string) (string, error) {
	encData, err := crypt.Encrypt(utils.Bytes(data))
	if err != nil {
		return "", err
	}
	return base64.Encode(encData), nil
}

func decrypt(data string) (map[string]string, error) {
	decodedBytes, err := base64.Decode(data)
	if err != nil {
		return nil, err
	}
	decData, err := crypt.Decrypt(decodedBytes)
	if err != nil {
		return nil, err
	}
	return utils.ExtractBytes(decData), nil
}
