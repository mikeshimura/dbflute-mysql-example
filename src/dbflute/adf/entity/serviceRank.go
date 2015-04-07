package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type ServiceRank struct {
	serviceRankCode string
	serviceRankName string
	servicePointIncidence df.Numeric
	newAcceptableFlg int64
	description string
	displayOrder int64
	df.BaseEntity
}

func CreateServiceRank() *ServiceRank{
	serviceRank :=new(ServiceRank)
	serviceRank.SetUp()
	return serviceRank 
}

func (l *ServiceRank) GetServiceRankCode () string {
	return l.serviceRankCode
}
func (l *ServiceRank) GetServiceRankName () string {
	return l.serviceRankName
}
func (l *ServiceRank) GetServicePointIncidence () df.Numeric {
	return l.servicePointIncidence
}
func (l *ServiceRank) GetNewAcceptableFlg () int64 {
	return l.newAcceptableFlg
}
func (l *ServiceRank) GetDescription () string {
	return l.description
}
func (l *ServiceRank) GetDisplayOrder () int64 {
	return l.displayOrder
}

func (t *ServiceRank) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 6)
	i[0] = &(t.serviceRankCode)
	i[1] = &(t.serviceRankName)
	i[2] = &(t.servicePointIncidence)
	i[3] = &(t.newAcceptableFlg)
	i[4] = &(t.description)
	i[5] = &(t.displayOrder)
	return i
}


func (t *ServiceRank) AsTableDbName() string {
	return "ServiceRank"
}

func (t *ServiceRank) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("serviceRankCode") == false {
            return false 
        }
        return true;
}
func (t *ServiceRank) SetServiceRankCode(serviceRankCode string) {
	t.AddPropertyName("serviceRankCode")
	t.serviceRankCode = serviceRankCode
}
func (t *ServiceRank) SetServiceRankName(serviceRankName string) {
	t.AddPropertyName("serviceRankName")
	t.serviceRankName = serviceRankName
}
func (t *ServiceRank) SetServicePointIncidence(servicePointIncidence df.Numeric) {
	t.AddPropertyName("servicePointIncidence")
	t.servicePointIncidence = servicePointIncidence
}
func (t *ServiceRank) SetNewAcceptableFlg(newAcceptableFlg int64) {
	t.AddPropertyName("newAcceptableFlg")
	t.newAcceptableFlg = newAcceptableFlg
}
func (t *ServiceRank) SetDescription(description string) {
	t.AddPropertyName("description")
	t.description = description
}
func (t *ServiceRank) SetDisplayOrder(displayOrder int64) {
	t.AddPropertyName("displayOrder")
	t.displayOrder = displayOrder
}

func (t *ServiceRank) SetUp(){
	t.servicePointIncidence.DecPoint = 3
	
}
func (t *ServiceRank)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}