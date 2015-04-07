package entity

import (
	"github.com/mikeshimura/dbflute/df"
	"database/sql"
)

type VendorLargeDataRef struct {
	largeDataRefId int64
	largeDataId int64
	dateIndex df.Date
	dateNoIndex df.Date
	timestampIndex df.Timestamp
	timestampNoIndex df.Timestamp
	nullableDecimalIndex df.NullNumeric
	nullableDecimalNoIndex df.NullNumeric
	selfParentId sql.NullInt64
	df.BaseEntity
}

func CreateVendorLargeDataRef() *VendorLargeDataRef{
	vendorLargeDataRef :=new(VendorLargeDataRef)
	vendorLargeDataRef.SetUp()
	return vendorLargeDataRef 
}

func (l *VendorLargeDataRef) GetLargeDataRefId () int64 {
	return l.largeDataRefId
}
func (l *VendorLargeDataRef) GetLargeDataId () int64 {
	return l.largeDataId
}
func (l *VendorLargeDataRef) GetDateIndex () df.Date {
	return l.dateIndex
}
func (l *VendorLargeDataRef) GetDateNoIndex () df.Date {
	return l.dateNoIndex
}
func (l *VendorLargeDataRef) GetTimestampIndex () df.Timestamp {
	return l.timestampIndex
}
func (l *VendorLargeDataRef) GetTimestampNoIndex () df.Timestamp {
	return l.timestampNoIndex
}
func (l *VendorLargeDataRef) GetNullableDecimalIndex () df.NullNumeric {
	return l.nullableDecimalIndex
}
func (l *VendorLargeDataRef) GetNullableDecimalNoIndex () df.NullNumeric {
	return l.nullableDecimalNoIndex
}
func (l *VendorLargeDataRef) GetSelfParentId () sql.NullInt64 {
	return l.selfParentId
}

func (t *VendorLargeDataRef) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 9)
	i[0] = &(t.largeDataRefId)
	i[1] = &(t.largeDataId)
	i[2] = &(t.dateIndex)
	i[3] = &(t.dateNoIndex)
	i[4] = &(t.timestampIndex)
	i[5] = &(t.timestampNoIndex)
	i[6] = &(t.nullableDecimalIndex)
	i[7] = &(t.nullableDecimalNoIndex)
	i[8] = &(t.selfParentId)
	return i
}


func (t *VendorLargeDataRef) AsTableDbName() string {
	return "VendorLargeDataRef"
}

func (t *VendorLargeDataRef) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("largeDataRefId") == false {
            return false 
        }
        return true;
}
func (t *VendorLargeDataRef) SetLargeDataRefId(largeDataRefId int64) {
	t.AddPropertyName("largeDataRefId")
	t.largeDataRefId = largeDataRefId
}
func (t *VendorLargeDataRef) SetLargeDataId(largeDataId int64) {
	t.AddPropertyName("largeDataId")
	t.largeDataId = largeDataId
}
func (t *VendorLargeDataRef) SetDateIndex(dateIndex df.Date) {
	t.AddPropertyName("dateIndex")
	t.dateIndex = dateIndex
}
func (t *VendorLargeDataRef) SetDateNoIndex(dateNoIndex df.Date) {
	t.AddPropertyName("dateNoIndex")
	t.dateNoIndex = dateNoIndex
}
func (t *VendorLargeDataRef) SetTimestampIndex(timestampIndex df.Timestamp) {
	t.AddPropertyName("timestampIndex")
	t.timestampIndex = timestampIndex
}
func (t *VendorLargeDataRef) SetTimestampNoIndex(timestampNoIndex df.Timestamp) {
	t.AddPropertyName("timestampNoIndex")
	t.timestampNoIndex = timestampNoIndex
}
func (t *VendorLargeDataRef) SetNullableDecimalIndex(nullableDecimalIndex df.NullNumeric) {
	t.AddPropertyName("nullableDecimalIndex")
	t.nullableDecimalIndex = nullableDecimalIndex
}
func (t *VendorLargeDataRef) SetNullableDecimalNoIndex(nullableDecimalNoIndex df.NullNumeric) {
	t.AddPropertyName("nullableDecimalNoIndex")
	t.nullableDecimalNoIndex = nullableDecimalNoIndex
}
func (t *VendorLargeDataRef) SetSelfParentId(selfParentId sql.NullInt64) {
	t.AddPropertyName("selfParentId")
	t.selfParentId = selfParentId
}

func (t *VendorLargeDataRef) SetUp(){
	t.nullableDecimalIndex.DecPoint = 3
	t.nullableDecimalNoIndex.DecPoint = 3
	
}
func (t *VendorLargeDataRef)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}