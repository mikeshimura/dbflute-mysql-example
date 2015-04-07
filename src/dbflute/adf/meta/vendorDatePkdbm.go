package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorDatePkDbm_T struct {
	df.BaseDBMeta
	ColumnFooDate *df.ColumnInfo
	ColumnFooName *df.ColumnInfo
}

func (b *VendorDatePkDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *VendorDatePkDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var VendorDatePkDbm *VendorDatePkDbm_T

func Create_VendorDatePkDbm() {
	VendorDatePkDbm = new(VendorDatePkDbm_T)
	VendorDatePkDbm.TableDbName = "vendor_date_pk"
	VendorDatePkDbm.TableDispName = "vendor_date_pk"
	VendorDatePkDbm.TablePropertyName = "vendorDatePk"
	VendorDatePkDbm.TableSqlName = new(df.TableSqlName)
	VendorDatePkDbm.TableSqlName.TableSqlName = "vendor_date_pk"
	VendorDatePkDbm.TableSqlName.CorrespondingDbName = VendorDatePkDbm.TableDbName

	var vendorDatePk df.DBMeta
	vendorDatePk = VendorDatePkDbm
	VendorDatePkDbm.DBMeta=&vendorDatePk
	fooDateSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo foo_date
	fooDateSqlName.ColumnSqlName = "foo_date"
	fooDateSqlName.IrregularChar = false
	VendorDatePkDbm.ColumnFooDate = df.CCI(&vendorDatePk, "foo_date", fooDateSqlName, "", "", "java.time.LocalDate.class", "fooDate", "", true, false,true, "date", 13, 0, "",false,"","", "","vendorDateFkList","",false,"df.Date")
	fooNameSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo foo_name
	fooNameSqlName.ColumnSqlName = "foo_name"
	fooNameSqlName.IrregularChar = false
	VendorDatePkDbm.ColumnFooName = df.CCI(&vendorDatePk, "foo_name", fooNameSqlName, "", "", "String.class", "fooName", "", false, false,true, "varchar", 2147483647, 0, "",false,"","", "","","",false,"string")

	VendorDatePkDbm.ColumnInfoList = new(df.List)
	VendorDatePkDbm.ColumnInfoList.Add(VendorDatePkDbm.ColumnFooDate)
	VendorDatePkDbm.ColumnInfoList.Add(VendorDatePkDbm.ColumnFooName)


	VendorDatePkDbm.ColumnInfoMap=make(map[string]int)
	VendorDatePkDbm.ColumnInfoMap["fooDate"]=0
		VendorDatePkDbm.ColumnInfoMap["fooName"]=1
	    VendorDatePkDbm.PrimaryKey = true
    VendorDatePkDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &vendorDatePk
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(VendorDatePkDbm.ColumnFooDate)

	VendorDatePkDbm.PrimaryInfo = new(df.PrimaryInfo)
	VendorDatePkDbm.PrimaryInfo.UniqueInfo = ui
	
	var vendorDatePkMeta df.DBMeta = VendorDatePkDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["VendorDatePk"] = &vendorDatePkMeta
}
