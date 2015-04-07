package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type WhiteXlsManBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *WhiteXlsManBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.WhiteXlsManDbm
	return &meta
}
func (l *WhiteXlsManBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *WhiteXlsManBhv) Update(whiteXlsMan *entity.WhiteXlsMan, tx *sql.Tx) (int64, error) {
	var entity df.Entity = whiteXlsMan
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *WhiteXlsManBhv) Insert(whiteXlsMan *entity.WhiteXlsMan, tx *sql.Tx) (int64, error) {
	var entity df.Entity = whiteXlsMan
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *WhiteXlsManBhv) Delete(whiteXlsMan *entity.WhiteXlsMan, tx *sql.Tx) (int64, error) {
	var entity df.Entity = whiteXlsMan
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *WhiteXlsManBhv) SelectList(cb *cb.WhiteXlsManCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "WhiteXlsMan", tx)
}
func (l *WhiteXlsManBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *WhiteXlsManBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var WhiteXlsManBhv_I *WhiteXlsManBhv
