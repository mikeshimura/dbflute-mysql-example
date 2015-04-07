package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type MemberDbm_T struct {
	df.BaseDBMeta
	ColumnMemberId *df.ColumnInfo
	ColumnMemberName *df.ColumnInfo
	ColumnMemberAccount *df.ColumnInfo
	ColumnMemberStatusCode *df.ColumnInfo
	ColumnFormalizedDatetime *df.ColumnInfo
	ColumnBirthdate *df.ColumnInfo
	ColumnRegisterDatetime *df.ColumnInfo
	ColumnRegisterUser *df.ColumnInfo
	ColumnRegisterProcess *df.ColumnInfo
	ColumnUpdateDatetime *df.ColumnInfo
	ColumnUpdateUser *df.ColumnInfo
	ColumnUpdateProcess *df.ColumnInfo
	ColumnVersionNo *df.ColumnInfo
}

func (b *MemberDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *MemberDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var MemberDbm *MemberDbm_T

func Create_MemberDbm() {
	MemberDbm = new(MemberDbm_T)
	MemberDbm.TableDbName = "member"
	MemberDbm.TableDispName = "member"
	MemberDbm.TablePropertyName = "member"
	MemberDbm.TableSqlName = new(df.TableSqlName)
	MemberDbm.TableSqlName.TableSqlName = "member"
	MemberDbm.TableSqlName.CorrespondingDbName = MemberDbm.TableDbName

	var member df.DBMeta
	member = MemberDbm
	MemberDbm.DBMeta=&member
	memberIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo member_id
	memberIdSqlName.ColumnSqlName = "member_id"
	memberIdSqlName.IrregularChar = false
	MemberDbm.ColumnMemberId = df.CCI(&member, "member_id", memberIdSqlName, "", "", "Integer.class", "memberId", "", true, true,true, "serial", 10, 0, "nextval('member_member_id_seq'::regclass)",false,"","", "","memberAddressList,memberLoginList,purchaseList","",false,"int64")
	memberNameSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo member_name
	memberNameSqlName.ColumnSqlName = "member_name"
	memberNameSqlName.IrregularChar = false
	MemberDbm.ColumnMemberName = df.CCI(&member, "member_name", memberNameSqlName, "", "", "String.class", "memberName", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	memberAccountSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo member_account
	memberAccountSqlName.ColumnSqlName = "member_account"
	memberAccountSqlName.IrregularChar = false
	MemberDbm.ColumnMemberAccount = df.CCI(&member, "member_account", memberAccountSqlName, "", "", "String.class", "memberAccount", "", false, false,true, "varchar", 50, 0, "",false,"","", "","","",false,"string")
	memberStatusCodeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo member_status_code
	memberStatusCodeSqlName.ColumnSqlName = "member_status_code"
	memberStatusCodeSqlName.IrregularChar = false
	MemberDbm.ColumnMemberStatusCode = df.CCI(&member, "member_status_code", memberStatusCodeSqlName, "", "", "String.class", "memberStatusCode", "", false, false,true, "bpchar", 3, 0, "",false,"","", "memberStatus","","",false,"string")
	formalizedDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo formalized_datetime
	formalizedDatetimeSqlName.ColumnSqlName = "formalized_datetime"
	formalizedDatetimeSqlName.IrregularChar = false
	MemberDbm.ColumnFormalizedDatetime = df.CCI(&member, "formalized_datetime", formalizedDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "formalizedDatetime", "", false, false,false, "timestamp", 26, 3, "",false,"","", "","","",false,"df.NullTimestamp")
	birthdateSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo birthdate
	birthdateSqlName.ColumnSqlName = "birthdate"
	birthdateSqlName.IrregularChar = false
	MemberDbm.ColumnBirthdate = df.CCI(&member, "birthdate", birthdateSqlName, "", "", "java.time.LocalDate.class", "birthdate", "", false, false,false, "date", 13, 0, "",false,"","", "","","",false,"df.NullDate")
	registerDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo register_datetime
	registerDatetimeSqlName.ColumnSqlName = "register_datetime"
	registerDatetimeSqlName.IrregularChar = false
	MemberDbm.ColumnRegisterDatetime = df.CCI(&member, "register_datetime", registerDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "registerDatetime", "", false, false,true, "timestamp", 26, 3, "",false,"","", "","","",false,"df.Timestamp")
	registerUserSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo register_user
	registerUserSqlName.ColumnSqlName = "register_user"
	registerUserSqlName.IrregularChar = false
	MemberDbm.ColumnRegisterUser = df.CCI(&member, "register_user", registerUserSqlName, "", "", "String.class", "registerUser", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	registerProcessSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo register_process
	registerProcessSqlName.ColumnSqlName = "register_process"
	registerProcessSqlName.IrregularChar = false
	MemberDbm.ColumnRegisterProcess = df.CCI(&member, "register_process", registerProcessSqlName, "", "", "String.class", "registerProcess", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	updateDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo update_datetime
	updateDatetimeSqlName.ColumnSqlName = "update_datetime"
	updateDatetimeSqlName.IrregularChar = false
	MemberDbm.ColumnUpdateDatetime = df.CCI(&member, "update_datetime", updateDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "updateDatetime", "", false, false,true, "timestamp", 26, 3, "",false,"","", "","","",false,"df.Timestamp")
	updateUserSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo update_user
	updateUserSqlName.ColumnSqlName = "update_user"
	updateUserSqlName.IrregularChar = false
	MemberDbm.ColumnUpdateUser = df.CCI(&member, "update_user", updateUserSqlName, "", "", "String.class", "updateUser", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	updateProcessSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo update_process
	updateProcessSqlName.ColumnSqlName = "update_process"
	updateProcessSqlName.IrregularChar = false
	MemberDbm.ColumnUpdateProcess = df.CCI(&member, "update_process", updateProcessSqlName, "", "", "String.class", "updateProcess", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	versionNoSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo version_no
	versionNoSqlName.ColumnSqlName = "version_no"
	versionNoSqlName.IrregularChar = false
	MemberDbm.ColumnVersionNo = df.CCI(&member, "version_no", versionNoSqlName, "", "", "Long.class", "versionNo", "", false, false,true, "int8", 19, 0, "",false,"OptimisticLockType.VERSION_NO","", "","","",false,"int64")

	MemberDbm.ColumnInfoList = new(df.List)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnMemberId)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnMemberName)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnMemberAccount)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnMemberStatusCode)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnFormalizedDatetime)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnBirthdate)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnRegisterDatetime)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnRegisterUser)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnRegisterProcess)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnUpdateDatetime)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnUpdateUser)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnUpdateProcess)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnVersionNo)


	MemberDbm.ColumnInfoMap=make(map[string]int)
	MemberDbm.ColumnInfoMap["memberId"]=0
		MemberDbm.ColumnInfoMap["memberName"]=1
		MemberDbm.ColumnInfoMap["memberAccount"]=2
		MemberDbm.ColumnInfoMap["memberStatusCode"]=3
		MemberDbm.ColumnInfoMap["formalizedDatetime"]=4
		MemberDbm.ColumnInfoMap["birthdate"]=5
		MemberDbm.ColumnInfoMap["registerDatetime"]=6
		MemberDbm.ColumnInfoMap["registerUser"]=7
		MemberDbm.ColumnInfoMap["registerProcess"]=8
		MemberDbm.ColumnInfoMap["updateDatetime"]=9
		MemberDbm.ColumnInfoMap["updateUser"]=10
		MemberDbm.ColumnInfoMap["updateProcess"]=11
		MemberDbm.ColumnInfoMap["versionNo"]=12
	    MemberDbm.PrimaryKey = true
    MemberDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &member
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(MemberDbm.ColumnMemberId)

	MemberDbm.PrimaryInfo = new(df.PrimaryInfo)
	MemberDbm.PrimaryInfo.UniqueInfo = ui
	
	MemberDbm.SequenceFlag = true
	MemberDbm.SequenceName = "member_member_id_seq"
	MemberDbm.SequenceIncrementSize = 1
	MemberDbm.SequenceCacheSize = 0
	MemberDbm.VersionNoFlag = true
	MemberDbm.VersionNoColumnInfo = MemberDbm.ColumnVersionNo
	
	var memberMeta df.DBMeta = MemberDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["Member"] = &memberMeta
}
