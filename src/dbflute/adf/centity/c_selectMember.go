package centity

import (
	"github.com/mikeshimura/dbflute/df"
	"database/sql"
)

type C_SelectMember  struct {
	MemberId sql.NullInt64
	MemberName df.NullString
	MemberAccount df.NullString
	Birthdate df.MysqlNullDate
	FormalizedDatetime df.MysqlNullTimestamp
	MemberStatusCode df.NullString
	MemberStatusName df.NullString
	Description df.NullString
	df.BaseEntity
}
func (l *C_SelectMember) GetMemberId () sql.NullInt64 {
	return l.MemberId
}
func (l *C_SelectMember) GetMemberName () df.NullString {
	return l.MemberName
}
func (l *C_SelectMember) GetMemberAccount () df.NullString {
	return l.MemberAccount
}
func (l *C_SelectMember) GetBirthdate () df.MysqlNullDate {
	return l.Birthdate
}
func (l *C_SelectMember) GetFormalizedDatetime () df.MysqlNullTimestamp {
	return l.FormalizedDatetime
}
func (l *C_SelectMember) GetMemberStatusCode () df.NullString {
	return l.MemberStatusCode
}
func (l *C_SelectMember) GetMemberStatusName () df.NullString {
	return l.MemberStatusName
}
func (l *C_SelectMember) GetDescription () df.NullString {
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
func (t *C_SelectMember) SetMemberName(memberName df.NullString) {
	t.AddPropertyName("memberName")
	t.MemberName = memberName
}
func (t *C_SelectMember) SetMemberAccount(memberAccount df.NullString) {
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
func (t *C_SelectMember) SetMemberStatusCode(memberStatusCode df.NullString) {
	t.AddPropertyName("memberStatusCode")
	t.MemberStatusCode = memberStatusCode
}
func (t *C_SelectMember) SetMemberStatusName(memberStatusName df.NullString) {
	t.AddPropertyName("memberStatusName")
	t.MemberStatusName = memberStatusName
}
func (t *C_SelectMember) SetDescription(description df.NullString) {
	t.AddPropertyName("description")
	t.Description = description
}


func (t *C_SelectMember) SetUp(){
}


func (t *C_SelectMember)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}