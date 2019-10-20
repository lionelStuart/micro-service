package handler

import (
	"context"
	"encoding/json"
	"github.com/go-log/log"
	"net/http"
	"strconv"
	"time"

	"github.com/micro/go-micro/client"
	auth "micro-service/auth/proto/auth"
	pay "micro-service/payment-srv/proto/payment"
)

var (
	serviceClient pay.PaymentService
	authClient    auth.Service
)

type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

func Init() {
	serviceClient = pay.NewPaymentService("mu.micro.book.srv.payment", client.DefaultClient)
	authClient = auth.NewService("mu.micro.book.srv.suth", client.DefaultClient)
}

func PayOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Logf("invalid request")
		http.Error(w, "invalid request ", 400)
	}

	r.ParseForm()

	orderId, _ := strconv.ParseInt(r.Form.Get("orderId"), 10, 10)

	_, err := serviceClient.PayOrder(context.TODO(), &pay.Request{
		OrderId: orderId,
	})
	response := map[string]interface{}{}

	response["ref"] = time.Now().Unix()
	if err != nil {
		response["success"] = false
		response["error"] = Error{
			Detail: err.Error(),
		}
	} else {
		response["success"] = true
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
