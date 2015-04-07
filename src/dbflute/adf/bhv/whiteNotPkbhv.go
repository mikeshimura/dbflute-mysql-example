package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type WhiteNotPkBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *WhiteNotPkBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.WhiteNotPkDbm
	return &meta
}
func (l *WhiteNotPkBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *WhiteNotPkBhv) SelectList(cb *cb.WhiteNotPkCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "WhiteNotPk", tx)
}
func (l *WhiteNotPkBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *WhiteNotPkBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var WhiteNotPkBhv_I *WhiteNotPkBhv
