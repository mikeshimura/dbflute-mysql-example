package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type VendorLargeDataBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *VendorLargeDataBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.VendorLargeDataDbm
	return &meta
}
func (l *VendorLargeDataBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *VendorLargeDataBhv) Update(vendorLargeData *entity.VendorLargeData, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorLargeData
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *VendorLargeDataBhv) Insert(vendorLargeData *entity.VendorLargeData, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorLargeData
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *VendorLargeDataBhv) Delete(vendorLargeData *entity.VendorLargeData, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorLargeData
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *VendorLargeDataBhv) SelectList(cb *cb.VendorLargeDataCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "VendorLargeData", tx)
}
func (l *VendorLargeDataBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *VendorLargeDataBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var VendorLargeDataBhv_I *VendorLargeDataBhv
