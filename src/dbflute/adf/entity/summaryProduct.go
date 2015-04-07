package entity

import (
	"github.com/mikeshimura/dbflute/df"
	"database/sql"
)

type SummaryProduct struct {
	productId sql.NullInt64
	productName sql.NullString
	productStatusCode sql.NullString
	latestPurchaseDatetime df.NullTimestamp
	df.BaseEntity
}

func CreateSummaryProduct() *SummaryProduct{
	summaryProduct :=new(SummaryProduct)
	summaryProduct.SetUp()
	return summaryProduct 
}

func (l *SummaryProduct) GetProductId () sql.NullInt64 {
	return l.productId
}
func (l *SummaryProduct) GetProductName () sql.NullString {
	return l.productName
}
func (l *SummaryProduct) GetProductStatusCode () sql.NullString {
	return l.productStatusCode
}
func (l *SummaryProduct) GetLatestPurchaseDatetime () df.NullTimestamp {
	return l.latestPurchaseDatetime
}

func (t *SummaryProduct) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 4)
	i[0] = &(t.productId)
	i[1] = &(t.productName)
	i[2] = &(t.productStatusCode)
	i[3] = &(t.latestPurchaseDatetime)
	return i
}


func (t *SummaryProduct) AsTableDbName() string {
	return "SummaryProduct"
}

func (t *SummaryProduct) HasPrimaryKeyValue() bool{
        return false;
}
func (t *SummaryProduct) SetProductId(productId sql.NullInt64) {
	t.AddPropertyName("productId")
	t.productId = productId
}
func (t *SummaryProduct) SetProductName(productName sql.NullString) {
	t.AddPropertyName("productName")
	t.productName = productName
}
func (t *SummaryProduct) SetProductStatusCode(productStatusCode sql.NullString) {
	t.AddPropertyName("productStatusCode")
	t.productStatusCode = productStatusCode
}
func (t *SummaryProduct) SetLatestPurchaseDatetime(latestPurchaseDatetime df.NullTimestamp) {
	t.AddPropertyName("latestPurchaseDatetime")
	t.latestPurchaseDatetime = latestPurchaseDatetime
}

func (t *SummaryProduct) SetUp(){
	
}
func (t *SummaryProduct)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}