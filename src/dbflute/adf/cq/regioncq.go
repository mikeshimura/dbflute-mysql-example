package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type RegionCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	RegionId *df.ConditionValue
	RegionName *df.ConditionValue
}

func (q *RegionCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *RegionCQ) getCValueRegionId() *df.ConditionValue {
	if q.RegionId == nil {
		q.RegionId = new(df.ConditionValue)
	}
	return q.RegionId
}



func (q *RegionCQ) SetRegionId_Equal(value int64) *RegionCQ {
	q.regRegionId(df.CK_EQ_C, value)
	return q
}
func (q *RegionCQ) SetRegionId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegionId(), "regionId")
}
func (q *RegionCQ) SetRegionId_NotEqual(value int64) *RegionCQ {
	q.regRegionId(df.CK_NE_C, value)
	return q
}

func (q *RegionCQ) SetRegionId_GreaterThan(value int64) *RegionCQ {
	q.regRegionId(df.CK_GT_C, value)
	return q
}

func (q *RegionCQ) SetRegionId_LessThan(value int64) *RegionCQ {
	q.regRegionId(df.CK_LT_C, value)
	return q
}

func (q *RegionCQ) SetRegionId_GreaterEqual(value int64) *RegionCQ {
	q.regRegionId(df.CK_GE_C, value)
	return q
}

func (q *RegionCQ) SetRegionId_LessEqual(value int64) *RegionCQ {
	q.regRegionId(df.CK_LE_C, value)
	return q
}
func (q *RegionCQ) SetRegionId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueRegionId(),"regionId",rangeOfOption)
}	


func (q *RegionCQ) SetRegionId_IsNull() *RegionCQ {
	q.regRegionId(df.CK_ISN_C, 0)
	return q
}
func (q *RegionCQ) SetRegionId_IsNotNull() *RegionCQ {
	q.regRegionId(df.CK_ISNN_C, 0)
	return q
}
func (q *RegionCQ) AddOrderBy_RegionId_Asc() *RegionCQ {
	q.BaseConditionQuery.RegOBA("regionId")
	return q
}
func (q *RegionCQ) AddOrderBy_RegionId_Desc() *RegionCQ {
	q.BaseConditionQuery.RegOBD("regionId")
	return q
}
func (q *RegionCQ) regRegionId(key *df.ConditionKey, value interface{}) {
	if q.RegionId == nil {
		q.RegionId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegionId, "regionId")
}

func (q *RegionCQ) getCValueRegionName() *df.ConditionValue {
	if q.RegionName == nil {
		q.RegionName = new(df.ConditionValue)
	}
	return q.RegionName
}


func (q *RegionCQ) SetRegionName_Equal(value string) *RegionCQ {
	q.regRegionName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *RegionCQ) SetRegionName_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegionName(), "regionName")
}
func (q *RegionCQ) SetRegionName_NotEqual(value string) *RegionCQ {
	q.regRegionName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *RegionCQ) SetRegionName_GreaterThan(value string) *RegionCQ {
	q.regRegionName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *RegionCQ) SetRegionName_LessThan(value string) *RegionCQ {
	q.regRegionName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *RegionCQ) SetRegionName_GreaterEqualThan(value string) *RegionCQ {
	q.regRegionName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *RegionCQ) SetRegionName_LessEqualThan(value string) *RegionCQ {
	q.regRegionName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *RegionCQ) SetRegionName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegionName(), "regionName", option)
}

func (q *RegionCQ) SetRegionName_PrefixSearch(value string) error {
	return q.SetRegionName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *RegionCQ) SetRegionName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegionName(), "regionName", option)
}



func (q *RegionCQ) AddOrderBy_RegionName_Asc() *RegionCQ {
	q.BaseConditionQuery.RegOBA("regionName")
	return q
}
func (q *RegionCQ) AddOrderBy_RegionName_Desc() *RegionCQ {
	q.BaseConditionQuery.RegOBD("regionName")
	return q
}
func (q *RegionCQ) regRegionName(key *df.ConditionKey, value interface{}) {
	if q.RegionName == nil {
		q.RegionName = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegionName, "regionName")
}


func CreateRegionCQ(referrerQuery *df.ConditionQuery, sqlClause *df.SqlClause, aliasName string, nestlevel int8) *RegionCQ {
	cq := new(RegionCQ)
	cq.BaseConditionQuery = new(df.BaseConditionQuery)
	cq.BaseConditionQuery.TableDbName = "Region"
	cq.BaseConditionQuery.ReferrerQuery = referrerQuery
	cq.BaseConditionQuery.SqlClause = sqlClause
	cq.BaseConditionQuery.AliasName = aliasName
	cq.BaseConditionQuery.NestLevel = nestlevel
	cq.BaseConditionQuery.DBMetaProvider = df.DBMetaProvider_I
	cq.BaseConditionQuery.CQ_PROPERTY = "Query"
	var cqi df.ConditionQuery = cq
	cq.BaseConditionQuery.ConditionQuery=&cqi
	return cq
}	