package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type WhiteSameNameBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *WhiteSameNameBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.WhiteSameNameDbm
	return &meta
}
func (l *WhiteSameNameBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *WhiteSameNameBhv) Update(whiteSameName *entity.WhiteSameName, tx *sql.Tx) (int64, error) {
	var entity df.Entity = whiteSameName
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *WhiteSameNameBhv) Insert(whiteSameName *entity.WhiteSameName, tx *sql.Tx) (int64, error) {
	var entity df.Entity = whiteSameName
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *WhiteSameNameBhv) Delete(whiteSameName *entity.WhiteSameName, tx *sql.Tx) (int64, error) {
	var entity df.Entity = whiteSameName
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *WhiteSameNameBhv) SelectList(cb *cb.WhiteSameNameCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "WhiteSameName", tx)
}
func (l *WhiteSameNameBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *WhiteSameNameBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var WhiteSameNameBhv_I *WhiteSameNameBhv
