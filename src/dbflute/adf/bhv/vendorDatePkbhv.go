package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type VendorDatePkBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *VendorDatePkBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.VendorDatePkDbm
	return &meta
}
func (l *VendorDatePkBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *VendorDatePkBhv) Update(vendorDatePk *entity.VendorDatePk, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorDatePk
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *VendorDatePkBhv) Insert(vendorDatePk *entity.VendorDatePk, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorDatePk
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *VendorDatePkBhv) Delete(vendorDatePk *entity.VendorDatePk, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorDatePk
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *VendorDatePkBhv) SelectList(cb *cb.VendorDatePkCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "VendorDatePk", tx)
}
func (l *VendorDatePkBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *VendorDatePkBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var VendorDatePkBhv_I *VendorDatePkBhv
