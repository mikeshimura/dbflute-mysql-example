package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type PurchasePaymentBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *PurchasePaymentBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.PurchasePaymentDbm
	return &meta
}
func (l *PurchasePaymentBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *PurchasePaymentBhv) Update(purchasePayment *entity.PurchasePayment, tx *sql.Tx) (int64, error) {
	var entity df.Entity = purchasePayment
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *PurchasePaymentBhv) Insert(purchasePayment *entity.PurchasePayment, tx *sql.Tx) (int64, error) {
	var entity df.Entity = purchasePayment
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *PurchasePaymentBhv) Delete(purchasePayment *entity.PurchasePayment, tx *sql.Tx) (int64, error) {
	var entity df.Entity = purchasePayment
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *PurchasePaymentBhv) SelectList(cb *cb.PurchasePaymentCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "PurchasePayment", tx)
}
func (l *PurchasePaymentBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *PurchasePaymentBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var PurchasePaymentBhv_I *PurchasePaymentBhv
