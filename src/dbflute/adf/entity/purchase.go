package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type Purchase struct {
	purchaseId int64
	memberId int64
	productId int64
	purchaseDatetime df.Timestamp
	purchaseCount int64
	purchasePrice int64
	paymentCompleteFlg int64
	registerDatetime df.Timestamp
	registerUser string
	registerProcess string
	updateDatetime df.Timestamp
	updateUser string
	updateProcess string
	versionNo int64
	df.BaseEntity
}

func CreatePurchase() *Purchase{
	purchase :=new(Purchase)
	purchase.SetUp()
	return purchase 
}

func (l *Purchase) GetPurchaseId () int64 {
	return l.purchaseId
}
func (l *Purchase) GetMemberId () int64 {
	return l.memberId
}
func (l *Purchase) GetProductId () int64 {
	return l.productId
}
func (l *Purchase) GetPurchaseDatetime () df.Timestamp {
	return l.purchaseDatetime
}
func (l *Purchase) GetPurchaseCount () int64 {
	return l.purchaseCount
}
func (l *Purchase) GetPurchasePrice () int64 {
	return l.purchasePrice
}
func (l *Purchase) GetPaymentCompleteFlg () int64 {
	return l.paymentCompleteFlg
}
func (l *Purchase) GetRegisterDatetime () df.Timestamp {
	return l.registerDatetime
}
func (l *Purchase) GetRegisterUser () string {
	return l.registerUser
}
func (l *Purchase) GetRegisterProcess () string {
	return l.registerProcess
}
func (l *Purchase) GetUpdateDatetime () df.Timestamp {
	return l.updateDatetime
}
func (l *Purchase) GetUpdateUser () string {
	return l.updateUser
}
func (l *Purchase) GetUpdateProcess () string {
	return l.updateProcess
}
func (l *Purchase) GetVersionNo () int64 {
	return l.versionNo
}

func (t *Purchase) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 14)
	i[0] = &(t.purchaseId)
	i[1] = &(t.memberId)
	i[2] = &(t.productId)
	i[3] = &(t.purchaseDatetime)
	i[4] = &(t.purchaseCount)
	i[5] = &(t.purchasePrice)
	i[6] = &(t.paymentCompleteFlg)
	i[7] = &(t.registerDatetime)
	i[8] = &(t.registerUser)
	i[9] = &(t.registerProcess)
	i[10] = &(t.updateDatetime)
	i[11] = &(t.updateUser)
	i[12] = &(t.updateProcess)
	i[13] = &(t.versionNo)
	return i
}


func (t *Purchase) AsTableDbName() string {
	return "Purchase"
}

func (t *Purchase) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("purchaseId") == false {
            return false 
        }
        return true;
}
func (t *Purchase) SetPurchaseId(purchaseId int64) {
	t.AddPropertyName("purchaseId")
	t.purchaseId = purchaseId
}
func (t *Purchase) SetMemberId(memberId int64) {
	t.AddPropertyName("memberId")
	t.memberId = memberId
}
func (t *Purchase) SetProductId(productId int64) {
	t.AddPropertyName("productId")
	t.productId = productId
}
func (t *Purchase) SetPurchaseDatetime(purchaseDatetime df.Timestamp) {
	t.AddPropertyName("purchaseDatetime")
	t.purchaseDatetime = purchaseDatetime
}
func (t *Purchase) SetPurchaseCount(purchaseCount int64) {
	t.AddPropertyName("purchaseCount")
	t.purchaseCount = purchaseCount
}
func (t *Purchase) SetPurchasePrice(purchasePrice int64) {
	t.AddPropertyName("purchasePrice")
	t.purchasePrice = purchasePrice
}
func (t *Purchase) SetPaymentCompleteFlg(paymentCompleteFlg int64) {
	t.AddPropertyName("paymentCompleteFlg")
	t.paymentCompleteFlg = paymentCompleteFlg
}
func (t *Purchase) SetRegisterDatetime(registerDatetime df.Timestamp) {
	t.AddPropertyName("registerDatetime")
	t.registerDatetime = registerDatetime
}
func (t *Purchase) SetRegisterUser(registerUser string) {
	t.AddPropertyName("registerUser")
	t.registerUser = registerUser
}
func (t *Purchase) SetRegisterProcess(registerProcess string) {
	t.AddPropertyName("registerProcess")
	t.registerProcess = registerProcess
}
func (t *Purchase) SetUpdateDatetime(updateDatetime df.Timestamp) {
	t.AddPropertyName("updateDatetime")
	t.updateDatetime = updateDatetime
}
func (t *Purchase) SetUpdateUser(updateUser string) {
	t.AddPropertyName("updateUser")
	t.updateUser = updateUser
}
func (t *Purchase) SetUpdateProcess(updateProcess string) {
	t.AddPropertyName("updateProcess")
	t.updateProcess = updateProcess
}
func (t *Purchase) SetVersionNo(versionNo int64) {
	t.AddPropertyName("versionNo")
	t.versionNo = versionNo
}

func (t *Purchase) SetUp(){
	
}
func (t *Purchase)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}