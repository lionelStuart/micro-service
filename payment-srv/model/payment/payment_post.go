package payment

import (
	"context"
	"fmt"
	"github.com/go-log/log"
	"micro-service/basic/common"
	"micro-service/basic/db"
	inv "micro-service/inventory-srv/proto/inventory"
	ord "micro-service/orders-srv/proto/order"
)

func (s *service) PayOrder(orderId int64) (err error) {
	orderRsp, err := orderClient.GetOrder(context.TODO(), &ord.Request{OrderId: orderId})
	if err != nil {
		log.Logf("[PayOrder] query order info err,%d , %s", orderId, err)
		return
	}

	if orderRsp == nil || !orderRsp.Success || orderRsp.Order == nil {
		err = fmt.Errorf("[PayOrder] order not exists %d, %s", orderId, err)
		log.Log(err.Error())
		return
	}

	tx, err := db.GetDB().Begin()
	if err != nil {
		log.Logf("[PayOrder] tx open fail ", err.Error())
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	insertSQL := `INSERT INTO payment (user_id, book_id, order_id, inv_his_id, state) VALUE(?,?,?,?,?)`
	_, err = tx.Exec(insertSQL, orderRsp.Order.UserId, orderRsp.Order.BookId, orderRsp.Order.Id, orderRsp.Order.InvHistoryId, common.InventoryHistoryStateOut)
	if err != nil {
		log.Logf("[New] add new order fail, %v, %s", orderRsp.Order, err)
		return
	}

	invRsp, err := invClient.Confirm(context.TODO(), &inv.Request{
		HistoryId: orderRsp.Order.InvHistoryId,
	})
	if err != nil || invRsp == nil || !invRsp.Success {
		err = fmt.Errorf("[PayOrder] cofirm order fail, %s", err)
		log.Logf("%s", err)
		return
	}

	s.sendPayDoneEvt(orderId, common.InventoryHistoryStateNotOut)

	tx.Commit()
	return
}
