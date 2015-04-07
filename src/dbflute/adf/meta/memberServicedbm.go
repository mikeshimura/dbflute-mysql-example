package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type MemberServiceDbm_T struct {
	df.BaseDBMeta
	ColumnMemberServiceId *df.ColumnInfo
	ColumnMemberId *df.ColumnInfo
	ColumnServicePointCount *df.ColumnInfo
	ColumnServiceRankCode *df.ColumnInfo
	ColumnRegisterDatetime *df.ColumnInfo
	ColumnRegisterProcess *df.ColumnInfo
	ColumnRegisterUser *df.ColumnInfo
	ColumnUpdateDatetime *df.ColumnInfo
	ColumnUpdateProcess *df.ColumnInfo
	ColumnUpdateUser *df.ColumnInfo
	ColumnVersionNo *df.ColumnInfo
}

func (b *MemberServiceDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *MemberServiceDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var MemberServiceDbm *MemberServiceDbm_T

func Create_MemberServiceDbm() {
	MemberServiceDbm = new(MemberServiceDbm_T)
	MemberServiceDbm.TableDbName = "member_service"
	MemberServiceDbm.TableDispName = "member_service"
	MemberServiceDbm.TablePropertyName = "memberService"
	MemberServiceDbm.TableSqlName = new(df.TableSqlName)
	MemberServiceDbm.TableSqlName.TableSqlName = "member_service"
	MemberServiceDbm.TableSqlName.CorrespondingDbName = MemberServiceDbm.TableDbName

	var memberService df.DBMeta
	memberService = MemberServiceDbm
	MemberServiceDbm.DBMeta=&memberService
	memberServiceIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo member_service_id
	memberServiceIdSqlName.ColumnSqlName = "member_service_id"
	memberServiceIdSqlName.IrregularChar = false
	MemberServiceDbm.ColumnMemberServiceId = df.CCI(&memberService, "member_service_id", memberServiceIdSqlName, "", "", "Integer.class", "memberServiceId", "", true, true,true, "serial", 10, 0, "nextval('member_service_member_service_id_seq'::regclass)",false,"","", "","","",false,"int64")
	memberIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo member_id
	memberIdSqlName.ColumnSqlName = "member_id"
	memberIdSqlName.IrregularChar = false
	MemberServiceDbm.ColumnMemberId = df.CCI(&memberService, "member_id", memberIdSqlName, "", "", "Integer.class", "memberId", "", false, false,true, "int4", 10, 0, "",false,"","", "member","","",false,"int64")
	servicePointCountSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo service_point_count
	servicePointCountSqlName.ColumnSqlName = "service_point_count"
	servicePointCountSqlName.IrregularChar = false
	MemberServiceDbm.ColumnServicePointCount = df.CCI(&memberService, "service_point_count", servicePointCountSqlName, "", "", "Integer.class", "servicePointCount", "", false, false,true, "int4", 10, 0, "",false,"","", "","","",false,"int64")
	serviceRankCodeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo service_rank_code
	serviceRankCodeSqlName.ColumnSqlName = "service_rank_code"
	serviceRankCodeSqlName.IrregularChar = false
	MemberServiceDbm.ColumnServiceRankCode = df.CCI(&memberService, "service_rank_code", serviceRankCodeSqlName, "", "", "String.class", "serviceRankCode", "", false, false,true, "bpchar", 3, 0, "",false,"","", "serviceRank","","",false,"string")
	registerDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo register_datetime
	registerDatetimeSqlName.ColumnSqlName = "register_datetime"
	registerDatetimeSqlName.IrregularChar = false
	MemberServiceDbm.ColumnRegisterDatetime = df.CCI(&memberService, "register_datetime", registerDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "registerDatetime", "", false, false,true, "timestamp", 26, 3, "",false,"","", "","","",false,"df.Timestamp")
	registerProcessSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo register_process
	registerProcessSqlName.ColumnSqlName = "register_process"
	registerProcessSqlName.IrregularChar = false
	MemberServiceDbm.ColumnRegisterProcess = df.CCI(&memberService, "register_process", registerProcessSqlName, "", "", "String.class", "registerProcess", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	registerUserSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo register_user
	registerUserSqlName.ColumnSqlName = "register_user"
	registerUserSqlName.IrregularChar = false
	MemberServiceDbm.ColumnRegisterUser = df.CCI(&memberService, "register_user", registerUserSqlName, "", "", "String.class", "registerUser", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	updateDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo update_datetime
	updateDatetimeSqlName.ColumnSqlName = "update_datetime"
	updateDatetimeSqlName.IrregularChar = false
	MemberServiceDbm.ColumnUpdateDatetime = df.CCI(&memberService, "update_datetime", updateDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "updateDatetime", "", false, false,true, "timestamp", 26, 3, "",false,"","", "","","",false,"df.Timestamp")
	updateProcessSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo update_process
	updateProcessSqlName.ColumnSqlName = "update_process"
	updateProcessSqlName.IrregularChar = false
	MemberServiceDbm.ColumnUpdateProcess = df.CCI(&memberService, "update_process", updateProcessSqlName, "", "", "String.class", "updateProcess", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	updateUserSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo update_user
	updateUserSqlName.ColumnSqlName = "update_user"
	updateUserSqlName.IrregularChar = false
	MemberServiceDbm.ColumnUpdateUser = df.CCI(&memberService, "update_user", updateUserSqlName, "", "", "String.class", "updateUser", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	versionNoSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo version_no
	versionNoSqlName.ColumnSqlName = "version_no"
	versionNoSqlName.IrregularChar = false
	MemberServiceDbm.ColumnVersionNo = df.CCI(&memberService, "version_no", versionNoSqlName, "", "", "Long.class", "versionNo", "", false, false,true, "int8", 19, 0, "",false,"OptimisticLockType.VERSION_NO","", "","","",false,"int64")

	MemberServiceDbm.ColumnInfoList = new(df.List)
	MemberServiceDbm.ColumnInfoList.Add(MemberServiceDbm.ColumnMemberServiceId)
	MemberServiceDbm.ColumnInfoList.Add(MemberServiceDbm.ColumnMemberId)
	MemberServiceDbm.ColumnInfoList.Add(MemberServiceDbm.ColumnServicePointCount)
	MemberServiceDbm.ColumnInfoList.Add(MemberServiceDbm.ColumnServiceRankCode)
	MemberServiceDbm.ColumnInfoList.Add(MemberServiceDbm.ColumnRegisterDatetime)
	MemberServiceDbm.ColumnInfoList.Add(MemberServiceDbm.ColumnRegisterProcess)
	MemberServiceDbm.ColumnInfoList.Add(MemberServiceDbm.ColumnRegisterUser)
	MemberServiceDbm.ColumnInfoList.Add(MemberServiceDbm.ColumnUpdateDatetime)
	MemberServiceDbm.ColumnInfoList.Add(MemberServiceDbm.ColumnUpdateProcess)
	MemberServiceDbm.ColumnInfoList.Add(MemberServiceDbm.ColumnUpdateUser)
	MemberServiceDbm.ColumnInfoList.Add(MemberServiceDbm.ColumnVersionNo)


	MemberServiceDbm.ColumnInfoMap=make(map[string]int)
	MemberServiceDbm.ColumnInfoMap["memberServiceId"]=0
		MemberServiceDbm.ColumnInfoMap["memberId"]=1
		MemberServiceDbm.ColumnInfoMap["servicePointCount"]=2
		MemberServiceDbm.ColumnInfoMap["serviceRankCode"]=3
		MemberServiceDbm.ColumnInfoMap["registerDatetime"]=4
		MemberServiceDbm.ColumnInfoMap["registerProcess"]=5
		MemberServiceDbm.ColumnInfoMap["registerUser"]=6
		MemberServiceDbm.ColumnInfoMap["updateDatetime"]=7
		MemberServiceDbm.ColumnInfoMap["updateProcess"]=8
		MemberServiceDbm.ColumnInfoMap["updateUser"]=9
		MemberServiceDbm.ColumnInfoMap["versionNo"]=10
	    MemberServiceDbm.PrimaryKey = true
    MemberServiceDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &memberService
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(MemberServiceDbm.ColumnMemberServiceId)

	MemberServiceDbm.PrimaryInfo = new(df.PrimaryInfo)
	MemberServiceDbm.PrimaryInfo.UniqueInfo = ui
	
	MemberServiceDbm.SequenceFlag = true
	MemberServiceDbm.SequenceName = "member_service_member_service_id_seq"
	MemberServiceDbm.SequenceIncrementSize = 1
	MemberServiceDbm.SequenceCacheSize = 0
	MemberServiceDbm.VersionNoFlag = true
	MemberServiceDbm.VersionNoColumnInfo = MemberServiceDbm.ColumnVersionNo
	
	var memberServiceMeta df.DBMeta = MemberServiceDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["MemberService"] = &memberServiceMeta
}
