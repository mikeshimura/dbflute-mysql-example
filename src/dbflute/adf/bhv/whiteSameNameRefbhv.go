package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type WhiteSameNameRefBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *WhiteSameNameRefBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.WhiteSameNameRefDbm
	return &meta
}
func (l *WhiteSameNameRefBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *WhiteSameNameRefBhv) Update(whiteSameNameRef *entity.WhiteSameNameRef, tx *sql.Tx) (int64, error) {
	var entity df.Entity = whiteSameNameRef
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *WhiteSameNameRefBhv) Insert(whiteSameNameRef *entity.WhiteSameNameRef, tx *sql.Tx) (int64, error) {
	var entity df.Entity = whiteSameNameRef
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *WhiteSameNameRefBhv) Delete(whiteSameNameRef *entity.WhiteSameNameRef, tx *sql.Tx) (int64, error) {
	var entity df.Entity = whiteSameNameRef
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *WhiteSameNameRefBhv) SelectList(cb *cb.WhiteSameNameRefCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "WhiteSameNameRef", tx)
}
func (l *WhiteSameNameRefBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *WhiteSameNameRefBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var WhiteSameNameRefBhv_I *WhiteSameNameRefBhv
