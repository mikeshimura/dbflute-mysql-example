package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type SummaryProductDbm_T struct {
	df.BaseDBMeta
	ColumnProductId *df.ColumnInfo
	ColumnProductName *df.ColumnInfo
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
	//colsqlname dayoo product_id
	productIdSqlName.ColumnSqlName = "product_id"
	productIdSqlName.IrregularChar = false
	SummaryProductDbm.ColumnProductId = df.CCI(&summaryProduct, "product_id", productIdSqlName, "", "", "Integer.class", "productId", "", false, false,false, "int4", 10, 0, "",false,"","", "","","",false,"sql.NullInt64")
	productNameSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo product_name
	productNameSqlName.ColumnSqlName = "product_name"
	productNameSqlName.IrregularChar = false
	SummaryProductDbm.ColumnProductName = df.CCI(&summaryProduct, "product_name", productNameSqlName, "", "", "String.class", "productName", "", false, false,false, "varchar", 50, 0, "",false,"","", "","","",false,"sql.NullString")
	productStatusCodeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo product_status_code
	productStatusCodeSqlName.ColumnSqlName = "product_status_code"
	productStatusCodeSqlName.IrregularChar = false
	SummaryProductDbm.ColumnProductStatusCode = df.CCI(&summaryProduct, "product_status_code", productStatusCodeSqlName, "", "", "String.class", "productStatusCode", "", false, false,false, "bpchar", 3, 0, "",false,"","", "","","",false,"sql.NullString")
	latestPurchaseDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo latest_purchase_datetime
	latestPurchaseDatetimeSqlName.ColumnSqlName = "latest_purchase_datetime"
	latestPurchaseDatetimeSqlName.IrregularChar = false
	SummaryProductDbm.ColumnLatestPurchaseDatetime = df.CCI(&summaryProduct, "latest_purchase_datetime", latestPurchaseDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "latestPurchaseDatetime", "", false, false,false, "timestamp", 29, 6, "",false,"","", "","","",false,"df.NullTimestamp")

	SummaryProductDbm.ColumnInfoList = new(df.List)
	SummaryProductDbm.ColumnInfoList.Add(SummaryProductDbm.ColumnProductId)
	SummaryProductDbm.ColumnInfoList.Add(SummaryProductDbm.ColumnProductName)
	SummaryProductDbm.ColumnInfoList.Add(SummaryProductDbm.ColumnProductStatusCode)
	SummaryProductDbm.ColumnInfoList.Add(SummaryProductDbm.ColumnLatestPurchaseDatetime)


	SummaryProductDbm.ColumnInfoMap=make(map[string]int)
	SummaryProductDbm.ColumnInfoMap["productId"]=0
		SummaryProductDbm.ColumnInfoMap["productName"]=1
		SummaryProductDbm.ColumnInfoMap["productStatusCode"]=2
		SummaryProductDbm.ColumnInfoMap["latestPurchaseDatetime"]=3
	    SummaryProductDbm.PrimaryKey = false
    SummaryProductDbm.CompoundPrimaryKey = false

	var summaryProductMeta df.DBMeta = SummaryProductDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["SummaryProduct"] = &summaryProductMeta
}
