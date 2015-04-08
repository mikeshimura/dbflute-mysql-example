package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type MemberLogin struct {
	memberLoginId int64
	memberId int64
	loginDatetime df.MysqlTimestamp
	mobileLoginFlg int64
	loginMemberStatusCode string
	df.BaseEntity
}

func CreateMemberLogin() *MemberLogin{
	memberLogin :=new(MemberLogin)
	memberLogin.SetUp()
	return memberLogin 
}

func (l *MemberLogin) GetMemberLoginId () int64 {
	return l.memberLoginId
}
func (l *MemberLogin) GetMemberId () int64 {
	return l.memberId
}
func (l *MemberLogin) GetLoginDatetime () df.MysqlTimestamp {
	return l.loginDatetime
}
func (l *MemberLogin) GetMobileLoginFlg () int64 {
	return l.mobileLoginFlg
}
func (l *MemberLogin) GetLoginMemberStatusCode () string {
	return l.loginMemberStatusCode
}

func (t *MemberLogin) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 5)
	i[0] = &(t.memberLoginId)
	i[1] = &(t.memberId)
	i[2] = &(t.loginDatetime)
	i[3] = &(t.mobileLoginFlg)
	i[4] = &(t.loginMemberStatusCode)
	return i
}


func (t *MemberLogin) AsTableDbName() string {
	return "MemberLogin"
}

func (t *MemberLogin) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("memberLoginId") == false {
            return false 
        }
        return true;
}
func (t *MemberLogin) SetMemberLoginId(memberLoginId int64) {
	t.AddPropertyName("memberLoginId")
	t.memberLoginId = memberLoginId
}
func (t *MemberLogin) SetMemberId(memberId int64) {
	t.AddPropertyName("memberId")
	t.memberId = memberId
}
func (t *MemberLogin) SetLoginDatetime(loginDatetime df.MysqlTimestamp) {
	t.AddPropertyName("loginDatetime")
	t.loginDatetime = loginDatetime
}
func (t *MemberLogin) SetMobileLoginFlg(mobileLoginFlg int64) {
	t.AddPropertyName("mobileLoginFlg")
	t.mobileLoginFlg = mobileLoginFlg
}
func (t *MemberLogin) SetLoginMemberStatusCode(loginMemberStatusCode string) {
	t.AddPropertyName("loginMemberStatusCode")
	t.loginMemberStatusCode = loginMemberStatusCode
}

func (t *MemberLogin) SetUp(){
	
}
func (t *MemberLogin)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}