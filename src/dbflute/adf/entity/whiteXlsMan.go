package entity

import (
	"github.com/mikeshimura/dbflute/df"
	"database/sql"
)

type WhiteXlsMan struct {
	xlsManId int64
	stringConverted sql.NullString
	timestampZeroDefaultMillis df.NullTimestamp
	timestampZeroPrefixMillis df.NullTimestamp
	timestampZeroSuffixMillis df.NullTimestamp
	df.BaseEntity
}

func CreateWhiteXlsMan() *WhiteXlsMan{
	whiteXlsMan :=new(WhiteXlsMan)
	whiteXlsMan.SetUp()
	return whiteXlsMan 
}

func (l *WhiteXlsMan) GetXlsManId () int64 {
	return l.xlsManId
}
func (l *WhiteXlsMan) GetStringConverted () sql.NullString {
	return l.stringConverted
}
func (l *WhiteXlsMan) GetTimestampZeroDefaultMillis () df.NullTimestamp {
	return l.timestampZeroDefaultMillis
}
func (l *WhiteXlsMan) GetTimestampZeroPrefixMillis () df.NullTimestamp {
	return l.timestampZeroPrefixMillis
}
func (l *WhiteXlsMan) GetTimestampZeroSuffixMillis () df.NullTimestamp {
	return l.timestampZeroSuffixMillis
}

func (t *WhiteXlsMan) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 5)
	i[0] = &(t.xlsManId)
	i[1] = &(t.stringConverted)
	i[2] = &(t.timestampZeroDefaultMillis)
	i[3] = &(t.timestampZeroPrefixMillis)
	i[4] = &(t.timestampZeroSuffixMillis)
	return i
}


func (t *WhiteXlsMan) AsTableDbName() string {
	return "WhiteXlsMan"
}

func (t *WhiteXlsMan) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("xlsManId") == false {
            return false 
        }
        return true;
}
func (t *WhiteXlsMan) SetXlsManId(xlsManId int64) {
	t.AddPropertyName("xlsManId")
	t.xlsManId = xlsManId
}
func (t *WhiteXlsMan) SetStringConverted(stringConverted sql.NullString) {
	t.AddPropertyName("stringConverted")
	t.stringConverted = stringConverted
}
func (t *WhiteXlsMan) SetTimestampZeroDefaultMillis(timestampZeroDefaultMillis df.NullTimestamp) {
	t.AddPropertyName("timestampZeroDefaultMillis")
	t.timestampZeroDefaultMillis = timestampZeroDefaultMillis
}
func (t *WhiteXlsMan) SetTimestampZeroPrefixMillis(timestampZeroPrefixMillis df.NullTimestamp) {
	t.AddPropertyName("timestampZeroPrefixMillis")
	t.timestampZeroPrefixMillis = timestampZeroPrefixMillis
}
func (t *WhiteXlsMan) SetTimestampZeroSuffixMillis(timestampZeroSuffixMillis df.NullTimestamp) {
	t.AddPropertyName("timestampZeroSuffixMillis")
	t.timestampZeroSuffixMillis = timestampZeroSuffixMillis
}

func (t *WhiteXlsMan) SetUp(){
	
}
func (t *WhiteXlsMan)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}