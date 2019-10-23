package order

import (
	"context"
	"github.com/go-log/log"
	"micro-service/basic/common"
	inv "micro-service/inventory-srv/proto/inventory"
	"micro-service/plugins/db"
)

func (s *service) New(bookId, userId int64) (orderId int64, err error) {
	rsp, err := invClient.Sell(context.TODO(), &inv.Request{
		BookId: bookId, UserId: userId})
	if err != nil {
		log.Logf("[New] sell call fail ,%s", err)
		return
	}

	o := db.GetDB()
	insertSql := `INSERT orders (user_id,book_id,inv_his_id,state) VALUE (?,?,?,?)`

	r, err := o.Exec(insertSql, userId, bookId, rsp.InvH.Id, common.InventoryHistoryStateNotOut)
	if err != nil {
		log.Logf("[NEW] add order fail err, %s", err)
		return
	}
	orderId, _ = r.LastInsertId()
	return
}

func (s *service) UpdateOrderState(orderId int64, state int) (err error) {
	updateSql := `UPDATE orders SET state = ? WHERE id = ?`

	o := db.GetDB()
	_, err = o.Exec(updateSql, state, orderId)
	if err != nil {
		log.Logf("[Confirm] update fail err,%s", err)
		return
	}

	return
}
