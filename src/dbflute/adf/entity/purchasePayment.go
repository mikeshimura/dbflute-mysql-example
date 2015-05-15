package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type PurchasePayment struct {
	purchasePaymentId int64
	purchaseId int64
	paymentAmount df.Numeric
	paymentDatetime df.MysqlTimestamp
	paymentMethodCode string
	registerDatetime df.MysqlTimestamp
	registerUser string
	updateDatetime df.MysqlTimestamp
	updateUser string
	df.BaseEntity
Purchase_R  *Purchase

}

func CreatePurchasePayment() *PurchasePayment{
	purchasePayment :=new(PurchasePayment)
	purchasePayment.SetUp()
	return purchasePayment 
}

func (l *PurchasePayment) GetPurchasePaymentId () int64 {
	return l.purchasePaymentId
}
func (l *PurchasePayment) GetPurchaseId () int64 {
	return l.purchaseId
}
func (l *PurchasePayment) GetPaymentAmount () df.Numeric {
	return l.paymentAmount
}
func (l *PurchasePayment) GetPaymentDatetime () df.MysqlTimestamp {
	return l.paymentDatetime
}
func (l *PurchasePayment) GetPaymentMethodCode () string {
	return l.paymentMethodCode
}
func (l *PurchasePayment) GetRegisterDatetime () df.MysqlTimestamp {
	return l.registerDatetime
}
func (l *PurchasePayment) GetRegisterUser () string {
	return l.registerUser
}
func (l *PurchasePayment) GetUpdateDatetime () df.MysqlTimestamp {
	return l.updateDatetime
}
func (l *PurchasePayment) GetUpdateUser () string {
	return l.updateUser
}

func (t *PurchasePayment) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 9)
	i[0] = &(t.purchasePaymentId)
	i[1] = &(t.purchaseId)
	i[2] = &(t.paymentAmount)
	i[3] = &(t.paymentDatetime)
	i[4] = &(t.paymentMethodCode)
	i[5] = &(t.registerDatetime)
	i[6] = &(t.registerUser)
	i[7] = &(t.updateDatetime)
	i[8] = &(t.updateUser)
	return i
}


func (t *PurchasePayment) AsTableDbName() string {
	return "PurchasePayment"
}

func (t *PurchasePayment) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("purchasePaymentId") == false {
            return false 
        }
        return true;
}
func (t *PurchasePayment) SetPurchasePaymentId(purchasePaymentId int64) {
	t.AddPropertyName("purchasePaymentId")
	t.purchasePaymentId = purchasePaymentId
}
func (t *PurchasePayment) SetPurchaseId(purchaseId int64) {
	t.AddPropertyName("purchaseId")
	t.purchaseId = purchaseId
}
func (t *PurchasePayment) SetPaymentAmount(paymentAmount df.Numeric) {
	t.AddPropertyName("paymentAmount")
	t.paymentAmount = paymentAmount
}
func (t *PurchasePayment) SetPaymentDatetime(paymentDatetime df.MysqlTimestamp) {
	t.AddPropertyName("paymentDatetime")
	t.paymentDatetime = paymentDatetime
}
func (t *PurchasePayment) SetPaymentMethodCode(paymentMethodCode string) {
	t.AddPropertyName("paymentMethodCode")
	t.paymentMethodCode = paymentMethodCode
}
func (t *PurchasePayment) SetRegisterDatetime(registerDatetime df.MysqlTimestamp) {
	t.AddPropertyName("registerDatetime")
	t.registerDatetime = registerDatetime
}
func (t *PurchasePayment) SetRegisterUser(registerUser string) {
	t.AddPropertyName("registerUser")
	t.registerUser = registerUser
}
func (t *PurchasePayment) SetUpdateDatetime(updateDatetime df.MysqlTimestamp) {
	t.AddPropertyName("updateDatetime")
	t.updateDatetime = updateDatetime
}
func (t *PurchasePayment) SetUpdateUser(updateUser string) {
	t.AddPropertyName("updateUser")
	t.updateUser = updateUser
}
func (t *PurchasePayment) GetPurchase_R() *Purchase{
	return t.Purchase_R
}
func (t *PurchasePayment) SetPurchase_R(value *Purchase) {
    t.Purchase_R = value
}
func (t *PurchasePayment) SetUp(){
	t.paymentAmount.DecPoint = 2
	
}
func (t *PurchasePayment)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}