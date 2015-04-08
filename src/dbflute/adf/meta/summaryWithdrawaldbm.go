package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type SummaryWithdrawalDbm_T struct {
	df.BaseDBMeta
	ColumnMemberId *df.ColumnInfo
	ColumnMemberName *df.ColumnInfo
	ColumnWithdrawalReasonCode *df.ColumnInfo
	ColumnWithdrawalReasonText *df.ColumnInfo
	ColumnWithdrawalReasonInputText *df.ColumnInfo
	ColumnWithdrawalDatetime *df.ColumnInfo
	ColumnMemberStatusCode *df.ColumnInfo
	ColumnMemberStatusName *df.ColumnInfo
	ColumnMaxPurchasePrice *df.ColumnInfo
}

func (b *SummaryWithdrawalDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *SummaryWithdrawalDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var SummaryWithdrawalDbm *SummaryWithdrawalDbm_T

func Create_SummaryWithdrawalDbm() {
	SummaryWithdrawalDbm = new(SummaryWithdrawalDbm_T)
	SummaryWithdrawalDbm.TableDbName = "summary_withdrawal"
	SummaryWithdrawalDbm.TableDispName = "summary_withdrawal"
	SummaryWithdrawalDbm.TablePropertyName = "summaryWithdrawal"
	SummaryWithdrawalDbm.TableSqlName = new(df.TableSqlName)
	SummaryWithdrawalDbm.TableSqlName.TableSqlName = "summary_withdrawal"
	SummaryWithdrawalDbm.TableSqlName.CorrespondingDbName = SummaryWithdrawalDbm.TableDbName

	var summaryWithdrawal df.DBMeta
	summaryWithdrawal = SummaryWithdrawalDbm
	SummaryWithdrawalDbm.DBMeta=&summaryWithdrawal
	memberIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo MEMBER_ID
	memberIdSqlName.ColumnSqlName = "MEMBER_ID"
	memberIdSqlName.IrregularChar = false
	SummaryWithdrawalDbm.ColumnMemberId = df.CCI(&summaryWithdrawal, "MEMBER_ID", memberIdSqlName, "", "", "Integer.class", "memberId", "", false, false,true, "INT", 10, 0, "",false,"","", "","","",false,"int64")
	memberNameSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo MEMBER_NAME
	memberNameSqlName.ColumnSqlName = "MEMBER_NAME"
	memberNameSqlName.IrregularChar = false
	SummaryWithdrawalDbm.ColumnMemberName = df.CCI(&summaryWithdrawal, "MEMBER_NAME", memberNameSqlName, "", "", "String.class", "memberName", "", false, false,false, "VARCHAR", 180, 0, "",false,"","", "","","",false,"df.NullString")
	withdrawalReasonCodeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo WITHDRAWAL_REASON_CODE
	withdrawalReasonCodeSqlName.ColumnSqlName = "WITHDRAWAL_REASON_CODE"
	withdrawalReasonCodeSqlName.IrregularChar = false
	SummaryWithdrawalDbm.ColumnWithdrawalReasonCode = df.CCI(&summaryWithdrawal, "WITHDRAWAL_REASON_CODE", withdrawalReasonCodeSqlName, "", "", "String.class", "withdrawalReasonCode", "", false, false,false, "CHAR", 3, 0, "",false,"","", "","","",false,"df.NullString")
	withdrawalReasonTextSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo WITHDRAWAL_REASON_TEXT
	withdrawalReasonTextSqlName.ColumnSqlName = "WITHDRAWAL_REASON_TEXT"
	withdrawalReasonTextSqlName.IrregularChar = false
	SummaryWithdrawalDbm.ColumnWithdrawalReasonText = df.CCI(&summaryWithdrawal, "WITHDRAWAL_REASON_TEXT", withdrawalReasonTextSqlName, "", "", "String.class", "withdrawalReasonText", "", false, false,false, "TEXT", 65535, 0, "",false,"","", "","","",false,"df.NullString")
	withdrawalReasonInputTextSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo WITHDRAWAL_REASON_INPUT_TEXT
	withdrawalReasonInputTextSqlName.ColumnSqlName = "WITHDRAWAL_REASON_INPUT_TEXT"
	withdrawalReasonInputTextSqlName.IrregularChar = false
	SummaryWithdrawalDbm.ColumnWithdrawalReasonInputText = df.CCI(&summaryWithdrawal, "WITHDRAWAL_REASON_INPUT_TEXT", withdrawalReasonInputTextSqlName, "", "", "String.class", "withdrawalReasonInputText", "", false, false,false, "TEXT", 65535, 0, "",false,"","", "","","",false,"df.NullString")
	withdrawalDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo WITHDRAWAL_DATETIME
	withdrawalDatetimeSqlName.ColumnSqlName = "WITHDRAWAL_DATETIME"
	withdrawalDatetimeSqlName.IrregularChar = false
	SummaryWithdrawalDbm.ColumnWithdrawalDatetime = df.CCI(&summaryWithdrawal, "WITHDRAWAL_DATETIME", withdrawalDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "withdrawalDatetime", "", false, false,true, "DATETIME", 19, 0, "",false,"","", "","","",false,"df.MysqlTimestamp")
	memberStatusCodeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo MEMBER_STATUS_CODE
	memberStatusCodeSqlName.ColumnSqlName = "MEMBER_STATUS_CODE"
	memberStatusCodeSqlName.IrregularChar = false
	SummaryWithdrawalDbm.ColumnMemberStatusCode = df.CCI(&summaryWithdrawal, "MEMBER_STATUS_CODE", memberStatusCodeSqlName, "", "", "String.class", "memberStatusCode", "", false, false,false, "CHAR", 3, 0, "",false,"","", "","","",false,"df.NullString")
	memberStatusNameSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo MEMBER_STATUS_NAME
	memberStatusNameSqlName.ColumnSqlName = "MEMBER_STATUS_NAME"
	memberStatusNameSqlName.IrregularChar = false
	SummaryWithdrawalDbm.ColumnMemberStatusName = df.CCI(&summaryWithdrawal, "MEMBER_STATUS_NAME", memberStatusNameSqlName, "", "", "String.class", "memberStatusName", "", false, false,false, "VARCHAR", 50, 0, "",false,"","", "","","",false,"df.NullString")
	maxPurchasePriceSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo MAX_PURCHASE_PRICE
	maxPurchasePriceSqlName.ColumnSqlName = "MAX_PURCHASE_PRICE"
	maxPurchasePriceSqlName.IrregularChar = false
	SummaryWithdrawalDbm.ColumnMaxPurchasePrice = df.CCI(&summaryWithdrawal, "MAX_PURCHASE_PRICE", maxPurchasePriceSqlName, "", "", "Long.class", "maxPurchasePrice", "", false, false,false, "BIGINT", 19, 0, "",false,"","", "","","",false,"sql.NullInt64")

	SummaryWithdrawalDbm.ColumnInfoList = new(df.List)
	SummaryWithdrawalDbm.ColumnInfoList.Add(SummaryWithdrawalDbm.ColumnMemberId)
	SummaryWithdrawalDbm.ColumnInfoList.Add(SummaryWithdrawalDbm.ColumnMemberName)
	SummaryWithdrawalDbm.ColumnInfoList.Add(SummaryWithdrawalDbm.ColumnWithdrawalReasonCode)
	SummaryWithdrawalDbm.ColumnInfoList.Add(SummaryWithdrawalDbm.ColumnWithdrawalReasonText)
	SummaryWithdrawalDbm.ColumnInfoList.Add(SummaryWithdrawalDbm.ColumnWithdrawalReasonInputText)
	SummaryWithdrawalDbm.ColumnInfoList.Add(SummaryWithdrawalDbm.ColumnWithdrawalDatetime)
	SummaryWithdrawalDbm.ColumnInfoList.Add(SummaryWithdrawalDbm.ColumnMemberStatusCode)
	SummaryWithdrawalDbm.ColumnInfoList.Add(SummaryWithdrawalDbm.ColumnMemberStatusName)
	SummaryWithdrawalDbm.ColumnInfoList.Add(SummaryWithdrawalDbm.ColumnMaxPurchasePrice)


	SummaryWithdrawalDbm.ColumnInfoMap=make(map[string]int)
	SummaryWithdrawalDbm.ColumnInfoMap["memberId"]=0
		SummaryWithdrawalDbm.ColumnInfoMap["memberName"]=1
		SummaryWithdrawalDbm.ColumnInfoMap["withdrawalReasonCode"]=2
		SummaryWithdrawalDbm.ColumnInfoMap["withdrawalReasonText"]=3
		SummaryWithdrawalDbm.ColumnInfoMap["withdrawalReasonInputText"]=4
		SummaryWithdrawalDbm.ColumnInfoMap["withdrawalDatetime"]=5
		SummaryWithdrawalDbm.ColumnInfoMap["memberStatusCode"]=6
		SummaryWithdrawalDbm.ColumnInfoMap["memberStatusName"]=7
		SummaryWithdrawalDbm.ColumnInfoMap["maxPurchasePrice"]=8
	    SummaryWithdrawalDbm.PrimaryKey = false
    SummaryWithdrawalDbm.CompoundPrimaryKey = false

	var summaryWithdrawalMeta df.DBMeta = SummaryWithdrawalDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["SummaryWithdrawal"] = &summaryWithdrawalMeta
}
