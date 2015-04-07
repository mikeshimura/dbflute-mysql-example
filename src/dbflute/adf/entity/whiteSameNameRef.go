package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type WhiteSameNameRef struct {
	sameNameRefId int64
	sameNameId int64
	nextSameNameId int64
	df.BaseEntity
}

func CreateWhiteSameNameRef() *WhiteSameNameRef{
	whiteSameNameRef :=new(WhiteSameNameRef)
	whiteSameNameRef.SetUp()
	return whiteSameNameRef 
}

func (l *WhiteSameNameRef) GetSameNameRefId () int64 {
	return l.sameNameRefId
}
func (l *WhiteSameNameRef) GetSameNameId () int64 {
	return l.sameNameId
}
func (l *WhiteSameNameRef) GetNextSameNameId () int64 {
	return l.nextSameNameId
}

func (t *WhiteSameNameRef) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 3)
	i[0] = &(t.sameNameRefId)
	i[1] = &(t.sameNameId)
	i[2] = &(t.nextSameNameId)
	return i
}


func (t *WhiteSameNameRef) AsTableDbName() string {
	return "WhiteSameNameRef"
}

func (t *WhiteSameNameRef) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("sameNameRefId") == false {
            return false 
        }
        return true;
}
func (t *WhiteSameNameRef) SetSameNameRefId(sameNameRefId int64) {
	t.AddPropertyName("sameNameRefId")
	t.sameNameRefId = sameNameRefId
}
func (t *WhiteSameNameRef) SetSameNameId(sameNameId int64) {
	t.AddPropertyName("sameNameId")
	t.sameNameId = sameNameId
}
func (t *WhiteSameNameRef) SetNextSameNameId(nextSameNameId int64) {
	t.AddPropertyName("nextSameNameId")
	t.nextSameNameId = nextSameNameId
}

func (t *WhiteSameNameRef) SetUp(){
	
}
func (t *WhiteSameNameRef)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}