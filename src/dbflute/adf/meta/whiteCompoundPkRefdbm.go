package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type WhiteCompoundPkRefDbm_T struct {
	df.BaseDBMeta
	ColumnMultipleFirstId *df.ColumnInfo
	ColumnMultipleSecondId *df.ColumnInfo
	ColumnRefFirstId *df.ColumnInfo
	ColumnRefSecondId *df.ColumnInfo
}

func (b *WhiteCompoundPkRefDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *WhiteCompoundPkRefDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var WhiteCompoundPkRefDbm *WhiteCompoundPkRefDbm_T

func Create_WhiteCompoundPkRefDbm() {
	WhiteCompoundPkRefDbm = new(WhiteCompoundPkRefDbm_T)
	WhiteCompoundPkRefDbm.TableDbName = "white_compound_pk_ref"
	WhiteCompoundPkRefDbm.TableDispName = "white_compound_pk_ref"
	WhiteCompoundPkRefDbm.TablePropertyName = "whiteCompoundPkRef"
	WhiteCompoundPkRefDbm.TableSqlName = new(df.TableSqlName)
	WhiteCompoundPkRefDbm.TableSqlName.TableSqlName = "white_compound_pk_ref"
	WhiteCompoundPkRefDbm.TableSqlName.CorrespondingDbName = WhiteCompoundPkRefDbm.TableDbName

	var whiteCompoundPkRef df.DBMeta
	whiteCompoundPkRef = WhiteCompoundPkRefDbm
	WhiteCompoundPkRefDbm.DBMeta=&whiteCompoundPkRef
	multipleFirstIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo multiple_first_id
	multipleFirstIdSqlName.ColumnSqlName = "multiple_first_id"
	multipleFirstIdSqlName.IrregularChar = false
	WhiteCompoundPkRefDbm.ColumnMultipleFirstId = df.CCI(&whiteCompoundPkRef, "multiple_first_id", multipleFirstIdSqlName, "", "", "Integer.class", "multipleFirstId", "", true, false,true, "int4", 10, 0, "",false,"","", "","","",false,"int64")
	multipleSecondIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo multiple_second_id
	multipleSecondIdSqlName.ColumnSqlName = "multiple_second_id"
	multipleSecondIdSqlName.IrregularChar = false
	WhiteCompoundPkRefDbm.ColumnMultipleSecondId = df.CCI(&whiteCompoundPkRef, "multiple_second_id", multipleSecondIdSqlName, "", "", "Integer.class", "multipleSecondId", "", true, false,true, "int4", 10, 0, "",false,"","", "","","",false,"int64")
	refFirstIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo ref_first_id
	refFirstIdSqlName.ColumnSqlName = "ref_first_id"
	refFirstIdSqlName.IrregularChar = false
	WhiteCompoundPkRefDbm.ColumnRefFirstId = df.CCI(&whiteCompoundPkRef, "ref_first_id", refFirstIdSqlName, "", "", "Integer.class", "refFirstId", "", false, false,true, "int4", 10, 0, "",false,"","", "whiteCompoundPk","","",false,"int64")
	refSecondIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo ref_second_id
	refSecondIdSqlName.ColumnSqlName = "ref_second_id"
	refSecondIdSqlName.IrregularChar = false
	WhiteCompoundPkRefDbm.ColumnRefSecondId = df.CCI(&whiteCompoundPkRef, "ref_second_id", refSecondIdSqlName, "", "", "Integer.class", "refSecondId", "", false, false,true, "int4", 10, 0, "",false,"","", "whiteCompoundPk","","",false,"int64")

	WhiteCompoundPkRefDbm.ColumnInfoList = new(df.List)
	WhiteCompoundPkRefDbm.ColumnInfoList.Add(WhiteCompoundPkRefDbm.ColumnMultipleFirstId)
	WhiteCompoundPkRefDbm.ColumnInfoList.Add(WhiteCompoundPkRefDbm.ColumnMultipleSecondId)
	WhiteCompoundPkRefDbm.ColumnInfoList.Add(WhiteCompoundPkRefDbm.ColumnRefFirstId)
	WhiteCompoundPkRefDbm.ColumnInfoList.Add(WhiteCompoundPkRefDbm.ColumnRefSecondId)


	WhiteCompoundPkRefDbm.ColumnInfoMap=make(map[string]int)
	WhiteCompoundPkRefDbm.ColumnInfoMap["multipleFirstId"]=0
		WhiteCompoundPkRefDbm.ColumnInfoMap["multipleSecondId"]=1
		WhiteCompoundPkRefDbm.ColumnInfoMap["refFirstId"]=2
		WhiteCompoundPkRefDbm.ColumnInfoMap["refSecondId"]=3
	    WhiteCompoundPkRefDbm.PrimaryKey = true
    WhiteCompoundPkRefDbm.CompoundPrimaryKey = true
	ui := new(df.UniqueInfo)
	ui.DbMeta = &whiteCompoundPkRef
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
    ui.UniqueColumnList.Add(WhiteCompoundPkRefDbm.ColumnMultipleFirstId)
    ui.UniqueColumnList.Add(WhiteCompoundPkRefDbm.ColumnMultipleSecondId)

	WhiteCompoundPkRefDbm.PrimaryInfo = new(df.PrimaryInfo)
	WhiteCompoundPkRefDbm.PrimaryInfo.UniqueInfo = ui
	
	var whiteCompoundPkRefMeta df.DBMeta = WhiteCompoundPkRefDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["WhiteCompoundPkRef"] = &whiteCompoundPkRefMeta
}
