package session

import (
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"net/http"
	"strings"
	"time"
)

var (
	sessionIdNamePrefix = "session-id-"
	store               *sessions.CookieStore
	defaultSessionKey   = "OnNUU5RUr6Ii2HMI0d6E54bXTS52tCCL"
)

func init() {
	store = sessions.NewCookieStore([]byte(defaultSessionKey))
}

func GetSession(w http.ResponseWriter, r *http.Request) *sessions.Session {
	var sId string

	// get sid if exists
	for _, c := range r.Cookies() {
		if strings.Index(c.Name, sessionIdNamePrefix) == 0 {
			sId = c.Name
			break
		}
	}

	// gen sid if not exists
	if sId == "" {
		sId = sessionIdNamePrefix + uuid.New().String()
	}

	// get sesssion from store
	ses, _ := store.Get(r, sId)

	// if id is new then gen cookie-sid and save
	if ses.ID == "" {
		cookie := &http.Cookie{Name: sId, Value: sId, Path: "/", Expires: time.Now().Add(30 * time.Second), MaxAge: 0}
		http.SetCookie(w, cookie)

		ses.ID = sId
		ses.Save(r, w)
	}
}
