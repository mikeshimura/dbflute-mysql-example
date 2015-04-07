package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type ProductStatusBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *ProductStatusBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.ProductStatusDbm
	return &meta
}
func (l *ProductStatusBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *ProductStatusBhv) Update(productStatus *entity.ProductStatus, tx *sql.Tx) (int64, error) {
	var entity df.Entity = productStatus
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *ProductStatusBhv) Insert(productStatus *entity.ProductStatus, tx *sql.Tx) (int64, error) {
	var entity df.Entity = productStatus
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *ProductStatusBhv) Delete(productStatus *entity.ProductStatus, tx *sql.Tx) (int64, error) {
	var entity df.Entity = productStatus
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *ProductStatusBhv) SelectList(cb *cb.ProductStatusCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "ProductStatus", tx)
}
func (l *ProductStatusBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *ProductStatusBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var ProductStatusBhv_I *ProductStatusBhv
