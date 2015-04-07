package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type ProductCategoryBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *ProductCategoryBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.ProductCategoryDbm
	return &meta
}
func (l *ProductCategoryBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *ProductCategoryBhv) Update(productCategory *entity.ProductCategory, tx *sql.Tx) (int64, error) {
	var entity df.Entity = productCategory
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *ProductCategoryBhv) Insert(productCategory *entity.ProductCategory, tx *sql.Tx) (int64, error) {
	var entity df.Entity = productCategory
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *ProductCategoryBhv) Delete(productCategory *entity.ProductCategory, tx *sql.Tx) (int64, error) {
	var entity df.Entity = productCategory
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *ProductCategoryBhv) SelectList(cb *cb.ProductCategoryCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "ProductCategory", tx)
}
func (l *ProductCategoryBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *ProductCategoryBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var ProductCategoryBhv_I *ProductCategoryBhv
