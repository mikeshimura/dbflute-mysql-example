package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type PurchaseBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *PurchaseBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.PurchaseDbm
	return &meta
}
func (l *PurchaseBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *PurchaseBhv) Update(purchase *entity.Purchase, tx *sql.Tx) (int64, error) {
	var entity df.Entity = purchase
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *PurchaseBhv) Insert(purchase *entity.Purchase, tx *sql.Tx) (int64, error) {
	var entity df.Entity = purchase
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *PurchaseBhv) Delete(purchase *entity.Purchase, tx *sql.Tx) (int64, error) {
	var entity df.Entity = purchase
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *PurchaseBhv) SelectList(cb *cb.PurchaseCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "Purchase", tx)
}
func (l *PurchaseBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *PurchaseBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var PurchaseBhv_I *PurchaseBhv
