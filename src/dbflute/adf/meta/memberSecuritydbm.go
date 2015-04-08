package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type MemberSecurityDbm_T struct {
	df.BaseDBMeta
	ColumnMemberId *df.ColumnInfo
	ColumnLoginPassword *df.ColumnInfo
	ColumnReminderQuestion *df.ColumnInfo
	ColumnReminderAnswer *df.ColumnInfo
	ColumnReminderUseCount *df.ColumnInfo
	ColumnRegisterDatetime *df.ColumnInfo
	ColumnRegisterUser *df.ColumnInfo
	ColumnUpdateDatetime *df.ColumnInfo
	ColumnUpdateUser *df.ColumnInfo
	ColumnVersionNo *df.ColumnInfo
}

func (b *MemberSecurityDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *MemberSecurityDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var MemberSecurityDbm *MemberSecurityDbm_T

func Create_MemberSecurityDbm() {
	MemberSecurityDbm = new(MemberSecurityDbm_T)
	MemberSecurityDbm.TableDbName = "member_security"
	MemberSecurityDbm.TableDispName = "member_security"
	MemberSecurityDbm.TablePropertyName = "memberSecurity"
	MemberSecurityDbm.TableSqlName = new(df.TableSqlName)
	MemberSecurityDbm.TableSqlName.TableSqlName = "member_security"
	MemberSecurityDbm.TableSqlName.CorrespondingDbName = MemberSecurityDbm.TableDbName

	var memberSecurity df.DBMeta
	memberSecurity = MemberSecurityDbm
	MemberSecurityDbm.DBMeta=&memberSecurity
	memberIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo MEMBER_ID
	memberIdSqlName.ColumnSqlName = "MEMBER_ID"
	memberIdSqlName.IrregularChar = false
	MemberSecurityDbm.ColumnMemberId = df.CCI(&memberSecurity, "MEMBER_ID", memberIdSqlName, "", "", "Integer.class", "memberId", "", true, false,true, "INT", 10, 0, "",false,"","", "member","","",false,"int64")
	loginPasswordSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo LOGIN_PASSWORD
	loginPasswordSqlName.ColumnSqlName = "LOGIN_PASSWORD"
	loginPasswordSqlName.IrregularChar = false
	MemberSecurityDbm.ColumnLoginPassword = df.CCI(&memberSecurity, "LOGIN_PASSWORD", loginPasswordSqlName, "", "", "String.class", "loginPassword", "", false, false,true, "VARCHAR", 50, 0, "",false,"","", "","","",false,"string")
	reminderQuestionSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo REMINDER_QUESTION
	reminderQuestionSqlName.ColumnSqlName = "REMINDER_QUESTION"
	reminderQuestionSqlName.IrregularChar = false
	MemberSecurityDbm.ColumnReminderQuestion = df.CCI(&memberSecurity, "REMINDER_QUESTION", reminderQuestionSqlName, "", "", "String.class", "reminderQuestion", "", false, false,true, "VARCHAR", 50, 0, "",false,"","", "","","",false,"string")
	reminderAnswerSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo REMINDER_ANSWER
	reminderAnswerSqlName.ColumnSqlName = "REMINDER_ANSWER"
	reminderAnswerSqlName.IrregularChar = false
	MemberSecurityDbm.ColumnReminderAnswer = df.CCI(&memberSecurity, "REMINDER_ANSWER", reminderAnswerSqlName, "", "", "String.class", "reminderAnswer", "", false, false,true, "VARCHAR", 50, 0, "",false,"","", "","","",false,"string")
	reminderUseCountSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo REMINDER_USE_COUNT
	reminderUseCountSqlName.ColumnSqlName = "REMINDER_USE_COUNT"
	reminderUseCountSqlName.IrregularChar = false
	MemberSecurityDbm.ColumnReminderUseCount = df.CCI(&memberSecurity, "REMINDER_USE_COUNT", reminderUseCountSqlName, "", "", "Integer.class", "reminderUseCount", "", false, false,true, "INT", 10, 0, "",false,"","", "","","",false,"int64")
	registerDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo REGISTER_DATETIME
	registerDatetimeSqlName.ColumnSqlName = "REGISTER_DATETIME"
	registerDatetimeSqlName.IrregularChar = false
	MemberSecurityDbm.ColumnRegisterDatetime = df.CCI(&memberSecurity, "REGISTER_DATETIME", registerDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "registerDatetime", "", false, false,true, "DATETIME", 19, 0, "",false,"","", "","","",false,"df.MysqlTimestamp")
	registerUserSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo REGISTER_USER
	registerUserSqlName.ColumnSqlName = "REGISTER_USER"
	registerUserSqlName.IrregularChar = false
	MemberSecurityDbm.ColumnRegisterUser = df.CCI(&memberSecurity, "REGISTER_USER", registerUserSqlName, "", "", "String.class", "registerUser", "", false, false,true, "VARCHAR", 200, 0, "",false,"","", "","","",false,"string")
	updateDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo UPDATE_DATETIME
	updateDatetimeSqlName.ColumnSqlName = "UPDATE_DATETIME"
	updateDatetimeSqlName.IrregularChar = false
	MemberSecurityDbm.ColumnUpdateDatetime = df.CCI(&memberSecurity, "UPDATE_DATETIME", updateDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "updateDatetime", "", false, false,true, "DATETIME", 19, 0, "",false,"","", "","","",false,"df.MysqlTimestamp")
	updateUserSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo UPDATE_USER
	updateUserSqlName.ColumnSqlName = "UPDATE_USER"
	updateUserSqlName.IrregularChar = false
	MemberSecurityDbm.ColumnUpdateUser = df.CCI(&memberSecurity, "UPDATE_USER", updateUserSqlName, "", "", "String.class", "updateUser", "", false, false,true, "VARCHAR", 200, 0, "",false,"","", "","","",false,"string")
	versionNoSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo VERSION_NO
	versionNoSqlName.ColumnSqlName = "VERSION_NO"
	versionNoSqlName.IrregularChar = false
	MemberSecurityDbm.ColumnVersionNo = df.CCI(&memberSecurity, "VERSION_NO", versionNoSqlName, "", "", "Long.class", "versionNo", "", false, false,true, "BIGINT", 19, 0, "",false,"OptimisticLockType.VERSION_NO","", "","","",false,"int64")

	MemberSecurityDbm.ColumnInfoList = new(df.List)
	MemberSecurityDbm.ColumnInfoList.Add(MemberSecurityDbm.ColumnMemberId)
	MemberSecurityDbm.ColumnInfoList.Add(MemberSecurityDbm.ColumnLoginPassword)
	MemberSecurityDbm.ColumnInfoList.Add(MemberSecurityDbm.ColumnReminderQuestion)
	MemberSecurityDbm.ColumnInfoList.Add(MemberSecurityDbm.ColumnReminderAnswer)
	MemberSecurityDbm.ColumnInfoList.Add(MemberSecurityDbm.ColumnReminderUseCount)
	MemberSecurityDbm.ColumnInfoList.Add(MemberSecurityDbm.ColumnRegisterDatetime)
	MemberSecurityDbm.ColumnInfoList.Add(MemberSecurityDbm.ColumnRegisterUser)
	MemberSecurityDbm.ColumnInfoList.Add(MemberSecurityDbm.ColumnUpdateDatetime)
	MemberSecurityDbm.ColumnInfoList.Add(MemberSecurityDbm.ColumnUpdateUser)
	MemberSecurityDbm.ColumnInfoList.Add(MemberSecurityDbm.ColumnVersionNo)


	MemberSecurityDbm.ColumnInfoMap=make(map[string]int)
	MemberSecurityDbm.ColumnInfoMap["memberId"]=0
		MemberSecurityDbm.ColumnInfoMap["loginPassword"]=1
		MemberSecurityDbm.ColumnInfoMap["reminderQuestion"]=2
		MemberSecurityDbm.ColumnInfoMap["reminderAnswer"]=3
		MemberSecurityDbm.ColumnInfoMap["reminderUseCount"]=4
		MemberSecurityDbm.ColumnInfoMap["registerDatetime"]=5
		MemberSecurityDbm.ColumnInfoMap["registerUser"]=6
		MemberSecurityDbm.ColumnInfoMap["updateDatetime"]=7
		MemberSecurityDbm.ColumnInfoMap["updateUser"]=8
		MemberSecurityDbm.ColumnInfoMap["versionNo"]=9
	    MemberSecurityDbm.PrimaryKey = true
    MemberSecurityDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &memberSecurity
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(MemberSecurityDbm.ColumnMemberId)

	MemberSecurityDbm.PrimaryInfo = new(df.PrimaryInfo)
	MemberSecurityDbm.PrimaryInfo.UniqueInfo = ui
	
	MemberSecurityDbm.VersionNoFlag = true
	MemberSecurityDbm.VersionNoColumnInfo = MemberSecurityDbm.ColumnVersionNo
	
	var memberSecurityMeta df.DBMeta = MemberSecurityDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["MemberSecurity"] = &memberSecurityMeta
}
