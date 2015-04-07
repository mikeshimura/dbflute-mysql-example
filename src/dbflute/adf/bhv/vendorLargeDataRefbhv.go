package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type VendorLargeDataRefBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *VendorLargeDataRefBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.VendorLargeDataRefDbm
	return &meta
}
func (l *VendorLargeDataRefBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *VendorLargeDataRefBhv) Update(vendorLargeDataRef *entity.VendorLargeDataRef, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorLargeDataRef
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *VendorLargeDataRefBhv) Insert(vendorLargeDataRef *entity.VendorLargeDataRef, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorLargeDataRef
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *VendorLargeDataRefBhv) Delete(vendorLargeDataRef *entity.VendorLargeDataRef, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorLargeDataRef
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *VendorLargeDataRefBhv) SelectList(cb *cb.VendorLargeDataRefCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "VendorLargeDataRef", tx)
}
func (l *VendorLargeDataRefBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *VendorLargeDataRefBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var VendorLargeDataRefBhv_I *VendorLargeDataRefBhv
