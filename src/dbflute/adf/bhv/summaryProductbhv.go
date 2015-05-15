package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type SummaryProductBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *SummaryProductBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.SummaryProductDbm
	return &meta
}
func (l *SummaryProductBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *SummaryProductBhv) SelectList(cb *cb.SummaryProductCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "SummaryProduct", tx)
}
func (l *SummaryProductBhv) SelectCount(cb *cb.SummaryProductCB, tx *sql.Tx) (int64, error) {

	return l.BaseBehavior.DoSelectCount(cb, tx)
}
func (l *SummaryProductBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *SummaryProductBhv) ReadNextVal(tx *sql.Tx) int64{
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var SummaryProductBhv_I *SummaryProductBhv
