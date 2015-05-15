package entity

import (
	"github.com/mikeshimura/dbflute/df"
	"database/sql"
)

type SummaryWithdrawal struct {
	memberId int64
	memberName sql.NullString
	withdrawalReasonCode sql.NullString
	withdrawalReasonText sql.NullString
	withdrawalReasonInputText sql.NullString
	withdrawalDatetime df.MysqlTimestamp
	memberStatusCode sql.NullString
	memberStatusName sql.NullString
	maxPurchasePrice sql.NullInt64
	df.BaseEntity
}

func CreateSummaryWithdrawal() *SummaryWithdrawal{
	summaryWithdrawal :=new(SummaryWithdrawal)
	summaryWithdrawal.SetUp()
	return summaryWithdrawal 
}

func (l *SummaryWithdrawal) GetMemberId () int64 {
	return l.memberId
}
func (l *SummaryWithdrawal) GetMemberName () sql.NullString {
	return l.memberName
}
func (l *SummaryWithdrawal) GetWithdrawalReasonCode () sql.NullString {
	return l.withdrawalReasonCode
}
func (l *SummaryWithdrawal) GetWithdrawalReasonText () sql.NullString {
	return l.withdrawalReasonText
}
func (l *SummaryWithdrawal) GetWithdrawalReasonInputText () sql.NullString {
	return l.withdrawalReasonInputText
}
func (l *SummaryWithdrawal) GetWithdrawalDatetime () df.MysqlTimestamp {
	return l.withdrawalDatetime
}
func (l *SummaryWithdrawal) GetMemberStatusCode () sql.NullString {
	return l.memberStatusCode
}
func (l *SummaryWithdrawal) GetMemberStatusName () sql.NullString {
	return l.memberStatusName
}
func (l *SummaryWithdrawal) GetMaxPurchasePrice () sql.NullInt64 {
	return l.maxPurchasePrice
}

func (t *SummaryWithdrawal) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 9)
	i[0] = &(t.memberId)
	i[1] = &(t.memberName)
	i[2] = &(t.withdrawalReasonCode)
	i[3] = &(t.withdrawalReasonText)
	i[4] = &(t.withdrawalReasonInputText)
	i[5] = &(t.withdrawalDatetime)
	i[6] = &(t.memberStatusCode)
	i[7] = &(t.memberStatusName)
	i[8] = &(t.maxPurchasePrice)
	return i
}


func (t *SummaryWithdrawal) AsTableDbName() string {
	return "SummaryWithdrawal"
}

func (t *SummaryWithdrawal) HasPrimaryKeyValue() bool{
        return false;
}
func (t *SummaryWithdrawal) SetMemberId(memberId int64) {
	t.AddPropertyName("memberId")
	t.memberId = memberId
}
func (t *SummaryWithdrawal) SetMemberName(memberName sql.NullString) {
	t.AddPropertyName("memberName")
	t.memberName = memberName
}
func (t *SummaryWithdrawal) SetWithdrawalReasonCode(withdrawalReasonCode sql.NullString) {
	t.AddPropertyName("withdrawalReasonCode")
	t.withdrawalReasonCode = withdrawalReasonCode
}
func (t *SummaryWithdrawal) SetWithdrawalReasonText(withdrawalReasonText sql.NullString) {
	t.AddPropertyName("withdrawalReasonText")
	t.withdrawalReasonText = withdrawalReasonText
}
func (t *SummaryWithdrawal) SetWithdrawalReasonInputText(withdrawalReasonInputText sql.NullString) {
	t.AddPropertyName("withdrawalReasonInputText")
	t.withdrawalReasonInputText = withdrawalReasonInputText
}
func (t *SummaryWithdrawal) SetWithdrawalDatetime(withdrawalDatetime df.MysqlTimestamp) {
	t.AddPropertyName("withdrawalDatetime")
	t.withdrawalDatetime = withdrawalDatetime
}
func (t *SummaryWithdrawal) SetMemberStatusCode(memberStatusCode sql.NullString) {
	t.AddPropertyName("memberStatusCode")
	t.memberStatusCode = memberStatusCode
}
func (t *SummaryWithdrawal) SetMemberStatusName(memberStatusName sql.NullString) {
	t.AddPropertyName("memberStatusName")
	t.memberStatusName = memberStatusName
}
func (t *SummaryWithdrawal) SetMaxPurchasePrice(maxPurchasePrice sql.NullInt64) {
	t.AddPropertyName("maxPurchasePrice")
	t.maxPurchasePrice = maxPurchasePrice
}
func (t *SummaryWithdrawal) SetUp(){
	
}
func (t *SummaryWithdrawal)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}