package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type MemberStatusCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	MemberStatusCode *df.ConditionValue
	MemberStatusName *df.ConditionValue
	Description *df.ConditionValue
	DisplayOrder *df.ConditionValue
}

func (q *MemberStatusCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *MemberStatusCQ) getCValueMemberStatusCode() *df.ConditionValue {
	if q.MemberStatusCode == nil {
		q.MemberStatusCode = new(df.ConditionValue)
	}
	return q.MemberStatusCode
}


func (q *MemberStatusCQ) SetMemberStatusCode_Equal(value string) *MemberStatusCQ {
	q.regMemberStatusCode(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *MemberStatusCQ) SetMemberStatusCode_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueMemberStatusCode(), "memberStatusCode")
}
func (q *MemberStatusCQ) SetMemberStatusCode_NotEqual(value string) *MemberStatusCQ {
	q.regMemberStatusCode(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberStatusCQ) SetMemberStatusCode_GreaterThan(value string) *MemberStatusCQ {
	q.regMemberStatusCode(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberStatusCQ) SetMemberStatusCode_LessThan(value string) *MemberStatusCQ {
	q.regMemberStatusCode(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberStatusCQ) SetMemberStatusCode_GreaterEqualThan(value string) *MemberStatusCQ {
	q.regMemberStatusCode(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *MemberStatusCQ) SetMemberStatusCode_LessEqualThan(value string) *MemberStatusCQ {
	q.regMemberStatusCode(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberStatusCQ) SetMemberStatusCode_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueMemberStatusCode(), "memberStatusCode", option)
}

func (q *MemberStatusCQ) SetMemberStatusCode_PrefixSearch(value string) error {
	return q.SetMemberStatusCode_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *MemberStatusCQ) SetMemberStatusCode_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueMemberStatusCode(), "memberStatusCode", option)
}



func (q *MemberStatusCQ) SetMemberStatusCode_IsNull() *MemberStatusCQ {
	q.regMemberStatusCode(df.CK_ISN_C, 0)
	return q
}
func (q *MemberStatusCQ) SetMemberStatusCode_IsNotNull() *MemberStatusCQ {
	q.regMemberStatusCode(df.CK_ISNN_C, 0)
	return q
}
func (q *MemberStatusCQ) AddOrderBy_MemberStatusCode_Asc() *MemberStatusCQ {
	q.BaseConditionQuery.RegOBA("memberStatusCode")
	return q
}
func (q *MemberStatusCQ) AddOrderBy_MemberStatusCode_Desc() *MemberStatusCQ {
	q.BaseConditionQuery.RegOBD("memberStatusCode")
	return q
}
func (q *MemberStatusCQ) regMemberStatusCode(key *df.ConditionKey, value interface{}) {
	if q.MemberStatusCode == nil {
		q.MemberStatusCode = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.MemberStatusCode, "memberStatusCode")
}

func (q *MemberStatusCQ) getCValueMemberStatusName() *df.ConditionValue {
	if q.MemberStatusName == nil {
		q.MemberStatusName = new(df.ConditionValue)
	}
	return q.MemberStatusName
}


func (q *MemberStatusCQ) SetMemberStatusName_Equal(value string) *MemberStatusCQ {
	q.regMemberStatusName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *MemberStatusCQ) SetMemberStatusName_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueMemberStatusName(), "memberStatusName")
}
func (q *MemberStatusCQ) SetMemberStatusName_NotEqual(value string) *MemberStatusCQ {
	q.regMemberStatusName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberStatusCQ) SetMemberStatusName_GreaterThan(value string) *MemberStatusCQ {
	q.regMemberStatusName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberStatusCQ) SetMemberStatusName_LessThan(value string) *MemberStatusCQ {
	q.regMemberStatusName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberStatusCQ) SetMemberStatusName_GreaterEqualThan(value string) *MemberStatusCQ {
	q.regMemberStatusName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *MemberStatusCQ) SetMemberStatusName_LessEqualThan(value string) *MemberStatusCQ {
	q.regMemberStatusName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberStatusCQ) SetMemberStatusName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueMemberStatusName(), "memberStatusName", option)
}

func (q *MemberStatusCQ) SetMemberStatusName_PrefixSearch(value string) error {
	return q.SetMemberStatusName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *MemberStatusCQ) SetMemberStatusName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueMemberStatusName(), "memberStatusName", option)
}



func (q *MemberStatusCQ) AddOrderBy_MemberStatusName_Asc() *MemberStatusCQ {
	q.BaseConditionQuery.RegOBA("memberStatusName")
	return q
}
func (q *MemberStatusCQ) AddOrderBy_MemberStatusName_Desc() *MemberStatusCQ {
	q.BaseConditionQuery.RegOBD("memberStatusName")
	return q
}
func (q *MemberStatusCQ) regMemberStatusName(key *df.ConditionKey, value interface{}) {
	if q.MemberStatusName == nil {
		q.MemberStatusName = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.MemberStatusName, "memberStatusName")
}

