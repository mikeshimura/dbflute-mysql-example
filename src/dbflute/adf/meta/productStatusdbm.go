package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type ProductStatusDbm_T struct {
	df.BaseDBMeta
	ColumnProductStatusCode *df.ColumnInfo
	ColumnProductStatusName *df.ColumnInfo
	ColumnDisplayOrder *df.ColumnInfo
}

func (b *ProductStatusDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}
func (b *ProductStatusDbm_T) CreateForeignInfoMap() {
	b.ForeignInfoMap = make(map[string]*df.ForeignInfo)
}

func (b *ProductStatusDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var ProductStatusDbm *ProductStatusDbm_T

func Create_ProductStatusDbm() {
	ProductStatusDbm = new(ProductStatusDbm_T)
	ProductStatusDbm.TableDbName = "product_status"
	ProductStatusDbm.TableDispName = "product_status"
	ProductStatusDbm.TablePropertyName = "productStatus"
	ProductStatusDbm.TableSqlName = new(df.TableSqlName)
	ProductStatusDbm.TableSqlName.TableSqlName = "product_status"
	ProductStatusDbm.TableSqlName.CorrespondingDbName = ProductStatusDbm.TableDbName

	var productStatus df.DBMeta
	productStatus = ProductStatusDbm
	ProductStatusDbm.DBMeta=&productStatus
	productStatusCodeSqlName := new(df.ColumnSqlName)
	productStatusCodeSqlName.ColumnSqlName = "PRODUCT_STATUS_CODE"
	productStatusCodeSqlName.IrregularChar = false
	ProductStatusDbm.ColumnProductStatusCode = df.CCI(&productStatus, "PRODUCT_STATUS_CODE", productStatusCodeSqlName, "", "", "String.class", "productStatusCode", "", true, false,true, "CHAR", 3, 0, "",false,"","", "","productList","",false,"string")
	productStatusNameSqlName := new(df.ColumnSqlName)
	productStatusNameSqlName.ColumnSqlName = "PRODUCT_STATUS_NAME"
	productStatusNameSqlName.IrregularChar = false
	ProductStatusDbm.ColumnProductStatusName = df.CCI(&productStatus, "PRODUCT_STATUS_NAME", productStatusNameSqlName, "", "", "String.class", "productStatusName", "", false, false,true, "VARCHAR", 50, 0, "",false,"","", "","","",false,"string")
	displayOrderSqlName := new(df.ColumnSqlName)
	displayOrderSqlName.ColumnSqlName = "DISPLAY_ORDER"
	displayOrderSqlName.IrregularChar = false
	ProductStatusDbm.ColumnDisplayOrder = df.CCI(&productStatus, "DISPLAY_ORDER", displayOrderSqlName, "", "", "Integer.class", "displayOrder", "", false, false,true, "INT", 10, 0, "",false,"","", "","","",false,"int64")

	ProductStatusDbm.ColumnInfoList = new(df.List)
	ProductStatusDbm.ColumnInfoList.Add(ProductStatusDbm.ColumnProductStatusCode)
	ProductStatusDbm.ColumnInfoList.Add(ProductStatusDbm.ColumnProductStatusName)
	ProductStatusDbm.ColumnInfoList.Add(ProductStatusDbm.ColumnDisplayOrder)


	ProductStatusDbm.ColumnInfoMap=make(map[string]int)
	ProductStatusDbm.ColumnInfoMap["productStatusCode"]=0
		ProductStatusDbm.ColumnInfoMap["productStatusName"]=1
		ProductStatusDbm.ColumnInfoMap["displayOrder"]=2
	    ProductStatusDbm.PrimaryKey = true
    ProductStatusDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &productStatus
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(ProductStatusDbm.ColumnProductStatusCode)

	ProductStatusDbm.PrimaryInfo = new(df.PrimaryInfo)
	ProductStatusDbm.PrimaryInfo.UniqueInfo = ui
	
	var productStatusMeta df.DBMeta = ProductStatusDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["ProductStatus"] = &productStatusMeta
}
