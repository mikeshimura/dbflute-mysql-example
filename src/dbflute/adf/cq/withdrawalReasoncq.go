package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type WithdrawalReasonCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	WithdrawalReasonCode *df.ConditionValue
	WithdrawalReasonText *df.ConditionValue
	DisplayOrder *df.ConditionValue
}

func (q *WithdrawalReasonCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *WithdrawalReasonCQ) getCValueWithdrawalReasonCode() *df.ConditionValue {
	if q.WithdrawalReasonCode == nil {
		q.WithdrawalReasonCode = new(df.ConditionValue)
	}
	return q.WithdrawalReasonCode
}


func (q *WithdrawalReasonCQ) SetWithdrawalReasonCode_Equal(value string) *WithdrawalReasonCQ {
	q.regWithdrawalReasonCode(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *WithdrawalReasonCQ) SetWithdrawalReasonCode_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueWithdrawalReasonCode(), "withdrawalReasonCode")
}
func (q *WithdrawalReasonCQ) SetWithdrawalReasonCode_NotEqual(value string) *WithdrawalReasonCQ {
	q.regWithdrawalReasonCode(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WithdrawalReasonCQ) SetWithdrawalReasonCode_GreaterThan(value string) *WithdrawalReasonCQ {
	q.regWithdrawalReasonCode(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WithdrawalReasonCQ) SetWithdrawalReasonCode_LessThan(value string) *WithdrawalReasonCQ {
	q.regWithdrawalReasonCode(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WithdrawalReasonCQ) SetWithdrawalReasonCode_GreaterEqualThan(value string) *WithdrawalReasonCQ {
	q.regWithdrawalReasonCode(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *WithdrawalReasonCQ) SetWithdrawalReasonCode_LessEqualThan(value string) *WithdrawalReasonCQ {
	q.regWithdrawalReasonCode(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WithdrawalReasonCQ) SetWithdrawalReasonCode_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueWithdrawalReasonCode(), "withdrawalReasonCode", option)
}

func (q *WithdrawalReasonCQ) SetWithdrawalReasonCode_PrefixSearch(value string) error {
	return q.SetWithdrawalReasonCode_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *WithdrawalReasonCQ) SetWithdrawalReasonCode_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueWithdrawalReasonCode(), "withdrawalReasonCode", option)
}



func (q *WithdrawalReasonCQ) SetWithdrawalReasonCode_IsNull() *WithdrawalReasonCQ {
	q.regWithdrawalReasonCode(df.CK_ISN_C, 0)
	return q
}
func (q *WithdrawalReasonCQ) SetWithdrawalReasonCode_IsNotNull() *WithdrawalReasonCQ {
	q.regWithdrawalReasonCode(df.CK_ISNN_C, 0)
	return q
}
func (q *WithdrawalReasonCQ) AddOrderBy_WithdrawalReasonCode_Asc() *WithdrawalReasonCQ {
	q.BaseConditionQuery.RegOBA("withdrawalReasonCode")
	return q
}
func (q *WithdrawalReasonCQ) AddOrderBy_WithdrawalReasonCode_Desc() *WithdrawalReasonCQ {
	q.BaseConditionQuery.RegOBD("withdrawalReasonCode")
	return q
}
func (q *WithdrawalReasonCQ) regWithdrawalReasonCode(key *df.ConditionKey, value interface{}) {
	if q.WithdrawalReasonCode == nil {
		q.WithdrawalReasonCode = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.WithdrawalReasonCode, "withdrawalReasonCode")
}

func (q *WithdrawalReasonCQ) getCValueWithdrawalReasonText() *df.ConditionValue {
	if q.WithdrawalReasonText == nil {
		q.WithdrawalReasonText = new(df.ConditionValue)
	}
	return q.WithdrawalReasonText
}


func (q *WithdrawalReasonCQ) SetWithdrawalReasonText_Equal(value string) *WithdrawalReasonCQ {
	q.regWithdrawalReasonText(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *WithdrawalReasonCQ) SetWithdrawalReasonText_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueWithdrawalReasonText(), "withdrawalReasonText")
}
func (q *WithdrawalReasonCQ) SetWithdrawalReasonText_NotEqual(value string) *WithdrawalReasonCQ {
	q.regWithdrawalReasonText(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WithdrawalReasonCQ) SetWithdrawalReasonText_GreaterThan(value string) *WithdrawalReasonCQ {
	q.regWithdrawalReasonText(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WithdrawalReasonCQ) SetWithdrawalReasonText_LessThan(value string) *WithdrawalReasonCQ {
	q.regWithdrawalReasonText(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WithdrawalReasonCQ) SetWithdrawalReasonText_GreaterEqualThan(value string) *WithdrawalReasonCQ {
	q.regWithdrawalReasonText(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *WithdrawalReasonCQ) SetWithdrawalReasonText_LessEqualThan(value string) *WithdrawalReasonCQ {
	q.regWithdrawalReasonText(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WithdrawalReasonCQ) SetWithdrawalReasonText_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueWithdrawalReasonText(), "withdrawalReasonText", option)
}

func (q *WithdrawalReasonCQ) SetWithdrawalReasonText_PrefixSearch(value string) error {
	return q.SetWithdrawalReasonText_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *WithdrawalReasonCQ) SetWithdrawalReasonText_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueWithdrawalReasonText(), "withdrawalReasonText", option)
}



func (q *WithdrawalReasonCQ) AddOrderBy_WithdrawalReasonText_Asc() *WithdrawalReasonCQ {
	q.BaseConditionQuery.RegOBA("withdrawalReasonText")
	return q
}
func (q *WithdrawalReasonCQ) AddOrderBy_WithdrawalReasonText_Desc() *WithdrawalReasonCQ {
	q.BaseConditionQuery.RegOBD("withdrawalReasonText")
	return q
}
func (q *WithdrawalReasonCQ) regWithdrawalReasonText(key *df.ConditionKey, value interface{}) {
	if q.WithdrawalReasonText == nil {
		q.WithdrawalReasonText = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.WithdrawalReasonText, "withdrawalReasonText")
}

func (q *WithdrawalReasonCQ) getCValueDisplayOrder() *df.ConditionValue {
	if q.DisplayOrder == nil {
		q.DisplayOrder = new(df.ConditionValue)
	}
	return q.DisplayOrder
}



func (q *WithdrawalReasonCQ) SetDisplayOrder_Equal(value int64) *WithdrawalReasonCQ {
	q.regDisplayOrder(df.CK_EQ_C, value)
	return q
}
func (q *WithdrawalReasonCQ) SetDisplayOrder_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueDisplayOrder(), "displayOrder")
}
func (q *WithdrawalReasonCQ) SetDisplayOrder_NotEqual(value int64) *WithdrawalReasonCQ {
	q.regDisplayOrder(df.CK_NE_C, value)
	return q
}

