package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type WhiteCompoundPkCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	PkFirstId *df.ConditionValue
	PkSecondId *df.ConditionValue
	PkName *df.ConditionValue
}

func (q *WhiteCompoundPkCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *WhiteCompoundPkCQ) getCValuePkFirstId() *df.ConditionValue {
	if q.PkFirstId == nil {
		q.PkFirstId = new(df.ConditionValue)
	}
	return q.PkFirstId
}



func (q *WhiteCompoundPkCQ) SetPkFirstId_Equal(value int64) *WhiteCompoundPkCQ {
	q.regPkFirstId(df.CK_EQ_C, value)
	return q
}

func (q *WhiteCompoundPkCQ) SetPkFirstId_NotEqual(value int64) *WhiteCompoundPkCQ {
	q.regPkFirstId(df.CK_NE_C, value)
	return q
}

func (q *WhiteCompoundPkCQ) SetPkFirstId_GreaterThan(value int64) *WhiteCompoundPkCQ {
	q.regPkFirstId(df.CK_GT_C, value)
	return q
}

func (q *WhiteCompoundPkCQ) SetPkFirstId_LessThan(value int64) *WhiteCompoundPkCQ {
	q.regPkFirstId(df.CK_LT_C, value)
	return q
}

func (q *WhiteCompoundPkCQ) SetPkFirstId_GreaterEqual(value int64) *WhiteCompoundPkCQ {
	q.regPkFirstId(df.CK_GE_C, value)
	return q
}

func (q *WhiteCompoundPkCQ) SetPkFirstId_LessEqual(value int64) *WhiteCompoundPkCQ {
	q.regPkFirstId(df.CK_LE_C, value)
	return q
}
func (q *WhiteCompoundPkCQ) SetPkFirstId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValuePkFirstId(),"pkFirstId",rangeOfOption)
}	


func (q *WhiteCompoundPkCQ) SetPkFirstId_IsNull() *WhiteCompoundPkCQ {
	q.regPkFirstId(df.CK_ISN_C, 0)
	return q
}
func (q *WhiteCompoundPkCQ) SetPkFirstId_IsNotNull() *WhiteCompoundPkCQ {
	q.regPkFirstId(df.CK_ISNN_C, 0)
	return q
}
func (q *WhiteCompoundPkCQ) AddOrderBy_PkFirstId_Asc() *WhiteCompoundPkCQ {
	q.BaseConditionQuery.RegOBA("pkFirstId")
	return q
}
func (q *WhiteCompoundPkCQ) AddOrderBy_PkFirstId_Desc() *WhiteCompoundPkCQ {
	q.BaseConditionQuery.RegOBD("pkFirstId")
	return q
}
func (q *WhiteCompoundPkCQ) regPkFirstId(key *df.ConditionKey, value interface{}) {
	if q.PkFirstId == nil {
		q.PkFirstId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PkFirstId, "pkFirstId")
}

func (q *WhiteCompoundPkCQ) getCValuePkSecondId() *df.ConditionValue {
	if q.PkSecondId == nil {
		q.PkSecondId = new(df.ConditionValue)
	}
	return q.PkSecondId
}



func (q *WhiteCompoundPkCQ) SetPkSecondId_Equal(value int64) *WhiteCompoundPkCQ {
	q.regPkSecondId(df.CK_EQ_C, value)
	return q
}

func (q *WhiteCompoundPkCQ) SetPkSecondId_NotEqual(value int64) *WhiteCompoundPkCQ {
	q.regPkSecondId(df.CK_NE_C, value)
	return q
}

func (q *WhiteCompoundPkCQ) SetPkSecondId_GreaterThan(value int64) *WhiteCompoundPkCQ {
	q.regPkSecondId(df.CK_GT_C, value)
	return q
}

func (q *WhiteCompoundPkCQ) SetPkSecondId_LessThan(value int64) *WhiteCompoundPkCQ {
	q.regPkSecondId(df.CK_LT_C, value)
	return q
}

func (q *WhiteCompoundPkCQ) SetPkSecondId_GreaterEqual(value int64) *WhiteCompoundPkCQ {
	q.regPkSecondId(df.CK_GE_C, value)
	return q
}

func (q *WhiteCompoundPkCQ) SetPkSecondId_LessEqual(value int64) *WhiteCompoundPkCQ {
	q.regPkSecondId(df.CK_LE_C, value)
	return q
}
func (q *WhiteCompoundPkCQ) SetPkSecondId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValuePkSecondId(),"pkSecondId",rangeOfOption)
}	


func (q *WhiteCompoundPkCQ) SetPkSecondId_IsNull() *WhiteCompoundPkCQ {
	q.regPkSecondId(df.CK_ISN_C, 0)
	return q
}
func (q *WhiteCompoundPkCQ) SetPkSecondId_IsNotNull() *WhiteCompoundPkCQ {
	q.regPkSecondId(df.CK_ISNN_C, 0)
	return q
}
func (q *WhiteCompoundPkCQ) AddOrderBy_PkSecondId_Asc() *WhiteCompoundPkCQ {
	q.BaseConditionQuery.RegOBA("pkSecondId")
	return q
}
func (q *WhiteCompoundPkCQ) AddOrderBy_PkSecondId_Desc() *WhiteCompoundPkCQ {
	q.BaseConditionQuery.RegOBD("pkSecondId")
	return q
}
func (q *WhiteCompoundPkCQ) regPkSecondId(key *df.ConditionKey, value interface{}) {
	if q.PkSecondId == nil {
		q.PkSecondId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PkSecondId, "pkSecondId")
}

func (q *WhiteCompoundPkCQ) getCValuePkName() *df.ConditionValue {
	if q.PkName == nil {
		q.PkName = new(df.ConditionValue)
	}
	return q.PkName
}


func (q *WhiteCompoundPkCQ) SetPkName_Equal(value string) *WhiteCompoundPkCQ {
	q.regPkName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *WhiteCompoundPkCQ) SetPkName_NotEqual(value string) *WhiteCompoundPkCQ {
	q.regPkName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WhiteCompoundPkCQ) SetPkName_GreaterThan(value string) *WhiteCompoundPkCQ {
	q.regPkName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WhiteCompoundPkCQ) SetPkName_LessThan(value string) *WhiteCompoundPkCQ {
	q.regPkName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WhiteCompoundPkCQ) SetPkName_GreaterEqualThan(value string) *WhiteCompoundPkCQ {
	q.regPkName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *WhiteCompoundPkCQ) SetPkName_LessEqualThan(value string) *WhiteCompoundPkCQ {
	q.regPkName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WhiteCompoundPkCQ) SetPkName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValuePkName(), "pkName", option)
}

func (q *WhiteCompoundPkCQ) SetPkName_PrefixSearch(value string) error {
	return q.SetPkName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *WhiteCompoundPkCQ) SetPkName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValuePkName(), "pkName", option)
}



func (q *WhiteCompoundPkCQ) AddOrderBy_PkName_Asc() *WhiteCompoundPkCQ {
	q.BaseConditionQuery.RegOBA("pkName")
	return q
}
func (q *WhiteCompoundPkCQ) AddOrderBy_PkName_Desc() *WhiteCompoundPkCQ {
	q.BaseConditionQuery.RegOBD("pkName")
	return q
}
func (q *WhiteCompoundPkCQ) regPkName(key *df.ConditionKey, value interface{}) {
	if q.PkName == nil {
		q.PkName = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PkName, "pkName")
}

