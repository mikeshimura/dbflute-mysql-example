package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type ServiceRankCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	ServiceRankCode *df.ConditionValue
	ServiceRankName *df.ConditionValue
	ServicePointIncidence *df.ConditionValue
	NewAcceptableFlg *df.ConditionValue
	Description *df.ConditionValue
	DisplayOrder *df.ConditionValue
}

func (q *ServiceRankCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *ServiceRankCQ) getCValueServiceRankCode() *df.ConditionValue {
	if q.ServiceRankCode == nil {
		q.ServiceRankCode = new(df.ConditionValue)
	}
	return q.ServiceRankCode
}


func (q *ServiceRankCQ) SetServiceRankCode_Equal(value string) *ServiceRankCQ {
	q.regServiceRankCode(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *ServiceRankCQ) SetServiceRankCode_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueServiceRankCode(), "serviceRankCode")
}
func (q *ServiceRankCQ) SetServiceRankCode_NotEqual(value string) *ServiceRankCQ {
	q.regServiceRankCode(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ServiceRankCQ) SetServiceRankCode_GreaterThan(value string) *ServiceRankCQ {
	q.regServiceRankCode(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ServiceRankCQ) SetServiceRankCode_LessThan(value string) *ServiceRankCQ {
	q.regServiceRankCode(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ServiceRankCQ) SetServiceRankCode_GreaterEqualThan(value string) *ServiceRankCQ {
	q.regServiceRankCode(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *ServiceRankCQ) SetServiceRankCode_LessEqualThan(value string) *ServiceRankCQ {
	q.regServiceRankCode(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ServiceRankCQ) SetServiceRankCode_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueServiceRankCode(), "serviceRankCode", option)
}

func (q *ServiceRankCQ) SetServiceRankCode_PrefixSearch(value string) error {
	return q.SetServiceRankCode_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *ServiceRankCQ) SetServiceRankCode_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueServiceRankCode(), "serviceRankCode", option)
}



func (q *ServiceRankCQ) SetServiceRankCode_IsNull() *ServiceRankCQ {
	q.regServiceRankCode(df.CK_ISN_C, 0)
	return q
}
func (q *ServiceRankCQ) SetServiceRankCode_IsNotNull() *ServiceRankCQ {
	q.regServiceRankCode(df.CK_ISNN_C, 0)
	return q
}
func (q *ServiceRankCQ) AddOrderBy_ServiceRankCode_Asc() *ServiceRankCQ {
	q.BaseConditionQuery.RegOBA("serviceRankCode")
	return q
}
func (q *ServiceRankCQ) AddOrderBy_ServiceRankCode_Desc() *ServiceRankCQ {
	q.BaseConditionQuery.RegOBD("serviceRankCode")
	return q
}
func (q *ServiceRankCQ) regServiceRankCode(key *df.ConditionKey, value interface{}) {
	if q.ServiceRankCode == nil {
		q.ServiceRankCode = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ServiceRankCode, "serviceRankCode")
}

func (q *ServiceRankCQ) getCValueServiceRankName() *df.ConditionValue {
	if q.ServiceRankName == nil {
		q.ServiceRankName = new(df.ConditionValue)
	}
	return q.ServiceRankName
}


func (q *ServiceRankCQ) SetServiceRankName_Equal(value string) *ServiceRankCQ {
	q.regServiceRankName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *ServiceRankCQ) SetServiceRankName_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueServiceRankName(), "serviceRankName")
}
func (q *ServiceRankCQ) SetServiceRankName_NotEqual(value string) *ServiceRankCQ {
	q.regServiceRankName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ServiceRankCQ) SetServiceRankName_GreaterThan(value string) *ServiceRankCQ {
	q.regServiceRankName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ServiceRankCQ) SetServiceRankName_LessThan(value string) *ServiceRankCQ {
	q.regServiceRankName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ServiceRankCQ) SetServiceRankName_GreaterEqualThan(value string) *ServiceRankCQ {
	q.regServiceRankName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *ServiceRankCQ) SetServiceRankName_LessEqualThan(value string) *ServiceRankCQ {
	q.regServiceRankName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ServiceRankCQ) SetServiceRankName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueServiceRankName(), "serviceRankName", option)
}

func (q *ServiceRankCQ) SetServiceRankName_PrefixSearch(value string) error {
	return q.SetServiceRankName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *ServiceRankCQ) SetServiceRankName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueServiceRankName(), "serviceRankName", option)
}



func (q *ServiceRankCQ) AddOrderBy_ServiceRankName_Asc() *ServiceRankCQ {
	q.BaseConditionQuery.RegOBA("serviceRankName")
	return q
}
func (q *ServiceRankCQ) AddOrderBy_ServiceRankName_Desc() *ServiceRankCQ {
	q.BaseConditionQuery.RegOBD("serviceRankName")
	return q
}
func (q *ServiceRankCQ) regServiceRankName(key *df.ConditionKey, value interface{}) {
	if q.ServiceRankName == nil {
		q.ServiceRankName = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ServiceRankName, "serviceRankName")
}

