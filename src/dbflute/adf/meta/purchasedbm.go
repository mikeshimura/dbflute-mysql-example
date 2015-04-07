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
	ColumnRegisterProcess *df.ColumnInfo
	ColumnUpdateDatetime *df.ColumnInfo
	ColumnUpdateUser *df.ColumnInfo
	ColumnUpdateProcess *df.ColumnInfo
	ColumnVersionNo *df.ColumnInfo
}

func (b *PurchaseDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
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

	var purchase df.DBMeta
	purchase = PurchaseDbm
	PurchaseDbm.DBMeta=&purchase
	purchaseIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo purchase_id
	purchaseIdSqlName.ColumnSqlName = "purchase_id"
	purchaseIdSqlName.IrregularChar = false
	PurchaseDbm.ColumnPurchaseId = df.CCI(&purchase, "purchase_id", purchaseIdSqlName, "", "", "Long.class", "purchaseId", "", true, true,true, "bigserial", 19, 0, "nextval('purchase_purchase_id_seq'::regclass)",false,"","", "","purchasePaymentList","",false,"int64")
	memberIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo member_id
	memberIdSqlName.ColumnSqlName = "member_id"
	memberIdSqlName.IrregularChar = false
	PurchaseDbm.ColumnMemberId = df.CCI(&purchase, "member_id", memberIdSqlName, "", "", "Integer.class", "memberId", "", false, false,true, "int4", 10, 0, "",false,"","", "member","","",false,"int64")
	productIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo product_id
	productIdSqlName.ColumnSqlName = "product_id"
	productIdSqlName.IrregularChar = false
	PurchaseDbm.ColumnProductId = df.CCI(&purchase, "product_id", productIdSqlName, "", "", "Integer.class", "productId", "", false, false,true, "int4", 10, 0, "",false,"","", "product","","",false,"int64")
	purchaseDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo purchase_datetime
	purchaseDatetimeSqlName.ColumnSqlName = "purchase_datetime"
	purchaseDatetimeSqlName.IrregularChar = false
	PurchaseDbm.ColumnPurchaseDatetime = df.CCI(&purchase, "purchase_datetime", purchaseDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "purchaseDatetime", "", false, false,true, "timestamp", 26, 3, "",false,"","", "","","",false,"df.Timestamp")
	purchaseCountSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo purchase_count
	purchaseCountSqlName.ColumnSqlName = "purchase_count"
	purchaseCountSqlName.IrregularChar = false
	PurchaseDbm.ColumnPurchaseCount = df.CCI(&purchase, "purchase_count", purchaseCountSqlName, "", "", "Integer.class", "purchaseCount", "", false, false,true, "int4", 10, 0, "",false,"","", "","","",false,"int64")
	purchasePriceSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo purchase_price
	purchasePriceSqlName.ColumnSqlName = "purchase_price"
	purchasePriceSqlName.IrregularChar = false
	PurchaseDbm.ColumnPurchasePrice = df.CCI(&purchase, "purchase_price", purchasePriceSqlName, "", "", "Integer.class", "purchasePrice", "", false, false,true, "int4", 10, 0, "",false,"","", "","","",false,"int64")
	paymentCompleteFlgSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo payment_complete_flg
	paymentCompleteFlgSqlName.ColumnSqlName = "payment_complete_flg"
	paymentCompleteFlgSqlName.IrregularChar = false
	PurchaseDbm.ColumnPaymentCompleteFlg = df.CCI(&purchase, "payment_complete_flg", paymentCompleteFlgSqlName, "", "", "Integer.class", "paymentCompleteFlg", "", false, false,true, "int4", 10, 0, "",false,"","", "","","",false,"int64")
	registerDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo register_datetime
	registerDatetimeSqlName.ColumnSqlName = "register_datetime"
	registerDatetimeSqlName.IrregularChar = false
	PurchaseDbm.ColumnRegisterDatetime = df.CCI(&purchase, "register_datetime", registerDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "registerDatetime", "", false, false,true, "timestamp", 26, 3, "",false,"","", "","","",false,"df.Timestamp")
	registerUserSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo register_user
	registerUserSqlName.ColumnSqlName = "register_user"
	registerUserSqlName.IrregularChar = false
	PurchaseDbm.ColumnRegisterUser = df.CCI(&purchase, "register_user", registerUserSqlName, "", "", "String.class", "registerUser", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	registerProcessSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo register_process
	registerProcessSqlName.ColumnSqlName = "register_process"
	registerProcessSqlName.IrregularChar = false
	PurchaseDbm.ColumnRegisterProcess = df.CCI(&purchase, "register_process", registerProcessSqlName, "", "", "String.class", "registerProcess", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	updateDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo update_datetime
	updateDatetimeSqlName.ColumnSqlName = "update_datetime"
	updateDatetimeSqlName.IrregularChar = false
	PurchaseDbm.ColumnUpdateDatetime = df.CCI(&purchase, "update_datetime", updateDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "updateDatetime", "", false, false,true, "timestamp", 26, 3, "",false,"","", "","","",false,"df.Timestamp")
	updateUserSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo update_user
	updateUserSqlName.ColumnSqlName = "update_user"
	updateUserSqlName.IrregularChar = false
	PurchaseDbm.ColumnUpdateUser = df.CCI(&purchase, "update_user", updateUserSqlName, "", "", "String.class", "updateUser", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	updateProcessSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo update_process
	updateProcessSqlName.ColumnSqlName = "update_process"
	updateProcessSqlName.IrregularChar = false
	PurchaseDbm.ColumnUpdateProcess = df.CCI(&purchase, "update_process", updateProcessSqlName, "", "", "String.class", "updateProcess", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	versionNoSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo version_no
	versionNoSqlName.ColumnSqlName = "version_no"
	versionNoSqlName.IrregularChar = false
	PurchaseDbm.ColumnVersionNo = df.CCI(&purchase, "version_no", versionNoSqlName, "", "", "Long.class", "versionNo", "", false, false,true, "int8", 19, 0, "",false,"OptimisticLockType.VERSION_NO","", "","","",false,"int64")

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
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnRegisterProcess)
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnUpdateDatetime)
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnUpdateUser)
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnUpdateProcess)
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
		PurchaseDbm.ColumnInfoMap["registerProcess"]=9
		PurchaseDbm.ColumnInfoMap["updateDatetime"]=10
		PurchaseDbm.ColumnInfoMap["updateUser"]=11
		PurchaseDbm.ColumnInfoMap["updateProcess"]=12
		PurchaseDbm.ColumnInfoMap["versionNo"]=13
	    PurchaseDbm.PrimaryKey = true
    PurchaseDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &purchase
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(PurchaseDbm.ColumnPurchaseId)

	PurchaseDbm.PrimaryInfo = new(df.PrimaryInfo)
	PurchaseDbm.PrimaryInfo.UniqueInfo = ui
	
	PurchaseDbm.SequenceFlag = true
	PurchaseDbm.SequenceName = "purchase_purchase_id_seq"
	PurchaseDbm.SequenceIncrementSize = 1
	PurchaseDbm.SequenceCacheSize = 0
	PurchaseDbm.VersionNoFlag = true
	PurchaseDbm.VersionNoColumnInfo = PurchaseDbm.ColumnVersionNo
	
	var purchaseMeta df.DBMeta = PurchaseDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["Purchase"] = &purchaseMeta
}
