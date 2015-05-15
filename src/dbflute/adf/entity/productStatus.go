package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type ProductStatus struct {
	productStatusCode string
	productStatusName string
	displayOrder int64
	df.BaseEntity
}

func CreateProductStatus() *ProductStatus{
	productStatus :=new(ProductStatus)
	productStatus.SetUp()
	return productStatus 
}

func (l *ProductStatus) GetProductStatusCode () string {
	return l.productStatusCode
}
func (l *ProductStatus) GetProductStatusName () string {
	return l.productStatusName
}
func (l *ProductStatus) GetDisplayOrder () int64 {
	return l.displayOrder
}

func (t *ProductStatus) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 3)
	i[0] = &(t.productStatusCode)
	i[1] = &(t.productStatusName)
	i[2] = &(t.displayOrder)
	return i
}


func (t *ProductStatus) AsTableDbName() string {
	return "ProductStatus"
}

func (t *ProductStatus) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("productStatusCode") == false {
            return false 
        }
        return true;
}
func (t *ProductStatus) SetProductStatusCode(productStatusCode string) {
	t.AddPropertyName("productStatusCode")
	t.productStatusCode = productStatusCode
}
func (t *ProductStatus) SetProductStatusName(productStatusName string) {
	t.AddPropertyName("productStatusName")
	t.productStatusName = productStatusName
}
func (t *ProductStatus) SetDisplayOrder(displayOrder int64) {
	t.AddPropertyName("displayOrder")
	t.displayOrder = displayOrder
}
func (t *ProductStatus) SetUp(){
	
}
func (t *ProductStatus)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}