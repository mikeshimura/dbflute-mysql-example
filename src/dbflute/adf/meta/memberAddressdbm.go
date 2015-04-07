package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type MemberAddressDbm_T struct {
	df.BaseDBMeta
	ColumnMemberAddressId *df.ColumnInfo
	ColumnMemberId *df.ColumnInfo
	ColumnValidBeginDate *df.ColumnInfo
	ColumnValidEndDate *df.ColumnInfo
	ColumnAddress *df.ColumnInfo
	ColumnRegionId *df.ColumnInfo
	ColumnRegisterDatetime *df.ColumnInfo
	ColumnRegisterProcess *df.ColumnInfo
	ColumnRegisterUser *df.ColumnInfo
	ColumnUpdateDatetime *df.ColumnInfo
	ColumnUpdateProcess *df.ColumnInfo
	ColumnUpdateUser *df.ColumnInfo
	ColumnVersionNo *df.ColumnInfo
}

func (b *MemberAddressDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *MemberAddressDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var MemberAddressDbm *MemberAddressDbm_T

func Create_MemberAddressDbm() {
	MemberAddressDbm = new(MemberAddressDbm_T)
	MemberAddressDbm.TableDbName = "member_address"
	MemberAddressDbm.TableDispName = "member_address"
	MemberAddressDbm.TablePropertyName = "memberAddress"
	MemberAddressDbm.TableSqlName = new(df.TableSqlName)
	MemberAddressDbm.TableSqlName.TableSqlName = "member_address"
	MemberAddressDbm.TableSqlName.CorrespondingDbName = MemberAddressDbm.TableDbName

	var memberAddress df.DBMeta
	memberAddress = MemberAddressDbm
	MemberAddressDbm.DBMeta=&memberAddress
	memberAddressIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo member_address_id
	memberAddressIdSqlName.ColumnSqlName = "member_address_id"
	memberAddressIdSqlName.IrregularChar = false
	MemberAddressDbm.ColumnMemberAddressId = df.CCI(&memberAddress, "member_address_id", memberAddressIdSqlName, "", "", "Integer.class", "memberAddressId", "", true, true,true, "serial", 10, 0, "nextval('member_address_member_address_id_seq'::regclass)",false,"","", "","","",false,"int64")
	memberIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo member_id
	memberIdSqlName.ColumnSqlName = "member_id"
	memberIdSqlName.IrregularChar = false
	MemberAddressDbm.ColumnMemberId = df.CCI(&memberAddress, "member_id", memberIdSqlName, "", "", "Integer.class", "memberId", "", false, false,true, "int4", 10, 0, "",false,"","", "member","","",false,"int64")
	validBeginDateSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo valid_begin_date
	validBeginDateSqlName.ColumnSqlName = "valid_begin_date"
	validBeginDateSqlName.IrregularChar = false
	MemberAddressDbm.ColumnValidBeginDate = df.CCI(&memberAddress, "valid_begin_date", validBeginDateSqlName, "", "", "java.time.LocalDate.class", "validBeginDate", "", false, false,true, "date", 13, 0, "",false,"","", "","","",false,"df.Date")
	validEndDateSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo valid_end_date
	validEndDateSqlName.ColumnSqlName = "valid_end_date"
	validEndDateSqlName.IrregularChar = false
	MemberAddressDbm.ColumnValidEndDate = df.CCI(&memberAddress, "valid_end_date", validEndDateSqlName, "", "", "java.time.LocalDate.class", "validEndDate", "", false, false,true, "date", 13, 0, "",false,"","", "","","",false,"df.Date")
	addressSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo address
	addressSqlName.ColumnSqlName = "address"
	addressSqlName.IrregularChar = false
	MemberAddressDbm.ColumnAddress = df.CCI(&memberAddress, "address", addressSqlName, "", "", "String.class", "address", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	regionIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo region_id
	regionIdSqlName.ColumnSqlName = "region_id"
	regionIdSqlName.IrregularChar = false
	MemberAddressDbm.ColumnRegionId = df.CCI(&memberAddress, "region_id", regionIdSqlName, "", "", "Integer.class", "regionId", "", false, false,true, "int4", 10, 0, "",false,"","", "region","","",false,"int64")
	registerDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo register_datetime
	registerDatetimeSqlName.ColumnSqlName = "register_datetime"
	registerDatetimeSqlName.IrregularChar = false
	MemberAddressDbm.ColumnRegisterDatetime = df.CCI(&memberAddress, "register_datetime", registerDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "registerDatetime", "", false, false,true, "timestamp", 26, 3, "",false,"","", "","","",false,"df.Timestamp")
	registerProcessSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo register_process
	registerProcessSqlName.ColumnSqlName = "register_process"
	registerProcessSqlName.IrregularChar = false
	MemberAddressDbm.ColumnRegisterProcess = df.CCI(&memberAddress, "register_process", registerProcessSqlName, "", "", "String.class", "registerProcess", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	registerUserSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo register_user
	registerUserSqlName.ColumnSqlName = "register_user"
	registerUserSqlName.IrregularChar = false
	MemberAddressDbm.ColumnRegisterUser = df.CCI(&memberAddress, "register_user", registerUserSqlName, "", "", "String.class", "registerUser", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	updateDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo update_datetime
	updateDatetimeSqlName.ColumnSqlName = "update_datetime"
	updateDatetimeSqlName.IrregularChar = false
	MemberAddressDbm.ColumnUpdateDatetime = df.CCI(&memberAddress, "update_datetime", updateDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "updateDatetime", "", false, false,true, "timestamp", 26, 3, "",false,"","", "","","",false,"df.Timestamp")
	updateProcessSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo update_process
	updateProcessSqlName.ColumnSqlName = "update_process"
	updateProcessSqlName.IrregularChar = false
	MemberAddressDbm.ColumnUpdateProcess = df.CCI(&memberAddress, "update_process", updateProcessSqlName, "", "", "String.class", "updateProcess", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	updateUserSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo update_user
	updateUserSqlName.ColumnSqlName = "update_user"
	updateUserSqlName.IrregularChar = false
	MemberAddressDbm.ColumnUpdateUser = df.CCI(&memberAddress, "update_user", updateUserSqlName, "", "", "String.class", "updateUser", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	versionNoSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo version_no
	versionNoSqlName.ColumnSqlName = "version_no"
	versionNoSqlName.IrregularChar = false
	MemberAddressDbm.ColumnVersionNo = df.CCI(&memberAddress, "version_no", versionNoSqlName, "", "", "Long.class", "versionNo", "", false, false,true, "int8", 19, 0, "",false,"OptimisticLockType.VERSION_NO","", "","","",false,"int64")

	MemberAddressDbm.ColumnInfoList = new(df.List)
	MemberAddressDbm.ColumnInfoList.Add(MemberAddressDbm.ColumnMemberAddressId)
	MemberAddressDbm.ColumnInfoList.Add(MemberAddressDbm.ColumnMemberId)
	MemberAddressDbm.ColumnInfoList.Add(MemberAddressDbm.ColumnValidBeginDate)
	MemberAddressDbm.ColumnInfoList.Add(MemberAddressDbm.ColumnValidEndDate)
	MemberAddressDbm.ColumnInfoList.Add(MemberAddressDbm.ColumnAddress)
	MemberAddressDbm.ColumnInfoList.Add(MemberAddressDbm.ColumnRegionId)
	MemberAddressDbm.ColumnInfoList.Add(MemberAddressDbm.ColumnRegisterDatetime)
	MemberAddressDbm.ColumnInfoList.Add(MemberAddressDbm.ColumnRegisterProcess)
	MemberAddressDbm.ColumnInfoList.Add(MemberAddressDbm.ColumnRegisterUser)
	MemberAddressDbm.ColumnInfoList.Add(MemberAddressDbm.ColumnUpdateDatetime)
	MemberAddressDbm.ColumnInfoList.Add(MemberAddressDbm.ColumnUpdateProcess)
	MemberAddressDbm.ColumnInfoList.Add(MemberAddressDbm.ColumnUpdateUser)
	MemberAddressDbm.ColumnInfoList.Add(MemberAddressDbm.ColumnVersionNo)


	MemberAddressDbm.ColumnInfoMap=make(map[string]int)
	MemberAddressDbm.ColumnInfoMap["memberAddressId"]=0
		MemberAddressDbm.ColumnInfoMap["memberId"]=1
		MemberAddressDbm.ColumnInfoMap["validBeginDate"]=2
		MemberAddressDbm.ColumnInfoMap["validEndDate"]=3
		MemberAddressDbm.ColumnInfoMap["address"]=4
		MemberAddressDbm.ColumnInfoMap["regionId"]=5
		MemberAddressDbm.ColumnInfoMap["registerDatetime"]=6
		MemberAddressDbm.ColumnInfoMap["registerProcess"]=7
		MemberAddressDbm.ColumnInfoMap["registerUser"]=8
		MemberAddressDbm.ColumnInfoMap["updateDatetime"]=9
		MemberAddressDbm.ColumnInfoMap["updateProcess"]=10
		MemberAddressDbm.ColumnInfoMap["updateUser"]=11
		MemberAddressDbm.ColumnInfoMap["versionNo"]=12
	    MemberAddressDbm.PrimaryKey = true
    MemberAddressDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &memberAddress
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(MemberAddressDbm.ColumnMemberAddressId)

	MemberAddressDbm.PrimaryInfo = new(df.PrimaryInfo)
	MemberAddressDbm.PrimaryInfo.UniqueInfo = ui
	
	MemberAddressDbm.SequenceFlag = true
	MemberAddressDbm.SequenceName = "member_address_member_address_id_seq"
	MemberAddressDbm.SequenceIncrementSize = 1
	MemberAddressDbm.SequenceCacheSize = 0
	MemberAddressDbm.VersionNoFlag = true
	MemberAddressDbm.VersionNoColumnInfo = MemberAddressDbm.ColumnVersionNo
	
	var memberAddressMeta df.DBMeta = MemberAddressDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["MemberAddress"] = &memberAddressMeta
}
