package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type VendorPartManBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *VendorPartManBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.VendorPartManDbm
	return &meta
}
func (l *VendorPartManBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *VendorPartManBhv) Update(vendorPartMan *entity.VendorPartMan, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorPartMan
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *VendorPartManBhv) Insert(vendorPartMan *entity.VendorPartMan, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorPartMan
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *VendorPartManBhv) Delete(vendorPartMan *entity.VendorPartMan, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorPartMan
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *VendorPartManBhv) SelectList(cb *cb.VendorPartManCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "VendorPartMan", tx)
}
func (l *VendorPartManBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *VendorPartManBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var VendorPartManBhv_I *VendorPartManBhv
