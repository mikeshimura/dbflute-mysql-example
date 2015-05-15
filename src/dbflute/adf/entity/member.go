package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type Member struct {
	memberId int64
	memberName string
	memberAccount string
	memberStatusCode string
	formalizedDatetime df.MysqlNullTimestamp
	birthdate df.MysqlNullDate
	registerDatetime df.MysqlTimestamp
	registerUser string
	updateDatetime df.MysqlTimestamp
	updateUser string
	versionNo int64
	df.BaseEntity
MemberStatus_R  *MemberStatus

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
func (l *Member) GetFormalizedDatetime () df.MysqlNullTimestamp {
	return l.formalizedDatetime
}
func (l *Member) GetBirthdate () df.MysqlNullDate {
	return l.birthdate
}
func (l *Member) GetRegisterDatetime () df.MysqlTimestamp {
	return l.registerDatetime
}
func (l *Member) GetRegisterUser () string {
	return l.registerUser
}
func (l *Member) GetUpdateDatetime () df.MysqlTimestamp {
	return l.updateDatetime
}
func (l *Member) GetUpdateUser () string {
	return l.updateUser
}
func (l *Member) GetVersionNo () int64 {
	return l.versionNo
}

func (t *Member) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 11)
	i[0] = &(t.memberId)
	i[1] = &(t.memberName)
	i[2] = &(t.memberAccount)
	i[3] = &(t.memberStatusCode)
	i[4] = &(t.formalizedDatetime)
	i[5] = &(t.birthdate)
	i[6] = &(t.registerDatetime)
	i[7] = &(t.registerUser)
	i[8] = &(t.updateDatetime)
	i[9] = &(t.updateUser)
	i[10] = &(t.versionNo)
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
func (t *Member) SetFormalizedDatetime(formalizedDatetime df.MysqlNullTimestamp) {
	t.AddPropertyName("formalizedDatetime")
	t.formalizedDatetime = formalizedDatetime
}
func (t *Member) SetBirthdate(birthdate df.MysqlNullDate) {
	t.AddPropertyName("birthdate")
	t.birthdate = birthdate
}
func (t *Member) SetRegisterDatetime(registerDatetime df.MysqlTimestamp) {
	t.AddPropertyName("registerDatetime")
	t.registerDatetime = registerDatetime
}
func (t *Member) SetRegisterUser(registerUser string) {
	t.AddPropertyName("registerUser")
	t.registerUser = registerUser
}
func (t *Member) SetUpdateDatetime(updateDatetime df.MysqlTimestamp) {
	t.AddPropertyName("updateDatetime")
	t.updateDatetime = updateDatetime
}
func (t *Member) SetUpdateUser(updateUser string) {
	t.AddPropertyName("updateUser")
	t.updateUser = updateUser
}
func (t *Member) SetVersionNo(versionNo int64) {
	t.AddPropertyName("versionNo")
	t.versionNo = versionNo
}
func (t *Member) GetMemberStatus_R() *MemberStatus{
	return t.MemberStatus_R
}
func (t *Member) SetMemberStatus_R(value *MemberStatus) {
    t.MemberStatus_R = value
}
func (t *Member) SetUp(){
	
}
func (t *Member)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}