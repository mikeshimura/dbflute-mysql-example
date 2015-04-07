package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type MemberWithdrawalBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *MemberWithdrawalBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.MemberWithdrawalDbm
	return &meta
}
func (l *MemberWithdrawalBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *MemberWithdrawalBhv) Update(memberWithdrawal *entity.MemberWithdrawal, tx *sql.Tx) (int64, error) {
	var entity df.Entity = memberWithdrawal
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *MemberWithdrawalBhv) Insert(memberWithdrawal *entity.MemberWithdrawal, tx *sql.Tx) (int64, error) {
	var entity df.Entity = memberWithdrawal
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *MemberWithdrawalBhv) Delete(memberWithdrawal *entity.MemberWithdrawal, tx *sql.Tx) (int64, error) {
	var entity df.Entity = memberWithdrawal
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *MemberWithdrawalBhv) SelectList(cb *cb.MemberWithdrawalCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "MemberWithdrawal", tx)
}
func (l *MemberWithdrawalBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *MemberWithdrawalBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var MemberWithdrawalBhv_I *MemberWithdrawalBhv
