package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type MemberSecurity struct {
	memberId int64
	loginPassword string
	reminderQuestion string
	reminderAnswer string
	reminderUseCount int64
	registerDatetime df.Timestamp
	registerProcess string
	registerUser string
	updateDatetime df.Timestamp
	updateProcess string
	updateUser string
	versionNo int64
	df.BaseEntity
}

func CreateMemberSecurity() *MemberSecurity{
	memberSecurity :=new(MemberSecurity)
	memberSecurity.SetUp()
	return memberSecurity 
}

func (l *MemberSecurity) GetMemberId () int64 {
	return l.memberId
}
func (l *MemberSecurity) GetLoginPassword () string {
	return l.loginPassword
}
func (l *MemberSecurity) GetReminderQuestion () string {
	return l.reminderQuestion
}
func (l *MemberSecurity) GetReminderAnswer () string {
	return l.reminderAnswer
}
func (l *MemberSecurity) GetReminderUseCount () int64 {
	return l.reminderUseCount
}
func (l *MemberSecurity) GetRegisterDatetime () df.Timestamp {
	return l.registerDatetime
}
func (l *MemberSecurity) GetRegisterProcess () string {
	return l.registerProcess
}
func (l *MemberSecurity) GetRegisterUser () string {
	return l.registerUser
}
func (l *MemberSecurity) GetUpdateDatetime () df.Timestamp {
	return l.updateDatetime
}
func (l *MemberSecurity) GetUpdateProcess () string {
	return l.updateProcess
}
func (l *MemberSecurity) GetUpdateUser () string {
	return l.updateUser
}
func (l *MemberSecurity) GetVersionNo () int64 {
	return l.versionNo
}

func (t *MemberSecurity) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 12)
	i[0] = &(t.memberId)
	i[1] = &(t.loginPassword)
	i[2] = &(t.reminderQuestion)
	i[3] = &(t.reminderAnswer)
	i[4] = &(t.reminderUseCount)
	i[5] = &(t.registerDatetime)
	i[6] = &(t.registerProcess)
	i[7] = &(t.registerUser)
	i[8] = &(t.updateDatetime)
	i[9] = &(t.updateProcess)
	i[10] = &(t.updateUser)
	i[11] = &(t.versionNo)
	return i
}


func (t *MemberSecurity) AsTableDbName() string {
	return "MemberSecurity"
}

func (t *MemberSecurity) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("memberId") == false {
            return false 
        }
        return true;
}
func (t *MemberSecurity) SetMemberId(memberId int64) {
	t.AddPropertyName("memberId")
	t.memberId = memberId
}
func (t *MemberSecurity) SetLoginPassword(loginPassword string) {
	t.AddPropertyName("loginPassword")
	t.loginPassword = loginPassword
}
func (t *MemberSecurity) SetReminderQuestion(reminderQuestion string) {
	t.AddPropertyName("reminderQuestion")
	t.reminderQuestion = reminderQuestion
}
func (t *MemberSecurity) SetReminderAnswer(reminderAnswer string) {
	t.AddPropertyName("reminderAnswer")
	t.reminderAnswer = reminderAnswer
}
func (t *MemberSecurity) SetReminderUseCount(reminderUseCount int64) {
	t.AddPropertyName("reminderUseCount")
	t.reminderUseCount = reminderUseCount
}
func (t *MemberSecurity) SetRegisterDatetime(registerDatetime df.Timestamp) {
	t.AddPropertyName("registerDatetime")
	t.registerDatetime = registerDatetime
}
func (t *MemberSecurity) SetRegisterProcess(registerProcess string) {
	t.AddPropertyName("registerProcess")
	t.registerProcess = registerProcess
}
func (t *MemberSecurity) SetRegisterUser(registerUser string) {
	t.AddPropertyName("registerUser")
	t.registerUser = registerUser
}
func (t *MemberSecurity) SetUpdateDatetime(updateDatetime df.Timestamp) {
	t.AddPropertyName("updateDatetime")
	t.updateDatetime = updateDatetime
}
func (t *MemberSecurity) SetUpdateProcess(updateProcess string) {
	t.AddPropertyName("updateProcess")
	t.updateProcess = updateProcess
}
func (t *MemberSecurity) SetUpdateUser(updateUser string) {
	t.AddPropertyName("updateUser")
	t.updateUser = updateUser
}
func (t *MemberSecurity) SetVersionNo(versionNo int64) {
	t.AddPropertyName("versionNo")
	t.versionNo = versionNo
}

func (t *MemberSecurity) SetUp(){
	
}
func (t *MemberSecurity)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}