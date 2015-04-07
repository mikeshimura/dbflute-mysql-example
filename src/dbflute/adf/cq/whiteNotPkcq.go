package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type WhiteNotPkCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	NotPkId *df.ConditionValue
	NotPkName *df.ConditionValue
	NotPkInteger *df.ConditionValue
}

func (q *WhiteNotPkCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *WhiteNotPkCQ) getCValueNotPkId() *df.ConditionValue {
	if q.NotPkId == nil {
		q.NotPkId = new(df.ConditionValue)
	}
	return q.NotPkId
}



func (q *WhiteNotPkCQ) SetNotPkId_Equal(value int64) *WhiteNotPkCQ {
	q.regNotPkId(df.CK_EQ_C, value)
	return q
}

func (q *WhiteNotPkCQ) SetNotPkId_NotEqual(value int64) *WhiteNotPkCQ {
	q.regNotPkId(df.CK_NE_C, value)
	return q
}

func (q *WhiteNotPkCQ) SetNotPkId_GreaterThan(value int64) *WhiteNotPkCQ {
	q.regNotPkId(df.CK_GT_C, value)
	return q
}

func (q *WhiteNotPkCQ) SetNotPkId_LessThan(value int64) *WhiteNotPkCQ {
	q.regNotPkId(df.CK_LT_C, value)
	return q
}

func (q *WhiteNotPkCQ) SetNotPkId_GreaterEqual(value int64) *WhiteNotPkCQ {
	q.regNotPkId(df.CK_GE_C, value)
	return q
}

func (q *WhiteNotPkCQ) SetNotPkId_LessEqual(value int64) *WhiteNotPkCQ {
	q.regNotPkId(df.CK_LE_C, value)
	return q
}
func (q *WhiteNotPkCQ) SetNotPkId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueNotPkId(),"notPkId",rangeOfOption)
}	


func (q *WhiteNotPkCQ) AddOrderBy_NotPkId_Asc() *WhiteNotPkCQ {
	q.BaseConditionQuery.RegOBA("notPkId")
	return q
}
func (q *WhiteNotPkCQ) AddOrderBy_NotPkId_Desc() *WhiteNotPkCQ {
	q.BaseConditionQuery.RegOBD("notPkId")
	return q
}
func (q *WhiteNotPkCQ) regNotPkId(key *df.ConditionKey, value interface{}) {
	if q.NotPkId == nil {
		q.NotPkId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.NotPkId, "notPkId")
}

func (q *WhiteNotPkCQ) getCValueNotPkName() *df.ConditionValue {
	if q.NotPkName == nil {
		q.NotPkName = new(df.ConditionValue)
	}
	return q.NotPkName
}


func (q *WhiteNotPkCQ) SetNotPkName_Equal(value string) *WhiteNotPkCQ {
	q.regNotPkName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *WhiteNotPkCQ) SetNotPkName_NotEqual(value string) *WhiteNotPkCQ {
	q.regNotPkName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WhiteNotPkCQ) SetNotPkName_GreaterThan(value string) *WhiteNotPkCQ {
	q.regNotPkName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WhiteNotPkCQ) SetNotPkName_LessThan(value string) *WhiteNotPkCQ {
	q.regNotPkName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WhiteNotPkCQ) SetNotPkName_GreaterEqualThan(value string) *WhiteNotPkCQ {
	q.regNotPkName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *WhiteNotPkCQ) SetNotPkName_LessEqualThan(value string) *WhiteNotPkCQ {
	q.regNotPkName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WhiteNotPkCQ) SetNotPkName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueNotPkName(), "notPkName", option)
}

func (q *WhiteNotPkCQ) SetNotPkName_PrefixSearch(value string) error {
	return q.SetNotPkName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *WhiteNotPkCQ) SetNotPkName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueNotPkName(), "notPkName", option)
}



func (q *WhiteNotPkCQ) SetNotPkName_IsNull() *WhiteNotPkCQ {
	q.regNotPkName(df.CK_ISN_C, 0)
	return q
}
func (q *WhiteNotPkCQ) SetNotPkName_IsNullOrEmpty() *WhiteNotPkCQ {
	q.regNotPkName(df.CK_ISNOE_C, 0)
	return q
}
func (q *WhiteNotPkCQ) SetNotPkName_IsNotNull() *WhiteNotPkCQ {
	q.regNotPkName(df.CK_ISNN_C, 0)
	return q
}
func (q *WhiteNotPkCQ) AddOrderBy_NotPkName_Asc() *WhiteNotPkCQ {
	q.BaseConditionQuery.RegOBA("notPkName")
	return q
}
func (q *WhiteNotPkCQ) AddOrderBy_NotPkName_Desc() *WhiteNotPkCQ {
	q.BaseConditionQuery.RegOBD("notPkName")
	return q
}
func (q *WhiteNotPkCQ) regNotPkName(key *df.ConditionKey, value interface{}) {
	if q.NotPkName == nil {
		q.NotPkName = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.NotPkName, "notPkName")
}

func (q *WhiteNotPkCQ) getCValueNotPkInteger() *df.ConditionValue {
	if q.NotPkInteger == nil {
		q.NotPkInteger = new(df.ConditionValue)
	}
	return q.NotPkInteger
}



func (q *WhiteNotPkCQ) SetNotPkInteger_Equal(value int64) *WhiteNotPkCQ {
	q.regNotPkInteger(df.CK_EQ_C, value)
	return q
}

func (q *WhiteNotPkCQ) SetNotPkInteger_NotEqual(value int64) *WhiteNotPkCQ {
	q.regNotPkInteger(df.CK_NE_C, value)
	return q
}

func (q *WhiteNotPkCQ) SetNotPkInteger_GreaterThan(value int64) *WhiteNotPkCQ {
	q.regNotPkInteger(df.CK_GT_C, value)
	return q
}

func (q *WhiteNotPkCQ) SetNotPkInteger_LessThan(value int64) *WhiteNotPkCQ {
	q.regNotPkInteger(df.CK_LT_C, value)
	return q
}

func (q *WhiteNotPkCQ) SetNotPkInteger_GreaterEqual(value int64) *WhiteNotPkCQ {
	q.regNotPkInteger(df.CK_GE_C, value)
	return q
}

func (q *WhiteNotPkCQ) SetNotPkInteger_LessEqual(value int64) *WhiteNotPkCQ {
	q.regNotPkInteger(df.CK_LE_C, value)
	return q
}
func (q *WhiteNotPkCQ) SetNotPkInteger_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueNotPkInteger(),"notPkInteger",rangeOfOption)
}	


func (q *WhiteNotPkCQ) SetNotPkInteger_IsNull() *WhiteNotPkCQ {
	q.regNotPkInteger(df.CK_ISN_C, 0)
	return q
}
func (q *WhiteNotPkCQ) SetNotPkInteger_IsNotNull() *WhiteNotPkCQ {
	q.regNotPkInteger(df.CK_ISNN_C, 0)
	return q
}
func (q *WhiteNotPkCQ) AddOrderBy_NotPkInteger_Asc() *WhiteNotPkCQ {
	q.BaseConditionQuery.RegOBA("notPkInteger")
	return q
}
func (q *WhiteNotPkCQ) AddOrderBy_NotPkInteger_Desc() *WhiteNotPkCQ {
	q.BaseConditionQuery.RegOBD("notPkInteger")
	return q
}
func (q *WhiteNotPkCQ) regNotPkInteger(key *df.ConditionKey, value interface{}) {
	if q.NotPkInteger == nil {
		q.NotPkInteger = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.NotPkInteger, "notPkInteger")
}

