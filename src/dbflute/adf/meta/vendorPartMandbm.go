package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorPartManDbm_T struct {
	df.BaseDBMeta
	ColumnPartManId *df.ColumnInfo
	ColumnPartManName *df.ColumnInfo
	ColumnPartManPoint *df.ColumnInfo
	ColumnPartManDate *df.ColumnInfo
}

func (b *VendorPartManDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *VendorPartManDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var VendorPartManDbm *VendorPartManDbm_T

func Create_VendorPartManDbm() {
	VendorPartManDbm = new(VendorPartManDbm_T)
	VendorPartManDbm.TableDbName = "vendor_part_man"
	VendorPartManDbm.TableDispName = "vendor_part_man"
	VendorPartManDbm.TablePropertyName = "vendorPartMan"
	VendorPartManDbm.TableSqlName = new(df.TableSqlName)
	VendorPartManDbm.TableSqlName.TableSqlName = "vendor_part_man"
	VendorPartManDbm.TableSqlName.CorrespondingDbName = VendorPartManDbm.TableDbName

	var vendorPartMan df.DBMeta
	vendorPartMan = VendorPartManDbm
	VendorPartManDbm.DBMeta=&vendorPartMan
	partManIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo part_man_id
	partManIdSqlName.ColumnSqlName = "part_man_id"
	partManIdSqlName.IrregularChar = false
	VendorPartManDbm.ColumnPartManId = df.CCI(&vendorPartMan, "part_man_id", partManIdSqlName, "", "", "Integer.class", "partManId", "", true, false,true, "int4", 10, 0, "",false,"","", "","","",false,"int64")
	partManNameSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo part_man_name
	partManNameSqlName.ColumnSqlName = "part_man_name"
	partManNameSqlName.IrregularChar = false
	VendorPartManDbm.ColumnPartManName = df.CCI(&vendorPartMan, "part_man_name", partManNameSqlName, "", "", "String.class", "partManName", "", false, false,true, "varchar", 2147483647, 0, "",false,"","", "","","",false,"string")
	partManPointSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo part_man_point
	partManPointSqlName.ColumnSqlName = "part_man_point"
	partManPointSqlName.IrregularChar = false
	VendorPartManDbm.ColumnPartManPoint = df.CCI(&vendorPartMan, "part_man_point", partManPointSqlName, "", "", "Integer.class", "partManPoint", "", false, false,true, "int4", 10, 0, "",false,"","", "","","",false,"int64")
	partManDateSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo part_man_date
	partManDateSqlName.ColumnSqlName = "part_man_date"
	partManDateSqlName.IrregularChar = false
	VendorPartManDbm.ColumnPartManDate = df.CCI(&vendorPartMan, "part_man_date", partManDateSqlName, "", "", "java.time.LocalDate.class", "partManDate", "", false, false,false, "date", 13, 0, "",false,"","", "","","",false,"df.NullDate")

	VendorPartManDbm.ColumnInfoList = new(df.List)
	VendorPartManDbm.ColumnInfoList.Add(VendorPartManDbm.ColumnPartManId)
	VendorPartManDbm.ColumnInfoList.Add(VendorPartManDbm.ColumnPartManName)
	VendorPartManDbm.ColumnInfoList.Add(VendorPartManDbm.ColumnPartManPoint)
	VendorPartManDbm.ColumnInfoList.Add(VendorPartManDbm.ColumnPartManDate)


	VendorPartManDbm.ColumnInfoMap=make(map[string]int)
	VendorPartManDbm.ColumnInfoMap["partManId"]=0
		VendorPartManDbm.ColumnInfoMap["partManName"]=1
		VendorPartManDbm.ColumnInfoMap["partManPoint"]=2
		VendorPartManDbm.ColumnInfoMap["partManDate"]=3
	    VendorPartManDbm.PrimaryKey = true
    VendorPartManDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &vendorPartMan
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(VendorPartManDbm.ColumnPartManId)

	VendorPartManDbm.PrimaryInfo = new(df.PrimaryInfo)
	VendorPartManDbm.PrimaryInfo.UniqueInfo = ui
	
	var vendorPartManMeta df.DBMeta = VendorPartManDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["VendorPartMan"] = &vendorPartManMeta
}
