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
	ColumnRegisterUser *df.ColumnInfo
	ColumnUpdateDatetime *df.ColumnInfo
	ColumnUpdateUser *df.ColumnInfo
}

func (b *MemberWithdrawalDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}
func (b *MemberWithdrawalDbm_T) foreignMember() *df.ForeignInfo {
	columns := []*df.ColumnInfo{
		MemberWithdrawalDbm.GetColumnInfoByPropertyName("memberId"),
		MemberDbm.GetColumnInfoByPropertyName("memberId"),
	}

	return b.BaseDBMeta.Cfi("FK_MEMBER_WITHDRAWAL_MEMBER", "Member",
		columns, 0, true, false, false, false,
		"", nil, false, "memberWithdrawalAsOne")
}	
func (b *MemberWithdrawalDbm_T) foreignWithdrawalReason() *df.ForeignInfo {
	columns := []*df.ColumnInfo{
		MemberWithdrawalDbm.GetColumnInfoByPropertyName("withdrawalReasonCode"),
		WithdrawalReasonDbm.GetColumnInfoByPropertyName("withdrawalReasonCode"),
	}

	return b.BaseDBMeta.Cfi("FK_MEMBER_WITHDRAWAL_WITHDRAWAL_REASON", "WithdrawalReason",
		columns, 1, false, false, false, false,
		"", nil, false, "memberWithdrawalList")
}	
func (b *MemberWithdrawalDbm_T) CreateForeignInfoMap() {
	b.ForeignInfoMap = make(map[string]*df.ForeignInfo)
	b.ForeignInfoMap["Member"] = b.foreignMember()
	b.ForeignInfoMap["WithdrawalReason"] = b.foreignWithdrawalReason()
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
	memberIdSqlName.ColumnSqlName = "MEMBER_ID"
	memberIdSqlName.IrregularChar = false
	MemberWithdrawalDbm.ColumnMemberId = df.CCI(&memberWithdrawal, "MEMBER_ID", memberIdSqlName, "", "", "Integer.class", "memberId", "", true, false,true, "INT", 10, 0, "",false,"","", "member","","",false,"int64")
	withdrawalReasonCodeSqlName := new(df.ColumnSqlName)
	withdrawalReasonCodeSqlName.ColumnSqlName = "WITHDRAWAL_REASON_CODE"
	withdrawalReasonCodeSqlName.IrregularChar = false
	MemberWithdrawalDbm.ColumnWithdrawalReasonCode = df.CCI(&memberWithdrawal, "WITHDRAWAL_REASON_CODE", withdrawalReasonCodeSqlName, "", "", "String.class", "withdrawalReasonCode", "", false, false,false, "CHAR", 3, 0, "",false,"","", "withdrawalReason","","",false,"sql.NullString")
	withdrawalReasonInputTextSqlName := new(df.ColumnSqlName)
	withdrawalReasonInputTextSqlName.ColumnSqlName = "WITHDRAWAL_REASON_INPUT_TEXT"
	withdrawalReasonInputTextSqlName.IrregularChar = false
	MemberWithdrawalDbm.ColumnWithdrawalReasonInputText = df.CCI(&memberWithdrawal, "WITHDRAWAL_REASON_INPUT_TEXT", withdrawalReasonInputTextSqlName, "", "", "String.class", "withdrawalReasonInputText", "", false, false,false, "TEXT", 65535, 0, "",false,"","", "","","",false,"sql.NullString")
	withdrawalDatetimeSqlName := new(df.ColumnSqlName)
	withdrawalDatetimeSqlName.ColumnSqlName = "WITHDRAWAL_DATETIME"
	withdrawalDatetimeSqlName.IrregularChar = false
	MemberWithdrawalDbm.ColumnWithdrawalDatetime = df.CCI(&memberWithdrawal, "WITHDRAWAL_DATETIME", withdrawalDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "withdrawalDatetime", "", false, false,true, "DATETIME", 19, 0, "",false,"","", "","","",false,"df.MysqlTimestamp")
	registerDatetimeSqlName := new(df.ColumnSqlName)
	registerDatetimeSqlName.ColumnSqlName = "REGISTER_DATETIME"
	registerDatetimeSqlName.IrregularChar = false
	MemberWithdrawalDbm.ColumnRegisterDatetime = df.CCI(&memberWithdrawal, "REGISTER_DATETIME", registerDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "registerDatetime", "", false, false,true, "DATETIME", 19, 0, "",false,"","", "","","",false,"df.MysqlTimestamp")
	registerUserSqlName := new(df.ColumnSqlName)
	registerUserSqlName.ColumnSqlName = "REGISTER_USER"
	registerUserSqlName.IrregularChar = false
	MemberWithdrawalDbm.ColumnRegisterUser = df.CCI(&memberWithdrawal, "REGISTER_USER", registerUserSqlName, "", "", "String.class", "registerUser", "", false, false,true, "VARCHAR", 200, 0, "",false,"","", "","","",false,"string")
	updateDatetimeSqlName := new(df.ColumnSqlName)
	updateDatetimeSqlName.ColumnSqlName = "UPDATE_DATETIME"
	updateDatetimeSqlName.IrregularChar = false
	MemberWithdrawalDbm.ColumnUpdateDatetime = df.CCI(&memberWithdrawal, "UPDATE_DATETIME", updateDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "updateDatetime", "", false, false,true, "DATETIME", 19, 0, "",false,"","", "","","",false,"df.MysqlTimestamp")
	updateUserSqlName := new(df.ColumnSqlName)
	updateUserSqlName.ColumnSqlName = "UPDATE_USER"
	updateUserSqlName.IrregularChar = false
	MemberWithdrawalDbm.ColumnUpdateUser = df.CCI(&memberWithdrawal, "UPDATE_USER", updateUserSqlName, "", "", "String.class", "updateUser", "", false, false,true, "VARCHAR", 200, 0, "",false,"","", "","","",false,"string")

	MemberWithdrawalDbm.ColumnInfoList = new(df.List)
	MemberWithdrawalDbm.ColumnInfoList.Add(MemberWithdrawalDbm.ColumnMemberId)
	MemberWithdrawalDbm.ColumnInfoList.Add(MemberWithdrawalDbm.ColumnWithdrawalReasonCode)
	MemberWithdrawalDbm.ColumnInfoList.Add(MemberWithdrawalDbm.ColumnWithdrawalReasonInputText)
	MemberWithdrawalDbm.ColumnInfoList.Add(MemberWithdrawalDbm.ColumnWithdrawalDatetime)
	MemberWithdrawalDbm.ColumnInfoList.Add(MemberWithdrawalDbm.ColumnRegisterDatetime)
	MemberWithdrawalDbm.ColumnInfoList.Add(MemberWithdrawalDbm.ColumnRegisterUser)
	MemberWithdrawalDbm.ColumnInfoList.Add(MemberWithdrawalDbm.ColumnUpdateDatetime)
	MemberWithdrawalDbm.ColumnInfoList.Add(MemberWithdrawalDbm.ColumnUpdateUser)


	MemberWithdrawalDbm.ColumnInfoMap=make(map[string]int)
	MemberWithdrawalDbm.ColumnInfoMap["memberId"]=0
		MemberWithdrawalDbm.ColumnInfoMap["withdrawalReasonCode"]=1
		MemberWithdrawalDbm.ColumnInfoMap["withdrawalReasonInputText"]=2
		MemberWithdrawalDbm.ColumnInfoMap["withdrawalDatetime"]=3
		MemberWithdrawalDbm.ColumnInfoMap["registerDatetime"]=4
		MemberWithdrawalDbm.ColumnInfoMap["registerUser"]=5
		MemberWithdrawalDbm.ColumnInfoMap["updateDatetime"]=6
		MemberWithdrawalDbm.ColumnInfoMap["updateUser"]=7
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
