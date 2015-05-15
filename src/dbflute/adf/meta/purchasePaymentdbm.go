package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type PurchasePaymentDbm_T struct {
	df.BaseDBMeta
	ColumnPurchasePaymentId *df.ColumnInfo
	ColumnPurchaseId *df.ColumnInfo
	ColumnPaymentAmount *df.ColumnInfo
	ColumnPaymentDatetime *df.ColumnInfo
	ColumnPaymentMethodCode *df.ColumnInfo
	ColumnRegisterDatetime *df.ColumnInfo
	ColumnRegisterUser *df.ColumnInfo
	ColumnUpdateDatetime *df.ColumnInfo
	ColumnUpdateUser *df.ColumnInfo
}

func (b *PurchasePaymentDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}
func (b *PurchasePaymentDbm_T) foreignPurchase() *df.ForeignInfo {
	columns := []*df.ColumnInfo{
		PurchasePaymentDbm.GetColumnInfoByPropertyName("purchaseId"),
		PurchaseDbm.GetColumnInfoByPropertyName("purchaseId"),
	}

	return b.BaseDBMeta.Cfi("FK_PURCHASE_PAYMENT_PURCHASE", "Purchase",
		columns, 0, false, false, false, false,
		"", nil, false, "purchasePaymentList")
}	
func (b *PurchasePaymentDbm_T) CreateForeignInfoMap() {
	b.ForeignInfoMap = make(map[string]*df.ForeignInfo)
	b.ForeignInfoMap["Purchase"] = b.foreignPurchase()
}

