package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type PurchaseDbm_T struct {
	df.BaseDBMeta
	ColumnPurchaseId *df.ColumnInfo
	ColumnMemberId *df.ColumnInfo
	ColumnProductId *df.ColumnInfo
	ColumnPurchaseDatetime *df.ColumnInfo
	ColumnPurchaseCount *df.ColumnInfo
	ColumnPurchasePrice *df.ColumnInfo
	ColumnPaymentCompleteFlg *df.ColumnInfo
	ColumnRegisterDatetime *df.ColumnInfo
	ColumnRegisterUser *df.ColumnInfo
	ColumnUpdateDatetime *df.ColumnInfo
	ColumnUpdateUser *df.ColumnInfo
	ColumnVersionNo *df.ColumnInfo
}

func (b *PurchaseDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}
func (b *PurchaseDbm_T) foreignMember() *df.ForeignInfo {
	columns := []*df.ColumnInfo{
		PurchaseDbm.GetColumnInfoByPropertyName("memberId"),
		MemberDbm.GetColumnInfoByPropertyName("memberId"),
	}

	return b.BaseDBMeta.Cfi("FK_PURCHASE_MEMBER", "Member",
		columns, 0, false, false, false, false,
		"", nil, false, "purchaseList")
}	
func (b *PurchaseDbm_T) foreignProduct() *df.ForeignInfo {
	columns := []*df.ColumnInfo{
		PurchaseDbm.GetColumnInfoByPropertyName("productId"),
		ProductDbm.GetColumnInfoByPropertyName("productId"),
	}

	return b.BaseDBMeta.Cfi("FK_PURCHASE_PRODUCT", "Product",
		columns, 1, false, false, false, false,
		"", nil, false, "purchaseList")
}	
func (b *PurchaseDbm_T) CreateForeignInfoMap() {
	b.ForeignInfoMap = make(map[string]*df.ForeignInfo)
	b.ForeignInfoMap["Member"] = b.foreignMember()
	b.ForeignInfoMap["Product"] = b.foreignProduct()
}

