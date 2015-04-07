package entity

import (
	"github.com/mikeshimura/dbflute/df"
	"database/sql"
)

type MemberWithdrawal struct {
	memberId int64
	withdrawalReasonCode sql.NullString
	withdrawalReasonInputText sql.NullString
	withdrawalDatetime df.Timestamp
	registerDatetime df.Timestamp
	registerProcess string
	registerUser string
	updateDatetime df.Timestamp
	updateProcess string
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
func (l *MemberWithdrawal) GetWithdrawalReasonCode () sql.NullString {
	return l.withdrawalReasonCode
}
func (l *MemberWithdrawal) GetWithdrawalReasonInputText () sql.NullString {
	return l.withdrawalReasonInputText
}
func (l *MemberWithdrawal) GetWithdrawalDatetime () df.Timestamp {
	return l.withdrawalDatetime
}
func (l *MemberWithdrawal) GetRegisterDatetime () df.Timestamp {
	return l.registerDatetime
}
func (l *MemberWithdrawal) GetRegisterProcess () string {
	return l.registerProcess
}
func (l *MemberWithdrawal) GetRegisterUser () string {
	return l.registerUser
}
func (l *MemberWithdrawal) GetUpdateDatetime () df.Timestamp {
	return l.updateDatetime
}
func (l *MemberWithdrawal) GetUpdateProcess () string {
	return l.updateProcess
}
func (l *MemberWithdrawal) GetUpdateUser () string {
	return l.updateUser
}

func (t *MemberWithdrawal) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 10)
	i[0] = &(t.memberId)
	i[1] = &(t.withdrawalReasonCode)
	i[2] = &(t.withdrawalReasonInputText)
	i[3] = &(t.withdrawalDatetime)
	i[4] = &(t.registerDatetime)
	i[5] = &(t.registerProcess)
	i[6] = &(t.registerUser)
	i[7] = &(t.updateDatetime)
	i[8] = &(t.updateProcess)
	i[9] = &(t.updateUser)
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
func (t *MemberWithdrawal) SetWithdrawalReasonCode(withdrawalReasonCode sql.NullString) {
	t.AddPropertyName("withdrawalReasonCode")
	t.withdrawalReasonCode = withdrawalReasonCode
}
func (t *MemberWithdrawal) SetWithdrawalReasonInputText(withdrawalReasonInputText sql.NullString) {
	t.AddPropertyName("withdrawalReasonInputText")
	t.withdrawalReasonInputText = withdrawalReasonInputText
}
func (t *MemberWithdrawal) SetWithdrawalDatetime(withdrawalDatetime df.Timestamp) {
	t.AddPropertyName("withdrawalDatetime")
	t.withdrawalDatetime = withdrawalDatetime
}
func (t *MemberWithdrawal) SetRegisterDatetime(registerDatetime df.Timestamp) {
	t.AddPropertyName("registerDatetime")
	t.registerDatetime = registerDatetime
}
func (t *MemberWithdrawal) SetRegisterProcess(registerProcess string) {
	t.AddPropertyName("registerProcess")
	t.registerProcess = registerProcess
}
func (t *MemberWithdrawal) SetRegisterUser(registerUser string) {
	t.AddPropertyName("registerUser")
	t.registerUser = registerUser
}
func (t *MemberWithdrawal) SetUpdateDatetime(updateDatetime df.Timestamp) {
	t.AddPropertyName("updateDatetime")
	t.updateDatetime = updateDatetime
}
func (t *MemberWithdrawal) SetUpdateProcess(updateProcess string) {
	t.AddPropertyName("updateProcess")
	t.updateProcess = updateProcess
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