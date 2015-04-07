package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type Member struct {
	memberId int64
	memberName string
	memberAccount string
	memberStatusCode string
	formalizedDatetime df.NullTimestamp
	birthdate df.NullDate
	registerDatetime df.Timestamp
	registerUser string
	registerProcess string
	updateDatetime df.Timestamp
	updateUser string
	updateProcess string
	versionNo int64
	df.BaseEntity
}

func CreateMember() *Member{
	member :=new(Member)
	member.SetUp()
	return member 
}

func (l *Member) GetMemberId () int64 {
	return l.memberId
}
func (l *Member) GetMemberName () string {
	return l.memberName
}
func (l *Member) GetMemberAccount () string {
	return l.memberAccount
}
func (l *Member) GetMemberStatusCode () string {
	return l.memberStatusCode
}
func (l *Member) GetFormalizedDatetime () df.NullTimestamp {
	return l.formalizedDatetime
}
func (l *Member) GetBirthdate () df.NullDate {
	return l.birthdate
}
func (l *Member) GetRegisterDatetime () df.Timestamp {
	return l.registerDatetime
}
func (l *Member) GetRegisterUser () string {
	return l.registerUser
}
func (l *Member) GetRegisterProcess () string {
	return l.registerProcess
}
func (l *Member) GetUpdateDatetime () df.Timestamp {
	return l.updateDatetime
}
func (l *Member) GetUpdateUser () string {
	return l.updateUser
}
func (l *Member) GetUpdateProcess () string {
	return l.updateProcess
}
func (l *Member) GetVersionNo () int64 {
	return l.versionNo
}

func (t *Member) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 13)
	i[0] = &(t.memberId)
	i[1] = &(t.memberName)
	i[2] = &(t.memberAccount)
	i[3] = &(t.memberStatusCode)
	i[4] = &(t.formalizedDatetime)
	i[5] = &(t.birthdate)
	i[6] = &(t.registerDatetime)
	i[7] = &(t.registerUser)
	i[8] = &(t.registerProcess)
	i[9] = &(t.updateDatetime)
	i[10] = &(t.updateUser)
	i[11] = &(t.updateProcess)
	i[12] = &(t.versionNo)
	return i
}


func (t *Member) AsTableDbName() string {
	return "Member"
}

func (t *Member) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("memberId") == false {
            return false 
        }
        return true;
}
func (t *Member) SetMemberId(memberId int64) {
	t.AddPropertyName("memberId")
	t.memberId = memberId
}
func (t *Member) SetMemberName(memberName string) {
	t.AddPropertyName("memberName")
	t.memberName = memberName
}
func (t *Member) SetMemberAccount(memberAccount string) {
	t.AddPropertyName("memberAccount")
	t.memberAccount = memberAccount
}
func (t *Member) SetMemberStatusCode(memberStatusCode string) {
	t.AddPropertyName("memberStatusCode")
	t.memberStatusCode = memberStatusCode
}
func (t *Member) SetFormalizedDatetime(formalizedDatetime df.NullTimestamp) {
	t.AddPropertyName("formalizedDatetime")
	t.formalizedDatetime = formalizedDatetime
}
func (t *Member) SetBirthdate(birthdate df.NullDate) {
	t.AddPropertyName("birthdate")
	t.birthdate = birthdate
}
func (t *Member) SetRegisterDatetime(registerDatetime df.Timestamp) {
	t.AddPropertyName("registerDatetime")
	t.registerDatetime = registerDatetime
}
func (t *Member) SetRegisterUser(registerUser string) {
	t.AddPropertyName("registerUser")
	t.registerUser = registerUser
}
func (t *Member) SetRegisterProcess(registerProcess string) {
	t.AddPropertyName("registerProcess")
	t.registerProcess = registerProcess
}
func (t *Member) SetUpdateDatetime(updateDatetime df.Timestamp) {
	t.AddPropertyName("updateDatetime")
	t.updateDatetime = updateDatetime
}
func (t *Member) SetUpdateUser(updateUser string) {
	t.AddPropertyName("updateUser")
	t.updateUser = updateUser
}
func (t *Member) SetUpdateProcess(updateProcess string) {
	t.AddPropertyName("updateProcess")
	t.updateProcess = updateProcess
}
func (t *Member) SetVersionNo(versionNo int64) {
	t.AddPropertyName("versionNo")
	t.versionNo = versionNo
}

func (t *Member) SetUp(){
	
}
func (t *Member)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}