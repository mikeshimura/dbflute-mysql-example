package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorInheritInu struct {
	inuId int64
	inuName string
	inuDate df.NullDate
	df.BaseEntity
}

func CreateVendorInheritInu() *VendorInheritInu{
	vendorInheritInu :=new(VendorInheritInu)
	vendorInheritInu.SetUp()
	return vendorInheritInu 
}

func (l *VendorInheritInu) GetInuId () int64 {
	return l.inuId
}
func (l *VendorInheritInu) GetInuName () string {
	return l.inuName
}
func (l *VendorInheritInu) GetInuDate () df.NullDate {
	return l.inuDate
}

func (t *VendorInheritInu) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 3)
	i[0] = &(t.inuId)
	i[1] = &(t.inuName)
	i[2] = &(t.inuDate)
	return i
}


func (t *VendorInheritInu) AsTableDbName() string {
	return "VendorInheritInu"
}

func (t *VendorInheritInu) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("inuId") == false {
            return false 
        }
        return true;
}
func (t *VendorInheritInu) SetInuId(inuId int64) {
	t.AddPropertyName("inuId")
	t.inuId = inuId
}
func (t *VendorInheritInu) SetInuName(inuName string) {
	t.AddPropertyName("inuName")
	t.inuName = inuName
}
func (t *VendorInheritInu) SetInuDate(inuDate df.NullDate) {
	t.AddPropertyName("inuDate")
	t.inuDate = inuDate
}

func (t *VendorInheritInu) SetUp(){
	
}
func (t *VendorInheritInu)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}