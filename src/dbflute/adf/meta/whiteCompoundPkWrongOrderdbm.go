package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type WhiteCompoundPkWrongOrderDbm_T struct {
	df.BaseDBMeta
	ColumnFirstId *df.ColumnInfo
	ColumnSecondId *df.ColumnInfo
	ColumnThirdId *df.ColumnInfo
	ColumnWrongName *df.ColumnInfo
}

func (b *WhiteCompoundPkWrongOrderDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *WhiteCompoundPkWrongOrderDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var WhiteCompoundPkWrongOrderDbm *WhiteCompoundPkWrongOrderDbm_T

func Create_WhiteCompoundPkWrongOrderDbm() {
	WhiteCompoundPkWrongOrderDbm = new(WhiteCompoundPkWrongOrderDbm_T)
	WhiteCompoundPkWrongOrderDbm.TableDbName = "white_compound_pk_wrong_order"
	WhiteCompoundPkWrongOrderDbm.TableDispName = "white_compound_pk_wrong_order"
	WhiteCompoundPkWrongOrderDbm.TablePropertyName = "whiteCompoundPkWrongOrder"
	WhiteCompoundPkWrongOrderDbm.TableSqlName = new(df.TableSqlName)
	WhiteCompoundPkWrongOrderDbm.TableSqlName.TableSqlName = "white_compound_pk_wrong_order"
	WhiteCompoundPkWrongOrderDbm.TableSqlName.CorrespondingDbName = WhiteCompoundPkWrongOrderDbm.TableDbName

	var whiteCompoundPkWrongOrder df.DBMeta
	whiteCompoundPkWrongOrder = WhiteCompoundPkWrongOrderDbm
	WhiteCompoundPkWrongOrderDbm.DBMeta=&whiteCompoundPkWrongOrder
	firstIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo first_id
	firstIdSqlName.ColumnSqlName = "first_id"
	firstIdSqlName.IrregularChar = false
	WhiteCompoundPkWrongOrderDbm.ColumnFirstId = df.CCI(&whiteCompoundPkWrongOrder, "first_id", firstIdSqlName, "", "", "Integer.class", "firstId", "", true, false,true, "int4", 10, 0, "",false,"","", "","","",false,"int64")
	secondIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo second_id
	secondIdSqlName.ColumnSqlName = "second_id"
	secondIdSqlName.IrregularChar = false
	WhiteCompoundPkWrongOrderDbm.ColumnSecondId = df.CCI(&whiteCompoundPkWrongOrder, "second_id", secondIdSqlName, "", "", "Integer.class", "secondId", "", true, false,true, "int4", 10, 0, "",false,"","", "","","",false,"int64")
	thirdIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo third_id
	thirdIdSqlName.ColumnSqlName = "third_id"
	thirdIdSqlName.IrregularChar = false
	WhiteCompoundPkWrongOrderDbm.ColumnThirdId = df.CCI(&whiteCompoundPkWrongOrder, "third_id", thirdIdSqlName, "", "", "Integer.class", "thirdId", "", true, false,true, "int4", 10, 0, "",false,"","", "","","",false,"int64")
	wrongNameSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo wrong_name
	wrongNameSqlName.ColumnSqlName = "wrong_name"
	wrongNameSqlName.IrregularChar = false
	WhiteCompoundPkWrongOrderDbm.ColumnWrongName = df.CCI(&whiteCompoundPkWrongOrder, "wrong_name", wrongNameSqlName, "", "", "String.class", "wrongName", "", false, false,false, "varchar", 200, 0, "",false,"","", "","","",false,"sql.NullString")

	WhiteCompoundPkWrongOrderDbm.ColumnInfoList = new(df.List)
	WhiteCompoundPkWrongOrderDbm.ColumnInfoList.Add(WhiteCompoundPkWrongOrderDbm.ColumnFirstId)
	WhiteCompoundPkWrongOrderDbm.ColumnInfoList.Add(WhiteCompoundPkWrongOrderDbm.ColumnSecondId)
	WhiteCompoundPkWrongOrderDbm.ColumnInfoList.Add(WhiteCompoundPkWrongOrderDbm.ColumnThirdId)
	WhiteCompoundPkWrongOrderDbm.ColumnInfoList.Add(WhiteCompoundPkWrongOrderDbm.ColumnWrongName)


	WhiteCompoundPkWrongOrderDbm.ColumnInfoMap=make(map[string]int)
	WhiteCompoundPkWrongOrderDbm.ColumnInfoMap["firstId"]=0
		WhiteCompoundPkWrongOrderDbm.ColumnInfoMap["secondId"]=1
		WhiteCompoundPkWrongOrderDbm.ColumnInfoMap["thirdId"]=2
		WhiteCompoundPkWrongOrderDbm.ColumnInfoMap["wrongName"]=3
	    WhiteCompoundPkWrongOrderDbm.PrimaryKey = true
    WhiteCompoundPkWrongOrderDbm.CompoundPrimaryKey = true
	ui := new(df.UniqueInfo)
	ui.DbMeta = &whiteCompoundPkWrongOrder
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
    ui.UniqueColumnList.Add(WhiteCompoundPkWrongOrderDbm.ColumnFirstId)
    ui.UniqueColumnList.Add(WhiteCompoundPkWrongOrderDbm.ColumnThirdId)
    ui.UniqueColumnList.Add(WhiteCompoundPkWrongOrderDbm.ColumnSecondId)

	WhiteCompoundPkWrongOrderDbm.PrimaryInfo = new(df.PrimaryInfo)
	WhiteCompoundPkWrongOrderDbm.PrimaryInfo.UniqueInfo = ui
	
	var whiteCompoundPkWrongOrderMeta df.DBMeta = WhiteCompoundPkWrongOrderDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["WhiteCompoundPkWrongOrder"] = &whiteCompoundPkWrongOrderMeta
}
