package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type Product struct {
	productId int64
	productName string
	productHandleCode string
	productCategoryCode string
	productStatusCode string
	regularPrice int64
	registerDatetime df.MysqlTimestamp
	registerUser string
	updateDatetime df.MysqlTimestamp
	updateUser string
	versionNo int64
	df.BaseEntity
}

func CreateProduct() *Product{
	product :=new(Product)
	product.SetUp()
	return product 
}

func (l *Product) GetProductId () int64 {
	return l.productId
}
func (l *Product) GetProductName () string {
	return l.productName
}
func (l *Product) GetProductHandleCode () string {
	return l.productHandleCode
}
func (l *Product) GetProductCategoryCode () string {
	return l.productCategoryCode
}
func (l *Product) GetProductStatusCode () string {
	return l.productStatusCode
}
func (l *Product) GetRegularPrice () int64 {
	return l.regularPrice
}
func (l *Product) GetRegisterDatetime () df.MysqlTimestamp {
	return l.registerDatetime
}
func (l *Product) GetRegisterUser () string {
	return l.registerUser
}
func (l *Product) GetUpdateDatetime () df.MysqlTimestamp {
	return l.updateDatetime
}
func (l *Product) GetUpdateUser () string {
	return l.updateUser
}
func (l *Product) GetVersionNo () int64 {
	return l.versionNo
}

func (t *Product) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 11)
	i[0] = &(t.productId)
	i[1] = &(t.productName)
	i[2] = &(t.productHandleCode)
	i[3] = &(t.productCategoryCode)
	i[4] = &(t.productStatusCode)
	i[5] = &(t.regularPrice)
	i[6] = &(t.registerDatetime)
	i[7] = &(t.registerUser)
	i[8] = &(t.updateDatetime)
	i[9] = &(t.updateUser)
	i[10] = &(t.versionNo)
	return i
}


func (t *Product) AsTableDbName() string {
	return "Product"
}

func (t *Product) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("productId") == false {
            return false 
        }
        return true;
}
func (t *Product) SetProductId(productId int64) {
	t.AddPropertyName("productId")
	t.productId = productId
}
func (t *Product) SetProductName(productName string) {
	t.AddPropertyName("productName")
	t.productName = productName
}
func (t *Product) SetProductHandleCode(productHandleCode string) {
	t.AddPropertyName("productHandleCode")
	t.productHandleCode = productHandleCode
}
func (t *Product) SetProductCategoryCode(productCategoryCode string) {
	t.AddPropertyName("productCategoryCode")
	t.productCategoryCode = productCategoryCode
}
func (t *Product) SetProductStatusCode(productStatusCode string) {
	t.AddPropertyName("productStatusCode")
	t.productStatusCode = productStatusCode
}
func (t *Product) SetRegularPrice(regularPrice int64) {
	t.AddPropertyName("regularPrice")
	t.regularPrice = regularPrice
}
func (t *Product) SetRegisterDatetime(registerDatetime df.MysqlTimestamp) {
	t.AddPropertyName("registerDatetime")
	t.registerDatetime = registerDatetime
}
func (t *Product) SetRegisterUser(registerUser string) {
	t.AddPropertyName("registerUser")
	t.registerUser = registerUser
}
func (t *Product) SetUpdateDatetime(updateDatetime df.MysqlTimestamp) {
	t.AddPropertyName("updateDatetime")
	t.updateDatetime = updateDatetime
}
func (t *Product) SetUpdateUser(updateUser string) {
	t.AddPropertyName("updateUser")
	t.updateUser = updateUser
}
func (t *Product) SetVersionNo(versionNo int64) {
	t.AddPropertyName("versionNo")
	t.versionNo = versionNo
}

func (t *Product) SetUp(){
	
}
func (t *Product)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}