func (q *MemberStatusCQ) getCValueDescription() *df.ConditionValue {
	if q.Description == nil {
		q.Description = new(df.ConditionValue)
	}
	return q.Description
}


func (q *MemberStatusCQ) SetDescription_Equal(value string) *MemberStatusCQ {
	q.regDescription(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *MemberStatusCQ) SetDescription_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueDescription(), "description")
}
func (q *MemberStatusCQ) SetDescription_NotEqual(value string) *MemberStatusCQ {
	q.regDescription(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberStatusCQ) SetDescription_GreaterThan(value string) *MemberStatusCQ {
	q.regDescription(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberStatusCQ) SetDescription_LessThan(value string) *MemberStatusCQ {
	q.regDescription(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberStatusCQ) SetDescription_GreaterEqualThan(value string) *MemberStatusCQ {
	q.regDescription(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *MemberStatusCQ) SetDescription_LessEqualThan(value string) *MemberStatusCQ {
	q.regDescription(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberStatusCQ) SetDescription_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueDescription(), "description", option)
}

func (q *MemberStatusCQ) SetDescription_PrefixSearch(value string) error {
	return q.SetDescription_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *MemberStatusCQ) SetDescription_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueDescription(), "description", option)
}



func (q *MemberStatusCQ) AddOrderBy_Description_Asc() *MemberStatusCQ {
	q.BaseConditionQuery.RegOBA("description")
	return q
}
func (q *MemberStatusCQ) AddOrderBy_Description_Desc() *MemberStatusCQ {
	q.BaseConditionQuery.RegOBD("description")
	return q
}
func (q *MemberStatusCQ) regDescription(key *df.ConditionKey, value interface{}) {
	if q.Description == nil {
		q.Description = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Description, "description")
}

func (q *MemberStatusCQ) getCValueDisplayOrder() *df.ConditionValue {
	if q.DisplayOrder == nil {
		q.DisplayOrder = new(df.ConditionValue)
	}
	return q.DisplayOrder
}



func (q *MemberStatusCQ) SetDisplayOrder_Equal(value int64) *MemberStatusCQ {
	q.regDisplayOrder(df.CK_EQ_C, value)
	return q
}
func (q *MemberStatusCQ) SetDisplayOrder_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueDisplayOrder(), "displayOrder")
}
func (q *MemberStatusCQ) SetDisplayOrder_NotEqual(value int64) *MemberStatusCQ {
	q.regDisplayOrder(df.CK_NE_C, value)
	return q
}

func (q *MemberStatusCQ) SetDisplayOrder_GreaterThan(value int64) *MemberStatusCQ {
	q.regDisplayOrder(df.CK_GT_C, value)
	return q
}

func (q *MemberStatusCQ) SetDisplayOrder_LessThan(value int64) *MemberStatusCQ {
	q.regDisplayOrder(df.CK_LT_C, value)
	return q
}

func (q *MemberStatusCQ) SetDisplayOrder_GreaterEqual(value int64) *MemberStatusCQ {
	q.regDisplayOrder(df.CK_GE_C, value)
	return q
}

func (q *MemberStatusCQ) SetDisplayOrder_LessEqual(value int64) *MemberStatusCQ {
	q.regDisplayOrder(df.CK_LE_C, value)
	return q
}
func (q *MemberStatusCQ) SetDisplayOrder_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueDisplayOrder(),"displayOrder",rangeOfOption)
}	


func (q *MemberStatusCQ) AddOrderBy_DisplayOrder_Asc() *MemberStatusCQ {
	q.BaseConditionQuery.RegOBA("displayOrder")
	return q
}
func (q *MemberStatusCQ) AddOrderBy_DisplayOrder_Desc() *MemberStatusCQ {
	q.BaseConditionQuery.RegOBD("displayOrder")
	return q
}
func (q *MemberStatusCQ) regDisplayOrder(key *df.ConditionKey, value interface{}) {
	if q.DisplayOrder == nil {
		q.DisplayOrder = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.DisplayOrder, "displayOrder")
}


func CreateMemberStatusCQ(referrerQuery *df.ConditionQuery, sqlClause *df.SqlClause, aliasName string, nestlevel int8) *MemberStatusCQ {
	cq := new(MemberStatusCQ)
	cq.BaseConditionQuery = new(df.BaseConditionQuery)
	cq.BaseConditionQuery.TableDbName = "MemberStatus"
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