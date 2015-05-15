package entity

import (
	"github.com/mikeshimura/dbflute/df"
	"database/sql"
)

type ProductCategory struct {
	productCategoryCode string
	productCategoryName string
	parentCategoryCode sql.NullString
	df.BaseEntity
ProductCategory_R  *ProductCategory

}

func CreateProductCategory() *ProductCategory{
	productCategory :=new(ProductCategory)
	productCategory.SetUp()
	return productCategory 
}

func (l *ProductCategory) GetProductCategoryCode () string {
	return l.productCategoryCode
}
func (l *ProductCategory) GetProductCategoryName () string {
	return l.productCategoryName
}
func (l *ProductCategory) GetParentCategoryCode () sql.NullString {
	return l.parentCategoryCode
}

func (t *ProductCategory) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 3)
	i[0] = &(t.productCategoryCode)
	i[1] = &(t.productCategoryName)
	i[2] = &(t.parentCategoryCode)
	return i
}


func (t *ProductCategory) AsTableDbName() string {
	return "ProductCategory"
}

func (t *ProductCategory) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("productCategoryCode") == false {
            return false 
        }
        return true;
}
func (t *ProductCategory) SetProductCategoryCode(productCategoryCode string) {
	t.AddPropertyName("productCategoryCode")
	t.productCategoryCode = productCategoryCode
}
func (t *ProductCategory) SetProductCategoryName(productCategoryName string) {
	t.AddPropertyName("productCategoryName")
	t.productCategoryName = productCategoryName
}
func (t *ProductCategory) SetParentCategoryCode(parentCategoryCode sql.NullString) {
	t.AddPropertyName("parentCategoryCode")
	t.parentCategoryCode = parentCategoryCode
}
func (t *ProductCategory) GetProductCategory_R() *ProductCategory{
	return t.ProductCategory_R
}
func (t *ProductCategory) SetProductCategory_R(value *ProductCategory) {
    t.ProductCategory_R = value
}
func (t *ProductCategory) SetUp(){
	
}
func (t *ProductCategory)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}