func (b *PurchasePaymentDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var PurchasePaymentDbm *PurchasePaymentDbm_T

func Create_PurchasePaymentDbm() {
	PurchasePaymentDbm = new(PurchasePaymentDbm_T)
	PurchasePaymentDbm.TableDbName = "purchase_payment"
	PurchasePaymentDbm.TableDispName = "purchase_payment"
	PurchasePaymentDbm.TablePropertyName = "purchasePayment"
	PurchasePaymentDbm.TableSqlName = new(df.TableSqlName)
	PurchasePaymentDbm.TableSqlName.TableSqlName = "purchase_payment"
	PurchasePaymentDbm.TableSqlName.CorrespondingDbName = PurchasePaymentDbm.TableDbName
	PurchasePaymentDbm.Identity=true

	var purchasePayment df.DBMeta
	purchasePayment = PurchasePaymentDbm
	PurchasePaymentDbm.DBMeta=&purchasePayment
	purchasePaymentIdSqlName := new(df.ColumnSqlName)
	purchasePaymentIdSqlName.ColumnSqlName = "PURCHASE_PAYMENT_ID"
	purchasePaymentIdSqlName.IrregularChar = false
	PurchasePaymentDbm.ColumnPurchasePaymentId = df.CCI(&purchasePayment, "PURCHASE_PAYMENT_ID", purchasePaymentIdSqlName, "", "", "Long.class", "purchasePaymentId", "", true, true,true, "BIGINT", 19, 0, "",false,"","", "","","",false,"int64")
	purchaseIdSqlName := new(df.ColumnSqlName)
	purchaseIdSqlName.ColumnSqlName = "PURCHASE_ID"
	purchaseIdSqlName.IrregularChar = false
	PurchasePaymentDbm.ColumnPurchaseId = df.CCI(&purchasePayment, "PURCHASE_ID", purchaseIdSqlName, "", "", "Long.class", "purchaseId", "", false, false,true, "BIGINT", 19, 0, "",false,"","", "purchase","","",false,"int64")
	paymentAmountSqlName := new(df.ColumnSqlName)
	paymentAmountSqlName.ColumnSqlName = "PAYMENT_AMOUNT"
	paymentAmountSqlName.IrregularChar = false
	PurchasePaymentDbm.ColumnPaymentAmount = df.CCI(&purchasePayment, "PAYMENT_AMOUNT", paymentAmountSqlName, "", "", "java.math.BigDecimal.class", "paymentAmount", "", false, false,true, "DECIMAL", 10, 2, "",false,"","", "","","",false,"df.Numeric")
	paymentDatetimeSqlName := new(df.ColumnSqlName)
	paymentDatetimeSqlName.ColumnSqlName = "PAYMENT_DATETIME"
	paymentDatetimeSqlName.IrregularChar = false
	PurchasePaymentDbm.ColumnPaymentDatetime = df.CCI(&purchasePayment, "PAYMENT_DATETIME", paymentDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "paymentDatetime", "", false, false,true, "DATETIME", 19, 0, "",false,"","", "","","",false,"df.MysqlTimestamp")
	paymentMethodCodeSqlName := new(df.ColumnSqlName)
	paymentMethodCodeSqlName.ColumnSqlName = "PAYMENT_METHOD_CODE"
	paymentMethodCodeSqlName.IrregularChar = false
	PurchasePaymentDbm.ColumnPaymentMethodCode = df.CCI(&purchasePayment, "PAYMENT_METHOD_CODE", paymentMethodCodeSqlName, "", "", "String.class", "paymentMethodCode", "", false, false,true, "CHAR", 3, 0, "",false,"","", "","","",false,"string")
	registerDatetimeSqlName := new(df.ColumnSqlName)
	registerDatetimeSqlName.ColumnSqlName = "REGISTER_DATETIME"
	registerDatetimeSqlName.IrregularChar = false
	PurchasePaymentDbm.ColumnRegisterDatetime = df.CCI(&purchasePayment, "REGISTER_DATETIME", registerDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "registerDatetime", "", false, false,true, "DATETIME", 19, 0, "",false,"","", "","","",false,"df.MysqlTimestamp")
	registerUserSqlName := new(df.ColumnSqlName)
	registerUserSqlName.ColumnSqlName = "REGISTER_USER"
	registerUserSqlName.IrregularChar = false
	PurchasePaymentDbm.ColumnRegisterUser = df.CCI(&purchasePayment, "REGISTER_USER", registerUserSqlName, "", "", "String.class", "registerUser", "", false, false,true, "VARCHAR", 200, 0, "",false,"","", "","","",false,"string")
	updateDatetimeSqlName := new(df.ColumnSqlName)
	updateDatetimeSqlName.ColumnSqlName = "UPDATE_DATETIME"
	updateDatetimeSqlName.IrregularChar = false
	PurchasePaymentDbm.ColumnUpdateDatetime = df.CCI(&purchasePayment, "UPDATE_DATETIME", updateDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "updateDatetime", "", false, false,true, "DATETIME", 19, 0, "",false,"","", "","","",false,"df.MysqlTimestamp")
	updateUserSqlName := new(df.ColumnSqlName)
	updateUserSqlName.ColumnSqlName = "UPDATE_USER"
	updateUserSqlName.IrregularChar = false
	PurchasePaymentDbm.ColumnUpdateUser = df.CCI(&purchasePayment, "UPDATE_USER", updateUserSqlName, "", "", "String.class", "updateUser", "", false, false,true, "VARCHAR", 200, 0, "",false,"","", "","","",false,"string")

	PurchasePaymentDbm.ColumnInfoList = new(df.List)
	PurchasePaymentDbm.ColumnInfoList.Add(PurchasePaymentDbm.ColumnPurchasePaymentId)
	PurchasePaymentDbm.ColumnInfoList.Add(PurchasePaymentDbm.ColumnPurchaseId)
	PurchasePaymentDbm.ColumnInfoList.Add(PurchasePaymentDbm.ColumnPaymentAmount)
	PurchasePaymentDbm.ColumnInfoList.Add(PurchasePaymentDbm.ColumnPaymentDatetime)
	PurchasePaymentDbm.ColumnInfoList.Add(PurchasePaymentDbm.ColumnPaymentMethodCode)
	PurchasePaymentDbm.ColumnInfoList.Add(PurchasePaymentDbm.ColumnRegisterDatetime)
	PurchasePaymentDbm.ColumnInfoList.Add(PurchasePaymentDbm.ColumnRegisterUser)
	PurchasePaymentDbm.ColumnInfoList.Add(PurchasePaymentDbm.ColumnUpdateDatetime)
	PurchasePaymentDbm.ColumnInfoList.Add(PurchasePaymentDbm.ColumnUpdateUser)


	PurchasePaymentDbm.ColumnInfoMap=make(map[string]int)
	PurchasePaymentDbm.ColumnInfoMap["purchasePaymentId"]=0
		PurchasePaymentDbm.ColumnInfoMap["purchaseId"]=1
		PurchasePaymentDbm.ColumnInfoMap["paymentAmount"]=2
		PurchasePaymentDbm.ColumnInfoMap["paymentDatetime"]=3
		PurchasePaymentDbm.ColumnInfoMap["paymentMethodCode"]=4
		PurchasePaymentDbm.ColumnInfoMap["registerDatetime"]=5
		PurchasePaymentDbm.ColumnInfoMap["registerUser"]=6
		PurchasePaymentDbm.ColumnInfoMap["updateDatetime"]=7
		PurchasePaymentDbm.ColumnInfoMap["updateUser"]=8
	    PurchasePaymentDbm.PrimaryKey = true
    PurchasePaymentDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &purchasePayment
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(PurchasePaymentDbm.ColumnPurchasePaymentId)

	PurchasePaymentDbm.PrimaryInfo = new(df.PrimaryInfo)
	PurchasePaymentDbm.PrimaryInfo.UniqueInfo = ui
	
	var purchasePaymentMeta df.DBMeta = PurchasePaymentDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["PurchasePayment"] = &purchasePaymentMeta
}
