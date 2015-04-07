package entity

import (
	"github.com/mikeshimura/dbflute/df"
	"database/sql"
)

type WhiteNotPk struct {
	notPkId int64
	notPkName sql.NullString
	notPkInteger sql.NullInt64
	df.BaseEntity
}

func CreateWhiteNotPk() *WhiteNotPk{
	whiteNotPk :=new(WhiteNotPk)
	whiteNotPk.SetUp()
	return whiteNotPk 
}

func (l *WhiteNotPk) GetNotPkId () int64 {
	return l.notPkId
}
func (l *WhiteNotPk) GetNotPkName () sql.NullString {
	return l.notPkName
}
func (l *WhiteNotPk) GetNotPkInteger () sql.NullInt64 {
	return l.notPkInteger
}

func (t *WhiteNotPk) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 3)
	i[0] = &(t.notPkId)
	i[1] = &(t.notPkName)
	i[2] = &(t.notPkInteger)
	return i
}


func (t *WhiteNotPk) AsTableDbName() string {
	return "WhiteNotPk"
}

func (t *WhiteNotPk) HasPrimaryKeyValue() bool{
        return false;
}
func (t *WhiteNotPk) SetNotPkId(notPkId int64) {
	t.AddPropertyName("notPkId")
	t.notPkId = notPkId
}
func (t *WhiteNotPk) SetNotPkName(notPkName sql.NullString) {
	t.AddPropertyName("notPkName")
	t.notPkName = notPkName
}
func (t *WhiteNotPk) SetNotPkInteger(notPkInteger sql.NullInt64) {
	t.AddPropertyName("notPkInteger")
	t.notPkInteger = notPkInteger
}

func (t *WhiteNotPk) SetUp(){
	
}
func (t *WhiteNotPk)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}