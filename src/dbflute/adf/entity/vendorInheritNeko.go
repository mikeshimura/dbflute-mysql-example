package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorInheritNeko struct {
	inuId int64
	inuName string
	inuDate df.NullDate
	nekoId int64
	nekoName string
	nekoDate df.NullDate
	df.BaseEntity
}

func CreateVendorInheritNeko() *VendorInheritNeko{
	vendorInheritNeko :=new(VendorInheritNeko)
	vendorInheritNeko.SetUp()
	return vendorInheritNeko 
}

func (l *VendorInheritNeko) GetInuId () int64 {
	return l.inuId
}
func (l *VendorInheritNeko) GetInuName () string {
	return l.inuName
}
func (l *VendorInheritNeko) GetInuDate () df.NullDate {
	return l.inuDate
}
func (l *VendorInheritNeko) GetNekoId () int64 {
	return l.nekoId
}
func (l *VendorInheritNeko) GetNekoName () string {
	return l.nekoName
}
func (l *VendorInheritNeko) GetNekoDate () df.NullDate {
	return l.nekoDate
}

func (t *VendorInheritNeko) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 6)
	i[0] = &(t.inuId)
	i[1] = &(t.inuName)
	i[2] = &(t.inuDate)
	i[3] = &(t.nekoId)
	i[4] = &(t.nekoName)
	i[5] = &(t.nekoDate)
	return i
}


func (t *VendorInheritNeko) AsTableDbName() string {
	return "VendorInheritNeko"
}

func (t *VendorInheritNeko) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("nekoId") == false {
            return false 
        }
        return true;
}
func (t *VendorInheritNeko) SetInuId(inuId int64) {
	t.AddPropertyName("inuId")
	t.inuId = inuId
}
func (t *VendorInheritNeko) SetInuName(inuName string) {
	t.AddPropertyName("inuName")
	t.inuName = inuName
}
func (t *VendorInheritNeko) SetInuDate(inuDate df.NullDate) {
	t.AddPropertyName("inuDate")
	t.inuDate = inuDate
}
func (t *VendorInheritNeko) SetNekoId(nekoId int64) {
	t.AddPropertyName("nekoId")
	t.nekoId = nekoId
}
func (t *VendorInheritNeko) SetNekoName(nekoName string) {
	t.AddPropertyName("nekoName")
	t.nekoName = nekoName
}
func (t *VendorInheritNeko) SetNekoDate(nekoDate df.NullDate) {
	t.AddPropertyName("nekoDate")
	t.nekoDate = nekoDate
}

func (t *VendorInheritNeko) SetUp(){
	
}
func (t *VendorInheritNeko)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}