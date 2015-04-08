package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type MemberWithdrawal struct {
	memberId int64
	withdrawalReasonCode df.NullString
	withdrawalReasonInputText df.NullString
	withdrawalDatetime df.MysqlTimestamp
	registerDatetime df.MysqlTimestamp
	registerUser string
	updateDatetime df.MysqlTimestamp
	updateUser string
	df.BaseEntity
}

func CreateMemberWithdrawal() *MemberWithdrawal{
	memberWithdrawal :=new(MemberWithdrawal)
	memberWithdrawal.SetUp()
	return memberWithdrawal 
}

func (l *MemberWithdrawal) GetMemberId () int64 {
	return l.memberId
}
func (l *MemberWithdrawal) GetWithdrawalReasonCode () df.NullString {
	return l.withdrawalReasonCode
}
func (l *MemberWithdrawal) GetWithdrawalReasonInputText () df.NullString {
	return l.withdrawalReasonInputText
}
func (l *MemberWithdrawal) GetWithdrawalDatetime () df.MysqlTimestamp {
	return l.withdrawalDatetime
}
func (l *MemberWithdrawal) GetRegisterDatetime () df.MysqlTimestamp {
	return l.registerDatetime
}
func (l *MemberWithdrawal) GetRegisterUser () string {
	return l.registerUser
}
func (l *MemberWithdrawal) GetUpdateDatetime () df.MysqlTimestamp {
	return l.updateDatetime
}
func (l *MemberWithdrawal) GetUpdateUser () string {
	return l.updateUser
}

func (t *MemberWithdrawal) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 8)
	i[0] = &(t.memberId)
	i[1] = &(t.withdrawalReasonCode)
	i[2] = &(t.withdrawalReasonInputText)
	i[3] = &(t.withdrawalDatetime)
	i[4] = &(t.registerDatetime)
	i[5] = &(t.registerUser)
	i[6] = &(t.updateDatetime)
	i[7] = &(t.updateUser)
	return i
}


func (t *MemberWithdrawal) AsTableDbName() string {
	return "MemberWithdrawal"
}

func (t *MemberWithdrawal) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("memberId") == false {
            return false 
        }
        return true;
}
func (t *MemberWithdrawal) SetMemberId(memberId int64) {
	t.AddPropertyName("memberId")
	t.memberId = memberId
}
func (t *MemberWithdrawal) SetWithdrawalReasonCode(withdrawalReasonCode df.NullString) {
	t.AddPropertyName("withdrawalReasonCode")
	t.withdrawalReasonCode = withdrawalReasonCode
}
func (t *MemberWithdrawal) SetWithdrawalReasonInputText(withdrawalReasonInputText df.NullString) {
	t.AddPropertyName("withdrawalReasonInputText")
	t.withdrawalReasonInputText = withdrawalReasonInputText
}
func (t *MemberWithdrawal) SetWithdrawalDatetime(withdrawalDatetime df.MysqlTimestamp) {
	t.AddPropertyName("withdrawalDatetime")
	t.withdrawalDatetime = withdrawalDatetime
}
func (t *MemberWithdrawal) SetRegisterDatetime(registerDatetime df.MysqlTimestamp) {
	t.AddPropertyName("registerDatetime")
	t.registerDatetime = registerDatetime
}
func (t *MemberWithdrawal) SetRegisterUser(registerUser string) {
	t.AddPropertyName("registerUser")
	t.registerUser = registerUser
}
func (t *MemberWithdrawal) SetUpdateDatetime(updateDatetime df.MysqlTimestamp) {
	t.AddPropertyName("updateDatetime")
	t.updateDatetime = updateDatetime
}
func (t *MemberWithdrawal) SetUpdateUser(updateUser string) {
	t.AddPropertyName("updateUser")
	t.updateUser = updateUser
}

func (t *MemberWithdrawal) SetUp(){
	
}
func (t *MemberWithdrawal)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}