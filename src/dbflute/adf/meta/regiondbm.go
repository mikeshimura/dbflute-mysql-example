package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type RegionDbm_T struct {
	df.BaseDBMeta
	ColumnRegionId *df.ColumnInfo
	ColumnRegionName *df.ColumnInfo
}

func (b *RegionDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}
func (b *RegionDbm_T) CreateForeignInfoMap() {
	b.ForeignInfoMap = make(map[string]*df.ForeignInfo)
}

func (b *RegionDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var RegionDbm *RegionDbm_T

func Create_RegionDbm() {
	RegionDbm = new(RegionDbm_T)
	RegionDbm.TableDbName = "region"
	RegionDbm.TableDispName = "region"
	RegionDbm.TablePropertyName = "region"
	RegionDbm.TableSqlName = new(df.TableSqlName)
	RegionDbm.TableSqlName.TableSqlName = "region"
	RegionDbm.TableSqlName.CorrespondingDbName = RegionDbm.TableDbName

	var region df.DBMeta
	region = RegionDbm
	RegionDbm.DBMeta=&region
	regionIdSqlName := new(df.ColumnSqlName)
	regionIdSqlName.ColumnSqlName = "REGION_ID"
	regionIdSqlName.IrregularChar = false
	RegionDbm.ColumnRegionId = df.CCI(&region, "REGION_ID", regionIdSqlName, "", "", "Integer.class", "regionId", "", true, false,true, "INT", 10, 0, "",false,"","", "","memberAddressList","",false,"int64")
	regionNameSqlName := new(df.ColumnSqlName)
	regionNameSqlName.ColumnSqlName = "REGION_NAME"
	regionNameSqlName.IrregularChar = false
	RegionDbm.ColumnRegionName = df.CCI(&region, "REGION_NAME", regionNameSqlName, "", "", "String.class", "regionName", "", false, false,true, "VARCHAR", 50, 0, "",false,"","", "","","",false,"string")

	RegionDbm.ColumnInfoList = new(df.List)
	RegionDbm.ColumnInfoList.Add(RegionDbm.ColumnRegionId)
	RegionDbm.ColumnInfoList.Add(RegionDbm.ColumnRegionName)


	RegionDbm.ColumnInfoMap=make(map[string]int)
	RegionDbm.ColumnInfoMap["regionId"]=0
		RegionDbm.ColumnInfoMap["regionName"]=1
	    RegionDbm.PrimaryKey = true
    RegionDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &region
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(RegionDbm.ColumnRegionId)

	RegionDbm.PrimaryInfo = new(df.PrimaryInfo)
	RegionDbm.PrimaryInfo.UniqueInfo = ui
	
	var regionMeta df.DBMeta = RegionDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["Region"] = &regionMeta
}
