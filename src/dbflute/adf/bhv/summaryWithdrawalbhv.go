package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type SummaryWithdrawalBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *SummaryWithdrawalBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.SummaryWithdrawalDbm
	return &meta
}
func (l *SummaryWithdrawalBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *SummaryWithdrawalBhv) SelectList(cb *cb.SummaryWithdrawalCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "SummaryWithdrawal", tx)
}
func (l *SummaryWithdrawalBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *SummaryWithdrawalBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var SummaryWithdrawalBhv_I *SummaryWithdrawalBhv
