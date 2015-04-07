package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type WhiteSameNameRefDbm_T struct {
	df.BaseDBMeta
	ColumnSameNameRefId *df.ColumnInfo
	ColumnSameNameId *df.ColumnInfo
	ColumnNextSameNameId *df.ColumnInfo
}

func (b *WhiteSameNameRefDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *WhiteSameNameRefDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var WhiteSameNameRefDbm *WhiteSameNameRefDbm_T

func Create_WhiteSameNameRefDbm() {
	WhiteSameNameRefDbm = new(WhiteSameNameRefDbm_T)
	WhiteSameNameRefDbm.TableDbName = "white_same_name_ref"
	WhiteSameNameRefDbm.TableDispName = "white_same_name_ref"
	WhiteSameNameRefDbm.TablePropertyName = "whiteSameNameRef"
	WhiteSameNameRefDbm.TableSqlName = new(df.TableSqlName)
	WhiteSameNameRefDbm.TableSqlName.TableSqlName = "white_same_name_ref"
	WhiteSameNameRefDbm.TableSqlName.CorrespondingDbName = WhiteSameNameRefDbm.TableDbName

	var whiteSameNameRef df.DBMeta
	whiteSameNameRef = WhiteSameNameRefDbm
	WhiteSameNameRefDbm.DBMeta=&whiteSameNameRef
	sameNameRefIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo same_name_ref_id
	sameNameRefIdSqlName.ColumnSqlName = "same_name_ref_id"
	sameNameRefIdSqlName.IrregularChar = false
	WhiteSameNameRefDbm.ColumnSameNameRefId = df.CCI(&whiteSameNameRef, "same_name_ref_id", sameNameRefIdSqlName, "", "", "Integer.class", "sameNameRefId", "", true, false,true, "int4", 10, 0, "",false,"","", "","","",false,"int64")
	sameNameIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo same_name_id
	sameNameIdSqlName.ColumnSqlName = "same_name_id"
	sameNameIdSqlName.IrregularChar = false
	WhiteSameNameRefDbm.ColumnSameNameId = df.CCI(&whiteSameNameRef, "same_name_id", sameNameIdSqlName, "", "", "Long.class", "sameNameId", "", false, false,true, "int8", 19, 0, "",false,"","", "whiteSameName","","",false,"int64")
	nextSameNameIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo next_same_name_id
	nextSameNameIdSqlName.ColumnSqlName = "next_same_name_id"
	nextSameNameIdSqlName.IrregularChar = false
	WhiteSameNameRefDbm.ColumnNextSameNameId = df.CCI(&whiteSameNameRef, "next_same_name_id", nextSameNameIdSqlName, "", "", "Long.class", "nextSameNameId", "", false, false,true, "int8", 19, 0, "",false,"","", "","","",false,"int64")

	WhiteSameNameRefDbm.ColumnInfoList = new(df.List)
	WhiteSameNameRefDbm.ColumnInfoList.Add(WhiteSameNameRefDbm.ColumnSameNameRefId)
	WhiteSameNameRefDbm.ColumnInfoList.Add(WhiteSameNameRefDbm.ColumnSameNameId)
	WhiteSameNameRefDbm.ColumnInfoList.Add(WhiteSameNameRefDbm.ColumnNextSameNameId)


	WhiteSameNameRefDbm.ColumnInfoMap=make(map[string]int)
	WhiteSameNameRefDbm.ColumnInfoMap["sameNameRefId"]=0
		WhiteSameNameRefDbm.ColumnInfoMap["sameNameId"]=1
		WhiteSameNameRefDbm.ColumnInfoMap["nextSameNameId"]=2
	    WhiteSameNameRefDbm.PrimaryKey = true
    WhiteSameNameRefDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &whiteSameNameRef
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(WhiteSameNameRefDbm.ColumnSameNameRefId)

	WhiteSameNameRefDbm.PrimaryInfo = new(df.PrimaryInfo)
	WhiteSameNameRefDbm.PrimaryInfo.UniqueInfo = ui
	
	var whiteSameNameRefMeta df.DBMeta = WhiteSameNameRefDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["WhiteSameNameRef"] = &whiteSameNameRefMeta
}
