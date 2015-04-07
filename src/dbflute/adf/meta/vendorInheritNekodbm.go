package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorInheritNekoDbm_T struct {
	df.BaseDBMeta
	ColumnInuId *df.ColumnInfo
	ColumnInuName *df.ColumnInfo
	ColumnInuDate *df.ColumnInfo
	ColumnNekoId *df.ColumnInfo
	ColumnNekoName *df.ColumnInfo
	ColumnNekoDate *df.ColumnInfo
}

func (b *VendorInheritNekoDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *VendorInheritNekoDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var VendorInheritNekoDbm *VendorInheritNekoDbm_T

func Create_VendorInheritNekoDbm() {
	VendorInheritNekoDbm = new(VendorInheritNekoDbm_T)
	VendorInheritNekoDbm.TableDbName = "vendor_inherit_neko"
	VendorInheritNekoDbm.TableDispName = "vendor_inherit_neko"
	VendorInheritNekoDbm.TablePropertyName = "vendorInheritNeko"
	VendorInheritNekoDbm.TableSqlName = new(df.TableSqlName)
	VendorInheritNekoDbm.TableSqlName.TableSqlName = "vendor_inherit_neko"
	VendorInheritNekoDbm.TableSqlName.CorrespondingDbName = VendorInheritNekoDbm.TableDbName

	var vendorInheritNeko df.DBMeta
	vendorInheritNeko = VendorInheritNekoDbm
	VendorInheritNekoDbm.DBMeta=&vendorInheritNeko
	inuIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo inu_id
	inuIdSqlName.ColumnSqlName = "inu_id"
	inuIdSqlName.IrregularChar = false
	VendorInheritNekoDbm.ColumnInuId = df.CCI(&vendorInheritNeko, "inu_id", inuIdSqlName, "", "", "Integer.class", "inuId", "", false, false,true, "int4", 10, 0, "",false,"","", "","","",false,"int64")
	inuNameSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo inu_name
	inuNameSqlName.ColumnSqlName = "inu_name"
	inuNameSqlName.IrregularChar = false
	VendorInheritNekoDbm.ColumnInuName = df.CCI(&vendorInheritNeko, "inu_name", inuNameSqlName, "", "", "String.class", "inuName", "", false, false,true, "varchar", 2147483647, 0, "",false,"","", "","","",false,"string")
	inuDateSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo inu_date
	inuDateSqlName.ColumnSqlName = "inu_date"
	inuDateSqlName.IrregularChar = false
	VendorInheritNekoDbm.ColumnInuDate = df.CCI(&vendorInheritNeko, "inu_date", inuDateSqlName, "", "", "java.time.LocalDate.class", "inuDate", "", false, false,false, "date", 13, 0, "",false,"","", "","","",false,"df.NullDate")
	nekoIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo neko_id
	nekoIdSqlName.ColumnSqlName = "neko_id"
	nekoIdSqlName.IrregularChar = false
	VendorInheritNekoDbm.ColumnNekoId = df.CCI(&vendorInheritNeko, "neko_id", nekoIdSqlName, "", "", "Integer.class", "nekoId", "", true, false,true, "int4", 10, 0, "",false,"","", "","","",false,"int64")
	nekoNameSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo neko_name
	nekoNameSqlName.ColumnSqlName = "neko_name"
	nekoNameSqlName.IrregularChar = false
	VendorInheritNekoDbm.ColumnNekoName = df.CCI(&vendorInheritNeko, "neko_name", nekoNameSqlName, "", "", "String.class", "nekoName", "", false, false,true, "varchar", 2147483647, 0, "",false,"","", "","","",false,"string")
	nekoDateSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo neko_date
	nekoDateSqlName.ColumnSqlName = "neko_date"
	nekoDateSqlName.IrregularChar = false
	VendorInheritNekoDbm.ColumnNekoDate = df.CCI(&vendorInheritNeko, "neko_date", nekoDateSqlName, "", "", "java.time.LocalDate.class", "nekoDate", "", false, false,false, "date", 13, 0, "",false,"","", "","","",false,"df.NullDate")

	VendorInheritNekoDbm.ColumnInfoList = new(df.List)
	VendorInheritNekoDbm.ColumnInfoList.Add(VendorInheritNekoDbm.ColumnInuId)
	VendorInheritNekoDbm.ColumnInfoList.Add(VendorInheritNekoDbm.ColumnInuName)
	VendorInheritNekoDbm.ColumnInfoList.Add(VendorInheritNekoDbm.ColumnInuDate)
	VendorInheritNekoDbm.ColumnInfoList.Add(VendorInheritNekoDbm.ColumnNekoId)
	VendorInheritNekoDbm.ColumnInfoList.Add(VendorInheritNekoDbm.ColumnNekoName)
	VendorInheritNekoDbm.ColumnInfoList.Add(VendorInheritNekoDbm.ColumnNekoDate)


	VendorInheritNekoDbm.ColumnInfoMap=make(map[string]int)
	VendorInheritNekoDbm.ColumnInfoMap["inuId"]=0
		VendorInheritNekoDbm.ColumnInfoMap["inuName"]=1
		VendorInheritNekoDbm.ColumnInfoMap["inuDate"]=2
		VendorInheritNekoDbm.ColumnInfoMap["nekoId"]=3
		VendorInheritNekoDbm.ColumnInfoMap["nekoName"]=4
		VendorInheritNekoDbm.ColumnInfoMap["nekoDate"]=5
	    VendorInheritNekoDbm.PrimaryKey = true
    VendorInheritNekoDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &vendorInheritNeko
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(VendorInheritNekoDbm.ColumnNekoId)

	VendorInheritNekoDbm.PrimaryInfo = new(df.PrimaryInfo)
	VendorInheritNekoDbm.PrimaryInfo.UniqueInfo = ui
	
	var vendorInheritNekoMeta df.DBMeta = VendorInheritNekoDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["VendorInheritNeko"] = &vendorInheritNekoMeta
}
