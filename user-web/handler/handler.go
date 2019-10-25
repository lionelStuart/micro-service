package handler

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/util/log"
	"micro-service/plugins/session"
	"net/http"
	"time"

	"github.com/micro/go-micro/client"
	auth "micro-service/auth/proto/auth"
	hystrix "micro-service/plugins/hystrix"
	us "micro-service/user-srv/proto/user"
)

var (
	serviceClient us.UserService
	authClient    auth.Service
)

type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

func Init() {
	hystrix.Init()
	cl := hystrix.WrapperClient(client.DefaultClient)
	serviceClient = us.NewUserService("mu.micro.book.srv.user", cl)
	authClient = auth.NewService("mu.micro.book.srv.auth", cl)
}

func Login(w http.ResponseWriter, r *http.Request) {
	// accept post only
	if r.Method != "POST" {
		log.Logf("invalid login request")
		http.Error(w, "invalid", 400)
		return
	}

	// parse form
	r.ParseForm()

	// query sql user by name
	rsp, err := serviceClient.QueryUserByName(context.TODO(), &us.Request{
		UserName: r.Form.Get("userName"),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// generate response by query result
	response := map[string]interface{}{
		"ref": time.Now().UnixNano(),
	}
	if rsp.User.Pwd == r.Form.Get("pwd") {
		response["success"] = true
		rsp.User.Pwd = ""
		response["data"] = rsp.User

		//gen cookies
		log.Logf("[Login] Secret Pass, Start Gen Token .. ")
		rsp2, err := authClient.MakeAccessToken(context.TODO(), &auth.Request{
			UserId:   rsp.User.Id,
			UserName: rsp.User.Name,
		})
		if err != nil {
			log.Logf("[Login] Fail Gen Token,err:%s", err)
			http.Error(w, err.Error(), 500)
			return
		}
		log.Logf("[Login] token %s", rsp2.Token)
		response["token"] = rsp2.Token

		//set token to cookies
		w.Header().Add("set-cookie", "application/json; charset=utf-8")
		expire := time.Now().Add(30 * time.Minute)
		cookie := http.Cookie{Name: "remember-me-token", Value: rsp2.Token, Path: "/", Expires: expire, MaxAge: 90000}
		http.SetCookie(w, &cookie)

		// update session
		sess := session.GetSession(w, r)
		sess.Values["userId"] = rsp.User.Id
		sess.Values["username"] = rsp.User.Name
		_ = sess.Save(r, w)

	} else {
		response["success"] = false
		response["error"] = &Error{
			Detail: "pwd error",
		}
	}

	// add header
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// encode result
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Logf("invalid request")
		http.Error(w, "invalid request", 400)
		return
	}

	tokenCookie, err := r.Cookie("remember-me-token")
	if err != nil {
		log.Logf("Token Request Failure")
		http.Error(w, "invalid request", 400)
		return
	}

	// del token
	_, err = authClient.DelUserAccessToken(context.TODO(), &auth.Request{
		Token: tokenCookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//set empty cookie
	cookie := http.Cookie{Name: "remember-me-token", Value: "", Path: "/", Expires: time.Now().Add(0 * time.Second), MaxAge: 0}
	http.SetCookie(w, &cookie)

	// set response
	w.Header().Add("Content-Type", "application/json;charset=utf-8")

	response := map[string]interface{}{
		"ref":     time.Now().UnixNano(),
		"success": true,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