func (b *PurchaseDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var PurchaseDbm *PurchaseDbm_T

func Create_PurchaseDbm() {
	PurchaseDbm = new(PurchaseDbm_T)
	PurchaseDbm.TableDbName = "purchase"
	PurchaseDbm.TableDispName = "purchase"
	PurchaseDbm.TablePropertyName = "purchase"
	PurchaseDbm.TableSqlName = new(df.TableSqlName)
	PurchaseDbm.TableSqlName.TableSqlName = "purchase"
	PurchaseDbm.TableSqlName.CorrespondingDbName = PurchaseDbm.TableDbName
	PurchaseDbm.Identity=true

	var purchase df.DBMeta
	purchase = PurchaseDbm
	PurchaseDbm.DBMeta=&purchase
	purchaseIdSqlName := new(df.ColumnSqlName)
	purchaseIdSqlName.ColumnSqlName = "PURCHASE_ID"
	purchaseIdSqlName.IrregularChar = false
	PurchaseDbm.ColumnPurchaseId = df.CCI(&purchase, "PURCHASE_ID", purchaseIdSqlName, "", "", "Long.class", "purchaseId", "", true, true,true, "BIGINT", 19, 0, "",false,"","", "","purchasePaymentList","",false,"int64")
	memberIdSqlName := new(df.ColumnSqlName)
	memberIdSqlName.ColumnSqlName = "MEMBER_ID"
	memberIdSqlName.IrregularChar = false
	PurchaseDbm.ColumnMemberId = df.CCI(&purchase, "MEMBER_ID", memberIdSqlName, "", "", "Integer.class", "memberId", "", false, false,true, "INT", 10, 0, "",false,"","", "member","","",false,"int64")
	productIdSqlName := new(df.ColumnSqlName)
	productIdSqlName.ColumnSqlName = "PRODUCT_ID"
	productIdSqlName.IrregularChar = false
	PurchaseDbm.ColumnProductId = df.CCI(&purchase, "PRODUCT_ID", productIdSqlName, "", "", "Integer.class", "productId", "", false, false,true, "INT", 10, 0, "",false,"","", "product","","",false,"int64")
	purchaseDatetimeSqlName := new(df.ColumnSqlName)
	purchaseDatetimeSqlName.ColumnSqlName = "PURCHASE_DATETIME"
	purchaseDatetimeSqlName.IrregularChar = false
	PurchaseDbm.ColumnPurchaseDatetime = df.CCI(&purchase, "PURCHASE_DATETIME", purchaseDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "purchaseDatetime", "", false, false,true, "DATETIME", 19, 0, "",false,"","", "","","",false,"df.MysqlTimestamp")
	purchaseCountSqlName := new(df.ColumnSqlName)
	purchaseCountSqlName.ColumnSqlName = "PURCHASE_COUNT"
	purchaseCountSqlName.IrregularChar = false
	PurchaseDbm.ColumnPurchaseCount = df.CCI(&purchase, "PURCHASE_COUNT", purchaseCountSqlName, "", "", "Integer.class", "purchaseCount", "", false, false,true, "INT", 10, 0, "",false,"","", "","","",false,"int64")
	purchasePriceSqlName := new(df.ColumnSqlName)
	purchasePriceSqlName.ColumnSqlName = "PURCHASE_PRICE"
	purchasePriceSqlName.IrregularChar = false
	PurchaseDbm.ColumnPurchasePrice = df.CCI(&purchase, "PURCHASE_PRICE", purchasePriceSqlName, "", "", "Integer.class", "purchasePrice", "", false, false,true, "INT", 10, 0, "",false,"","", "","","",false,"int64")
	paymentCompleteFlgSqlName := new(df.ColumnSqlName)
	paymentCompleteFlgSqlName.ColumnSqlName = "PAYMENT_COMPLETE_FLG"
	paymentCompleteFlgSqlName.IrregularChar = false
	PurchaseDbm.ColumnPaymentCompleteFlg = df.CCI(&purchase, "PAYMENT_COMPLETE_FLG", paymentCompleteFlgSqlName, "", "", "Integer.class", "paymentCompleteFlg", "", false, false,true, "INT", 10, 0, "",false,"","", "","","",false,"int64")
	registerDatetimeSqlName := new(df.ColumnSqlName)
	registerDatetimeSqlName.ColumnSqlName = "REGISTER_DATETIME"
	registerDatetimeSqlName.IrregularChar = false
	PurchaseDbm.ColumnRegisterDatetime = df.CCI(&purchase, "REGISTER_DATETIME", registerDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "registerDatetime", "", false, false,true, "DATETIME", 19, 0, "",false,"","", "","","",false,"df.MysqlTimestamp")
	registerUserSqlName := new(df.ColumnSqlName)
	registerUserSqlName.ColumnSqlName = "REGISTER_USER"
	registerUserSqlName.IrregularChar = false
	PurchaseDbm.ColumnRegisterUser = df.CCI(&purchase, "REGISTER_USER", registerUserSqlName, "", "", "String.class", "registerUser", "", false, false,true, "VARCHAR", 200, 0, "",false,"","", "","","",false,"string")
	updateDatetimeSqlName := new(df.ColumnSqlName)
	updateDatetimeSqlName.ColumnSqlName = "UPDATE_DATETIME"
	updateDatetimeSqlName.IrregularChar = false
	PurchaseDbm.ColumnUpdateDatetime = df.CCI(&purchase, "UPDATE_DATETIME", updateDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "updateDatetime", "", false, false,true, "DATETIME", 19, 0, "",false,"","", "","","",false,"df.MysqlTimestamp")
	updateUserSqlName := new(df.ColumnSqlName)
	updateUserSqlName.ColumnSqlName = "UPDATE_USER"
	updateUserSqlName.IrregularChar = false
	PurchaseDbm.ColumnUpdateUser = df.CCI(&purchase, "UPDATE_USER", updateUserSqlName, "", "", "String.class", "updateUser", "", false, false,true, "VARCHAR", 200, 0, "",false,"","", "","","",false,"string")
	versionNoSqlName := new(df.ColumnSqlName)
	versionNoSqlName.ColumnSqlName = "VERSION_NO"
	versionNoSqlName.IrregularChar = false
	PurchaseDbm.ColumnVersionNo = df.CCI(&purchase, "VERSION_NO", versionNoSqlName, "", "", "Long.class", "versionNo", "", false, false,true, "BIGINT", 19, 0, "",false,"OptimisticLockType.VERSION_NO","", "","","",false,"int64")

	PurchaseDbm.ColumnInfoList = new(df.List)
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnPurchaseId)
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnMemberId)
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnProductId)
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnPurchaseDatetime)
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnPurchaseCount)
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnPurchasePrice)
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnPaymentCompleteFlg)
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnRegisterDatetime)
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnRegisterUser)
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnUpdateDatetime)
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnUpdateUser)
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnVersionNo)


	PurchaseDbm.ColumnInfoMap=make(map[string]int)
	PurchaseDbm.ColumnInfoMap["purchaseId"]=0
		PurchaseDbm.ColumnInfoMap["memberId"]=1
		PurchaseDbm.ColumnInfoMap["productId"]=2
		PurchaseDbm.ColumnInfoMap["purchaseDatetime"]=3
		PurchaseDbm.ColumnInfoMap["purchaseCount"]=4
		PurchaseDbm.ColumnInfoMap["purchasePrice"]=5
		PurchaseDbm.ColumnInfoMap["paymentCompleteFlg"]=6
		PurchaseDbm.ColumnInfoMap["registerDatetime"]=7
		PurchaseDbm.ColumnInfoMap["registerUser"]=8
		PurchaseDbm.ColumnInfoMap["updateDatetime"]=9
		PurchaseDbm.ColumnInfoMap["updateUser"]=10
		PurchaseDbm.ColumnInfoMap["versionNo"]=11
	    PurchaseDbm.PrimaryKey = true
    PurchaseDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &purchase
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(PurchaseDbm.ColumnPurchaseId)

	PurchaseDbm.PrimaryInfo = new(df.PrimaryInfo)
	PurchaseDbm.PrimaryInfo.UniqueInfo = ui
	
	PurchaseDbm.VersionNoFlag = true
	PurchaseDbm.VersionNoColumnInfo = PurchaseDbm.ColumnVersionNo
	
	var purchaseMeta df.DBMeta = PurchaseDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["Purchase"] = &purchaseMeta
}
