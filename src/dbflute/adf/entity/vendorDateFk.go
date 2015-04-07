package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorDateFk struct {
	barId int64
	barDate df.Date
	df.BaseEntity
}

func CreateVendorDateFk() *VendorDateFk{
	vendorDateFk :=new(VendorDateFk)
	vendorDateFk.SetUp()
	return vendorDateFk 
}

func (l *VendorDateFk) GetBarId () int64 {
	return l.barId
}
func (l *VendorDateFk) GetBarDate () df.Date {
	return l.barDate
}

func (t *VendorDateFk) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 2)
	i[0] = &(t.barId)
	i[1] = &(t.barDate)
	return i
}


func (t *VendorDateFk) AsTableDbName() string {
	return "VendorDateFk"
}

func (t *VendorDateFk) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("barId") == false {
            return false 
        }
        return true;
}
func (t *VendorDateFk) SetBarId(barId int64) {
	t.AddPropertyName("barId")
	t.barId = barId
}
func (t *VendorDateFk) SetBarDate(barDate df.Date) {
	t.AddPropertyName("barDate")
	t.barDate = barDate
}

func (t *VendorDateFk) SetUp(){
	
}
func (t *VendorDateFk)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}