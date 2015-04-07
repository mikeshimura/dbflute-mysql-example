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

	var memberLogin df.DBMeta
	memberLogin = MemberLoginDbm
	MemberLoginDbm.DBMeta=&memberLogin
	memberLoginIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo member_login_id
	memberLoginIdSqlName.ColumnSqlName = "member_login_id"
	memberLoginIdSqlName.IrregularChar = false
	MemberLoginDbm.ColumnMemberLoginId = df.CCI(&memberLogin, "member_login_id", memberLoginIdSqlName, "", "", "Long.class", "memberLoginId", "", true, true,true, "bigserial", 19, 0, "nextval('member_login_member_login_id_seq'::regclass)",false,"","", "","","",false,"int64")
	memberIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo member_id
	memberIdSqlName.ColumnSqlName = "member_id"
	memberIdSqlName.IrregularChar = false
	MemberLoginDbm.ColumnMemberId = df.CCI(&memberLogin, "member_id", memberIdSqlName, "", "", "Integer.class", "memberId", "", false, false,true, "int4", 10, 0, "",false,"","", "member","","",false,"int64")
	loginDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo login_datetime
	loginDatetimeSqlName.ColumnSqlName = "login_datetime"
	loginDatetimeSqlName.IrregularChar = false
	MemberLoginDbm.ColumnLoginDatetime = df.CCI(&memberLogin, "login_datetime", loginDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "loginDatetime", "", false, false,true, "timestamp", 26, 3, "",false,"","", "","","",false,"df.Timestamp")
	mobileLoginFlgSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo mobile_login_flg
	mobileLoginFlgSqlName.ColumnSqlName = "mobile_login_flg"
	mobileLoginFlgSqlName.IrregularChar = false
	MemberLoginDbm.ColumnMobileLoginFlg = df.CCI(&memberLogin, "mobile_login_flg", mobileLoginFlgSqlName, "", "", "Integer.class", "mobileLoginFlg", "", false, false,true, "int4", 10, 0, "",false,"","", "","","",false,"int64")
	loginMemberStatusCodeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo login_member_status_code
	loginMemberStatusCodeSqlName.ColumnSqlName = "login_member_status_code"
	loginMemberStatusCodeSqlName.IrregularChar = false
	MemberLoginDbm.ColumnLoginMemberStatusCode = df.CCI(&memberLogin, "login_member_status_code", loginMemberStatusCodeSqlName, "", "", "String.class", "loginMemberStatusCode", "", false, false,true, "bpchar", 3, 0, "",false,"","", "memberStatus","","",false,"string")

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
	
	MemberLoginDbm.SequenceFlag = true
	MemberLoginDbm.SequenceName = "member_login_member_login_id_seq"
	MemberLoginDbm.SequenceIncrementSize = 1
	MemberLoginDbm.SequenceCacheSize = 0
	var memberLoginMeta df.DBMeta = MemberLoginDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["MemberLogin"] = &memberLoginMeta
}
