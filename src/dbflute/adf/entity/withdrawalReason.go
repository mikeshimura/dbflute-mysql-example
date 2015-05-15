package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type WithdrawalReason struct {
	withdrawalReasonCode string
	withdrawalReasonText string
	displayOrder int64
	df.BaseEntity
}

func CreateWithdrawalReason() *WithdrawalReason{
	withdrawalReason :=new(WithdrawalReason)
	withdrawalReason.SetUp()
	return withdrawalReason 
}

func (l *WithdrawalReason) GetWithdrawalReasonCode () string {
	return l.withdrawalReasonCode
}
func (l *WithdrawalReason) GetWithdrawalReasonText () string {
	return l.withdrawalReasonText
}
func (l *WithdrawalReason) GetDisplayOrder () int64 {
	return l.displayOrder
}

func (t *WithdrawalReason) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 3)
	i[0] = &(t.withdrawalReasonCode)
	i[1] = &(t.withdrawalReasonText)
	i[2] = &(t.displayOrder)
	return i
}


func (t *WithdrawalReason) AsTableDbName() string {
	return "WithdrawalReason"
}

func (t *WithdrawalReason) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("withdrawalReasonCode") == false {
            return false 
        }
        return true;
}
func (t *WithdrawalReason) SetWithdrawalReasonCode(withdrawalReasonCode string) {
	t.AddPropertyName("withdrawalReasonCode")
	t.withdrawalReasonCode = withdrawalReasonCode
}
func (t *WithdrawalReason) SetWithdrawalReasonText(withdrawalReasonText string) {
	t.AddPropertyName("withdrawalReasonText")
	t.withdrawalReasonText = withdrawalReasonText
}
func (t *WithdrawalReason) SetDisplayOrder(displayOrder int64) {
	t.AddPropertyName("displayOrder")
	t.displayOrder = displayOrder
}
func (t *WithdrawalReason) SetUp(){
	
}
func (t *WithdrawalReason)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}