func (q *ServiceRankCQ) getCValueServicePointIncidence() *df.ConditionValue {
	if q.ServicePointIncidence == nil {
		q.ServicePointIncidence = new(df.ConditionValue)
	}
	return q.ServicePointIncidence
}



func (q *ServiceRankCQ) SetServicePointIncidence_Equal(value df.Numeric) *ServiceRankCQ {
	q.regServicePointIncidence(df.CK_EQ_C, value)
	return q
}
func (q *ServiceRankCQ) SetServicePointIncidence_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueServicePointIncidence(), "servicePointIncidence")
}
func (q *ServiceRankCQ) SetServicePointIncidence_NotEqual(value df.Numeric) *ServiceRankCQ {
	q.regServicePointIncidence(df.CK_NE_C, value)
	return q
}

func (q *ServiceRankCQ) SetServicePointIncidence_GreaterThan(value df.Numeric) *ServiceRankCQ {
	q.regServicePointIncidence(df.CK_GT_C, value)
	return q
}

func (q *ServiceRankCQ) SetServicePointIncidence_LessThan(value df.Numeric) *ServiceRankCQ {
	q.regServicePointIncidence(df.CK_LT_C, value)
	return q
}

func (q *ServiceRankCQ) SetServicePointIncidence_GreaterEqual(value df.Numeric) *ServiceRankCQ {
	q.regServicePointIncidence(df.CK_GE_C, value)
	return q
}

func (q *ServiceRankCQ) SetServicePointIncidence_LessEqual(value df.Numeric) *ServiceRankCQ {
	q.regServicePointIncidence(df.CK_LE_C, value)
	return q
}
func (q *ServiceRankCQ) SetServicePointIncidence_RangeOf(minNumber df.Numeric, maxNumber df.Numeric, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueServicePointIncidence(),"servicePointIncidence",rangeOfOption)
}	


func (q *ServiceRankCQ) AddOrderBy_ServicePointIncidence_Asc() *ServiceRankCQ {
	q.BaseConditionQuery.RegOBA("servicePointIncidence")
	return q
}
func (q *ServiceRankCQ) AddOrderBy_ServicePointIncidence_Desc() *ServiceRankCQ {
	q.BaseConditionQuery.RegOBD("servicePointIncidence")
	return q
}
func (q *ServiceRankCQ) regServicePointIncidence(key *df.ConditionKey, value interface{}) {
	if q.ServicePointIncidence == nil {
		q.ServicePointIncidence = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ServicePointIncidence, "servicePointIncidence")
}

func (q *ServiceRankCQ) getCValueNewAcceptableFlg() *df.ConditionValue {
	if q.NewAcceptableFlg == nil {
		q.NewAcceptableFlg = new(df.ConditionValue)
	}
	return q.NewAcceptableFlg
}



func (q *ServiceRankCQ) SetNewAcceptableFlg_Equal(value int64) *ServiceRankCQ {
	q.regNewAcceptableFlg(df.CK_EQ_C, value)
	return q
}
func (q *ServiceRankCQ) SetNewAcceptableFlg_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueNewAcceptableFlg(), "newAcceptableFlg")
}
func (q *ServiceRankCQ) SetNewAcceptableFlg_NotEqual(value int64) *ServiceRankCQ {
	q.regNewAcceptableFlg(df.CK_NE_C, value)
	return q
}

func (q *ServiceRankCQ) SetNewAcceptableFlg_GreaterThan(value int64) *ServiceRankCQ {
	q.regNewAcceptableFlg(df.CK_GT_C, value)
	return q
}

func (q *ServiceRankCQ) SetNewAcceptableFlg_LessThan(value int64) *ServiceRankCQ {
	q.regNewAcceptableFlg(df.CK_LT_C, value)
	return q
}

func (q *ServiceRankCQ) SetNewAcceptableFlg_GreaterEqual(value int64) *ServiceRankCQ {
	q.regNewAcceptableFlg(df.CK_GE_C, value)
	return q
}

func (q *ServiceRankCQ) SetNewAcceptableFlg_LessEqual(value int64) *ServiceRankCQ {
	q.regNewAcceptableFlg(df.CK_LE_C, value)
	return q
}
func (q *ServiceRankCQ) SetNewAcceptableFlg_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueNewAcceptableFlg(),"newAcceptableFlg",rangeOfOption)
}	