func (q *WithdrawalReasonCQ) SetDisplayOrder_GreaterThan(value int64) *WithdrawalReasonCQ {
	q.regDisplayOrder(df.CK_GT_C, value)
	return q
}

func (q *WithdrawalReasonCQ) SetDisplayOrder_LessThan(value int64) *WithdrawalReasonCQ {
	q.regDisplayOrder(df.CK_LT_C, value)
	return q
}

func (q *WithdrawalReasonCQ) SetDisplayOrder_GreaterEqual(value int64) *WithdrawalReasonCQ {
	q.regDisplayOrder(df.CK_GE_C, value)
	return q
}

func (q *WithdrawalReasonCQ) SetDisplayOrder_LessEqual(value int64) *WithdrawalReasonCQ {
	q.regDisplayOrder(df.CK_LE_C, value)
	return q
}
func (q *WithdrawalReasonCQ) SetDisplayOrder_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueDisplayOrder(),"displayOrder",rangeOfOption)
}	


func (q *WithdrawalReasonCQ) AddOrderBy_DisplayOrder_Asc() *WithdrawalReasonCQ {
	q.BaseConditionQuery.RegOBA("displayOrder")
	return q
}
func (q *WithdrawalReasonCQ) AddOrderBy_DisplayOrder_Desc() *WithdrawalReasonCQ {
	q.BaseConditionQuery.RegOBD("displayOrder")
	return q
}
func (q *WithdrawalReasonCQ) regDisplayOrder(key *df.ConditionKey, value interface{}) {
	if q.DisplayOrder == nil {
		q.DisplayOrder = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.DisplayOrder, "displayOrder")
}


func CreateWithdrawalReasonCQ(referrerQuery *df.ConditionQuery, sqlClause *df.SqlClause, aliasName string, nestlevel int8) *WithdrawalReasonCQ {
	cq := new(WithdrawalReasonCQ)
	cq.BaseConditionQuery = new(df.BaseConditionQuery)
	cq.BaseConditionQuery.TableDbName = "WithdrawalReason"
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