package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type Region struct {
	regionId int64
	regionName string
	df.BaseEntity
}

func CreateRegion() *Region{
	region :=new(Region)
	region.SetUp()
	return region 
}

func (l *Region) GetRegionId () int64 {
	return l.regionId
}
func (l *Region) GetRegionName () string {
	return l.regionName
}

func (t *Region) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 2)
	i[0] = &(t.regionId)
	i[1] = &(t.regionName)
	return i
}


func (t *Region) AsTableDbName() string {
	return "Region"
}

func (t *Region) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("regionId") == false {
            return false 
        }
        return true;
}
func (t *Region) SetRegionId(regionId int64) {
	t.AddPropertyName("regionId")
	t.regionId = regionId
}
func (t *Region) SetRegionName(regionName string) {
	t.AddPropertyName("regionName")
	t.regionName = regionName
}

func (t *Region) SetUp(){
	
}
func (t *Region)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}