func (q *ServiceRankCQ) AddOrderBy_NewAcceptableFlg_Asc() *ServiceRankCQ {
	q.BaseConditionQuery.RegOBA("newAcceptableFlg")
	return q
}
func (q *ServiceRankCQ) AddOrderBy_NewAcceptableFlg_Desc() *ServiceRankCQ {
	q.BaseConditionQuery.RegOBD("newAcceptableFlg")
	return q
}
func (q *ServiceRankCQ) regNewAcceptableFlg(key *df.ConditionKey, value interface{}) {
	if q.NewAcceptableFlg == nil {
		q.NewAcceptableFlg = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.NewAcceptableFlg, "newAcceptableFlg")
}

func (q *ServiceRankCQ) getCValueDescription() *df.ConditionValue {
	if q.Description == nil {
		q.Description = new(df.ConditionValue)
	}
	return q.Description
}


func (q *ServiceRankCQ) SetDescription_Equal(value string) *ServiceRankCQ {
	q.regDescription(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *ServiceRankCQ) SetDescription_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueDescription(), "description")
}
func (q *ServiceRankCQ) SetDescription_NotEqual(value string) *ServiceRankCQ {
	q.regDescription(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ServiceRankCQ) SetDescription_GreaterThan(value string) *ServiceRankCQ {
	q.regDescription(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ServiceRankCQ) SetDescription_LessThan(value string) *ServiceRankCQ {
	q.regDescription(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ServiceRankCQ) SetDescription_GreaterEqualThan(value string) *ServiceRankCQ {
	q.regDescription(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *ServiceRankCQ) SetDescription_LessEqualThan(value string) *ServiceRankCQ {
	q.regDescription(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ServiceRankCQ) SetDescription_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueDescription(), "description", option)
}

func (q *ServiceRankCQ) SetDescription_PrefixSearch(value string) error {
	return q.SetDescription_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *ServiceRankCQ) SetDescription_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueDescription(), "description", option)
}



func (q *ServiceRankCQ) AddOrderBy_Description_Asc() *ServiceRankCQ {
	q.BaseConditionQuery.RegOBA("description")
	return q
}
func (q *ServiceRankCQ) AddOrderBy_Description_Desc() *ServiceRankCQ {
	q.BaseConditionQuery.RegOBD("description")
	return q
}
func (q *ServiceRankCQ) regDescription(key *df.ConditionKey, value interface{}) {
	if q.Description == nil {
		q.Description = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Description, "description")
}

func (q *ServiceRankCQ) getCValueDisplayOrder() *df.ConditionValue {
	if q.DisplayOrder == nil {
		q.DisplayOrder = new(df.ConditionValue)
	}
	return q.DisplayOrder
}



func (q *ServiceRankCQ) SetDisplayOrder_Equal(value int64) *ServiceRankCQ {
	q.regDisplayOrder(df.CK_EQ_C, value)
	return q
}
func (q *ServiceRankCQ) SetDisplayOrder_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueDisplayOrder(), "displayOrder")
}
func (q *ServiceRankCQ) SetDisplayOrder_NotEqual(value int64) *ServiceRankCQ {
	q.regDisplayOrder(df.CK_NE_C, value)
	return q
}

func (q *ServiceRankCQ) SetDisplayOrder_GreaterThan(value int64) *ServiceRankCQ {
	q.regDisplayOrder(df.CK_GT_C, value)
	return q
}

func (q *ServiceRankCQ) SetDisplayOrder_LessThan(value int64) *ServiceRankCQ {
	q.regDisplayOrder(df.CK_LT_C, value)
	return q
}

func (q *ServiceRankCQ) SetDisplayOrder_GreaterEqual(value int64) *ServiceRankCQ {
	q.regDisplayOrder(df.CK_GE_C, value)
	return q
}

func (q *ServiceRankCQ) SetDisplayOrder_LessEqual(value int64) *ServiceRankCQ {
	q.regDisplayOrder(df.CK_LE_C, value)
	return q
}
func (q *ServiceRankCQ) SetDisplayOrder_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueDisplayOrder(),"displayOrder",rangeOfOption)
}	


func (q *ServiceRankCQ) AddOrderBy_DisplayOrder_Asc() *ServiceRankCQ {
	q.BaseConditionQuery.RegOBA("displayOrder")
	return q
}
func (q *ServiceRankCQ) AddOrderBy_DisplayOrder_Desc() *ServiceRankCQ {
	q.BaseConditionQuery.RegOBD("displayOrder")
	return q
}
func (q *ServiceRankCQ) regDisplayOrder(key *df.ConditionKey, value interface{}) {
	if q.DisplayOrder == nil {
		q.DisplayOrder = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.DisplayOrder, "displayOrder")
}


func CreateServiceRankCQ(referrerQuery *df.ConditionQuery, sqlClause *df.SqlClause, aliasName string, nestlevel int8) *ServiceRankCQ {
	cq := new(ServiceRankCQ)
	cq.BaseConditionQuery = new(df.BaseConditionQuery)
	cq.BaseConditionQuery.TableDbName = "ServiceRank"
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