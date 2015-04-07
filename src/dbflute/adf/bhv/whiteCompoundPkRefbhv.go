package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type WhiteCompoundPkRefBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *WhiteCompoundPkRefBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.WhiteCompoundPkRefDbm
	return &meta
}
func (l *WhiteCompoundPkRefBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *WhiteCompoundPkRefBhv) Update(whiteCompoundPkRef *entity.WhiteCompoundPkRef, tx *sql.Tx) (int64, error) {
	var entity df.Entity = whiteCompoundPkRef
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *WhiteCompoundPkRefBhv) Insert(whiteCompoundPkRef *entity.WhiteCompoundPkRef, tx *sql.Tx) (int64, error) {
	var entity df.Entity = whiteCompoundPkRef
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *WhiteCompoundPkRefBhv) Delete(whiteCompoundPkRef *entity.WhiteCompoundPkRef, tx *sql.Tx) (int64, error) {
	var entity df.Entity = whiteCompoundPkRef
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *WhiteCompoundPkRefBhv) SelectList(cb *cb.WhiteCompoundPkRefCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "WhiteCompoundPkRef", tx)
}
func (l *WhiteCompoundPkRefBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *WhiteCompoundPkRefBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var WhiteCompoundPkRefBhv_I *WhiteCompoundPkRefBhv
