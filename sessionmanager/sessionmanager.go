package sessionmanager

import (
	"github.com/guiceolin/goub/jwt"
	"net/http"
	"time"
)

type SessionManager struct {
	domain         string
	jwtSecret      []byte
	name           string
	expirationTime time.Time
}

func (sm *SessionManager) SetPayload(w http.ResponseWriter, payload map[string]interface{}) error {
	tokenString, err := jwt.BuildJWT(sm.jwtSecret, payload, sm.expirationTime)

	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Path:    "/",
		Expires: sm.expirationTime,
		Domain:  sm.domain,
	})

	return nil
}

func (sm *SessionManager) GetPayload(r *http.Request) map[string]interface{} {
	c, err := r.Cookie(sm.name)
	if err != nil {
		return nil
	}

	payload, err := jwt.ValidateJWT(sm.jwtSecret, c.Value)
	if err != nil {
		return nil
	}

	return payload
}
