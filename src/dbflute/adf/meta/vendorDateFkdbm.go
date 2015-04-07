package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorDateFkDbm_T struct {
	df.BaseDBMeta
	ColumnBarId *df.ColumnInfo
	ColumnBarDate *df.ColumnInfo
}

func (b *VendorDateFkDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *VendorDateFkDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var VendorDateFkDbm *VendorDateFkDbm_T

func Create_VendorDateFkDbm() {
	VendorDateFkDbm = new(VendorDateFkDbm_T)
	VendorDateFkDbm.TableDbName = "vendor_date_fk"
	VendorDateFkDbm.TableDispName = "vendor_date_fk"
	VendorDateFkDbm.TablePropertyName = "vendorDateFk"
	VendorDateFkDbm.TableSqlName = new(df.TableSqlName)
	VendorDateFkDbm.TableSqlName.TableSqlName = "vendor_date_fk"
	VendorDateFkDbm.TableSqlName.CorrespondingDbName = VendorDateFkDbm.TableDbName

	var vendorDateFk df.DBMeta
	vendorDateFk = VendorDateFkDbm
	VendorDateFkDbm.DBMeta=&vendorDateFk
	barIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo bar_id
	barIdSqlName.ColumnSqlName = "bar_id"
	barIdSqlName.IrregularChar = false
	VendorDateFkDbm.ColumnBarId = df.CCI(&vendorDateFk, "bar_id", barIdSqlName, "", "", "Integer.class", "barId", "", true, false,true, "int4", 10, 0, "",false,"","", "","","",false,"int64")
	barDateSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo bar_date
	barDateSqlName.ColumnSqlName = "bar_date"
	barDateSqlName.IrregularChar = false
	VendorDateFkDbm.ColumnBarDate = df.CCI(&vendorDateFk, "bar_date", barDateSqlName, "", "", "java.time.LocalDate.class", "barDate", "", false, false,true, "date", 13, 0, "",false,"","", "vendorDatePk","","",false,"df.Date")

	VendorDateFkDbm.ColumnInfoList = new(df.List)
	VendorDateFkDbm.ColumnInfoList.Add(VendorDateFkDbm.ColumnBarId)
	VendorDateFkDbm.ColumnInfoList.Add(VendorDateFkDbm.ColumnBarDate)


	VendorDateFkDbm.ColumnInfoMap=make(map[string]int)
	VendorDateFkDbm.ColumnInfoMap["barId"]=0
		VendorDateFkDbm.ColumnInfoMap["barDate"]=1
	    VendorDateFkDbm.PrimaryKey = true
    VendorDateFkDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &vendorDateFk
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(VendorDateFkDbm.ColumnBarId)

	VendorDateFkDbm.PrimaryInfo = new(df.PrimaryInfo)
	VendorDateFkDbm.PrimaryInfo.UniqueInfo = ui
	
	var vendorDateFkMeta df.DBMeta = VendorDateFkDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["VendorDateFk"] = &vendorDateFkMeta
}
