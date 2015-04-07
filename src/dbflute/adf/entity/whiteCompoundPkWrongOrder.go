package entity

import (
	"github.com/mikeshimura/dbflute/df"
	"database/sql"
)

type WhiteCompoundPkWrongOrder struct {
	firstId int64
	secondId int64
	thirdId int64
	wrongName sql.NullString
	df.BaseEntity
}

func CreateWhiteCompoundPkWrongOrder() *WhiteCompoundPkWrongOrder{
	whiteCompoundPkWrongOrder :=new(WhiteCompoundPkWrongOrder)
	whiteCompoundPkWrongOrder.SetUp()
	return whiteCompoundPkWrongOrder 
}

func (l *WhiteCompoundPkWrongOrder) GetFirstId () int64 {
	return l.firstId
}
func (l *WhiteCompoundPkWrongOrder) GetSecondId () int64 {
	return l.secondId
}
func (l *WhiteCompoundPkWrongOrder) GetThirdId () int64 {
	return l.thirdId
}
func (l *WhiteCompoundPkWrongOrder) GetWrongName () sql.NullString {
	return l.wrongName
}

func (t *WhiteCompoundPkWrongOrder) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 4)
	i[0] = &(t.firstId)
	i[1] = &(t.secondId)
	i[2] = &(t.thirdId)
	i[3] = &(t.wrongName)
	return i
}


func (t *WhiteCompoundPkWrongOrder) AsTableDbName() string {
	return "WhiteCompoundPkWrongOrder"
}

func (t *WhiteCompoundPkWrongOrder) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("firstId") == false {
            return false 
        }
        if t.IsModifiedProperty("thirdId") == false {
            return false 
        }
        if t.IsModifiedProperty("secondId") == false {
            return false 
        }
        return true;
}
func (t *WhiteCompoundPkWrongOrder) SetFirstId(firstId int64) {
	t.AddPropertyName("firstId")
	t.firstId = firstId
}
func (t *WhiteCompoundPkWrongOrder) SetSecondId(secondId int64) {
	t.AddPropertyName("secondId")
	t.secondId = secondId
}
func (t *WhiteCompoundPkWrongOrder) SetThirdId(thirdId int64) {
	t.AddPropertyName("thirdId")
	t.thirdId = thirdId
}
func (t *WhiteCompoundPkWrongOrder) SetWrongName(wrongName sql.NullString) {
	t.AddPropertyName("wrongName")
	t.wrongName = wrongName
}

func (t *WhiteCompoundPkWrongOrder) SetUp(){
	
}
func (t *WhiteCompoundPkWrongOrder)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}