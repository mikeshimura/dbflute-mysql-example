package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type WhiteSameNameDbm_T struct {
	df.BaseDBMeta
	ColumnSameNameId *df.ColumnInfo
	ColumnSameNameName *df.ColumnInfo
	ColumnSameNameInteger *df.ColumnInfo
	ColumnNextSchemaProductId *df.ColumnInfo
}

func (b *WhiteSameNameDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *WhiteSameNameDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var WhiteSameNameDbm *WhiteSameNameDbm_T

func Create_WhiteSameNameDbm() {
	WhiteSameNameDbm = new(WhiteSameNameDbm_T)
	WhiteSameNameDbm.TableDbName = "white_same_name"
	WhiteSameNameDbm.TableDispName = "white_same_name"
	WhiteSameNameDbm.TablePropertyName = "whiteSameName"
	WhiteSameNameDbm.TableSqlName = new(df.TableSqlName)
	WhiteSameNameDbm.TableSqlName.TableSqlName = "white_same_name"
	WhiteSameNameDbm.TableSqlName.CorrespondingDbName = WhiteSameNameDbm.TableDbName

	var whiteSameName df.DBMeta
	whiteSameName = WhiteSameNameDbm
	WhiteSameNameDbm.DBMeta=&whiteSameName
	sameNameIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo same_name_id
	sameNameIdSqlName.ColumnSqlName = "same_name_id"
	sameNameIdSqlName.IrregularChar = false
	WhiteSameNameDbm.ColumnSameNameId = df.CCI(&whiteSameName, "same_name_id", sameNameIdSqlName, "", "", "Long.class", "sameNameId", "", true, false,true, "int8", 19, 0, "",false,"","", "","whiteSameNameRefList","",false,"int64")
	sameNameNameSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo same_name_name
	sameNameNameSqlName.ColumnSqlName = "same_name_name"
	sameNameNameSqlName.IrregularChar = false
	WhiteSameNameDbm.ColumnSameNameName = df.CCI(&whiteSameName, "same_name_name", sameNameNameSqlName, "", "", "String.class", "sameNameName", "", false, false,false, "varchar", 200, 0, "",false,"","", "","","",false,"sql.NullString")
	sameNameIntegerSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo same_name_integer
	sameNameIntegerSqlName.ColumnSqlName = "same_name_integer"
	sameNameIntegerSqlName.IrregularChar = false
	WhiteSameNameDbm.ColumnSameNameInteger = df.CCI(&whiteSameName, "same_name_integer", sameNameIntegerSqlName, "", "", "Integer.class", "sameNameInteger", "", false, false,false, "int4", 10, 0, "",false,"","", "","","",false,"sql.NullInt64")
	nextSchemaProductIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo next_schema_product_id
	nextSchemaProductIdSqlName.ColumnSqlName = "next_schema_product_id"
	nextSchemaProductIdSqlName.IrregularChar = false
	WhiteSameNameDbm.ColumnNextSchemaProductId = df.CCI(&whiteSameName, "next_schema_product_id", nextSchemaProductIdSqlName, "", "", "Integer.class", "nextSchemaProductId", "", false, false,false, "int4", 10, 0, "",false,"","", "","","",false,"sql.NullInt64")

	WhiteSameNameDbm.ColumnInfoList = new(df.List)
	WhiteSameNameDbm.ColumnInfoList.Add(WhiteSameNameDbm.ColumnSameNameId)
	WhiteSameNameDbm.ColumnInfoList.Add(WhiteSameNameDbm.ColumnSameNameName)
	WhiteSameNameDbm.ColumnInfoList.Add(WhiteSameNameDbm.ColumnSameNameInteger)
	WhiteSameNameDbm.ColumnInfoList.Add(WhiteSameNameDbm.ColumnNextSchemaProductId)


	WhiteSameNameDbm.ColumnInfoMap=make(map[string]int)
	WhiteSameNameDbm.ColumnInfoMap["sameNameId"]=0
		WhiteSameNameDbm.ColumnInfoMap["sameNameName"]=1
		WhiteSameNameDbm.ColumnInfoMap["sameNameInteger"]=2
		WhiteSameNameDbm.ColumnInfoMap["nextSchemaProductId"]=3
	    WhiteSameNameDbm.PrimaryKey = true
    WhiteSameNameDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &whiteSameName
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(WhiteSameNameDbm.ColumnSameNameId)

	WhiteSameNameDbm.PrimaryInfo = new(df.PrimaryInfo)
	WhiteSameNameDbm.PrimaryInfo.UniqueInfo = ui
	
	var whiteSameNameMeta df.DBMeta = WhiteSameNameDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["WhiteSameName"] = &whiteSameNameMeta
}
