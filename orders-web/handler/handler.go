package handler

import (
	"context"
	"encoding/json"
	"github.com/go-log/log"
	"micro-service/plugins/session"
	"net/http"
	"strconv"
	"time"

	"github.com/micro/go-micro/client"
	auth "micro-service/auth/proto/auth"
	orders "micro-service/orders-srv/proto/order"
)

var (
	serviceClient orders.OrdersService
	authClient    auth.Service
)

type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

func Init() {
	serviceClient = orders.NewOrdersService("mu.micro.book.srv.orders", client.DefaultClient)
	authClient = auth.NewService("mu.micro.book.srv.auth", client.DefaultClient)
}

func New(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Logf("invalid request")
		http.Error(w, "invalid request", 400)
		return
	}

	r.ParseForm()
	bookId, _ := strconv.ParseInt(r.Form.Get("bookId"), 10, 10)

	response := map[string]interface{}{}

	rsp, err := serviceClient.New(context.TODO(), &orders.Request{
		BookId: bookId,
		UserId: session.GetSession(w, r).Values["userId"].(int64),
	})

	response["ref"] = time.Now().UnixNano()
	if err != nil {
		response["success"] = false
		response["error"] = Error{
			Detail: err.Error(),
		}
	} else {
		response["success"] = true
		response["orderId"] = rsp.Order.Id
	}

	w.Header().Add("Content-Type", "application/json;charset=utf-8")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
