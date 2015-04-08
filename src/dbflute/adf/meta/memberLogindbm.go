package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type MemberLoginDbm_T struct {
	df.BaseDBMeta
	ColumnMemberLoginId *df.ColumnInfo
	ColumnMemberId *df.ColumnInfo
	ColumnLoginDatetime *df.ColumnInfo
	ColumnMobileLoginFlg *df.ColumnInfo
	ColumnLoginMemberStatusCode *df.ColumnInfo
}

func (b *MemberLoginDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *MemberLoginDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var MemberLoginDbm *MemberLoginDbm_T

func Create_MemberLoginDbm() {
	MemberLoginDbm = new(MemberLoginDbm_T)
	MemberLoginDbm.TableDbName = "member_login"
	MemberLoginDbm.TableDispName = "member_login"
	MemberLoginDbm.TablePropertyName = "memberLogin"
	MemberLoginDbm.TableSqlName = new(df.TableSqlName)
	MemberLoginDbm.TableSqlName.TableSqlName = "member_login"
	MemberLoginDbm.TableSqlName.CorrespondingDbName = MemberLoginDbm.TableDbName
	MemberLoginDbm.Identity=true

	var memberLogin df.DBMeta
	memberLogin = MemberLoginDbm
	MemberLoginDbm.DBMeta=&memberLogin
	memberLoginIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo MEMBER_LOGIN_ID
	memberLoginIdSqlName.ColumnSqlName = "MEMBER_LOGIN_ID"
	memberLoginIdSqlName.IrregularChar = false
	MemberLoginDbm.ColumnMemberLoginId = df.CCI(&memberLogin, "MEMBER_LOGIN_ID", memberLoginIdSqlName, "", "", "Long.class", "memberLoginId", "", true, true,true, "BIGINT", 19, 0, "",false,"","", "","","",false,"int64")
	memberIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo MEMBER_ID
	memberIdSqlName.ColumnSqlName = "MEMBER_ID"
	memberIdSqlName.IrregularChar = false
	MemberLoginDbm.ColumnMemberId = df.CCI(&memberLogin, "MEMBER_ID", memberIdSqlName, "", "", "Integer.class", "memberId", "", false, false,true, "INT", 10, 0, "",false,"","", "member","","",false,"int64")
	loginDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo LOGIN_DATETIME
	loginDatetimeSqlName.ColumnSqlName = "LOGIN_DATETIME"
	loginDatetimeSqlName.IrregularChar = false
	MemberLoginDbm.ColumnLoginDatetime = df.CCI(&memberLogin, "LOGIN_DATETIME", loginDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "loginDatetime", "", false, false,true, "DATETIME", 19, 0, "",false,"","", "","","",false,"df.MysqlTimestamp")
	mobileLoginFlgSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo MOBILE_LOGIN_FLG
	mobileLoginFlgSqlName.ColumnSqlName = "MOBILE_LOGIN_FLG"
	mobileLoginFlgSqlName.IrregularChar = false
	MemberLoginDbm.ColumnMobileLoginFlg = df.CCI(&memberLogin, "MOBILE_LOGIN_FLG", mobileLoginFlgSqlName, "", "", "Integer.class", "mobileLoginFlg", "", false, false,true, "INT", 10, 0, "",false,"","", "","","",false,"int64")
	loginMemberStatusCodeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo LOGIN_MEMBER_STATUS_CODE
	loginMemberStatusCodeSqlName.ColumnSqlName = "LOGIN_MEMBER_STATUS_CODE"
	loginMemberStatusCodeSqlName.IrregularChar = false
	MemberLoginDbm.ColumnLoginMemberStatusCode = df.CCI(&memberLogin, "LOGIN_MEMBER_STATUS_CODE", loginMemberStatusCodeSqlName, "", "", "String.class", "loginMemberStatusCode", "", false, false,true, "CHAR", 3, 0, "",false,"","", "memberStatus","","",false,"string")

	MemberLoginDbm.ColumnInfoList = new(df.List)
	MemberLoginDbm.ColumnInfoList.Add(MemberLoginDbm.ColumnMemberLoginId)
	MemberLoginDbm.ColumnInfoList.Add(MemberLoginDbm.ColumnMemberId)
	MemberLoginDbm.ColumnInfoList.Add(MemberLoginDbm.ColumnLoginDatetime)
	MemberLoginDbm.ColumnInfoList.Add(MemberLoginDbm.ColumnMobileLoginFlg)
	MemberLoginDbm.ColumnInfoList.Add(MemberLoginDbm.ColumnLoginMemberStatusCode)


	MemberLoginDbm.ColumnInfoMap=make(map[string]int)
	MemberLoginDbm.ColumnInfoMap["memberLoginId"]=0
		MemberLoginDbm.ColumnInfoMap["memberId"]=1
		MemberLoginDbm.ColumnInfoMap["loginDatetime"]=2
		MemberLoginDbm.ColumnInfoMap["mobileLoginFlg"]=3
		MemberLoginDbm.ColumnInfoMap["loginMemberStatusCode"]=4
	    MemberLoginDbm.PrimaryKey = true
    MemberLoginDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &memberLogin
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(MemberLoginDbm.ColumnMemberLoginId)

	MemberLoginDbm.PrimaryInfo = new(df.PrimaryInfo)
	MemberLoginDbm.PrimaryInfo.UniqueInfo = ui
	
	var memberLoginMeta df.DBMeta = MemberLoginDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["MemberLogin"] = &memberLoginMeta
}
