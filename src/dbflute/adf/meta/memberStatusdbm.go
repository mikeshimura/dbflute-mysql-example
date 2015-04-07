package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type MemberStatusDbm_T struct {
	df.BaseDBMeta
	ColumnMemberStatusCode *df.ColumnInfo
	ColumnMemberStatusName *df.ColumnInfo
	ColumnDescription *df.ColumnInfo
	ColumnDisplayOrder *df.ColumnInfo
}

func (b *MemberStatusDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *MemberStatusDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var MemberStatusDbm *MemberStatusDbm_T

func Create_MemberStatusDbm() {
	MemberStatusDbm = new(MemberStatusDbm_T)
	MemberStatusDbm.TableDbName = "member_status"
	MemberStatusDbm.TableDispName = "member_status"
	MemberStatusDbm.TablePropertyName = "memberStatus"
	MemberStatusDbm.TableSqlName = new(df.TableSqlName)
	MemberStatusDbm.TableSqlName.TableSqlName = "member_status"
	MemberStatusDbm.TableSqlName.CorrespondingDbName = MemberStatusDbm.TableDbName

	var memberStatus df.DBMeta
	memberStatus = MemberStatusDbm
	MemberStatusDbm.DBMeta=&memberStatus
	memberStatusCodeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo member_status_code
	memberStatusCodeSqlName.ColumnSqlName = "member_status_code"
	memberStatusCodeSqlName.IrregularChar = false
	MemberStatusDbm.ColumnMemberStatusCode = df.CCI(&memberStatus, "member_status_code", memberStatusCodeSqlName, "", "", "String.class", "memberStatusCode", "", true, false,true, "bpchar", 3, 0, "",false,"","", "","memberList,memberLoginList","",false,"string")
	memberStatusNameSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo member_status_name
	memberStatusNameSqlName.ColumnSqlName = "member_status_name"
	memberStatusNameSqlName.IrregularChar = false
	MemberStatusDbm.ColumnMemberStatusName = df.CCI(&memberStatus, "member_status_name", memberStatusNameSqlName, "", "", "String.class", "memberStatusName", "", false, false,true, "varchar", 50, 0, "",false,"","", "","","",false,"string")
	descriptionSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo description
	descriptionSqlName.ColumnSqlName = "description"
	descriptionSqlName.IrregularChar = false
	MemberStatusDbm.ColumnDescription = df.CCI(&memberStatus, "description", descriptionSqlName, "", "", "String.class", "description", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	displayOrderSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo display_order
	displayOrderSqlName.ColumnSqlName = "display_order"
	displayOrderSqlName.IrregularChar = false
	MemberStatusDbm.ColumnDisplayOrder = df.CCI(&memberStatus, "display_order", displayOrderSqlName, "", "", "Integer.class", "displayOrder", "", false, false,true, "int4", 10, 0, "",false,"","", "","","",false,"int64")

	MemberStatusDbm.ColumnInfoList = new(df.List)
	MemberStatusDbm.ColumnInfoList.Add(MemberStatusDbm.ColumnMemberStatusCode)
	MemberStatusDbm.ColumnInfoList.Add(MemberStatusDbm.ColumnMemberStatusName)
	MemberStatusDbm.ColumnInfoList.Add(MemberStatusDbm.ColumnDescription)
	MemberStatusDbm.ColumnInfoList.Add(MemberStatusDbm.ColumnDisplayOrder)


	MemberStatusDbm.ColumnInfoMap=make(map[string]int)
	MemberStatusDbm.ColumnInfoMap["memberStatusCode"]=0
		MemberStatusDbm.ColumnInfoMap["memberStatusName"]=1
		MemberStatusDbm.ColumnInfoMap["description"]=2
		MemberStatusDbm.ColumnInfoMap["displayOrder"]=3
	    MemberStatusDbm.PrimaryKey = true
    MemberStatusDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &memberStatus
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(MemberStatusDbm.ColumnMemberStatusCode)

	MemberStatusDbm.PrimaryInfo = new(df.PrimaryInfo)
	MemberStatusDbm.PrimaryInfo.UniqueInfo = ui
	
	var memberStatusMeta df.DBMeta = MemberStatusDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["MemberStatus"] = &memberStatusMeta
}
