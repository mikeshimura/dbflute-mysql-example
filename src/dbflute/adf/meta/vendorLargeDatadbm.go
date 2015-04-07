package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorLargeDataDbm_T struct {
	df.BaseDBMeta
	ColumnLargeDataId *df.ColumnInfo
	ColumnStringIndex *df.ColumnInfo
	ColumnStringNoIndex *df.ColumnInfo
	ColumnStringUniqueIndex *df.ColumnInfo
	ColumnIntflgIndex *df.ColumnInfo
	ColumnNumericIntegerIndex *df.ColumnInfo
	ColumnNumericIntegerNoIndex *df.ColumnInfo
}

func (b *VendorLargeDataDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *VendorLargeDataDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var VendorLargeDataDbm *VendorLargeDataDbm_T

func Create_VendorLargeDataDbm() {
	VendorLargeDataDbm = new(VendorLargeDataDbm_T)
	VendorLargeDataDbm.TableDbName = "vendor_large_data"
	VendorLargeDataDbm.TableDispName = "vendor_large_data"
	VendorLargeDataDbm.TablePropertyName = "vendorLargeData"
	VendorLargeDataDbm.TableSqlName = new(df.TableSqlName)
	VendorLargeDataDbm.TableSqlName.TableSqlName = "vendor_large_data"
	VendorLargeDataDbm.TableSqlName.CorrespondingDbName = VendorLargeDataDbm.TableDbName

	var vendorLargeData df.DBMeta
	vendorLargeData = VendorLargeDataDbm
	VendorLargeDataDbm.DBMeta=&vendorLargeData
	largeDataIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo large_data_id
	largeDataIdSqlName.ColumnSqlName = "large_data_id"
	largeDataIdSqlName.IrregularChar = false
	VendorLargeDataDbm.ColumnLargeDataId = df.CCI(&vendorLargeData, "large_data_id", largeDataIdSqlName, "", "", "Long.class", "largeDataId", "", true, false,true, "int8", 19, 0, "",false,"","", "","vendorLargeDataRefList","",false,"int64")
	stringIndexSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo string_index
	stringIndexSqlName.ColumnSqlName = "string_index"
	stringIndexSqlName.IrregularChar = false
	VendorLargeDataDbm.ColumnStringIndex = df.CCI(&vendorLargeData, "string_index", stringIndexSqlName, "", "", "String.class", "stringIndex", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	stringNoIndexSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo string_no_index
	stringNoIndexSqlName.ColumnSqlName = "string_no_index"
	stringNoIndexSqlName.IrregularChar = false
	VendorLargeDataDbm.ColumnStringNoIndex = df.CCI(&vendorLargeData, "string_no_index", stringNoIndexSqlName, "", "", "String.class", "stringNoIndex", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	stringUniqueIndexSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo string_unique_index
	stringUniqueIndexSqlName.ColumnSqlName = "string_unique_index"
	stringUniqueIndexSqlName.IrregularChar = false
	VendorLargeDataDbm.ColumnStringUniqueIndex = df.CCI(&vendorLargeData, "string_unique_index", stringUniqueIndexSqlName, "", "", "String.class", "stringUniqueIndex", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	intflgIndexSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo intflg_index
	intflgIndexSqlName.ColumnSqlName = "intflg_index"
	intflgIndexSqlName.IrregularChar = false
	VendorLargeDataDbm.ColumnIntflgIndex = df.CCI(&vendorLargeData, "intflg_index", intflgIndexSqlName, "", "", "Integer.class", "intflgIndex", "", false, false,true, "int4", 10, 0, "",false,"","", "","","",false,"int64")
	numericIntegerIndexSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo numeric_integer_index
	numericIntegerIndexSqlName.ColumnSqlName = "numeric_integer_index"
	numericIntegerIndexSqlName.IrregularChar = false
	VendorLargeDataDbm.ColumnNumericIntegerIndex = df.CCI(&vendorLargeData, "numeric_integer_index", numericIntegerIndexSqlName, "", "", "Integer.class", "numericIntegerIndex", "", false, false,true, "numeric", 8, 0, "",false,"","", "","","",false,"df.Numeric")
	numericIntegerNoIndexSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo numeric_integer_no_index
	numericIntegerNoIndexSqlName.ColumnSqlName = "numeric_integer_no_index"
	numericIntegerNoIndexSqlName.IrregularChar = false
	VendorLargeDataDbm.ColumnNumericIntegerNoIndex = df.CCI(&vendorLargeData, "numeric_integer_no_index", numericIntegerNoIndexSqlName, "", "", "Integer.class", "numericIntegerNoIndex", "", false, false,true, "numeric", 8, 0, "",false,"","", "","","",false,"df.Numeric")

	VendorLargeDataDbm.ColumnInfoList = new(df.List)
	VendorLargeDataDbm.ColumnInfoList.Add(VendorLargeDataDbm.ColumnLargeDataId)
	VendorLargeDataDbm.ColumnInfoList.Add(VendorLargeDataDbm.ColumnStringIndex)
	VendorLargeDataDbm.ColumnInfoList.Add(VendorLargeDataDbm.ColumnStringNoIndex)
	VendorLargeDataDbm.ColumnInfoList.Add(VendorLargeDataDbm.ColumnStringUniqueIndex)
	VendorLargeDataDbm.ColumnInfoList.Add(VendorLargeDataDbm.ColumnIntflgIndex)
	VendorLargeDataDbm.ColumnInfoList.Add(VendorLargeDataDbm.ColumnNumericIntegerIndex)
	VendorLargeDataDbm.ColumnInfoList.Add(VendorLargeDataDbm.ColumnNumericIntegerNoIndex)


	VendorLargeDataDbm.ColumnInfoMap=make(map[string]int)
	VendorLargeDataDbm.ColumnInfoMap["largeDataId"]=0
		VendorLargeDataDbm.ColumnInfoMap["stringIndex"]=1
		VendorLargeDataDbm.ColumnInfoMap["stringNoIndex"]=2
		VendorLargeDataDbm.ColumnInfoMap["stringUniqueIndex"]=3
		VendorLargeDataDbm.ColumnInfoMap["intflgIndex"]=4
		VendorLargeDataDbm.ColumnInfoMap["numericIntegerIndex"]=5
		VendorLargeDataDbm.ColumnInfoMap["numericIntegerNoIndex"]=6
	    VendorLargeDataDbm.PrimaryKey = true
    VendorLargeDataDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &vendorLargeData
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(VendorLargeDataDbm.ColumnLargeDataId)

	VendorLargeDataDbm.PrimaryInfo = new(df.PrimaryInfo)
	VendorLargeDataDbm.PrimaryInfo.UniqueInfo = ui
	
	var vendorLargeDataMeta df.DBMeta = VendorLargeDataDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["VendorLargeData"] = &vendorLargeDataMeta
}
