package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorDatePk struct {
	fooDate df.Date
	fooName string
	df.BaseEntity
}

func CreateVendorDatePk() *VendorDatePk{
	vendorDatePk :=new(VendorDatePk)
	vendorDatePk.SetUp()
	return vendorDatePk 
}

func (l *VendorDatePk) GetFooDate () df.Date {
	return l.fooDate
}
func (l *VendorDatePk) GetFooName () string {
	return l.fooName
}

func (t *VendorDatePk) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 2)
	i[0] = &(t.fooDate)
	i[1] = &(t.fooName)
	return i
}


func (t *VendorDatePk) AsTableDbName() string {
	return "VendorDatePk"
}

func (t *VendorDatePk) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("fooDate") == false {
            return false 
        }
        return true;
}
func (t *VendorDatePk) SetFooDate(fooDate df.Date) {
	t.AddPropertyName("fooDate")
	t.fooDate = fooDate
}
func (t *VendorDatePk) SetFooName(fooName string) {
	t.AddPropertyName("fooName")
	t.fooName = fooName
}

func (t *VendorDatePk) SetUp(){
	
}
func (t *VendorDatePk)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}