package order

import (
	"github.com/go-log/log"
	proto "micro-service/orders-srv/proto/order"
	"micro-service/plugins/db"
)

func (s *service) GetOrder(orderId int64) (order *proto.Order, err error) {
	order = &proto.Order{}

	querySql := `SELECT id,user_id,book_id,inv_his_id,state FROM orders WHERE id = ?`
	o := db.GetDB()

	err = o.QueryRow(querySql, orderId).Scan(
		&order.Id, &order.UserId, &order.BookId, &order.InvHistoryId, &order.State)
	if err != nil {
		log.Logf("[GetOrder] query fail,err,%s", err)
		return
	}

	return
}
