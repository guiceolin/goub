package sessionmanager

import (
	"github.com/guiceolin/goub/jwt"
	"net/http"
	"time"
)

type Session struct {
	payload map[string]interface{}
}

func (s *Session) Get(name string) interface{} {
	return s.payload[name]
}

func (s *Session) Set(name string, value interface{}) {
	s.payload[name] = value
}

type SessionManager struct {
	domain         string
	jwtSecret      []byte
	name           string
	expirationTime time.Time
}

func (sm *SessionManager) SetSession(w http.ResponseWriter, session Session) error {
	tokenString, err := jwt.BuildJWT(sm.jwtSecret, session.payload, sm.expirationTime)

	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:    sm.name,
		Value:   tokenString,
		Path:    "/",
		Expires: sm.expirationTime,
		Domain:  sm.domain,
	})

	return nil
}

func (sm *SessionManager) GetSession(r *http.Request) *Session {
	c, err := r.Cookie(sm.name)
	if err != nil {
		return nil
	}

	payload, err := jwt.ValidateJWT(sm.jwtSecret, c.Value)
	if err != nil {
		return nil
	}

	return &Session{payload: payload}
}
