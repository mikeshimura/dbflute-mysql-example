package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type SummaryProduct struct {
	productId int64
	productName string
	productHandleCode string
	productStatusCode string
	latestPurchaseDatetime df.MysqlNullTimestamp
	df.BaseEntity
}

func CreateSummaryProduct() *SummaryProduct{
	summaryProduct :=new(SummaryProduct)
	summaryProduct.SetUp()
	return summaryProduct 
}

func (l *SummaryProduct) GetProductId () int64 {
	return l.productId
}
func (l *SummaryProduct) GetProductName () string {
	return l.productName
}
func (l *SummaryProduct) GetProductHandleCode () string {
	return l.productHandleCode
}
func (l *SummaryProduct) GetProductStatusCode () string {
	return l.productStatusCode
}
func (l *SummaryProduct) GetLatestPurchaseDatetime () df.MysqlNullTimestamp {
	return l.latestPurchaseDatetime
}

func (t *SummaryProduct) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 5)
	i[0] = &(t.productId)
	i[1] = &(t.productName)
	i[2] = &(t.productHandleCode)
	i[3] = &(t.productStatusCode)
	i[4] = &(t.latestPurchaseDatetime)
	return i
}


func (t *SummaryProduct) AsTableDbName() string {
	return "SummaryProduct"
}

func (t *SummaryProduct) HasPrimaryKeyValue() bool{
        return false;
}
func (t *SummaryProduct) SetProductId(productId int64) {
	t.AddPropertyName("productId")
	t.productId = productId
}
func (t *SummaryProduct) SetProductName(productName string) {
	t.AddPropertyName("productName")
	t.productName = productName
}
func (t *SummaryProduct) SetProductHandleCode(productHandleCode string) {
	t.AddPropertyName("productHandleCode")
	t.productHandleCode = productHandleCode
}
func (t *SummaryProduct) SetProductStatusCode(productStatusCode string) {
	t.AddPropertyName("productStatusCode")
	t.productStatusCode = productStatusCode
}
func (t *SummaryProduct) SetLatestPurchaseDatetime(latestPurchaseDatetime df.MysqlNullTimestamp) {
	t.AddPropertyName("latestPurchaseDatetime")
	t.latestPurchaseDatetime = latestPurchaseDatetime
}

func (t *SummaryProduct) SetUp(){
	
}
func (t *SummaryProduct)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}