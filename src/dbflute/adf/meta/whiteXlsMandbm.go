package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type WhiteXlsManDbm_T struct {
	df.BaseDBMeta
	ColumnXlsManId *df.ColumnInfo
	ColumnStringConverted *df.ColumnInfo
	ColumnTimestampZeroDefaultMillis *df.ColumnInfo
	ColumnTimestampZeroPrefixMillis *df.ColumnInfo
	ColumnTimestampZeroSuffixMillis *df.ColumnInfo
}

func (b *WhiteXlsManDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *WhiteXlsManDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var WhiteXlsManDbm *WhiteXlsManDbm_T

func Create_WhiteXlsManDbm() {
	WhiteXlsManDbm = new(WhiteXlsManDbm_T)
	WhiteXlsManDbm.TableDbName = "white_xls_man"
	WhiteXlsManDbm.TableDispName = "white_xls_man"
	WhiteXlsManDbm.TablePropertyName = "whiteXlsMan"
	WhiteXlsManDbm.TableSqlName = new(df.TableSqlName)
	WhiteXlsManDbm.TableSqlName.TableSqlName = "white_xls_man"
	WhiteXlsManDbm.TableSqlName.CorrespondingDbName = WhiteXlsManDbm.TableDbName

	var whiteXlsMan df.DBMeta
	whiteXlsMan = WhiteXlsManDbm
	WhiteXlsManDbm.DBMeta=&whiteXlsMan
	xlsManIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo xls_man_id
	xlsManIdSqlName.ColumnSqlName = "xls_man_id"
	xlsManIdSqlName.IrregularChar = false
	WhiteXlsManDbm.ColumnXlsManId = df.CCI(&whiteXlsMan, "xls_man_id", xlsManIdSqlName, "", "", "Long.class", "xlsManId", "", true, false,true, "int8", 19, 0, "",false,"","", "","","",false,"int64")
	stringConvertedSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo string_converted
	stringConvertedSqlName.ColumnSqlName = "string_converted"
	stringConvertedSqlName.IrregularChar = false
	WhiteXlsManDbm.ColumnStringConverted = df.CCI(&whiteXlsMan, "string_converted", stringConvertedSqlName, "", "", "String.class", "stringConverted", "", false, false,false, "varchar", 2147483647, 0, "",false,"","", "","","",false,"sql.NullString")
	timestampZeroDefaultMillisSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo timestamp_zero_default_millis
	timestampZeroDefaultMillisSqlName.ColumnSqlName = "timestamp_zero_default_millis"
	timestampZeroDefaultMillisSqlName.IrregularChar = false
	WhiteXlsManDbm.ColumnTimestampZeroDefaultMillis = df.CCI(&whiteXlsMan, "timestamp_zero_default_millis", timestampZeroDefaultMillisSqlName, "", "", "java.time.LocalDateTime.class", "timestampZeroDefaultMillis", "", false, false,false, "timestamp", 26, 3, "",false,"","", "","","",false,"df.NullTimestamp")
	timestampZeroPrefixMillisSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo timestamp_zero_prefix_millis
	timestampZeroPrefixMillisSqlName.ColumnSqlName = "timestamp_zero_prefix_millis"
	timestampZeroPrefixMillisSqlName.IrregularChar = false
	WhiteXlsManDbm.ColumnTimestampZeroPrefixMillis = df.CCI(&whiteXlsMan, "timestamp_zero_prefix_millis", timestampZeroPrefixMillisSqlName, "", "", "java.time.LocalDateTime.class", "timestampZeroPrefixMillis", "", false, false,false, "timestamp", 26, 3, "",false,"","", "","","",false,"df.NullTimestamp")
	timestampZeroSuffixMillisSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo timestamp_zero_suffix_millis
	timestampZeroSuffixMillisSqlName.ColumnSqlName = "timestamp_zero_suffix_millis"
	timestampZeroSuffixMillisSqlName.IrregularChar = false
	WhiteXlsManDbm.ColumnTimestampZeroSuffixMillis = df.CCI(&whiteXlsMan, "timestamp_zero_suffix_millis", timestampZeroSuffixMillisSqlName, "", "", "java.time.LocalDateTime.class", "timestampZeroSuffixMillis", "", false, false,false, "timestamp", 26, 3, "",false,"","", "","","",false,"df.NullTimestamp")

	WhiteXlsManDbm.ColumnInfoList = new(df.List)
	WhiteXlsManDbm.ColumnInfoList.Add(WhiteXlsManDbm.ColumnXlsManId)
	WhiteXlsManDbm.ColumnInfoList.Add(WhiteXlsManDbm.ColumnStringConverted)
	WhiteXlsManDbm.ColumnInfoList.Add(WhiteXlsManDbm.ColumnTimestampZeroDefaultMillis)
	WhiteXlsManDbm.ColumnInfoList.Add(WhiteXlsManDbm.ColumnTimestampZeroPrefixMillis)
	WhiteXlsManDbm.ColumnInfoList.Add(WhiteXlsManDbm.ColumnTimestampZeroSuffixMillis)


	WhiteXlsManDbm.ColumnInfoMap=make(map[string]int)
	WhiteXlsManDbm.ColumnInfoMap["xlsManId"]=0
		WhiteXlsManDbm.ColumnInfoMap["stringConverted"]=1
		WhiteXlsManDbm.ColumnInfoMap["timestampZeroDefaultMillis"]=2
		WhiteXlsManDbm.ColumnInfoMap["timestampZeroPrefixMillis"]=3
		WhiteXlsManDbm.ColumnInfoMap["timestampZeroSuffixMillis"]=4
	    WhiteXlsManDbm.PrimaryKey = true
    WhiteXlsManDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &whiteXlsMan
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(WhiteXlsManDbm.ColumnXlsManId)

	WhiteXlsManDbm.PrimaryInfo = new(df.PrimaryInfo)
	WhiteXlsManDbm.PrimaryInfo.UniqueInfo = ui
	
	var whiteXlsManMeta df.DBMeta = WhiteXlsManDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["WhiteXlsMan"] = &whiteXlsManMeta
}
