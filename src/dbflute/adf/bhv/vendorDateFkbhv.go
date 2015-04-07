package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type VendorDateFkBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *VendorDateFkBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.VendorDateFkDbm
	return &meta
}
func (l *VendorDateFkBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *VendorDateFkBhv) Update(vendorDateFk *entity.VendorDateFk, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorDateFk
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *VendorDateFkBhv) Insert(vendorDateFk *entity.VendorDateFk, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorDateFk
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *VendorDateFkBhv) Delete(vendorDateFk *entity.VendorDateFk, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorDateFk
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *VendorDateFkBhv) SelectList(cb *cb.VendorDateFkCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "VendorDateFk", tx)
}
func (l *VendorDateFkBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *VendorDateFkBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var VendorDateFkBhv_I *VendorDateFkBhv
