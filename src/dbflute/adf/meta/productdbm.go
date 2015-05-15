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
	ColumnUpdateDatetime *df.ColumnInfo
	ColumnUpdateUser *df.ColumnInfo
	ColumnVersionNo *df.ColumnInfo
}

func (b *ProductDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}
func (b *ProductDbm_T) foreignProductCategory() *df.ForeignInfo {
	columns := []*df.ColumnInfo{
		ProductDbm.GetColumnInfoByPropertyName("productCategoryCode"),
		ProductCategoryDbm.GetColumnInfoByPropertyName("productCategoryCode"),
	}

	return b.BaseDBMeta.Cfi("FK_PRODUCT_PRODUCT_CATEGORY", "ProductCategory",
		columns, 0, false, false, false, false,
		"", nil, false, "productList")
}	
func (b *ProductDbm_T) foreignProductStatus() *df.ForeignInfo {
	columns := []*df.ColumnInfo{
		ProductDbm.GetColumnInfoByPropertyName("productStatusCode"),
		ProductStatusDbm.GetColumnInfoByPropertyName("productStatusCode"),
	}

	return b.BaseDBMeta.Cfi("FK_PRODUCT_PRODUCT_STATUS", "ProductStatus",
		columns, 1, false, false, false, false,
		"", nil, false, "productList")
}	
func (b *ProductDbm_T) CreateForeignInfoMap() {
	b.ForeignInfoMap = make(map[string]*df.ForeignInfo)
	b.ForeignInfoMap["ProductCategory"] = b.foreignProductCategory()
	b.ForeignInfoMap["ProductStatus"] = b.foreignProductStatus()
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
	ProductDbm.Identity=true

	var product df.DBMeta
	product = ProductDbm
	ProductDbm.DBMeta=&product
	productIdSqlName := new(df.ColumnSqlName)
	productIdSqlName.ColumnSqlName = "PRODUCT_ID"
	productIdSqlName.IrregularChar = false
	ProductDbm.ColumnProductId = df.CCI(&product, "PRODUCT_ID", productIdSqlName, "", "", "Integer.class", "productId", "", true, true,true, "INT", 10, 0, "",false,"","", "","purchaseList","",false,"int64")
	productNameSqlName := new(df.ColumnSqlName)
	productNameSqlName.ColumnSqlName = "PRODUCT_NAME"
	productNameSqlName.IrregularChar = false
	ProductDbm.ColumnProductName = df.CCI(&product, "PRODUCT_NAME", productNameSqlName, "", "", "String.class", "productName", "", false, false,true, "VARCHAR", 50, 0, "",false,"","", "","","",false,"string")
	productHandleCodeSqlName := new(df.ColumnSqlName)
	productHandleCodeSqlName.ColumnSqlName = "PRODUCT_HANDLE_CODE"
	productHandleCodeSqlName.IrregularChar = false
	ProductDbm.ColumnProductHandleCode = df.CCI(&product, "PRODUCT_HANDLE_CODE", productHandleCodeSqlName, "", "", "String.class", "productHandleCode", "", false, false,true, "VARCHAR", 100, 0, "",false,"","", "","","",false,"string")
	productCategoryCodeSqlName := new(df.ColumnSqlName)
	productCategoryCodeSqlName.ColumnSqlName = "PRODUCT_CATEGORY_CODE"
	productCategoryCodeSqlName.IrregularChar = false
	ProductDbm.ColumnProductCategoryCode = df.CCI(&product, "PRODUCT_CATEGORY_CODE", productCategoryCodeSqlName, "", "", "String.class", "productCategoryCode", "", false, false,true, "CHAR", 3, 0, "",false,"","", "productCategory","","",false,"string")
	productStatusCodeSqlName := new(df.ColumnSqlName)
	productStatusCodeSqlName.ColumnSqlName = "PRODUCT_STATUS_CODE"
	productStatusCodeSqlName.IrregularChar = false
	ProductDbm.ColumnProductStatusCode = df.CCI(&product, "PRODUCT_STATUS_CODE", productStatusCodeSqlName, "", "", "String.class", "productStatusCode", "", false, false,true, "CHAR", 3, 0, "",false,"","", "productStatus","","",false,"string")
	regularPriceSqlName := new(df.ColumnSqlName)
	regularPriceSqlName.ColumnSqlName = "REGULAR_PRICE"
	regularPriceSqlName.IrregularChar = false
	ProductDbm.ColumnRegularPrice = df.CCI(&product, "REGULAR_PRICE", regularPriceSqlName, "", "", "Integer.class", "regularPrice", "", false, false,true, "INT", 10, 0, "",false,"","", "","","",false,"int64")
	registerDatetimeSqlName := new(df.ColumnSqlName)
	registerDatetimeSqlName.ColumnSqlName = "REGISTER_DATETIME"
	registerDatetimeSqlName.IrregularChar = false
	ProductDbm.ColumnRegisterDatetime = df.CCI(&product, "REGISTER_DATETIME", registerDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "registerDatetime", "", false, false,true, "DATETIME", 19, 0, "",false,"","", "","","",false,"df.MysqlTimestamp")
	registerUserSqlName := new(df.ColumnSqlName)
	registerUserSqlName.ColumnSqlName = "REGISTER_USER"
	registerUserSqlName.IrregularChar = false
	ProductDbm.ColumnRegisterUser = df.CCI(&product, "REGISTER_USER", registerUserSqlName, "", "", "String.class", "registerUser", "", false, false,true, "VARCHAR", 200, 0, "",false,"","", "","","",false,"string")
	updateDatetimeSqlName := new(df.ColumnSqlName)
	updateDatetimeSqlName.ColumnSqlName = "UPDATE_DATETIME"
	updateDatetimeSqlName.IrregularChar = false
	ProductDbm.ColumnUpdateDatetime = df.CCI(&product, "UPDATE_DATETIME", updateDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "updateDatetime", "", false, false,true, "DATETIME", 19, 0, "",false,"","", "","","",false,"df.MysqlTimestamp")
	updateUserSqlName := new(df.ColumnSqlName)
	updateUserSqlName.ColumnSqlName = "UPDATE_USER"
	updateUserSqlName.IrregularChar = false
	ProductDbm.ColumnUpdateUser = df.CCI(&product, "UPDATE_USER", updateUserSqlName, "", "", "String.class", "updateUser", "", false, false,true, "VARCHAR", 200, 0, "",false,"","", "","","",false,"string")
	versionNoSqlName := new(df.ColumnSqlName)
	versionNoSqlName.ColumnSqlName = "VERSION_NO"
	versionNoSqlName.IrregularChar = false
	ProductDbm.ColumnVersionNo = df.CCI(&product, "VERSION_NO", versionNoSqlName, "", "", "Long.class", "versionNo", "", false, false,true, "BIGINT", 19, 0, "",false,"OptimisticLockType.VERSION_NO","", "","","",false,"int64")

	ProductDbm.ColumnInfoList = new(df.List)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnProductId)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnProductName)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnProductHandleCode)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnProductCategoryCode)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnProductStatusCode)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnRegularPrice)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnRegisterDatetime)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnRegisterUser)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnUpdateDatetime)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnUpdateUser)
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
		ProductDbm.ColumnInfoMap["updateDatetime"]=8
		ProductDbm.ColumnInfoMap["updateUser"]=9
		ProductDbm.ColumnInfoMap["versionNo"]=10
	    ProductDbm.PrimaryKey = true
    ProductDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &product
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(ProductDbm.ColumnProductId)

	ProductDbm.PrimaryInfo = new(df.PrimaryInfo)
	ProductDbm.PrimaryInfo.UniqueInfo = ui
	
	ProductDbm.VersionNoFlag = true
	ProductDbm.VersionNoColumnInfo = ProductDbm.ColumnVersionNo
	
	var productMeta df.DBMeta = ProductDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["Product"] = &productMeta
}
