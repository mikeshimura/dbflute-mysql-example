package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorInheritInuDbm_T struct {
	df.BaseDBMeta
	ColumnInuId *df.ColumnInfo
	ColumnInuName *df.ColumnInfo
	ColumnInuDate *df.ColumnInfo
}

func (b *VendorInheritInuDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *VendorInheritInuDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var VendorInheritInuDbm *VendorInheritInuDbm_T

func Create_VendorInheritInuDbm() {
	VendorInheritInuDbm = new(VendorInheritInuDbm_T)
	VendorInheritInuDbm.TableDbName = "vendor_inherit_inu"
	VendorInheritInuDbm.TableDispName = "vendor_inherit_inu"
	VendorInheritInuDbm.TablePropertyName = "vendorInheritInu"
	VendorInheritInuDbm.TableSqlName = new(df.TableSqlName)
	VendorInheritInuDbm.TableSqlName.TableSqlName = "vendor_inherit_inu"
	VendorInheritInuDbm.TableSqlName.CorrespondingDbName = VendorInheritInuDbm.TableDbName

	var vendorInheritInu df.DBMeta
	vendorInheritInu = VendorInheritInuDbm
	VendorInheritInuDbm.DBMeta=&vendorInheritInu
	inuIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo inu_id
	inuIdSqlName.ColumnSqlName = "inu_id"
	inuIdSqlName.IrregularChar = false
	VendorInheritInuDbm.ColumnInuId = df.CCI(&vendorInheritInu, "inu_id", inuIdSqlName, "", "", "Integer.class", "inuId", "", true, false,true, "int4", 10, 0, "",false,"","", "","","",false,"int64")
	inuNameSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo inu_name
	inuNameSqlName.ColumnSqlName = "inu_name"
	inuNameSqlName.IrregularChar = false
	VendorInheritInuDbm.ColumnInuName = df.CCI(&vendorInheritInu, "inu_name", inuNameSqlName, "", "", "String.class", "inuName", "", false, false,true, "varchar", 2147483647, 0, "",false,"","", "","","",false,"string")
	inuDateSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo inu_date
	inuDateSqlName.ColumnSqlName = "inu_date"
	inuDateSqlName.IrregularChar = false
	VendorInheritInuDbm.ColumnInuDate = df.CCI(&vendorInheritInu, "inu_date", inuDateSqlName, "", "", "java.time.LocalDate.class", "inuDate", "", false, false,false, "date", 13, 0, "",false,"","", "","","",false,"df.NullDate")

	VendorInheritInuDbm.ColumnInfoList = new(df.List)
	VendorInheritInuDbm.ColumnInfoList.Add(VendorInheritInuDbm.ColumnInuId)
	VendorInheritInuDbm.ColumnInfoList.Add(VendorInheritInuDbm.ColumnInuName)
	VendorInheritInuDbm.ColumnInfoList.Add(VendorInheritInuDbm.ColumnInuDate)


	VendorInheritInuDbm.ColumnInfoMap=make(map[string]int)
	VendorInheritInuDbm.ColumnInfoMap["inuId"]=0
		VendorInheritInuDbm.ColumnInfoMap["inuName"]=1
		VendorInheritInuDbm.ColumnInfoMap["inuDate"]=2
	    VendorInheritInuDbm.PrimaryKey = true
    VendorInheritInuDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &vendorInheritInu
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(VendorInheritInuDbm.ColumnInuId)

	VendorInheritInuDbm.PrimaryInfo = new(df.PrimaryInfo)
	VendorInheritInuDbm.PrimaryInfo.UniqueInfo = ui
	
	var vendorInheritInuMeta df.DBMeta = VendorInheritInuDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["VendorInheritInu"] = &vendorInheritInuMeta
}
