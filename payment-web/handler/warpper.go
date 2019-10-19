package handler

import (
	"context"
	"github.com/go-log/log"
	auth "micro-service/auth/proto/auth"
	"micro-service/basic/common"
	"micro-service/plugins/session"
	"net/http"
)

func AuthWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ck, _ := r.Cookie(common.RememberMeCookieName)
		if ck == nil {
			http.Error(w, "invalid request", 400)
			return
		}

		sess := session.GetSession(w, r)
		if sess.ID == "" {
			http.Error(w, "invalid request", 400)
		}

		if sess.Values["valid"] != nil {
			h.ServeHTTP(w, r)
			return
		}

		userId := sess.Values["userId"].(int64)
		if userId == 0 {
			log.Logf("[AuthWarpper], session invalid, no user id")
			return
		}

		rsp, err := authClient.GetCachedAccessToken(context.TODO(),
			&auth.Request{UserId: userId})
		if err != nil {
			log.Logf("[AuthWarpper],err:%s", err)
			http.Error(w, "invalid request", 400)
			return
		}

		if rsp.Token != ck.Value {
			log.Logf("[AuthWrapper], token incorrect")
			http.Error(w, "invalid request", 400)
			return
		}

		h.ServeHTTP(w, r)
	})
}
