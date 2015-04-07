package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type VendorInheritInuBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *VendorInheritInuBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.VendorInheritInuDbm
	return &meta
}
func (l *VendorInheritInuBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *VendorInheritInuBhv) Update(vendorInheritInu *entity.VendorInheritInu, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorInheritInu
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *VendorInheritInuBhv) Insert(vendorInheritInu *entity.VendorInheritInu, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorInheritInu
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *VendorInheritInuBhv) Delete(vendorInheritInu *entity.VendorInheritInu, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorInheritInu
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *VendorInheritInuBhv) SelectList(cb *cb.VendorInheritInuCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "VendorInheritInu", tx)
}
func (l *VendorInheritInuBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *VendorInheritInuBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var VendorInheritInuBhv_I *VendorInheritInuBhv
