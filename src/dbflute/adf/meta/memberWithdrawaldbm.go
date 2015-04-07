package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type MemberWithdrawalDbm_T struct {
	df.BaseDBMeta
	ColumnMemberId *df.ColumnInfo
	ColumnWithdrawalReasonCode *df.ColumnInfo
	ColumnWithdrawalReasonInputText *df.ColumnInfo
	ColumnWithdrawalDatetime *df.ColumnInfo
	ColumnRegisterDatetime *df.ColumnInfo
	ColumnRegisterProcess *df.ColumnInfo
	ColumnRegisterUser *df.ColumnInfo
	ColumnUpdateDatetime *df.ColumnInfo
	ColumnUpdateProcess *df.ColumnInfo
	ColumnUpdateUser *df.ColumnInfo
}

func (b *MemberWithdrawalDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *MemberWithdrawalDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var MemberWithdrawalDbm *MemberWithdrawalDbm_T

func Create_MemberWithdrawalDbm() {
	MemberWithdrawalDbm = new(MemberWithdrawalDbm_T)
	MemberWithdrawalDbm.TableDbName = "member_withdrawal"
	MemberWithdrawalDbm.TableDispName = "member_withdrawal"
	MemberWithdrawalDbm.TablePropertyName = "memberWithdrawal"
	MemberWithdrawalDbm.TableSqlName = new(df.TableSqlName)
	MemberWithdrawalDbm.TableSqlName.TableSqlName = "member_withdrawal"
	MemberWithdrawalDbm.TableSqlName.CorrespondingDbName = MemberWithdrawalDbm.TableDbName

	var memberWithdrawal df.DBMeta
	memberWithdrawal = MemberWithdrawalDbm
	MemberWithdrawalDbm.DBMeta=&memberWithdrawal
	memberIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo member_id
	memberIdSqlName.ColumnSqlName = "member_id"
	memberIdSqlName.IrregularChar = false
	MemberWithdrawalDbm.ColumnMemberId = df.CCI(&memberWithdrawal, "member_id", memberIdSqlName, "", "", "Integer.class", "memberId", "", true, false,true, "int4", 10, 0, "",false,"","", "member","","",false,"int64")
	withdrawalReasonCodeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo withdrawal_reason_code
	withdrawalReasonCodeSqlName.ColumnSqlName = "withdrawal_reason_code"
	withdrawalReasonCodeSqlName.IrregularChar = false
	MemberWithdrawalDbm.ColumnWithdrawalReasonCode = df.CCI(&memberWithdrawal, "withdrawal_reason_code", withdrawalReasonCodeSqlName, "", "", "String.class", "withdrawalReasonCode", "", false, false,false, "bpchar", 3, 0, "",false,"","", "withdrawalReason","","",false,"sql.NullString")
	withdrawalReasonInputTextSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo withdrawal_reason_input_text
	withdrawalReasonInputTextSqlName.ColumnSqlName = "withdrawal_reason_input_text"
	withdrawalReasonInputTextSqlName.IrregularChar = false
	MemberWithdrawalDbm.ColumnWithdrawalReasonInputText = df.CCI(&memberWithdrawal, "withdrawal_reason_input_text", withdrawalReasonInputTextSqlName, "", "", "String.class", "withdrawalReasonInputText", "", false, false,false, "text", 2147483647, 0, "",false,"","", "","","",false,"sql.NullString")
	withdrawalDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo withdrawal_datetime
	withdrawalDatetimeSqlName.ColumnSqlName = "withdrawal_datetime"
	withdrawalDatetimeSqlName.IrregularChar = false
	MemberWithdrawalDbm.ColumnWithdrawalDatetime = df.CCI(&memberWithdrawal, "withdrawal_datetime", withdrawalDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "withdrawalDatetime", "", false, false,true, "timestamp", 26, 3, "",false,"","", "","","",false,"df.Timestamp")
	registerDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo register_datetime
	registerDatetimeSqlName.ColumnSqlName = "register_datetime"
	registerDatetimeSqlName.IrregularChar = false
	MemberWithdrawalDbm.ColumnRegisterDatetime = df.CCI(&memberWithdrawal, "register_datetime", registerDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "registerDatetime", "", false, false,true, "timestamp", 26, 3, "",false,"","", "","","",false,"df.Timestamp")
	registerProcessSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo register_process
	registerProcessSqlName.ColumnSqlName = "register_process"
	registerProcessSqlName.IrregularChar = false
	MemberWithdrawalDbm.ColumnRegisterProcess = df.CCI(&memberWithdrawal, "register_process", registerProcessSqlName, "", "", "String.class", "registerProcess", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	registerUserSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo register_user
	registerUserSqlName.ColumnSqlName = "register_user"
	registerUserSqlName.IrregularChar = false
	MemberWithdrawalDbm.ColumnRegisterUser = df.CCI(&memberWithdrawal, "register_user", registerUserSqlName, "", "", "String.class", "registerUser", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	updateDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo update_datetime
	updateDatetimeSqlName.ColumnSqlName = "update_datetime"
	updateDatetimeSqlName.IrregularChar = false
	MemberWithdrawalDbm.ColumnUpdateDatetime = df.CCI(&memberWithdrawal, "update_datetime", updateDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "updateDatetime", "", false, false,true, "timestamp", 26, 3, "",false,"","", "","","",false,"df.Timestamp")
	updateProcessSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo update_process
	updateProcessSqlName.ColumnSqlName = "update_process"
	updateProcessSqlName.IrregularChar = false
	MemberWithdrawalDbm.ColumnUpdateProcess = df.CCI(&memberWithdrawal, "update_process", updateProcessSqlName, "", "", "String.class", "updateProcess", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	updateUserSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo update_user
	updateUserSqlName.ColumnSqlName = "update_user"
	updateUserSqlName.IrregularChar = false
	MemberWithdrawalDbm.ColumnUpdateUser = df.CCI(&memberWithdrawal, "update_user", updateUserSqlName, "", "", "String.class", "updateUser", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")

	MemberWithdrawalDbm.ColumnInfoList = new(df.List)
	MemberWithdrawalDbm.ColumnInfoList.Add(MemberWithdrawalDbm.ColumnMemberId)
	MemberWithdrawalDbm.ColumnInfoList.Add(MemberWithdrawalDbm.ColumnWithdrawalReasonCode)
	MemberWithdrawalDbm.ColumnInfoList.Add(MemberWithdrawalDbm.ColumnWithdrawalReasonInputText)
	MemberWithdrawalDbm.ColumnInfoList.Add(MemberWithdrawalDbm.ColumnWithdrawalDatetime)
	MemberWithdrawalDbm.ColumnInfoList.Add(MemberWithdrawalDbm.ColumnRegisterDatetime)
	MemberWithdrawalDbm.ColumnInfoList.Add(MemberWithdrawalDbm.ColumnRegisterProcess)
	MemberWithdrawalDbm.ColumnInfoList.Add(MemberWithdrawalDbm.ColumnRegisterUser)
	MemberWithdrawalDbm.ColumnInfoList.Add(MemberWithdrawalDbm.ColumnUpdateDatetime)
	MemberWithdrawalDbm.ColumnInfoList.Add(MemberWithdrawalDbm.ColumnUpdateProcess)
	MemberWithdrawalDbm.ColumnInfoList.Add(MemberWithdrawalDbm.ColumnUpdateUser)


	MemberWithdrawalDbm.ColumnInfoMap=make(map[string]int)
	MemberWithdrawalDbm.ColumnInfoMap["memberId"]=0
		MemberWithdrawalDbm.ColumnInfoMap["withdrawalReasonCode"]=1
		MemberWithdrawalDbm.ColumnInfoMap["withdrawalReasonInputText"]=2
		MemberWithdrawalDbm.ColumnInfoMap["withdrawalDatetime"]=3
		MemberWithdrawalDbm.ColumnInfoMap["registerDatetime"]=4
		MemberWithdrawalDbm.ColumnInfoMap["registerProcess"]=5
		MemberWithdrawalDbm.ColumnInfoMap["registerUser"]=6
		MemberWithdrawalDbm.ColumnInfoMap["updateDatetime"]=7
		MemberWithdrawalDbm.ColumnInfoMap["updateProcess"]=8
		MemberWithdrawalDbm.ColumnInfoMap["updateUser"]=9
	    MemberWithdrawalDbm.PrimaryKey = true
    MemberWithdrawalDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &memberWithdrawal
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(MemberWithdrawalDbm.ColumnMemberId)

	MemberWithdrawalDbm.PrimaryInfo = new(df.PrimaryInfo)
	MemberWithdrawalDbm.PrimaryInfo.UniqueInfo = ui
	
	var memberWithdrawalMeta df.DBMeta = MemberWithdrawalDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["MemberWithdrawal"] = &memberWithdrawalMeta
}
