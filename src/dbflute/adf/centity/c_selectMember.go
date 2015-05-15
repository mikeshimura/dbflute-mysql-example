package centity

import (
	"github.com/mikeshimura/dbflute/df"
	"database/sql"
)

type C_SelectMember  struct {
	MemberId sql.NullInt64
	MemberName sql.NullString
	MemberAccount sql.NullString
	Birthdate df.MysqlNullDate
	FormalizedDatetime df.MysqlNullTimestamp
	MemberStatusCode sql.NullString
	MemberStatusName sql.NullString
	Description sql.NullString
	df.BaseEntity
}
func (l *C_SelectMember) GetMemberId () sql.NullInt64 {
	return l.MemberId
}
func (l *C_SelectMember) GetMemberName () sql.NullString {
	return l.MemberName
}
func (l *C_SelectMember) GetMemberAccount () sql.NullString {
	return l.MemberAccount
}
func (l *C_SelectMember) GetBirthdate () df.MysqlNullDate {
	return l.Birthdate
}
func (l *C_SelectMember) GetFormalizedDatetime () df.MysqlNullTimestamp {
	return l.FormalizedDatetime
}
func (l *C_SelectMember) GetMemberStatusCode () sql.NullString {
	return l.MemberStatusCode
}
func (l *C_SelectMember) GetMemberStatusName () sql.NullString {
	return l.MemberStatusName
}
func (l *C_SelectMember) GetDescription () sql.NullString {
	return l.Description
}


func (t *C_SelectMember) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 8)
	i[0] = &(t.MemberId)
	i[1] = &(t.MemberName)
	i[2] = &(t.MemberAccount)
	i[3] = &(t.Birthdate)
	i[4] = &(t.FormalizedDatetime)
	i[5] = &(t.MemberStatusCode)
	i[6] = &(t.MemberStatusName)
	i[7] = &(t.Description)
	return i
}

func (t *C_SelectMember) AsTableDbName() string {
	return "c_SelectMember"
}

func (t *C_SelectMember) HasPrimaryKeyValue() bool{
        return false;
}
func (t *C_SelectMember) SetMemberId(memberId sql.NullInt64) {
	t.AddPropertyName("memberId")
	t.MemberId = memberId
}
func (t *C_SelectMember) SetMemberName(memberName sql.NullString) {
	t.AddPropertyName("memberName")
	t.MemberName = memberName
}
func (t *C_SelectMember) SetMemberAccount(memberAccount sql.NullString) {
	t.AddPropertyName("memberAccount")
	t.MemberAccount = memberAccount
}
func (t *C_SelectMember) SetBirthdate(birthdate df.MysqlNullDate) {
	t.AddPropertyName("birthdate")
	t.Birthdate = birthdate
}
func (t *C_SelectMember) SetFormalizedDatetime(formalizedDatetime df.MysqlNullTimestamp) {
	t.AddPropertyName("formalizedDatetime")
	t.FormalizedDatetime = formalizedDatetime
}
func (t *C_SelectMember) SetMemberStatusCode(memberStatusCode sql.NullString) {
	t.AddPropertyName("memberStatusCode")
	t.MemberStatusCode = memberStatusCode
}
func (t *C_SelectMember) SetMemberStatusName(memberStatusName sql.NullString) {
	t.AddPropertyName("memberStatusName")
	t.MemberStatusName = memberStatusName
}
func (t *C_SelectMember) SetDescription(description sql.NullString) {
	t.AddPropertyName("description")
	t.Description = description
}


func (t *C_SelectMember) SetUp(){
}


func (t *C_SelectMember)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}