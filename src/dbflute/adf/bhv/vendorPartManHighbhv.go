package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type VendorPartManHighBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *VendorPartManHighBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.VendorPartManHighDbm
	return &meta
}
func (l *VendorPartManHighBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *VendorPartManHighBhv) Update(vendorPartManHigh *entity.VendorPartManHigh, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorPartManHigh
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *VendorPartManHighBhv) Insert(vendorPartManHigh *entity.VendorPartManHigh, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorPartManHigh
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *VendorPartManHighBhv) Delete(vendorPartManHigh *entity.VendorPartManHigh, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorPartManHigh
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *VendorPartManHighBhv) SelectList(cb *cb.VendorPartManHighCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "VendorPartManHigh", tx)
}
func (l *VendorPartManHighBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *VendorPartManHighBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var VendorPartManHighBhv_I *VendorPartManHighBhv
