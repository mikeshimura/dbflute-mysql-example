package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type WithdrawalReasonDbm_T struct {
	df.BaseDBMeta
	ColumnWithdrawalReasonCode *df.ColumnInfo
	ColumnWithdrawalReasonText *df.ColumnInfo
	ColumnDisplayOrder *df.ColumnInfo
}

func (b *WithdrawalReasonDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *WithdrawalReasonDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var WithdrawalReasonDbm *WithdrawalReasonDbm_T

func Create_WithdrawalReasonDbm() {
	WithdrawalReasonDbm = new(WithdrawalReasonDbm_T)
	WithdrawalReasonDbm.TableDbName = "withdrawal_reason"
	WithdrawalReasonDbm.TableDispName = "withdrawal_reason"
	WithdrawalReasonDbm.TablePropertyName = "withdrawalReason"
	WithdrawalReasonDbm.TableSqlName = new(df.TableSqlName)
	WithdrawalReasonDbm.TableSqlName.TableSqlName = "withdrawal_reason"
	WithdrawalReasonDbm.TableSqlName.CorrespondingDbName = WithdrawalReasonDbm.TableDbName

	var withdrawalReason df.DBMeta
	withdrawalReason = WithdrawalReasonDbm
	WithdrawalReasonDbm.DBMeta=&withdrawalReason
	withdrawalReasonCodeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo WITHDRAWAL_REASON_CODE
	withdrawalReasonCodeSqlName.ColumnSqlName = "WITHDRAWAL_REASON_CODE"
	withdrawalReasonCodeSqlName.IrregularChar = false
	WithdrawalReasonDbm.ColumnWithdrawalReasonCode = df.CCI(&withdrawalReason, "WITHDRAWAL_REASON_CODE", withdrawalReasonCodeSqlName, "", "", "String.class", "withdrawalReasonCode", "", true, false,true, "CHAR", 3, 0, "",false,"","", "","memberWithdrawalList","",false,"string")
	withdrawalReasonTextSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo WITHDRAWAL_REASON_TEXT
	withdrawalReasonTextSqlName.ColumnSqlName = "WITHDRAWAL_REASON_TEXT"
	withdrawalReasonTextSqlName.IrregularChar = false
	WithdrawalReasonDbm.ColumnWithdrawalReasonText = df.CCI(&withdrawalReason, "WITHDRAWAL_REASON_TEXT", withdrawalReasonTextSqlName, "", "", "String.class", "withdrawalReasonText", "", false, false,true, "TEXT", 65535, 0, "",false,"","", "","","",false,"string")
	displayOrderSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo DISPLAY_ORDER
	displayOrderSqlName.ColumnSqlName = "DISPLAY_ORDER"
	displayOrderSqlName.IrregularChar = false
	WithdrawalReasonDbm.ColumnDisplayOrder = df.CCI(&withdrawalReason, "DISPLAY_ORDER", displayOrderSqlName, "", "", "Integer.class", "displayOrder", "", false, false,true, "INT", 10, 0, "",false,"","", "","","",false,"int64")

	WithdrawalReasonDbm.ColumnInfoList = new(df.List)
	WithdrawalReasonDbm.ColumnInfoList.Add(WithdrawalReasonDbm.ColumnWithdrawalReasonCode)
	WithdrawalReasonDbm.ColumnInfoList.Add(WithdrawalReasonDbm.ColumnWithdrawalReasonText)
	WithdrawalReasonDbm.ColumnInfoList.Add(WithdrawalReasonDbm.ColumnDisplayOrder)


	WithdrawalReasonDbm.ColumnInfoMap=make(map[string]int)
	WithdrawalReasonDbm.ColumnInfoMap["withdrawalReasonCode"]=0
		WithdrawalReasonDbm.ColumnInfoMap["withdrawalReasonText"]=1
		WithdrawalReasonDbm.ColumnInfoMap["displayOrder"]=2
	    WithdrawalReasonDbm.PrimaryKey = true
    WithdrawalReasonDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &withdrawalReason
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(WithdrawalReasonDbm.ColumnWithdrawalReasonCode)

	WithdrawalReasonDbm.PrimaryInfo = new(df.PrimaryInfo)
	WithdrawalReasonDbm.PrimaryInfo.UniqueInfo = ui
	
	var withdrawalReasonMeta df.DBMeta = WithdrawalReasonDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["WithdrawalReason"] = &withdrawalReasonMeta
}
