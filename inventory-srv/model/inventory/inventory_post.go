package inventory

import (
	"fmt"
	"github.com/micro/go-micro/util/log"
	"micro-service/basic/common"
	"micro-service/basic/db"
	proto "micro-service/inventory-srv/proto/inventory"
)

func (s service) Sell(bookId, userId int64) (id int64, err error) {
	tx, err := db.GetDB().Begin()
	if err != nil {
		log.Logf("[Sell] Session Open Failed %s", err.Error())
		return
	}

	// session should roll back when shit happens
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	querySQL := `SELECT id, book_id, unit_price, stock, version FROM inventory WHERE book_id = ?`
	inv := &proto.Inv{}

	updateSQL := `UPDATE inventory SET stock = ?, version = ? WHERE book_id = ? AND version = ?`

	// deduct Inventory method
	var deductInv func() error
	deductInv = func() (errIn error) {
		// exec query
		errIn = tx.QueryRow(querySQL, bookId).Scan(&inv.Id, &inv.BookId, &inv.UnitPrice, &inv.Stock, &inv.Version)
		if errIn != nil {
			log.Logf("[SELL] query data fail, err: %s", errIn)
			return errIn
		}

		if inv.Stock < 1 {
			errIn = fmt.Errorf("[SELL] Inventory shortage")
			log.Logf(errIn.Error())
			return errIn
		}

		// exec update
		r, errIn := tx.Exec(updateSQL, inv.Stock-1, inv.Version+1, bookId, inv.Version)
		if errIn != nil {
			log.Logf("[SELL] Update Database Failure, %s", errIn)
			return
		}

		//try until , danger: should set try-out times
		if affected, _ := r.RowsAffected(); affected == 0 {
			log.Logf("[SELL] Update Database Failure, Version %s Overdue")
			deductInv()
		}

		return
	}

	// do deduct inventory
	err = deductInv()
	if err != nil {
		log.Logf("[SELL] deductInv Failed, %s", err)
		return
	}

	insertSQL := `INSERT inventory_history (book_id, user_id, state) VALUE (?,?,?)`
	r, err := tx.Exec(insertSQL, bookId, userId, common.InventoryHistoryStateNotOut)
	if err != nil {
		log.Logf("[SELL] insert deduInv Failure err, %s", err)
		return
	}

	id, _ = r.LastInsertId()
	tx.Commit()
	return
}

func (s service) Confirm(id int64, state int) (err error) {
	updateSQL := `UPDATE inventory_history SET state = ? WHERE id = ?`

	o := db.GetDB()

	_, err = o.Exec(updateSQL, state, id)
	if err != nil {
		log.Logf("[Confirm] update failed err, %s", err)
		return
	}

	return
}
