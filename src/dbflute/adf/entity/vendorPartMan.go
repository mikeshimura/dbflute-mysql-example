package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorPartMan struct {
	partManId int64
	partManName string
	partManPoint int64
	partManDate df.NullDate
	df.BaseEntity
}

func CreateVendorPartMan() *VendorPartMan{
	vendorPartMan :=new(VendorPartMan)
	vendorPartMan.SetUp()
	return vendorPartMan 
}

func (l *VendorPartMan) GetPartManId () int64 {
	return l.partManId
}
func (l *VendorPartMan) GetPartManName () string {
	return l.partManName
}
func (l *VendorPartMan) GetPartManPoint () int64 {
	return l.partManPoint
}
func (l *VendorPartMan) GetPartManDate () df.NullDate {
	return l.partManDate
}

func (t *VendorPartMan) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 4)
	i[0] = &(t.partManId)
	i[1] = &(t.partManName)
	i[2] = &(t.partManPoint)
	i[3] = &(t.partManDate)
	return i
}


func (t *VendorPartMan) AsTableDbName() string {
	return "VendorPartMan"
}

func (t *VendorPartMan) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("partManId") == false {
            return false 
        }
        return true;
}
func (t *VendorPartMan) SetPartManId(partManId int64) {
	t.AddPropertyName("partManId")
	t.partManId = partManId
}
func (t *VendorPartMan) SetPartManName(partManName string) {
	t.AddPropertyName("partManName")
	t.partManName = partManName
}
func (t *VendorPartMan) SetPartManPoint(partManPoint int64) {
	t.AddPropertyName("partManPoint")
	t.partManPoint = partManPoint
}
func (t *VendorPartMan) SetPartManDate(partManDate df.NullDate) {
	t.AddPropertyName("partManDate")
	t.partManDate = partManDate
}

func (t *VendorPartMan) SetUp(){
	
}
func (t *VendorPartMan)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}