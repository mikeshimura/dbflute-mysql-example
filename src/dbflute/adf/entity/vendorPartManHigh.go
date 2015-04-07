package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorPartManHigh struct {
	partManId int64
	partManName string
	partManPoint int64
	partManDate df.NullDate
	df.BaseEntity
}

func CreateVendorPartManHigh() *VendorPartManHigh{
	vendorPartManHigh :=new(VendorPartManHigh)
	vendorPartManHigh.SetUp()
	return vendorPartManHigh 
}

func (l *VendorPartManHigh) GetPartManId () int64 {
	return l.partManId
}
func (l *VendorPartManHigh) GetPartManName () string {
	return l.partManName
}
func (l *VendorPartManHigh) GetPartManPoint () int64 {
	return l.partManPoint
}
func (l *VendorPartManHigh) GetPartManDate () df.NullDate {
	return l.partManDate
}

func (t *VendorPartManHigh) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 4)
	i[0] = &(t.partManId)
	i[1] = &(t.partManName)
	i[2] = &(t.partManPoint)
	i[3] = &(t.partManDate)
	return i
}


func (t *VendorPartManHigh) AsTableDbName() string {
	return "VendorPartManHigh"
}

func (t *VendorPartManHigh) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("partManId") == false {
            return false 
        }
        return true;
}
func (t *VendorPartManHigh) SetPartManId(partManId int64) {
	t.AddPropertyName("partManId")
	t.partManId = partManId
}
func (t *VendorPartManHigh) SetPartManName(partManName string) {
	t.AddPropertyName("partManName")
	t.partManName = partManName
}
func (t *VendorPartManHigh) SetPartManPoint(partManPoint int64) {
	t.AddPropertyName("partManPoint")
	t.partManPoint = partManPoint
}
func (t *VendorPartManHigh) SetPartManDate(partManDate df.NullDate) {
	t.AddPropertyName("partManDate")
	t.partManDate = partManDate
}

func (t *VendorPartManHigh) SetUp(){
	
}
func (t *VendorPartManHigh)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}