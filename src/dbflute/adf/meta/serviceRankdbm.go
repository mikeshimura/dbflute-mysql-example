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
func (b *ServiceRankDbm_T) CreateForeignInfoMap() {
	b.ForeignInfoMap = make(map[string]*df.ForeignInfo)
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
	serviceRankCodeSqlName.ColumnSqlName = "SERVICE_RANK_CODE"
	serviceRankCodeSqlName.IrregularChar = false
	ServiceRankDbm.ColumnServiceRankCode = df.CCI(&serviceRank, "SERVICE_RANK_CODE", serviceRankCodeSqlName, "", "", "String.class", "serviceRankCode", "", true, false,true, "CHAR", 3, 0, "",false,"","", "","memberServiceList","",false,"string")
	serviceRankNameSqlName := new(df.ColumnSqlName)
	serviceRankNameSqlName.ColumnSqlName = "SERVICE_RANK_NAME"
	serviceRankNameSqlName.IrregularChar = false
	ServiceRankDbm.ColumnServiceRankName = df.CCI(&serviceRank, "SERVICE_RANK_NAME", serviceRankNameSqlName, "", "", "String.class", "serviceRankName", "", false, false,true, "VARCHAR", 50, 0, "",false,"","", "","","",false,"string")
	servicePointIncidenceSqlName := new(df.ColumnSqlName)
	servicePointIncidenceSqlName.ColumnSqlName = "SERVICE_POINT_INCIDENCE"
	servicePointIncidenceSqlName.IrregularChar = false
	ServiceRankDbm.ColumnServicePointIncidence = df.CCI(&serviceRank, "SERVICE_POINT_INCIDENCE", servicePointIncidenceSqlName, "", "", "java.math.BigDecimal.class", "servicePointIncidence", "", false, false,true, "DECIMAL", 5, 3, "",false,"","", "","","",false,"df.Numeric")
	newAcceptableFlgSqlName := new(df.ColumnSqlName)
	newAcceptableFlgSqlName.ColumnSqlName = "NEW_ACCEPTABLE_FLG"
	newAcceptableFlgSqlName.IrregularChar = false
	ServiceRankDbm.ColumnNewAcceptableFlg = df.CCI(&serviceRank, "NEW_ACCEPTABLE_FLG", newAcceptableFlgSqlName, "", "", "Integer.class", "newAcceptableFlg", "", false, false,true, "INT", 10, 0, "",false,"","", "","","",false,"int64")
	descriptionSqlName := new(df.ColumnSqlName)
	descriptionSqlName.ColumnSqlName = "DESCRIPTION"
	descriptionSqlName.IrregularChar = false
	ServiceRankDbm.ColumnDescription = df.CCI(&serviceRank, "DESCRIPTION", descriptionSqlName, "", "", "String.class", "description", "", false, false,true, "VARCHAR", 200, 0, "",false,"","", "","","",false,"string")
	displayOrderSqlName := new(df.ColumnSqlName)
	displayOrderSqlName.ColumnSqlName = "DISPLAY_ORDER"
	displayOrderSqlName.IrregularChar = false
	ServiceRankDbm.ColumnDisplayOrder = df.CCI(&serviceRank, "DISPLAY_ORDER", displayOrderSqlName, "", "", "Integer.class", "displayOrder", "", false, false,true, "INT", 10, 0, "",false,"","", "","","",false,"int64")

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
