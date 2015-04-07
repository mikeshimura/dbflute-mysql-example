package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type WhiteCompoundPkWrongOrderBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *WhiteCompoundPkWrongOrderBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.WhiteCompoundPkWrongOrderDbm
	return &meta
}
func (l *WhiteCompoundPkWrongOrderBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *WhiteCompoundPkWrongOrderBhv) Update(whiteCompoundPkWrongOrder *entity.WhiteCompoundPkWrongOrder, tx *sql.Tx) (int64, error) {
	var entity df.Entity = whiteCompoundPkWrongOrder
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *WhiteCompoundPkWrongOrderBhv) Insert(whiteCompoundPkWrongOrder *entity.WhiteCompoundPkWrongOrder, tx *sql.Tx) (int64, error) {
	var entity df.Entity = whiteCompoundPkWrongOrder
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *WhiteCompoundPkWrongOrderBhv) Delete(whiteCompoundPkWrongOrder *entity.WhiteCompoundPkWrongOrder, tx *sql.Tx) (int64, error) {
	var entity df.Entity = whiteCompoundPkWrongOrder
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *WhiteCompoundPkWrongOrderBhv) SelectList(cb *cb.WhiteCompoundPkWrongOrderCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "WhiteCompoundPkWrongOrder", tx)
}
func (l *WhiteCompoundPkWrongOrderBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *WhiteCompoundPkWrongOrderBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var WhiteCompoundPkWrongOrderBhv_I *WhiteCompoundPkWrongOrderBhv
