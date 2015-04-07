package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type WhiteNotPkDbm_T struct {
	df.BaseDBMeta
	ColumnNotPkId *df.ColumnInfo
	ColumnNotPkName *df.ColumnInfo
	ColumnNotPkInteger *df.ColumnInfo
}

func (b *WhiteNotPkDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *WhiteNotPkDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var WhiteNotPkDbm *WhiteNotPkDbm_T

func Create_WhiteNotPkDbm() {
	WhiteNotPkDbm = new(WhiteNotPkDbm_T)
	WhiteNotPkDbm.TableDbName = "white_not_pk"
	WhiteNotPkDbm.TableDispName = "white_not_pk"
	WhiteNotPkDbm.TablePropertyName = "whiteNotPk"
	WhiteNotPkDbm.TableSqlName = new(df.TableSqlName)
	WhiteNotPkDbm.TableSqlName.TableSqlName = "white_not_pk"
	WhiteNotPkDbm.TableSqlName.CorrespondingDbName = WhiteNotPkDbm.TableDbName

	var whiteNotPk df.DBMeta
	whiteNotPk = WhiteNotPkDbm
	WhiteNotPkDbm.DBMeta=&whiteNotPk
	notPkIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo not_pk_id
	notPkIdSqlName.ColumnSqlName = "not_pk_id"
	notPkIdSqlName.IrregularChar = false
	WhiteNotPkDbm.ColumnNotPkId = df.CCI(&whiteNotPk, "not_pk_id", notPkIdSqlName, "", "", "Long.class", "notPkId", "", false, false,true, "int8", 19, 0, "",false,"","", "","","",false,"int64")
	notPkNameSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo not_pk_name
	notPkNameSqlName.ColumnSqlName = "not_pk_name"
	notPkNameSqlName.IrregularChar = false
	WhiteNotPkDbm.ColumnNotPkName = df.CCI(&whiteNotPk, "not_pk_name", notPkNameSqlName, "", "", "String.class", "notPkName", "", false, false,false, "varchar", 2147483647, 0, "",false,"","", "","","",false,"sql.NullString")
	notPkIntegerSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo not_pk_integer
	notPkIntegerSqlName.ColumnSqlName = "not_pk_integer"
	notPkIntegerSqlName.IrregularChar = false
	WhiteNotPkDbm.ColumnNotPkInteger = df.CCI(&whiteNotPk, "not_pk_integer", notPkIntegerSqlName, "", "", "Integer.class", "notPkInteger", "", false, false,false, "int4", 10, 0, "",false,"","", "","","",false,"sql.NullInt64")

	WhiteNotPkDbm.ColumnInfoList = new(df.List)
	WhiteNotPkDbm.ColumnInfoList.Add(WhiteNotPkDbm.ColumnNotPkId)
	WhiteNotPkDbm.ColumnInfoList.Add(WhiteNotPkDbm.ColumnNotPkName)
	WhiteNotPkDbm.ColumnInfoList.Add(WhiteNotPkDbm.ColumnNotPkInteger)


	WhiteNotPkDbm.ColumnInfoMap=make(map[string]int)
	WhiteNotPkDbm.ColumnInfoMap["notPkId"]=0
		WhiteNotPkDbm.ColumnInfoMap["notPkName"]=1
		WhiteNotPkDbm.ColumnInfoMap["notPkInteger"]=2
	    WhiteNotPkDbm.PrimaryKey = false
    WhiteNotPkDbm.CompoundPrimaryKey = false

	var whiteNotPkMeta df.DBMeta = WhiteNotPkDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["WhiteNotPk"] = &whiteNotPkMeta
}
