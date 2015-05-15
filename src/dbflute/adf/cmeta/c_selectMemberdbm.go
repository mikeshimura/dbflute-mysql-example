package cmeta

import (
	"github.com/mikeshimura/dbflute/df"
)

type C_SelectMemberDbm_T struct {
	df.BaseDBMeta
	ColumnMemberId *df.ColumnInfo
	ColumnMemberName *df.ColumnInfo
	ColumnMemberAccount *df.ColumnInfo
	ColumnBirthdate *df.ColumnInfo
	ColumnFormalizedDatetime *df.ColumnInfo
	ColumnMemberStatusCode *df.ColumnInfo
	ColumnMemberStatusName *df.ColumnInfo
	ColumnDescription *df.ColumnInfo
}

func (b *C_SelectMemberDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *C_SelectMemberDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var C_SelectMemberDbm *C_SelectMemberDbm_T

func Create_C_SelectMemberDbm() {
	C_SelectMemberDbm = new(C_SelectMemberDbm_T)
	C_SelectMemberDbm.TableDbName = "c_SelectMember"
	C_SelectMemberDbm.TableDispName = "c_SelectMember"
	C_SelectMemberDbm.TablePropertyName = "c_SelectMember"
	C_SelectMemberDbm.TableSqlName = new(df.TableSqlName)
	C_SelectMemberDbm.TableSqlName.TableSqlName = "c_SelectMember"
	C_SelectMemberDbm.TableSqlName.CorrespondingDbName = C_SelectMemberDbm.TableDbName


	var selectMember df.DBMeta
	selectMember = C_SelectMemberDbm
	C_SelectMemberDbm.DBMeta=&selectMember
	memberIdSqlName := new(df.ColumnSqlName)
	memberIdSqlName.ColumnSqlName = "member_id"
	memberIdSqlName.IrregularChar = false
	C_SelectMemberDbm.ColumnMemberId = df.CCI(&selectMember, "member_id", memberIdSqlName, "", "", "Integer.class", "memberId", "", false, false,false, "INT", 11, 0, "",false,"","", "","","",false,"sql.NullInt64")
	memberNameSqlName := new(df.ColumnSqlName)
	memberNameSqlName.ColumnSqlName = "member_name"
	memberNameSqlName.IrregularChar = false
	C_SelectMemberDbm.ColumnMemberName = df.CCI(&selectMember, "member_name", memberNameSqlName, "", "", "String.class", "memberName", "", false, false,false, "VARCHAR", 180, 0, "",false,"","", "","","",false,"sql.NullString")
	memberAccountSqlName := new(df.ColumnSqlName)
	memberAccountSqlName.ColumnSqlName = "member_account"
	memberAccountSqlName.IrregularChar = false
	C_SelectMemberDbm.ColumnMemberAccount = df.CCI(&selectMember, "member_account", memberAccountSqlName, "", "", "String.class", "memberAccount", "", false, false,false, "VARCHAR", 50, 0, "",false,"","", "","","",false,"sql.NullString")
	birthdateSqlName := new(df.ColumnSqlName)
	birthdateSqlName.ColumnSqlName = "birthdate"
	birthdateSqlName.IrregularChar = false
	C_SelectMemberDbm.ColumnBirthdate = df.CCI(&selectMember, "birthdate", birthdateSqlName, "", "", "java.time.LocalDate.class", "birthdate", "", false, false,false, "DATE", 10, 0, "",false,"","", "","","",false,"df.MysqlNullDate")
	formalizedDatetimeSqlName := new(df.ColumnSqlName)
	formalizedDatetimeSqlName.ColumnSqlName = "formalized_datetime"
	formalizedDatetimeSqlName.IrregularChar = false
	C_SelectMemberDbm.ColumnFormalizedDatetime = df.CCI(&selectMember, "formalized_datetime", formalizedDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "formalizedDatetime", "", false, false,false, "DATETIME", 19, 0, "",false,"","", "","","",false,"df.MysqlNullTimestamp")
	memberStatusCodeSqlName := new(df.ColumnSqlName)
	memberStatusCodeSqlName.ColumnSqlName = "member_status_code"
	memberStatusCodeSqlName.IrregularChar = false
	C_SelectMemberDbm.ColumnMemberStatusCode = df.CCI(&selectMember, "member_status_code", memberStatusCodeSqlName, "", "", "String.class", "memberStatusCode", "", false, false,false, "CHAR", 3, 0, "",false,"","", "","","",false,"sql.NullString")
	memberStatusNameSqlName := new(df.ColumnSqlName)
	memberStatusNameSqlName.ColumnSqlName = "member_status_name"
	memberStatusNameSqlName.IrregularChar = false
	C_SelectMemberDbm.ColumnMemberStatusName = df.CCI(&selectMember, "member_status_name", memberStatusNameSqlName, "", "", "String.class", "memberStatusName", "", false, false,false, "VARCHAR", 50, 0, "",false,"","", "","","",false,"sql.NullString")
	descriptionSqlName := new(df.ColumnSqlName)
	descriptionSqlName.ColumnSqlName = "description"
	descriptionSqlName.IrregularChar = false
	C_SelectMemberDbm.ColumnDescription = df.CCI(&selectMember, "description", descriptionSqlName, "", "", "String.class", "description", "", false, false,false, "VARCHAR", 200, 0, "",false,"","", "","","",false,"sql.NullString")

	C_SelectMemberDbm.ColumnInfoList = new(df.List)
	C_SelectMemberDbm.ColumnInfoList.Add(C_SelectMemberDbm.ColumnMemberId)
	C_SelectMemberDbm.ColumnInfoList.Add(C_SelectMemberDbm.ColumnMemberName)
	C_SelectMemberDbm.ColumnInfoList.Add(C_SelectMemberDbm.ColumnMemberAccount)
	C_SelectMemberDbm.ColumnInfoList.Add(C_SelectMemberDbm.ColumnBirthdate)
	C_SelectMemberDbm.ColumnInfoList.Add(C_SelectMemberDbm.ColumnFormalizedDatetime)
	C_SelectMemberDbm.ColumnInfoList.Add(C_SelectMemberDbm.ColumnMemberStatusCode)
	C_SelectMemberDbm.ColumnInfoList.Add(C_SelectMemberDbm.ColumnMemberStatusName)
	C_SelectMemberDbm.ColumnInfoList.Add(C_SelectMemberDbm.ColumnDescription)


	C_SelectMemberDbm.ColumnInfoMap=make(map[string]int)
	C_SelectMemberDbm.ColumnInfoMap["memberId"]=0
		C_SelectMemberDbm.ColumnInfoMap["memberName"]=1
		C_SelectMemberDbm.ColumnInfoMap["memberAccount"]=2
		C_SelectMemberDbm.ColumnInfoMap["birthdate"]=3
		C_SelectMemberDbm.ColumnInfoMap["formalizedDatetime"]=4
		C_SelectMemberDbm.ColumnInfoMap["memberStatusCode"]=5
		C_SelectMemberDbm.ColumnInfoMap["memberStatusName"]=6
		C_SelectMemberDbm.ColumnInfoMap["description"]=7
	    C_SelectMemberDbm.PrimaryKey = false
    C_SelectMemberDbm.CompoundPrimaryKey = false

	var selectMemberMeta df.DBMeta = C_SelectMemberDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["c_SelectMember"] = &selectMemberMeta
}
