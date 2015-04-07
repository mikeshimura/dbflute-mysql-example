package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type WhiteCompoundPkDbm_T struct {
	df.BaseDBMeta
	ColumnPkFirstId *df.ColumnInfo
	ColumnPkSecondId *df.ColumnInfo
	ColumnPkName *df.ColumnInfo
}

func (b *WhiteCompoundPkDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *WhiteCompoundPkDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var WhiteCompoundPkDbm *WhiteCompoundPkDbm_T

func Create_WhiteCompoundPkDbm() {
	WhiteCompoundPkDbm = new(WhiteCompoundPkDbm_T)
	WhiteCompoundPkDbm.TableDbName = "white_compound_pk"
	WhiteCompoundPkDbm.TableDispName = "white_compound_pk"
	WhiteCompoundPkDbm.TablePropertyName = "whiteCompoundPk"
	WhiteCompoundPkDbm.TableSqlName = new(df.TableSqlName)
	WhiteCompoundPkDbm.TableSqlName.TableSqlName = "white_compound_pk"
	WhiteCompoundPkDbm.TableSqlName.CorrespondingDbName = WhiteCompoundPkDbm.TableDbName

	var whiteCompoundPk df.DBMeta
	whiteCompoundPk = WhiteCompoundPkDbm
	WhiteCompoundPkDbm.DBMeta=&whiteCompoundPk
	pkFirstIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo pk_first_id
	pkFirstIdSqlName.ColumnSqlName = "pk_first_id"
	pkFirstIdSqlName.IrregularChar = false
	WhiteCompoundPkDbm.ColumnPkFirstId = df.CCI(&whiteCompoundPk, "pk_first_id", pkFirstIdSqlName, "", "", "Integer.class", "pkFirstId", "", true, false,true, "int4", 10, 0, "",false,"","", "","whiteCompoundPkRefList","",false,"int64")
	pkSecondIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo pk_second_id
	pkSecondIdSqlName.ColumnSqlName = "pk_second_id"
	pkSecondIdSqlName.IrregularChar = false
	WhiteCompoundPkDbm.ColumnPkSecondId = df.CCI(&whiteCompoundPk, "pk_second_id", pkSecondIdSqlName, "", "", "Integer.class", "pkSecondId", "", true, false,true, "int4", 10, 0, "",false,"","", "","whiteCompoundPkRefList","",false,"int64")
	pkNameSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo pk_name
	pkNameSqlName.ColumnSqlName = "pk_name"
	pkNameSqlName.IrregularChar = false
	WhiteCompoundPkDbm.ColumnPkName = df.CCI(&whiteCompoundPk, "pk_name", pkNameSqlName, "", "", "String.class", "pkName", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")

	WhiteCompoundPkDbm.ColumnInfoList = new(df.List)
	WhiteCompoundPkDbm.ColumnInfoList.Add(WhiteCompoundPkDbm.ColumnPkFirstId)
	WhiteCompoundPkDbm.ColumnInfoList.Add(WhiteCompoundPkDbm.ColumnPkSecondId)
	WhiteCompoundPkDbm.ColumnInfoList.Add(WhiteCompoundPkDbm.ColumnPkName)


	WhiteCompoundPkDbm.ColumnInfoMap=make(map[string]int)
	WhiteCompoundPkDbm.ColumnInfoMap["pkFirstId"]=0
		WhiteCompoundPkDbm.ColumnInfoMap["pkSecondId"]=1
		WhiteCompoundPkDbm.ColumnInfoMap["pkName"]=2
	    WhiteCompoundPkDbm.PrimaryKey = true
    WhiteCompoundPkDbm.CompoundPrimaryKey = true
	ui := new(df.UniqueInfo)
	ui.DbMeta = &whiteCompoundPk
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
    ui.UniqueColumnList.Add(WhiteCompoundPkDbm.ColumnPkFirstId)
    ui.UniqueColumnList.Add(WhiteCompoundPkDbm.ColumnPkSecondId)

	WhiteCompoundPkDbm.PrimaryInfo = new(df.PrimaryInfo)
	WhiteCompoundPkDbm.PrimaryInfo.UniqueInfo = ui
	
	var whiteCompoundPkMeta df.DBMeta = WhiteCompoundPkDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["WhiteCompoundPk"] = &whiteCompoundPkMeta
}
