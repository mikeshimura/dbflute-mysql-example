package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type MemberAddress struct {
	memberAddressId int64
	memberId int64
	validBeginDate df.MysqlDate
	validEndDate df.MysqlDate
	address string
	regionId int64
	registerDatetime df.MysqlTimestamp
	registerUser string
	updateDatetime df.MysqlTimestamp
	updateUser string
	versionNo int64
	df.BaseEntity
}

func CreateMemberAddress() *MemberAddress{
	memberAddress :=new(MemberAddress)
	memberAddress.SetUp()
	return memberAddress 
}

func (l *MemberAddress) GetMemberAddressId () int64 {
	return l.memberAddressId
}
func (l *MemberAddress) GetMemberId () int64 {
	return l.memberId
}
func (l *MemberAddress) GetValidBeginDate () df.MysqlDate {
	return l.validBeginDate
}
func (l *MemberAddress) GetValidEndDate () df.MysqlDate {
	return l.validEndDate
}
func (l *MemberAddress) GetAddress () string {
	return l.address
}
func (l *MemberAddress) GetRegionId () int64 {
	return l.regionId
}
func (l *MemberAddress) GetRegisterDatetime () df.MysqlTimestamp {
	return l.registerDatetime
}
func (l *MemberAddress) GetRegisterUser () string {
	return l.registerUser
}
func (l *MemberAddress) GetUpdateDatetime () df.MysqlTimestamp {
	return l.updateDatetime
}
func (l *MemberAddress) GetUpdateUser () string {
	return l.updateUser
}
func (l *MemberAddress) GetVersionNo () int64 {
	return l.versionNo
}

func (t *MemberAddress) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 11)
	i[0] = &(t.memberAddressId)
	i[1] = &(t.memberId)
	i[2] = &(t.validBeginDate)
	i[3] = &(t.validEndDate)
	i[4] = &(t.address)
	i[5] = &(t.regionId)
	i[6] = &(t.registerDatetime)
	i[7] = &(t.registerUser)
	i[8] = &(t.updateDatetime)
	i[9] = &(t.updateUser)
	i[10] = &(t.versionNo)
	return i
}


func (t *MemberAddress) AsTableDbName() string {
	return "MemberAddress"
}

func (t *MemberAddress) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("memberAddressId") == false {
            return false 
        }
        return true;
}
func (t *MemberAddress) SetMemberAddressId(memberAddressId int64) {
	t.AddPropertyName("memberAddressId")
	t.memberAddressId = memberAddressId
}
func (t *MemberAddress) SetMemberId(memberId int64) {
	t.AddPropertyName("memberId")
	t.memberId = memberId
}
func (t *MemberAddress) SetValidBeginDate(validBeginDate df.MysqlDate) {
	t.AddPropertyName("validBeginDate")
	t.validBeginDate = validBeginDate
}
func (t *MemberAddress) SetValidEndDate(validEndDate df.MysqlDate) {
	t.AddPropertyName("validEndDate")
	t.validEndDate = validEndDate
}
func (t *MemberAddress) SetAddress(address string) {
	t.AddPropertyName("address")
	t.address = address
}
func (t *MemberAddress) SetRegionId(regionId int64) {
	t.AddPropertyName("regionId")
	t.regionId = regionId
}
func (t *MemberAddress) SetRegisterDatetime(registerDatetime df.MysqlTimestamp) {
	t.AddPropertyName("registerDatetime")
	t.registerDatetime = registerDatetime
}
func (t *MemberAddress) SetRegisterUser(registerUser string) {
	t.AddPropertyName("registerUser")
	t.registerUser = registerUser
}
func (t *MemberAddress) SetUpdateDatetime(updateDatetime df.MysqlTimestamp) {
	t.AddPropertyName("updateDatetime")
	t.updateDatetime = updateDatetime
}
func (t *MemberAddress) SetUpdateUser(updateUser string) {
	t.AddPropertyName("updateUser")
	t.updateUser = updateUser
}
func (t *MemberAddress) SetVersionNo(versionNo int64) {
	t.AddPropertyName("versionNo")
	t.versionNo = versionNo
}

func (t *MemberAddress) SetUp(){
	
}
func (t *MemberAddress)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}