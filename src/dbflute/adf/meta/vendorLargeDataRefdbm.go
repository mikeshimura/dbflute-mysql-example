package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorLargeDataRefDbm_T struct {
	df.BaseDBMeta
	ColumnLargeDataRefId *df.ColumnInfo
	ColumnLargeDataId *df.ColumnInfo
	ColumnDateIndex *df.ColumnInfo
	ColumnDateNoIndex *df.ColumnInfo
	ColumnTimestampIndex *df.ColumnInfo
	ColumnTimestampNoIndex *df.ColumnInfo
	ColumnNullableDecimalIndex *df.ColumnInfo
	ColumnNullableDecimalNoIndex *df.ColumnInfo
	ColumnSelfParentId *df.ColumnInfo
}

func (b *VendorLargeDataRefDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *VendorLargeDataRefDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var VendorLargeDataRefDbm *VendorLargeDataRefDbm_T

func Create_VendorLargeDataRefDbm() {
	VendorLargeDataRefDbm = new(VendorLargeDataRefDbm_T)
	VendorLargeDataRefDbm.TableDbName = "vendor_large_data_ref"
	VendorLargeDataRefDbm.TableDispName = "vendor_large_data_ref"
	VendorLargeDataRefDbm.TablePropertyName = "vendorLargeDataRef"
	VendorLargeDataRefDbm.TableSqlName = new(df.TableSqlName)
	VendorLargeDataRefDbm.TableSqlName.TableSqlName = "vendor_large_data_ref"
	VendorLargeDataRefDbm.TableSqlName.CorrespondingDbName = VendorLargeDataRefDbm.TableDbName

	var vendorLargeDataRef df.DBMeta
	vendorLargeDataRef = VendorLargeDataRefDbm
	VendorLargeDataRefDbm.DBMeta=&vendorLargeDataRef
	largeDataRefIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo large_data_ref_id
	largeDataRefIdSqlName.ColumnSqlName = "large_data_ref_id"
	largeDataRefIdSqlName.IrregularChar = false
	VendorLargeDataRefDbm.ColumnLargeDataRefId = df.CCI(&vendorLargeDataRef, "large_data_ref_id", largeDataRefIdSqlName, "", "", "Long.class", "largeDataRefId", "", true, false,true, "int8", 19, 0, "",false,"","", "","vendorLargeDataRefSelfList","",false,"int64")
	largeDataIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo large_data_id
	largeDataIdSqlName.ColumnSqlName = "large_data_id"
	largeDataIdSqlName.IrregularChar = false
	VendorLargeDataRefDbm.ColumnLargeDataId = df.CCI(&vendorLargeDataRef, "large_data_id", largeDataIdSqlName, "", "", "Long.class", "largeDataId", "", false, false,true, "int8", 19, 0, "",false,"","", "vendorLargeData","","",false,"int64")
	dateIndexSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo date_index
	dateIndexSqlName.ColumnSqlName = "date_index"
	dateIndexSqlName.IrregularChar = false
	VendorLargeDataRefDbm.ColumnDateIndex = df.CCI(&vendorLargeDataRef, "date_index", dateIndexSqlName, "", "", "java.time.LocalDate.class", "dateIndex", "", false, false,true, "date", 13, 0, "",false,"","", "","","",false,"df.Date")
	dateNoIndexSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo date_no_index
	dateNoIndexSqlName.ColumnSqlName = "date_no_index"
	dateNoIndexSqlName.IrregularChar = false
	VendorLargeDataRefDbm.ColumnDateNoIndex = df.CCI(&vendorLargeDataRef, "date_no_index", dateNoIndexSqlName, "", "", "java.time.LocalDate.class", "dateNoIndex", "", false, false,true, "date", 13, 0, "",false,"","", "","","",false,"df.Date")
	timestampIndexSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo timestamp_index
	timestampIndexSqlName.ColumnSqlName = "timestamp_index"
	timestampIndexSqlName.IrregularChar = false
	VendorLargeDataRefDbm.ColumnTimestampIndex = df.CCI(&vendorLargeDataRef, "timestamp_index", timestampIndexSqlName, "", "", "java.time.LocalDateTime.class", "timestampIndex", "", false, false,true, "timestamp", 26, 3, "",false,"","", "","","",false,"df.Timestamp")
	timestampNoIndexSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo timestamp_no_index
	timestampNoIndexSqlName.ColumnSqlName = "timestamp_no_index"
	timestampNoIndexSqlName.IrregularChar = false
	VendorLargeDataRefDbm.ColumnTimestampNoIndex = df.CCI(&vendorLargeDataRef, "timestamp_no_index", timestampNoIndexSqlName, "", "", "java.time.LocalDateTime.class", "timestampNoIndex", "", false, false,true, "timestamp", 26, 3, "",false,"","", "","","",false,"df.Timestamp")
	nullableDecimalIndexSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo nullable_decimal_index
	nullableDecimalIndexSqlName.ColumnSqlName = "nullable_decimal_index"
	nullableDecimalIndexSqlName.IrregularChar = false
	VendorLargeDataRefDbm.ColumnNullableDecimalIndex = df.CCI(&vendorLargeDataRef, "nullable_decimal_index", nullableDecimalIndexSqlName, "", "", "java.math.BigDecimal.class", "nullableDecimalIndex", "", false, false,false, "numeric", 12, 3, "",false,"","", "","","",false,"df.NullNumeric")
	nullableDecimalNoIndexSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo nullable_decimal_no_index
	nullableDecimalNoIndexSqlName.ColumnSqlName = "nullable_decimal_no_index"
	nullableDecimalNoIndexSqlName.IrregularChar = false
	VendorLargeDataRefDbm.ColumnNullableDecimalNoIndex = df.CCI(&vendorLargeDataRef, "nullable_decimal_no_index", nullableDecimalNoIndexSqlName, "", "", "java.math.BigDecimal.class", "nullableDecimalNoIndex", "", false, false,false, "numeric", 12, 3, "",false,"","", "","","",false,"df.NullNumeric")
	selfParentIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo self_parent_id
	selfParentIdSqlName.ColumnSqlName = "self_parent_id"
	selfParentIdSqlName.IrregularChar = false
	VendorLargeDataRefDbm.ColumnSelfParentId = df.CCI(&vendorLargeDataRef, "self_parent_id", selfParentIdSqlName, "", "", "Long.class", "selfParentId", "", false, false,false, "int8", 19, 0, "",false,"","", "vendorLargeDataRefSelf","","",false,"sql.NullInt64")

	VendorLargeDataRefDbm.ColumnInfoList = new(df.List)
	VendorLargeDataRefDbm.ColumnInfoList.Add(VendorLargeDataRefDbm.ColumnLargeDataRefId)
	VendorLargeDataRefDbm.ColumnInfoList.Add(VendorLargeDataRefDbm.ColumnLargeDataId)
	VendorLargeDataRefDbm.ColumnInfoList.Add(VendorLargeDataRefDbm.ColumnDateIndex)
	VendorLargeDataRefDbm.ColumnInfoList.Add(VendorLargeDataRefDbm.ColumnDateNoIndex)
	VendorLargeDataRefDbm.ColumnInfoList.Add(VendorLargeDataRefDbm.ColumnTimestampIndex)
	VendorLargeDataRefDbm.ColumnInfoList.Add(VendorLargeDataRefDbm.ColumnTimestampNoIndex)
	VendorLargeDataRefDbm.ColumnInfoList.Add(VendorLargeDataRefDbm.ColumnNullableDecimalIndex)
	VendorLargeDataRefDbm.ColumnInfoList.Add(VendorLargeDataRefDbm.ColumnNullableDecimalNoIndex)
	VendorLargeDataRefDbm.ColumnInfoList.Add(VendorLargeDataRefDbm.ColumnSelfParentId)


	VendorLargeDataRefDbm.ColumnInfoMap=make(map[string]int)
	VendorLargeDataRefDbm.ColumnInfoMap["largeDataRefId"]=0
		VendorLargeDataRefDbm.ColumnInfoMap["largeDataId"]=1
		VendorLargeDataRefDbm.ColumnInfoMap["dateIndex"]=2
		VendorLargeDataRefDbm.ColumnInfoMap["dateNoIndex"]=3
		VendorLargeDataRefDbm.ColumnInfoMap["timestampIndex"]=4
		VendorLargeDataRefDbm.ColumnInfoMap["timestampNoIndex"]=5
		VendorLargeDataRefDbm.ColumnInfoMap["nullableDecimalIndex"]=6
		VendorLargeDataRefDbm.ColumnInfoMap["nullableDecimalNoIndex"]=7
		VendorLargeDataRefDbm.ColumnInfoMap["selfParentId"]=8
	    VendorLargeDataRefDbm.PrimaryKey = true
    VendorLargeDataRefDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &vendorLargeDataRef
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(VendorLargeDataRefDbm.ColumnLargeDataRefId)

	VendorLargeDataRefDbm.PrimaryInfo = new(df.PrimaryInfo)
	VendorLargeDataRefDbm.PrimaryInfo.UniqueInfo = ui
	
	var vendorLargeDataRefMeta df.DBMeta = VendorLargeDataRefDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["VendorLargeDataRef"] = &vendorLargeDataRefMeta
}
