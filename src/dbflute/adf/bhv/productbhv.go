package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type ProductBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *ProductBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.ProductDbm
	return &meta
}
func (l *ProductBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *ProductBhv) Update(product *entity.Product, tx *sql.Tx) (int64, error) {
	var entity df.Entity = product
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *ProductBhv) Insert(product *entity.Product, tx *sql.Tx) (int64, error) {
	var entity df.Entity = product
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *ProductBhv) Delete(product *entity.Product, tx *sql.Tx) (int64, error) {
	var entity df.Entity = product
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *ProductBhv) SelectList(cb *cb.ProductCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "Product", tx)
}
func (l *ProductBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *ProductBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var ProductBhv_I *ProductBhv
