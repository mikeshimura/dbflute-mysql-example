package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type WithdrawalReasonBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *WithdrawalReasonBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.WithdrawalReasonDbm
	return &meta
}
func (l *WithdrawalReasonBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *WithdrawalReasonBhv) Update(withdrawalReason *entity.WithdrawalReason, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = withdrawalReason
	return l.BaseBehavior.DoUpdate(&entity, nil, tx, ctx)
}
func (l *WithdrawalReasonBhv) Insert(withdrawalReason *entity.WithdrawalReason, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = withdrawalReason
	return l.BaseBehavior.DoInsert(&entity, nil, tx, ctx)
}
func (l *WithdrawalReasonBhv) Delete(withdrawalReason *entity.WithdrawalReason, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = withdrawalReason
	return l.BaseBehavior.DoDelete(&entity, nil, tx, ctx)
}
func (l *WithdrawalReasonBhv) SelectList(cb *cb.WithdrawalReasonCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "WithdrawalReason", tx)
}
func (l *WithdrawalReasonBhv) SelectCount(cb *cb.WithdrawalReasonCB, tx *sql.Tx) (int64, error) {

	return l.BaseBehavior.DoSelectCount(cb, tx)
}
func (l *WithdrawalReasonBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *WithdrawalReasonBhv) ReadNextVal(tx *sql.Tx) int64{
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var WithdrawalReasonBhv_I *WithdrawalReasonBhv
