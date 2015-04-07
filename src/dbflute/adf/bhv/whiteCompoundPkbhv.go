package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type WhiteCompoundPkBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *WhiteCompoundPkBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.WhiteCompoundPkDbm
	return &meta
}
func (l *WhiteCompoundPkBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *WhiteCompoundPkBhv) Update(whiteCompoundPk *entity.WhiteCompoundPk, tx *sql.Tx) (int64, error) {
	var entity df.Entity = whiteCompoundPk
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *WhiteCompoundPkBhv) Insert(whiteCompoundPk *entity.WhiteCompoundPk, tx *sql.Tx) (int64, error) {
	var entity df.Entity = whiteCompoundPk
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *WhiteCompoundPkBhv) Delete(whiteCompoundPk *entity.WhiteCompoundPk, tx *sql.Tx) (int64, error) {
	var entity df.Entity = whiteCompoundPk
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *WhiteCompoundPkBhv) SelectList(cb *cb.WhiteCompoundPkCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "WhiteCompoundPk", tx)
}
func (l *WhiteCompoundPkBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *WhiteCompoundPkBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var WhiteCompoundPkBhv_I *WhiteCompoundPkBhv
