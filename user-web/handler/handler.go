package handler

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/util/log"
	"net/http"
	"time"

	"github.com/micro/go-micro/client"
	proto "user-web/proto/user"
)

var (
	serviceClient proto.UserService
)

type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

func Init() {
	serviceClient = proto.NewUserService("mu.micro.book.srv.user", client.DefaultClient)
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
	rsp, err := serviceClient.QueryUserByName(context.TODO(), &proto.Request{
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
