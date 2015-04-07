package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type MemberStatus struct {
	memberStatusCode string
	memberStatusName string
	description string
	displayOrder int64
	df.BaseEntity
}

func CreateMemberStatus() *MemberStatus{
	memberStatus :=new(MemberStatus)
	memberStatus.SetUp()
	return memberStatus 
}

func (l *MemberStatus) GetMemberStatusCode () string {
	return l.memberStatusCode
}
func (l *MemberStatus) GetMemberStatusName () string {
	return l.memberStatusName
}
func (l *MemberStatus) GetDescription () string {
	return l.description
}
func (l *MemberStatus) GetDisplayOrder () int64 {
	return l.displayOrder
}

func (t *MemberStatus) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 4)
	i[0] = &(t.memberStatusCode)
	i[1] = &(t.memberStatusName)
	i[2] = &(t.description)
	i[3] = &(t.displayOrder)
	return i
}


func (t *MemberStatus) AsTableDbName() string {
	return "MemberStatus"
}

func (t *MemberStatus) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("memberStatusCode") == false {
            return false 
        }
        return true;
}
func (t *MemberStatus) SetMemberStatusCode(memberStatusCode string) {
	t.AddPropertyName("memberStatusCode")
	t.memberStatusCode = memberStatusCode
}
func (t *MemberStatus) SetMemberStatusName(memberStatusName string) {
	t.AddPropertyName("memberStatusName")
	t.memberStatusName = memberStatusName
}
func (t *MemberStatus) SetDescription(description string) {
	t.AddPropertyName("description")
	t.description = description
}
func (t *MemberStatus) SetDisplayOrder(displayOrder int64) {
	t.AddPropertyName("displayOrder")
	t.displayOrder = displayOrder
}

func (t *MemberStatus) SetUp(){
	
}
func (t *MemberStatus)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}