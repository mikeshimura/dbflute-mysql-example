package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type MemberService struct {
	memberServiceId int64
	memberId int64
	servicePointCount int64
	serviceRankCode string
	registerDatetime df.MysqlTimestamp
	registerUser string
	updateDatetime df.MysqlTimestamp
	updateUser string
	versionNo int64
	df.BaseEntity
Member_R  *Member

ServiceRank_R  *ServiceRank

}

func CreateMemberService() *MemberService{
	memberService :=new(MemberService)
	memberService.SetUp()
	return memberService 
}

func (l *MemberService) GetMemberServiceId () int64 {
	return l.memberServiceId
}
func (l *MemberService) GetMemberId () int64 {
	return l.memberId
}
func (l *MemberService) GetServicePointCount () int64 {
	return l.servicePointCount
}
func (l *MemberService) GetServiceRankCode () string {
	return l.serviceRankCode
}
func (l *MemberService) GetRegisterDatetime () df.MysqlTimestamp {
	return l.registerDatetime
}
func (l *MemberService) GetRegisterUser () string {
	return l.registerUser
}
func (l *MemberService) GetUpdateDatetime () df.MysqlTimestamp {
	return l.updateDatetime
}
func (l *MemberService) GetUpdateUser () string {
	return l.updateUser
}
func (l *MemberService) GetVersionNo () int64 {
	return l.versionNo
}

func (t *MemberService) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 9)
	i[0] = &(t.memberServiceId)
	i[1] = &(t.memberId)
	i[2] = &(t.servicePointCount)
	i[3] = &(t.serviceRankCode)
	i[4] = &(t.registerDatetime)
	i[5] = &(t.registerUser)
	i[6] = &(t.updateDatetime)
	i[7] = &(t.updateUser)
	i[8] = &(t.versionNo)
	return i
}


func (t *MemberService) AsTableDbName() string {
	return "MemberService"
}

func (t *MemberService) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("memberServiceId") == false {
            return false 
        }
        return true;
}
func (t *MemberService) SetMemberServiceId(memberServiceId int64) {
	t.AddPropertyName("memberServiceId")
	t.memberServiceId = memberServiceId
}
func (t *MemberService) SetMemberId(memberId int64) {
	t.AddPropertyName("memberId")
	t.memberId = memberId
}
func (t *MemberService) SetServicePointCount(servicePointCount int64) {
	t.AddPropertyName("servicePointCount")
	t.servicePointCount = servicePointCount
}
func (t *MemberService) SetServiceRankCode(serviceRankCode string) {
	t.AddPropertyName("serviceRankCode")
	t.serviceRankCode = serviceRankCode
}
func (t *MemberService) SetRegisterDatetime(registerDatetime df.MysqlTimestamp) {
	t.AddPropertyName("registerDatetime")
	t.registerDatetime = registerDatetime
}
func (t *MemberService) SetRegisterUser(registerUser string) {
	t.AddPropertyName("registerUser")
	t.registerUser = registerUser
}
func (t *MemberService) SetUpdateDatetime(updateDatetime df.MysqlTimestamp) {
	t.AddPropertyName("updateDatetime")
	t.updateDatetime = updateDatetime
}
func (t *MemberService) SetUpdateUser(updateUser string) {
	t.AddPropertyName("updateUser")
	t.updateUser = updateUser
}
func (t *MemberService) SetVersionNo(versionNo int64) {
	t.AddPropertyName("versionNo")
	t.versionNo = versionNo
}
func (t *MemberService) GetMember_R() *Member{
	return t.Member_R
}
func (t *MemberService) SetMember_R(value *Member) {
    t.Member_R = value
}
func (t *MemberService) GetServiceRank_R() *ServiceRank{
	return t.ServiceRank_R
}
func (t *MemberService) SetServiceRank_R(value *ServiceRank) {
    t.ServiceRank_R = value
}
func (t *MemberService) SetUp(){
	
}
func (t *MemberService)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}