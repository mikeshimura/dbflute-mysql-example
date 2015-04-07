package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type ServiceRankDbm_T struct {
	df.BaseDBMeta
	ColumnServiceRankCode *df.ColumnInfo
	ColumnServiceRankName *df.ColumnInfo
	ColumnServicePointIncidence *df.ColumnInfo
	ColumnNewAcceptableFlg *df.ColumnInfo
	ColumnDescription *df.ColumnInfo
	ColumnDisplayOrder *df.ColumnInfo
}

func (b *ServiceRankDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *ServiceRankDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var ServiceRankDbm *ServiceRankDbm_T

func Create_ServiceRankDbm() {
	ServiceRankDbm = new(ServiceRankDbm_T)
	ServiceRankDbm.TableDbName = "service_rank"
	ServiceRankDbm.TableDispName = "service_rank"
	ServiceRankDbm.TablePropertyName = "serviceRank"
	ServiceRankDbm.TableSqlName = new(df.TableSqlName)
	ServiceRankDbm.TableSqlName.TableSqlName = "service_rank"
	ServiceRankDbm.TableSqlName.CorrespondingDbName = ServiceRankDbm.TableDbName

	var serviceRank df.DBMeta
	serviceRank = ServiceRankDbm
	ServiceRankDbm.DBMeta=&serviceRank
	serviceRankCodeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo service_rank_code
	serviceRankCodeSqlName.ColumnSqlName = "service_rank_code"
	serviceRankCodeSqlName.IrregularChar = false
	ServiceRankDbm.ColumnServiceRankCode = df.CCI(&serviceRank, "service_rank_code", serviceRankCodeSqlName, "", "", "String.class", "serviceRankCode", "", true, false,true, "bpchar", 3, 0, "",false,"","", "","memberServiceList","",false,"string")
	serviceRankNameSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo service_rank_name
	serviceRankNameSqlName.ColumnSqlName = "service_rank_name"
	serviceRankNameSqlName.IrregularChar = false
	ServiceRankDbm.ColumnServiceRankName = df.CCI(&serviceRank, "service_rank_name", serviceRankNameSqlName, "", "", "String.class", "serviceRankName", "", false, false,true, "varchar", 50, 0, "",false,"","", "","","",false,"string")
	servicePointIncidenceSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo service_point_incidence
	servicePointIncidenceSqlName.ColumnSqlName = "service_point_incidence"
	servicePointIncidenceSqlName.IrregularChar = false
	ServiceRankDbm.ColumnServicePointIncidence = df.CCI(&serviceRank, "service_point_incidence", servicePointIncidenceSqlName, "", "", "java.math.BigDecimal.class", "servicePointIncidence", "", false, false,true, "numeric", 5, 3, "",false,"","", "","","",false,"df.Numeric")
	newAcceptableFlgSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo new_acceptable_flg
	newAcceptableFlgSqlName.ColumnSqlName = "new_acceptable_flg"
	newAcceptableFlgSqlName.IrregularChar = false
	ServiceRankDbm.ColumnNewAcceptableFlg = df.CCI(&serviceRank, "new_acceptable_flg", newAcceptableFlgSqlName, "", "", "Integer.class", "newAcceptableFlg", "", false, false,true, "int4", 10, 0, "",false,"","", "","","",false,"int64")
	descriptionSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo description
	descriptionSqlName.ColumnSqlName = "description"
	descriptionSqlName.IrregularChar = false
	ServiceRankDbm.ColumnDescription = df.CCI(&serviceRank, "description", descriptionSqlName, "", "", "String.class", "description", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	displayOrderSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo display_order
	displayOrderSqlName.ColumnSqlName = "display_order"
	displayOrderSqlName.IrregularChar = false
	ServiceRankDbm.ColumnDisplayOrder = df.CCI(&serviceRank, "display_order", displayOrderSqlName, "", "", "Integer.class", "displayOrder", "", false, false,true, "int4", 10, 0, "",false,"","", "","","",false,"int64")

	ServiceRankDbm.ColumnInfoList = new(df.List)
	ServiceRankDbm.ColumnInfoList.Add(ServiceRankDbm.ColumnServiceRankCode)
	ServiceRankDbm.ColumnInfoList.Add(ServiceRankDbm.ColumnServiceRankName)
	ServiceRankDbm.ColumnInfoList.Add(ServiceRankDbm.ColumnServicePointIncidence)
	ServiceRankDbm.ColumnInfoList.Add(ServiceRankDbm.ColumnNewAcceptableFlg)
	ServiceRankDbm.ColumnInfoList.Add(ServiceRankDbm.ColumnDescription)
	ServiceRankDbm.ColumnInfoList.Add(ServiceRankDbm.ColumnDisplayOrder)


	ServiceRankDbm.ColumnInfoMap=make(map[string]int)
	ServiceRankDbm.ColumnInfoMap["serviceRankCode"]=0
		ServiceRankDbm.ColumnInfoMap["serviceRankName"]=1
		ServiceRankDbm.ColumnInfoMap["servicePointIncidence"]=2
		ServiceRankDbm.ColumnInfoMap["newAcceptableFlg"]=3
		ServiceRankDbm.ColumnInfoMap["description"]=4
		ServiceRankDbm.ColumnInfoMap["displayOrder"]=5
	    ServiceRankDbm.PrimaryKey = true
    ServiceRankDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &serviceRank
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(ServiceRankDbm.ColumnServiceRankCode)

	ServiceRankDbm.PrimaryInfo = new(df.PrimaryInfo)
	ServiceRankDbm.PrimaryInfo.UniqueInfo = ui
	
	var serviceRankMeta df.DBMeta = ServiceRankDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["ServiceRank"] = &serviceRankMeta
}
