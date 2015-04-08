package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type SummaryProductDbm_T struct {
	df.BaseDBMeta
	ColumnProductId *df.ColumnInfo
	ColumnProductName *df.ColumnInfo
	ColumnProductHandleCode *df.ColumnInfo
	ColumnProductStatusCode *df.ColumnInfo
	ColumnLatestPurchaseDatetime *df.ColumnInfo
}

func (b *SummaryProductDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *SummaryProductDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var SummaryProductDbm *SummaryProductDbm_T

func Create_SummaryProductDbm() {
	SummaryProductDbm = new(SummaryProductDbm_T)
	SummaryProductDbm.TableDbName = "summary_product"
	SummaryProductDbm.TableDispName = "summary_product"
	SummaryProductDbm.TablePropertyName = "summaryProduct"
	SummaryProductDbm.TableSqlName = new(df.TableSqlName)
	SummaryProductDbm.TableSqlName.TableSqlName = "summary_product"
	SummaryProductDbm.TableSqlName.CorrespondingDbName = SummaryProductDbm.TableDbName

	var summaryProduct df.DBMeta
	summaryProduct = SummaryProductDbm
	SummaryProductDbm.DBMeta=&summaryProduct
	productIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo PRODUCT_ID
	productIdSqlName.ColumnSqlName = "PRODUCT_ID"
	productIdSqlName.IrregularChar = false
	SummaryProductDbm.ColumnProductId = df.CCI(&summaryProduct, "PRODUCT_ID", productIdSqlName, "", "", "Integer.class", "productId", "", false, false,true, "INT", 10, 0, "0",false,"","", "","","",false,"int64")
	productNameSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo PRODUCT_NAME
	productNameSqlName.ColumnSqlName = "PRODUCT_NAME"
	productNameSqlName.IrregularChar = false
	SummaryProductDbm.ColumnProductName = df.CCI(&summaryProduct, "PRODUCT_NAME", productNameSqlName, "", "", "String.class", "productName", "", false, false,true, "VARCHAR", 50, 0, "",false,"","", "","","",false,"string")
	productHandleCodeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo PRODUCT_HANDLE_CODE
	productHandleCodeSqlName.ColumnSqlName = "PRODUCT_HANDLE_CODE"
	productHandleCodeSqlName.IrregularChar = false
	SummaryProductDbm.ColumnProductHandleCode = df.CCI(&summaryProduct, "PRODUCT_HANDLE_CODE", productHandleCodeSqlName, "", "", "String.class", "productHandleCode", "", false, false,true, "VARCHAR", 100, 0, "",false,"","", "","","",false,"string")
	productStatusCodeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo PRODUCT_STATUS_CODE
	productStatusCodeSqlName.ColumnSqlName = "PRODUCT_STATUS_CODE"
	productStatusCodeSqlName.IrregularChar = false
	SummaryProductDbm.ColumnProductStatusCode = df.CCI(&summaryProduct, "PRODUCT_STATUS_CODE", productStatusCodeSqlName, "", "", "String.class", "productStatusCode", "", false, false,true, "CHAR", 3, 0, "",false,"","", "","","",false,"string")
	latestPurchaseDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo LATEST_PURCHASE_DATETIME
	latestPurchaseDatetimeSqlName.ColumnSqlName = "LATEST_PURCHASE_DATETIME"
	latestPurchaseDatetimeSqlName.IrregularChar = false
	SummaryProductDbm.ColumnLatestPurchaseDatetime = df.CCI(&summaryProduct, "LATEST_PURCHASE_DATETIME", latestPurchaseDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "latestPurchaseDatetime", "", false, false,false, "DATETIME", 19, 0, "",false,"","", "","","",false,"df.MysqlNullTimestamp")

	SummaryProductDbm.ColumnInfoList = new(df.List)
	SummaryProductDbm.ColumnInfoList.Add(SummaryProductDbm.ColumnProductId)
	SummaryProductDbm.ColumnInfoList.Add(SummaryProductDbm.ColumnProductName)
	SummaryProductDbm.ColumnInfoList.Add(SummaryProductDbm.ColumnProductHandleCode)
	SummaryProductDbm.ColumnInfoList.Add(SummaryProductDbm.ColumnProductStatusCode)
	SummaryProductDbm.ColumnInfoList.Add(SummaryProductDbm.ColumnLatestPurchaseDatetime)


	SummaryProductDbm.ColumnInfoMap=make(map[string]int)
	SummaryProductDbm.ColumnInfoMap["productId"]=0
		SummaryProductDbm.ColumnInfoMap["productName"]=1
		SummaryProductDbm.ColumnInfoMap["productHandleCode"]=2
		SummaryProductDbm.ColumnInfoMap["productStatusCode"]=3
		SummaryProductDbm.ColumnInfoMap["latestPurchaseDatetime"]=4
	    SummaryProductDbm.PrimaryKey = false
    SummaryProductDbm.CompoundPrimaryKey = false

	var summaryProductMeta df.DBMeta = SummaryProductDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["SummaryProduct"] = &summaryProductMeta
}
