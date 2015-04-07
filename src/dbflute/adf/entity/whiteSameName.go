package entity

import (
	"github.com/mikeshimura/dbflute/df"
	"database/sql"
)

type WhiteSameName struct {
	sameNameId int64
	sameNameName sql.NullString
	sameNameInteger sql.NullInt64
	nextSchemaProductId sql.NullInt64
	df.BaseEntity
}

func CreateWhiteSameName() *WhiteSameName{
	whiteSameName :=new(WhiteSameName)
	whiteSameName.SetUp()
	return whiteSameName 
}

func (l *WhiteSameName) GetSameNameId () int64 {
	return l.sameNameId
}
func (l *WhiteSameName) GetSameNameName () sql.NullString {
	return l.sameNameName
}
func (l *WhiteSameName) GetSameNameInteger () sql.NullInt64 {
	return l.sameNameInteger
}
func (l *WhiteSameName) GetNextSchemaProductId () sql.NullInt64 {
	return l.nextSchemaProductId
}

func (t *WhiteSameName) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 4)
	i[0] = &(t.sameNameId)
	i[1] = &(t.sameNameName)
	i[2] = &(t.sameNameInteger)
	i[3] = &(t.nextSchemaProductId)
	return i
}


func (t *WhiteSameName) AsTableDbName() string {
	return "WhiteSameName"
}

func (t *WhiteSameName) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("sameNameId") == false {
            return false 
        }
        return true;
}
func (t *WhiteSameName) SetSameNameId(sameNameId int64) {
	t.AddPropertyName("sameNameId")
	t.sameNameId = sameNameId
}
func (t *WhiteSameName) SetSameNameName(sameNameName sql.NullString) {
	t.AddPropertyName("sameNameName")
	t.sameNameName = sameNameName
}
func (t *WhiteSameName) SetSameNameInteger(sameNameInteger sql.NullInt64) {
	t.AddPropertyName("sameNameInteger")
	t.sameNameInteger = sameNameInteger
}
func (t *WhiteSameName) SetNextSchemaProductId(nextSchemaProductId sql.NullInt64) {
	t.AddPropertyName("nextSchemaProductId")
	t.nextSchemaProductId = nextSchemaProductId
}

func (t *WhiteSameName) SetUp(){
	
}
func (t *WhiteSameName)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}