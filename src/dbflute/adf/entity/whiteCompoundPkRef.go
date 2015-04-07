package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type WhiteCompoundPkRef struct {
	multipleFirstId int64
	multipleSecondId int64
	refFirstId int64
	refSecondId int64
	df.BaseEntity
}

func CreateWhiteCompoundPkRef() *WhiteCompoundPkRef{
	whiteCompoundPkRef :=new(WhiteCompoundPkRef)
	whiteCompoundPkRef.SetUp()
	return whiteCompoundPkRef 
}

func (l *WhiteCompoundPkRef) GetMultipleFirstId () int64 {
	return l.multipleFirstId
}
func (l *WhiteCompoundPkRef) GetMultipleSecondId () int64 {
	return l.multipleSecondId
}
func (l *WhiteCompoundPkRef) GetRefFirstId () int64 {
	return l.refFirstId
}
func (l *WhiteCompoundPkRef) GetRefSecondId () int64 {
	return l.refSecondId
}

func (t *WhiteCompoundPkRef) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 4)
	i[0] = &(t.multipleFirstId)
	i[1] = &(t.multipleSecondId)
	i[2] = &(t.refFirstId)
	i[3] = &(t.refSecondId)
	return i
}


func (t *WhiteCompoundPkRef) AsTableDbName() string {
	return "WhiteCompoundPkRef"
}

func (t *WhiteCompoundPkRef) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("multipleFirstId") == false {
            return false 
        }
        if t.IsModifiedProperty("multipleSecondId") == false {
            return false 
        }
        return true;
}
func (t *WhiteCompoundPkRef) SetMultipleFirstId(multipleFirstId int64) {
	t.AddPropertyName("multipleFirstId")
	t.multipleFirstId = multipleFirstId
}
func (t *WhiteCompoundPkRef) SetMultipleSecondId(multipleSecondId int64) {
	t.AddPropertyName("multipleSecondId")
	t.multipleSecondId = multipleSecondId
}
func (t *WhiteCompoundPkRef) SetRefFirstId(refFirstId int64) {
	t.AddPropertyName("refFirstId")
	t.refFirstId = refFirstId
}
func (t *WhiteCompoundPkRef) SetRefSecondId(refSecondId int64) {
	t.AddPropertyName("refSecondId")
	t.refSecondId = refSecondId
}

func (t *WhiteCompoundPkRef) SetUp(){
	
}
func (t *WhiteCompoundPkRef)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}