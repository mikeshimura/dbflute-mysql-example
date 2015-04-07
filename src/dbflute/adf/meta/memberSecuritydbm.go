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
	ColumnRegisterProcess *df.ColumnInfo
	ColumnRegisterUser *df.ColumnInfo
	ColumnUpdateDatetime *df.ColumnInfo
	ColumnUpdateProcess *df.ColumnInfo
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
	//colsqlname dayoo member_id
	memberIdSqlName.ColumnSqlName = "member_id"
	memberIdSqlName.IrregularChar = false
	MemberSecurityDbm.ColumnMemberId = df.CCI(&memberSecurity, "member_id", memberIdSqlName, "", "", "Integer.class", "memberId", "", true, false,true, "int4", 10, 0, "",false,"","", "member","","",false,"int64")
	loginPasswordSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo login_password
	loginPasswordSqlName.ColumnSqlName = "login_password"
	loginPasswordSqlName.IrregularChar = false
	MemberSecurityDbm.ColumnLoginPassword = df.CCI(&memberSecurity, "login_password", loginPasswordSqlName, "", "", "String.class", "loginPassword", "", false, false,true, "varchar", 50, 0, "",false,"","", "","","",false,"string")
	reminderQuestionSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo reminder_question
	reminderQuestionSqlName.ColumnSqlName = "reminder_question"
	reminderQuestionSqlName.IrregularChar = false
	MemberSecurityDbm.ColumnReminderQuestion = df.CCI(&memberSecurity, "reminder_question", reminderQuestionSqlName, "", "", "String.class", "reminderQuestion", "", false, false,true, "varchar", 50, 0, "",false,"","", "","","",false,"string")
	reminderAnswerSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo reminder_answer
	reminderAnswerSqlName.ColumnSqlName = "reminder_answer"
	reminderAnswerSqlName.IrregularChar = false
	MemberSecurityDbm.ColumnReminderAnswer = df.CCI(&memberSecurity, "reminder_answer", reminderAnswerSqlName, "", "", "String.class", "reminderAnswer", "", false, false,true, "varchar", 50, 0, "",false,"","", "","","",false,"string")
	reminderUseCountSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo reminder_use_count
	reminderUseCountSqlName.ColumnSqlName = "reminder_use_count"
	reminderUseCountSqlName.IrregularChar = false
	MemberSecurityDbm.ColumnReminderUseCount = df.CCI(&memberSecurity, "reminder_use_count", reminderUseCountSqlName, "", "", "Integer.class", "reminderUseCount", "", false, false,true, "int4", 10, 0, "",false,"","", "","","",false,"int64")
	registerDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo register_datetime
	registerDatetimeSqlName.ColumnSqlName = "register_datetime"
	registerDatetimeSqlName.IrregularChar = false
	MemberSecurityDbm.ColumnRegisterDatetime = df.CCI(&memberSecurity, "register_datetime", registerDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "registerDatetime", "", false, false,true, "timestamp", 26, 3, "",false,"","", "","","",false,"df.Timestamp")
	registerProcessSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo register_process
	registerProcessSqlName.ColumnSqlName = "register_process"
	registerProcessSqlName.IrregularChar = false
	MemberSecurityDbm.ColumnRegisterProcess = df.CCI(&memberSecurity, "register_process", registerProcessSqlName, "", "", "String.class", "registerProcess", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	registerUserSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo register_user
	registerUserSqlName.ColumnSqlName = "register_user"
	registerUserSqlName.IrregularChar = false
	MemberSecurityDbm.ColumnRegisterUser = df.CCI(&memberSecurity, "register_user", registerUserSqlName, "", "", "String.class", "registerUser", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	updateDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo update_datetime
	updateDatetimeSqlName.ColumnSqlName = "update_datetime"
	updateDatetimeSqlName.IrregularChar = false
	MemberSecurityDbm.ColumnUpdateDatetime = df.CCI(&memberSecurity, "update_datetime", updateDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "updateDatetime", "", false, false,true, "timestamp", 26, 3, "",false,"","", "","","",false,"df.Timestamp")
	updateProcessSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo update_process
	updateProcessSqlName.ColumnSqlName = "update_process"
	updateProcessSqlName.IrregularChar = false
	MemberSecurityDbm.ColumnUpdateProcess = df.CCI(&memberSecurity, "update_process", updateProcessSqlName, "", "", "String.class", "updateProcess", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	updateUserSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo update_user
	updateUserSqlName.ColumnSqlName = "update_user"
	updateUserSqlName.IrregularChar = false
	MemberSecurityDbm.ColumnUpdateUser = df.CCI(&memberSecurity, "update_user", updateUserSqlName, "", "", "String.class", "updateUser", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	versionNoSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo version_no
	versionNoSqlName.ColumnSqlName = "version_no"
	versionNoSqlName.IrregularChar = false
	MemberSecurityDbm.ColumnVersionNo = df.CCI(&memberSecurity, "version_no", versionNoSqlName, "", "", "Long.class", "versionNo", "", false, false,true, "int8", 19, 0, "",false,"OptimisticLockType.VERSION_NO","", "","","",false,"int64")

	MemberSecurityDbm.ColumnInfoList = new(df.List)
	MemberSecurityDbm.ColumnInfoList.Add(MemberSecurityDbm.ColumnMemberId)
	MemberSecurityDbm.ColumnInfoList.Add(MemberSecurityDbm.ColumnLoginPassword)
	MemberSecurityDbm.ColumnInfoList.Add(MemberSecurityDbm.ColumnReminderQuestion)
	MemberSecurityDbm.ColumnInfoList.Add(MemberSecurityDbm.ColumnReminderAnswer)
	MemberSecurityDbm.ColumnInfoList.Add(MemberSecurityDbm.ColumnReminderUseCount)
	MemberSecurityDbm.ColumnInfoList.Add(MemberSecurityDbm.ColumnRegisterDatetime)
	MemberSecurityDbm.ColumnInfoList.Add(MemberSecurityDbm.ColumnRegisterProcess)
	MemberSecurityDbm.ColumnInfoList.Add(MemberSecurityDbm.ColumnRegisterUser)
	MemberSecurityDbm.ColumnInfoList.Add(MemberSecurityDbm.ColumnUpdateDatetime)
	MemberSecurityDbm.ColumnInfoList.Add(MemberSecurityDbm.ColumnUpdateProcess)
	MemberSecurityDbm.ColumnInfoList.Add(MemberSecurityDbm.ColumnUpdateUser)
	MemberSecurityDbm.ColumnInfoList.Add(MemberSecurityDbm.ColumnVersionNo)


	MemberSecurityDbm.ColumnInfoMap=make(map[string]int)
	MemberSecurityDbm.ColumnInfoMap["memberId"]=0
		MemberSecurityDbm.ColumnInfoMap["loginPassword"]=1
		MemberSecurityDbm.ColumnInfoMap["reminderQuestion"]=2
		MemberSecurityDbm.ColumnInfoMap["reminderAnswer"]=3
		MemberSecurityDbm.ColumnInfoMap["reminderUseCount"]=4
		MemberSecurityDbm.ColumnInfoMap["registerDatetime"]=5
		MemberSecurityDbm.ColumnInfoMap["registerProcess"]=6
		MemberSecurityDbm.ColumnInfoMap["registerUser"]=7
		MemberSecurityDbm.ColumnInfoMap["updateDatetime"]=8
		MemberSecurityDbm.ColumnInfoMap["updateProcess"]=9
		MemberSecurityDbm.ColumnInfoMap["updateUser"]=10
		MemberSecurityDbm.ColumnInfoMap["versionNo"]=11
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
