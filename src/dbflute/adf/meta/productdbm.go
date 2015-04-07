package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type ProductDbm_T struct {
	df.BaseDBMeta
	ColumnProductId *df.ColumnInfo
	ColumnProductName *df.ColumnInfo
	ColumnProductHandleCode *df.ColumnInfo
	ColumnProductCategoryCode *df.ColumnInfo
	ColumnProductStatusCode *df.ColumnInfo
	ColumnRegularPrice *df.ColumnInfo
	ColumnRegisterDatetime *df.ColumnInfo
	ColumnRegisterUser *df.ColumnInfo
	ColumnRegisterProcess *df.ColumnInfo
	ColumnUpdateDatetime *df.ColumnInfo
	ColumnUpdateUser *df.ColumnInfo
	ColumnUpdateProcess *df.ColumnInfo
	ColumnVersionNo *df.ColumnInfo
}

func (b *ProductDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *ProductDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var ProductDbm *ProductDbm_T

func Create_ProductDbm() {
	ProductDbm = new(ProductDbm_T)
	ProductDbm.TableDbName = "product"
	ProductDbm.TableDispName = "product"
	ProductDbm.TablePropertyName = "product"
	ProductDbm.TableSqlName = new(df.TableSqlName)
	ProductDbm.TableSqlName.TableSqlName = "product"
	ProductDbm.TableSqlName.CorrespondingDbName = ProductDbm.TableDbName

	var product df.DBMeta
	product = ProductDbm
	ProductDbm.DBMeta=&product
	productIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo product_id
	productIdSqlName.ColumnSqlName = "product_id"
	productIdSqlName.IrregularChar = false
	ProductDbm.ColumnProductId = df.CCI(&product, "product_id", productIdSqlName, "", "", "Integer.class", "productId", "", true, true,true, "serial", 10, 0, "nextval('product_product_id_seq'::regclass)",false,"","", "","purchaseList","",false,"int64")
	productNameSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo product_name
	productNameSqlName.ColumnSqlName = "product_name"
	productNameSqlName.IrregularChar = false
	ProductDbm.ColumnProductName = df.CCI(&product, "product_name", productNameSqlName, "", "", "String.class", "productName", "", false, false,true, "varchar", 50, 0, "",false,"","", "","","",false,"string")
	productHandleCodeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo product_handle_code
	productHandleCodeSqlName.ColumnSqlName = "product_handle_code"
	productHandleCodeSqlName.IrregularChar = false
	ProductDbm.ColumnProductHandleCode = df.CCI(&product, "product_handle_code", productHandleCodeSqlName, "", "", "String.class", "productHandleCode", "", false, false,true, "varchar", 100, 0, "",false,"","", "","","",false,"string")
	productCategoryCodeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo product_category_code
	productCategoryCodeSqlName.ColumnSqlName = "product_category_code"
	productCategoryCodeSqlName.IrregularChar = false
	ProductDbm.ColumnProductCategoryCode = df.CCI(&product, "product_category_code", productCategoryCodeSqlName, "", "", "String.class", "productCategoryCode", "", false, false,true, "bpchar", 3, 0, "",false,"","", "productCategory","","",false,"string")
	productStatusCodeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo product_status_code
	productStatusCodeSqlName.ColumnSqlName = "product_status_code"
	productStatusCodeSqlName.IrregularChar = false
	ProductDbm.ColumnProductStatusCode = df.CCI(&product, "product_status_code", productStatusCodeSqlName, "", "", "String.class", "productStatusCode", "", false, false,true, "bpchar", 3, 0, "",false,"","", "productStatus","","",false,"string")
	regularPriceSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo regular_price
	regularPriceSqlName.ColumnSqlName = "regular_price"
	regularPriceSqlName.IrregularChar = false
	ProductDbm.ColumnRegularPrice = df.CCI(&product, "regular_price", regularPriceSqlName, "", "", "Integer.class", "regularPrice", "", false, false,true, "int4", 10, 0, "",false,"","", "","","",false,"int64")
	registerDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo register_datetime
	registerDatetimeSqlName.ColumnSqlName = "register_datetime"
	registerDatetimeSqlName.IrregularChar = false
	ProductDbm.ColumnRegisterDatetime = df.CCI(&product, "register_datetime", registerDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "registerDatetime", "", false, false,true, "timestamp", 26, 3, "",false,"","", "","","",false,"df.Timestamp")
	registerUserSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo register_user
	registerUserSqlName.ColumnSqlName = "register_user"
	registerUserSqlName.IrregularChar = false
	ProductDbm.ColumnRegisterUser = df.CCI(&product, "register_user", registerUserSqlName, "", "", "String.class", "registerUser", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	registerProcessSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo register_process
	registerProcessSqlName.ColumnSqlName = "register_process"
	registerProcessSqlName.IrregularChar = false
	ProductDbm.ColumnRegisterProcess = df.CCI(&product, "register_process", registerProcessSqlName, "", "", "String.class", "registerProcess", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	updateDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo update_datetime
	updateDatetimeSqlName.ColumnSqlName = "update_datetime"
	updateDatetimeSqlName.IrregularChar = false
	ProductDbm.ColumnUpdateDatetime = df.CCI(&product, "update_datetime", updateDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "updateDatetime", "", false, false,true, "timestamp", 26, 3, "",false,"","", "","","",false,"df.Timestamp")
	updateUserSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo update_user
	updateUserSqlName.ColumnSqlName = "update_user"
	updateUserSqlName.IrregularChar = false
	ProductDbm.ColumnUpdateUser = df.CCI(&product, "update_user", updateUserSqlName, "", "", "String.class", "updateUser", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	updateProcessSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo update_process
	updateProcessSqlName.ColumnSqlName = "update_process"
	updateProcessSqlName.IrregularChar = false
	ProductDbm.ColumnUpdateProcess = df.CCI(&product, "update_process", updateProcessSqlName, "", "", "String.class", "updateProcess", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	versionNoSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo version_no
	versionNoSqlName.ColumnSqlName = "version_no"
	versionNoSqlName.IrregularChar = false
	ProductDbm.ColumnVersionNo = df.CCI(&product, "version_no", versionNoSqlName, "", "", "Long.class", "versionNo", "", false, false,true, "int8", 19, 0, "",false,"OptimisticLockType.VERSION_NO","", "","","",false,"int64")

	ProductDbm.ColumnInfoList = new(df.List)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnProductId)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnProductName)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnProductHandleCode)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnProductCategoryCode)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnProductStatusCode)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnRegularPrice)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnRegisterDatetime)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnRegisterUser)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnRegisterProcess)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnUpdateDatetime)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnUpdateUser)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnUpdateProcess)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnVersionNo)


	ProductDbm.ColumnInfoMap=make(map[string]int)
	ProductDbm.ColumnInfoMap["productId"]=0
		ProductDbm.ColumnInfoMap["productName"]=1
		ProductDbm.ColumnInfoMap["productHandleCode"]=2
		ProductDbm.ColumnInfoMap["productCategoryCode"]=3
		ProductDbm.ColumnInfoMap["productStatusCode"]=4
		ProductDbm.ColumnInfoMap["regularPrice"]=5
		ProductDbm.ColumnInfoMap["registerDatetime"]=6
		ProductDbm.ColumnInfoMap["registerUser"]=7
		ProductDbm.ColumnInfoMap["registerProcess"]=8
		ProductDbm.ColumnInfoMap["updateDatetime"]=9
		ProductDbm.ColumnInfoMap["updateUser"]=10
		ProductDbm.ColumnInfoMap["updateProcess"]=11
		ProductDbm.ColumnInfoMap["versionNo"]=12
	    ProductDbm.PrimaryKey = true
    ProductDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &product
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(ProductDbm.ColumnProductId)

	ProductDbm.PrimaryInfo = new(df.PrimaryInfo)
	ProductDbm.PrimaryInfo.UniqueInfo = ui
	
	ProductDbm.SequenceFlag = true
	ProductDbm.SequenceName = "product_product_id_seq"
	ProductDbm.SequenceIncrementSize = 1
	ProductDbm.SequenceCacheSize = 0
	ProductDbm.VersionNoFlag = true
	ProductDbm.VersionNoColumnInfo = ProductDbm.ColumnVersionNo
	
	var productMeta df.DBMeta = ProductDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["Product"] = &productMeta
}
