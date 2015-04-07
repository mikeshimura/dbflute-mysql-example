package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type VendorInheritNekoBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *VendorInheritNekoBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.VendorInheritNekoDbm
	return &meta
}
func (l *VendorInheritNekoBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *VendorInheritNekoBhv) Update(vendorInheritNeko *entity.VendorInheritNeko, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorInheritNeko
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *VendorInheritNekoBhv) Insert(vendorInheritNeko *entity.VendorInheritNeko, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorInheritNeko
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *VendorInheritNekoBhv) Delete(vendorInheritNeko *entity.VendorInheritNeko, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorInheritNeko
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *VendorInheritNekoBhv) SelectList(cb *cb.VendorInheritNekoCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "VendorInheritNeko", tx)
}
func (l *VendorInheritNekoBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *VendorInheritNekoBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var VendorInheritNekoBhv_I *VendorInheritNekoBhv
