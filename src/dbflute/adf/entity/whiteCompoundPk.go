package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type WhiteCompoundPk struct {
	pkFirstId int64
	pkSecondId int64
	pkName string
	df.BaseEntity
}

func CreateWhiteCompoundPk() *WhiteCompoundPk{
	whiteCompoundPk :=new(WhiteCompoundPk)
	whiteCompoundPk.SetUp()
	return whiteCompoundPk 
}

func (l *WhiteCompoundPk) GetPkFirstId () int64 {
	return l.pkFirstId
}
func (l *WhiteCompoundPk) GetPkSecondId () int64 {
	return l.pkSecondId
}
func (l *WhiteCompoundPk) GetPkName () string {
	return l.pkName
}

func (t *WhiteCompoundPk) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 3)
	i[0] = &(t.pkFirstId)
	i[1] = &(t.pkSecondId)
	i[2] = &(t.pkName)
	return i
}


func (t *WhiteCompoundPk) AsTableDbName() string {
	return "WhiteCompoundPk"
}

func (t *WhiteCompoundPk) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("pkFirstId") == false {
            return false 
        }
        if t.IsModifiedProperty("pkSecondId") == false {
            return false 
        }
        return true;
}
func (t *WhiteCompoundPk) SetPkFirstId(pkFirstId int64) {
	t.AddPropertyName("pkFirstId")
	t.pkFirstId = pkFirstId
}
func (t *WhiteCompoundPk) SetPkSecondId(pkSecondId int64) {
	t.AddPropertyName("pkSecondId")
	t.pkSecondId = pkSecondId
}
func (t *WhiteCompoundPk) SetPkName(pkName string) {
	t.AddPropertyName("pkName")
	t.pkName = pkName
}

func (t *WhiteCompoundPk) SetUp(){
	
}
func (t *WhiteCompoundPk)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}