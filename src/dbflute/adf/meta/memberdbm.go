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
	ColumnUpdateDatetime *df.ColumnInfo
	ColumnUpdateUser *df.ColumnInfo
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
	MemberDbm.Identity=true

	var member df.DBMeta
	member = MemberDbm
	MemberDbm.DBMeta=&member
	memberIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo MEMBER_ID
	memberIdSqlName.ColumnSqlName = "MEMBER_ID"
	memberIdSqlName.IrregularChar = false
	MemberDbm.ColumnMemberId = df.CCI(&member, "MEMBER_ID", memberIdSqlName, "", "", "Integer.class", "memberId", "", true, true,true, "INT", 10, 0, "",false,"","", "","memberAddressList,memberLoginList,purchaseList","",false,"int64")
	memberNameSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo MEMBER_NAME
	memberNameSqlName.ColumnSqlName = "MEMBER_NAME"
	memberNameSqlName.IrregularChar = false
	MemberDbm.ColumnMemberName = df.CCI(&member, "MEMBER_NAME", memberNameSqlName, "", "", "String.class", "memberName", "", false, false,true, "VARCHAR", 180, 0, "",false,"","", "","","",false,"string")
	memberAccountSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo MEMBER_ACCOUNT
	memberAccountSqlName.ColumnSqlName = "MEMBER_ACCOUNT"
	memberAccountSqlName.IrregularChar = false
	MemberDbm.ColumnMemberAccount = df.CCI(&member, "MEMBER_ACCOUNT", memberAccountSqlName, "", "", "String.class", "memberAccount", "", false, false,true, "VARCHAR", 50, 0, "",false,"","", "","","",false,"string")
	memberStatusCodeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo MEMBER_STATUS_CODE
	memberStatusCodeSqlName.ColumnSqlName = "MEMBER_STATUS_CODE"
	memberStatusCodeSqlName.IrregularChar = false
	MemberDbm.ColumnMemberStatusCode = df.CCI(&member, "MEMBER_STATUS_CODE", memberStatusCodeSqlName, "", "", "String.class", "memberStatusCode", "", false, false,true, "CHAR", 3, 0, "",false,"","", "memberStatus","","",false,"string")
	formalizedDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo FORMALIZED_DATETIME
	formalizedDatetimeSqlName.ColumnSqlName = "FORMALIZED_DATETIME"
	formalizedDatetimeSqlName.IrregularChar = false
	MemberDbm.ColumnFormalizedDatetime = df.CCI(&member, "FORMALIZED_DATETIME", formalizedDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "formalizedDatetime", "", false, false,false, "DATETIME", 19, 0, "",false,"","", "","","",false,"df.MysqlNullTimestamp")
	birthdateSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo BIRTHDATE
	birthdateSqlName.ColumnSqlName = "BIRTHDATE"
	birthdateSqlName.IrregularChar = false
	MemberDbm.ColumnBirthdate = df.CCI(&member, "BIRTHDATE", birthdateSqlName, "", "", "java.time.LocalDate.class", "birthdate", "", false, false,false, "DATE", 10, 0, "",false,"","", "","","",false,"df.MysqlNullDate")
	registerDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo REGISTER_DATETIME
	registerDatetimeSqlName.ColumnSqlName = "REGISTER_DATETIME"
	registerDatetimeSqlName.IrregularChar = false
	MemberDbm.ColumnRegisterDatetime = df.CCI(&member, "REGISTER_DATETIME", registerDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "registerDatetime", "", false, false,true, "DATETIME", 19, 0, "",false,"","", "","","",false,"df.MysqlTimestamp")
	registerUserSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo REGISTER_USER
	registerUserSqlName.ColumnSqlName = "REGISTER_USER"
	registerUserSqlName.IrregularChar = false
	MemberDbm.ColumnRegisterUser = df.CCI(&member, "REGISTER_USER", registerUserSqlName, "", "", "String.class", "registerUser", "", false, false,true, "VARCHAR", 200, 0, "",false,"","", "","","",false,"string")
	updateDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo UPDATE_DATETIME
	updateDatetimeSqlName.ColumnSqlName = "UPDATE_DATETIME"
	updateDatetimeSqlName.IrregularChar = false
	MemberDbm.ColumnUpdateDatetime = df.CCI(&member, "UPDATE_DATETIME", updateDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "updateDatetime", "", false, false,true, "DATETIME", 19, 0, "",false,"","", "","","",false,"df.MysqlTimestamp")
	updateUserSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo UPDATE_USER
	updateUserSqlName.ColumnSqlName = "UPDATE_USER"
	updateUserSqlName.IrregularChar = false
	MemberDbm.ColumnUpdateUser = df.CCI(&member, "UPDATE_USER", updateUserSqlName, "", "", "String.class", "updateUser", "", false, false,true, "VARCHAR", 200, 0, "",false,"","", "","","",false,"string")
	versionNoSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo VERSION_NO
	versionNoSqlName.ColumnSqlName = "VERSION_NO"
	versionNoSqlName.IrregularChar = false
	MemberDbm.ColumnVersionNo = df.CCI(&member, "VERSION_NO", versionNoSqlName, "", "", "Long.class", "versionNo", "", false, false,true, "BIGINT", 19, 0, "",false,"OptimisticLockType.VERSION_NO","", "","","",false,"int64")

	MemberDbm.ColumnInfoList = new(df.List)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnMemberId)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnMemberName)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnMemberAccount)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnMemberStatusCode)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnFormalizedDatetime)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnBirthdate)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnRegisterDatetime)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnRegisterUser)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnUpdateDatetime)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnUpdateUser)
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
		MemberDbm.ColumnInfoMap["updateDatetime"]=8
		MemberDbm.ColumnInfoMap["updateUser"]=9
		MemberDbm.ColumnInfoMap["versionNo"]=10
	    MemberDbm.PrimaryKey = true
    MemberDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &member
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(MemberDbm.ColumnMemberId)

	MemberDbm.PrimaryInfo = new(df.PrimaryInfo)
	MemberDbm.PrimaryInfo.UniqueInfo = ui
	
	MemberDbm.VersionNoFlag = true
	MemberDbm.VersionNoColumnInfo = MemberDbm.ColumnVersionNo
	
	var memberMeta df.DBMeta = MemberDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["Member"] = &memberMeta
}
