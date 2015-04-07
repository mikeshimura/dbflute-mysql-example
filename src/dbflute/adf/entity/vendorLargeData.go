package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorLargeData struct {
	largeDataId int64
	stringIndex string
	stringNoIndex string
	stringUniqueIndex string
	intflgIndex int64
	numericIntegerIndex df.Numeric
	numericIntegerNoIndex df.Numeric
	df.BaseEntity
}

func CreateVendorLargeData() *VendorLargeData{
	vendorLargeData :=new(VendorLargeData)
	vendorLargeData.SetUp()
	return vendorLargeData 
}

func (l *VendorLargeData) GetLargeDataId () int64 {
	return l.largeDataId
}
func (l *VendorLargeData) GetStringIndex () string {
	return l.stringIndex
}
func (l *VendorLargeData) GetStringNoIndex () string {
	return l.stringNoIndex
}
func (l *VendorLargeData) GetStringUniqueIndex () string {
	return l.stringUniqueIndex
}
func (l *VendorLargeData) GetIntflgIndex () int64 {
	return l.intflgIndex
}
func (l *VendorLargeData) GetNumericIntegerIndex () df.Numeric {
	return l.numericIntegerIndex
}
func (l *VendorLargeData) GetNumericIntegerNoIndex () df.Numeric {
	return l.numericIntegerNoIndex
}

func (t *VendorLargeData) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 7)
	i[0] = &(t.largeDataId)
	i[1] = &(t.stringIndex)
	i[2] = &(t.stringNoIndex)
	i[3] = &(t.stringUniqueIndex)
	i[4] = &(t.intflgIndex)
	i[5] = &(t.numericIntegerIndex)
	i[6] = &(t.numericIntegerNoIndex)
	return i
}


func (t *VendorLargeData) AsTableDbName() string {
	return "VendorLargeData"
}

func (t *VendorLargeData) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("largeDataId") == false {
            return false 
        }
        return true;
}
func (t *VendorLargeData) SetLargeDataId(largeDataId int64) {
	t.AddPropertyName("largeDataId")
	t.largeDataId = largeDataId
}
func (t *VendorLargeData) SetStringIndex(stringIndex string) {
	t.AddPropertyName("stringIndex")
	t.stringIndex = stringIndex
}
func (t *VendorLargeData) SetStringNoIndex(stringNoIndex string) {
	t.AddPropertyName("stringNoIndex")
	t.stringNoIndex = stringNoIndex
}
func (t *VendorLargeData) SetStringUniqueIndex(stringUniqueIndex string) {
	t.AddPropertyName("stringUniqueIndex")
	t.stringUniqueIndex = stringUniqueIndex
}
func (t *VendorLargeData) SetIntflgIndex(intflgIndex int64) {
	t.AddPropertyName("intflgIndex")
	t.intflgIndex = intflgIndex
}
func (t *VendorLargeData) SetNumericIntegerIndex(numericIntegerIndex df.Numeric) {
	t.AddPropertyName("numericIntegerIndex")
	t.numericIntegerIndex = numericIntegerIndex
}
func (t *VendorLargeData) SetNumericIntegerNoIndex(numericIntegerNoIndex df.Numeric) {
	t.AddPropertyName("numericIntegerNoIndex")
	t.numericIntegerNoIndex = numericIntegerNoIndex
}

func (t *VendorLargeData) SetUp(){
	t.numericIntegerIndex.DecPoint = 0
	t.numericIntegerNoIndex.DecPoint = 0
	
}
func (t *VendorLargeData)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}