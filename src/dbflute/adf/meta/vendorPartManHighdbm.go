package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorPartManHighDbm_T struct {
	df.BaseDBMeta
	ColumnPartManId *df.ColumnInfo
	ColumnPartManName *df.ColumnInfo
	ColumnPartManPoint *df.ColumnInfo
	ColumnPartManDate *df.ColumnInfo
}

func (b *VendorPartManHighDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *VendorPartManHighDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var VendorPartManHighDbm *VendorPartManHighDbm_T

func Create_VendorPartManHighDbm() {
	VendorPartManHighDbm = new(VendorPartManHighDbm_T)
	VendorPartManHighDbm.TableDbName = "vendor_part_man_high"
	VendorPartManHighDbm.TableDispName = "vendor_part_man_high"
	VendorPartManHighDbm.TablePropertyName = "vendorPartManHigh"
	VendorPartManHighDbm.TableSqlName = new(df.TableSqlName)
	VendorPartManHighDbm.TableSqlName.TableSqlName = "vendor_part_man_high"
	VendorPartManHighDbm.TableSqlName.CorrespondingDbName = VendorPartManHighDbm.TableDbName

	var vendorPartManHigh df.DBMeta
	vendorPartManHigh = VendorPartManHighDbm
	VendorPartManHighDbm.DBMeta=&vendorPartManHigh
	partManIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo part_man_id
	partManIdSqlName.ColumnSqlName = "part_man_id"
	partManIdSqlName.IrregularChar = false
	VendorPartManHighDbm.ColumnPartManId = df.CCI(&vendorPartManHigh, "part_man_id", partManIdSqlName, "", "", "Integer.class", "partManId", "", true, false,true, "int4", 10, 0, "",false,"","", "","","",false,"int64")
	partManNameSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo part_man_name
	partManNameSqlName.ColumnSqlName = "part_man_name"
	partManNameSqlName.IrregularChar = false
	VendorPartManHighDbm.ColumnPartManName = df.CCI(&vendorPartManHigh, "part_man_name", partManNameSqlName, "", "", "String.class", "partManName", "", false, false,true, "varchar", 2147483647, 0, "",false,"","", "","","",false,"string")
	partManPointSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo part_man_point
	partManPointSqlName.ColumnSqlName = "part_man_point"
	partManPointSqlName.IrregularChar = false
	VendorPartManHighDbm.ColumnPartManPoint = df.CCI(&vendorPartManHigh, "part_man_point", partManPointSqlName, "", "", "Integer.class", "partManPoint", "", false, false,true, "int4", 10, 0, "",false,"","", "","","",false,"int64")
	partManDateSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo part_man_date
	partManDateSqlName.ColumnSqlName = "part_man_date"
	partManDateSqlName.IrregularChar = false
	VendorPartManHighDbm.ColumnPartManDate = df.CCI(&vendorPartManHigh, "part_man_date", partManDateSqlName, "", "", "java.time.LocalDate.class", "partManDate", "", false, false,false, "date", 13, 0, "",false,"","", "","","",false,"df.NullDate")

	VendorPartManHighDbm.ColumnInfoList = new(df.List)
	VendorPartManHighDbm.ColumnInfoList.Add(VendorPartManHighDbm.ColumnPartManId)
	VendorPartManHighDbm.ColumnInfoList.Add(VendorPartManHighDbm.ColumnPartManName)
	VendorPartManHighDbm.ColumnInfoList.Add(VendorPartManHighDbm.ColumnPartManPoint)
	VendorPartManHighDbm.ColumnInfoList.Add(VendorPartManHighDbm.ColumnPartManDate)


	VendorPartManHighDbm.ColumnInfoMap=make(map[string]int)
	VendorPartManHighDbm.ColumnInfoMap["partManId"]=0
		VendorPartManHighDbm.ColumnInfoMap["partManName"]=1
		VendorPartManHighDbm.ColumnInfoMap["partManPoint"]=2
		VendorPartManHighDbm.ColumnInfoMap["partManDate"]=3
	    VendorPartManHighDbm.PrimaryKey = true
    VendorPartManHighDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &vendorPartManHigh
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(VendorPartManHighDbm.ColumnPartManId)

	VendorPartManHighDbm.PrimaryInfo = new(df.PrimaryInfo)
	VendorPartManHighDbm.PrimaryInfo.UniqueInfo = ui
	
	var vendorPartManHighMeta df.DBMeta = VendorPartManHighDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["VendorPartManHigh"] = &vendorPartManHighMeta
}
