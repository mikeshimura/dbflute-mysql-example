package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type WhiteSameNameCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	SameNameId *df.ConditionValue
	SameNameName *df.ConditionValue
	SameNameInteger *df.ConditionValue
	NextSchemaProductId *df.ConditionValue
}

func (q *WhiteSameNameCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *WhiteSameNameCQ) getCValueSameNameId() *df.ConditionValue {
	if q.SameNameId == nil {
		q.SameNameId = new(df.ConditionValue)
	}
	return q.SameNameId
}



func (q *WhiteSameNameCQ) SetSameNameId_Equal(value int64) *WhiteSameNameCQ {
	q.regSameNameId(df.CK_EQ_C, value)
	return q
}

func (q *WhiteSameNameCQ) SetSameNameId_NotEqual(value int64) *WhiteSameNameCQ {
	q.regSameNameId(df.CK_NE_C, value)
	return q
}

func (q *WhiteSameNameCQ) SetSameNameId_GreaterThan(value int64) *WhiteSameNameCQ {
	q.regSameNameId(df.CK_GT_C, value)
	return q
}

func (q *WhiteSameNameCQ) SetSameNameId_LessThan(value int64) *WhiteSameNameCQ {
	q.regSameNameId(df.CK_LT_C, value)
	return q
}

func (q *WhiteSameNameCQ) SetSameNameId_GreaterEqual(value int64) *WhiteSameNameCQ {
	q.regSameNameId(df.CK_GE_C, value)
	return q
}

func (q *WhiteSameNameCQ) SetSameNameId_LessEqual(value int64) *WhiteSameNameCQ {
	q.regSameNameId(df.CK_LE_C, value)
	return q
}
func (q *WhiteSameNameCQ) SetSameNameId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueSameNameId(),"sameNameId",rangeOfOption)
}	


func (q *WhiteSameNameCQ) SetSameNameId_IsNull() *WhiteSameNameCQ {
	q.regSameNameId(df.CK_ISN_C, 0)
	return q
}
func (q *WhiteSameNameCQ) SetSameNameId_IsNotNull() *WhiteSameNameCQ {
	q.regSameNameId(df.CK_ISNN_C, 0)
	return q
}
func (q *WhiteSameNameCQ) AddOrderBy_SameNameId_Asc() *WhiteSameNameCQ {
	q.BaseConditionQuery.RegOBA("sameNameId")
	return q
}
func (q *WhiteSameNameCQ) AddOrderBy_SameNameId_Desc() *WhiteSameNameCQ {
	q.BaseConditionQuery.RegOBD("sameNameId")
	return q
}
func (q *WhiteSameNameCQ) regSameNameId(key *df.ConditionKey, value interface{}) {
	if q.SameNameId == nil {
		q.SameNameId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.SameNameId, "sameNameId")
}

func (q *WhiteSameNameCQ) getCValueSameNameName() *df.ConditionValue {
	if q.SameNameName == nil {
		q.SameNameName = new(df.ConditionValue)
	}
	return q.SameNameName
}


func (q *WhiteSameNameCQ) SetSameNameName_Equal(value string) *WhiteSameNameCQ {
	q.regSameNameName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *WhiteSameNameCQ) SetSameNameName_NotEqual(value string) *WhiteSameNameCQ {
	q.regSameNameName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WhiteSameNameCQ) SetSameNameName_GreaterThan(value string) *WhiteSameNameCQ {
	q.regSameNameName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WhiteSameNameCQ) SetSameNameName_LessThan(value string) *WhiteSameNameCQ {
	q.regSameNameName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WhiteSameNameCQ) SetSameNameName_GreaterEqualThan(value string) *WhiteSameNameCQ {
	q.regSameNameName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *WhiteSameNameCQ) SetSameNameName_LessEqualThan(value string) *WhiteSameNameCQ {
	q.regSameNameName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WhiteSameNameCQ) SetSameNameName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueSameNameName(), "sameNameName", option)
}

func (q *WhiteSameNameCQ) SetSameNameName_PrefixSearch(value string) error {
	return q.SetSameNameName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *WhiteSameNameCQ) SetSameNameName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueSameNameName(), "sameNameName", option)
}



func (q *WhiteSameNameCQ) SetSameNameName_IsNull() *WhiteSameNameCQ {
	q.regSameNameName(df.CK_ISN_C, 0)
	return q
}
func (q *WhiteSameNameCQ) SetSameNameName_IsNullOrEmpty() *WhiteSameNameCQ {
	q.regSameNameName(df.CK_ISNOE_C, 0)
	return q
}
func (q *WhiteSameNameCQ) SetSameNameName_IsNotNull() *WhiteSameNameCQ {
	q.regSameNameName(df.CK_ISNN_C, 0)
	return q
}
func (q *WhiteSameNameCQ) AddOrderBy_SameNameName_Asc() *WhiteSameNameCQ {
	q.BaseConditionQuery.RegOBA("sameNameName")
	return q
}
func (q *WhiteSameNameCQ) AddOrderBy_SameNameName_Desc() *WhiteSameNameCQ {
	q.BaseConditionQuery.RegOBD("sameNameName")
	return q
}
func (q *WhiteSameNameCQ) regSameNameName(key *df.ConditionKey, value interface{}) {
	if q.SameNameName == nil {
		q.SameNameName = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.SameNameName, "sameNameName")
}

func (q *WhiteSameNameCQ) getCValueSameNameInteger() *df.ConditionValue {
	if q.SameNameInteger == nil {
		q.SameNameInteger = new(df.ConditionValue)
	}
	return q.SameNameInteger
}



func (q *WhiteSameNameCQ) SetSameNameInteger_Equal(value int64) *WhiteSameNameCQ {
	q.regSameNameInteger(df.CK_EQ_C, value)
	return q
}

func (q *WhiteSameNameCQ) SetSameNameInteger_NotEqual(value int64) *WhiteSameNameCQ {
	q.regSameNameInteger(df.CK_NE_C, value)
	return q
}

func (q *WhiteSameNameCQ) SetSameNameInteger_GreaterThan(value int64) *WhiteSameNameCQ {
	q.regSameNameInteger(df.CK_GT_C, value)
	return q
}

func (q *WhiteSameNameCQ) SetSameNameInteger_LessThan(value int64) *WhiteSameNameCQ {
	q.regSameNameInteger(df.CK_LT_C, value)
	return q
}

func (q *WhiteSameNameCQ) SetSameNameInteger_GreaterEqual(value int64) *WhiteSameNameCQ {
	q.regSameNameInteger(df.CK_GE_C, value)
	return q
}

func (q *WhiteSameNameCQ) SetSameNameInteger_LessEqual(value int64) *WhiteSameNameCQ {
	q.regSameNameInteger(df.CK_LE_C, value)
	return q
}
func (q *WhiteSameNameCQ) SetSameNameInteger_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueSameNameInteger(),"sameNameInteger",rangeOfOption)
}	


func (q *WhiteSameNameCQ) SetSameNameInteger_IsNull() *WhiteSameNameCQ {
	q.regSameNameInteger(df.CK_ISN_C, 0)
	return q
}
func (q *WhiteSameNameCQ) SetSameNameInteger_IsNotNull() *WhiteSameNameCQ {
	q.regSameNameInteger(df.CK_ISNN_C, 0)
	return q
}
func (q *WhiteSameNameCQ) AddOrderBy_SameNameInteger_Asc() *WhiteSameNameCQ {
	q.BaseConditionQuery.RegOBA("sameNameInteger")
	return q
}
func (q *WhiteSameNameCQ) AddOrderBy_SameNameInteger_Desc() *WhiteSameNameCQ {
	q.BaseConditionQuery.RegOBD("sameNameInteger")
	return q
}
func (q *WhiteSameNameCQ) regSameNameInteger(key *df.ConditionKey, value interface{}) {
	if q.SameNameInteger == nil {
		q.SameNameInteger = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.SameNameInteger, "sameNameInteger")
}

func (q *WhiteSameNameCQ) getCValueNextSchemaProductId() *df.ConditionValue {
	if q.NextSchemaProductId == nil {
		q.NextSchemaProductId = new(df.ConditionValue)
	}
	return q.NextSchemaProductId
}



func (q *WhiteSameNameCQ) SetNextSchemaProductId_Equal(value int64) *WhiteSameNameCQ {
	q.regNextSchemaProductId(df.CK_EQ_C, value)
	return q
}

func (q *WhiteSameNameCQ) SetNextSchemaProductId_NotEqual(value int64) *WhiteSameNameCQ {
	q.regNextSchemaProductId(df.CK_NE_C, value)
	return q
}

func (q *WhiteSameNameCQ) SetNextSchemaProductId_GreaterThan(value int64) *WhiteSameNameCQ {
	q.regNextSchemaProductId(df.CK_GT_C, value)
	return q
}

func (q *WhiteSameNameCQ) SetNextSchemaProductId_LessThan(value int64) *WhiteSameNameCQ {
	q.regNextSchemaProductId(df.CK_LT_C, value)
	return q
}

func (q *WhiteSameNameCQ) SetNextSchemaProductId_GreaterEqual(value int64) *WhiteSameNameCQ {
	q.regNextSchemaProductId(df.CK_GE_C, value)
	return q
}

func (q *WhiteSameNameCQ) SetNextSchemaProductId_LessEqual(value int64) *WhiteSameNameCQ {
	q.regNextSchemaProductId(df.CK_LE_C, value)
	return q
}
func (q *WhiteSameNameCQ) SetNextSchemaProductId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueNextSchemaProductId(),"nextSchemaProductId",rangeOfOption)
}	


func (q *WhiteSameNameCQ) SetNextSchemaProductId_IsNull() *WhiteSameNameCQ {
	q.regNextSchemaProductId(df.CK_ISN_C, 0)
	return q
}
func (q *WhiteSameNameCQ) SetNextSchemaProductId_IsNotNull() *WhiteSameNameCQ {
	q.regNextSchemaProductId(df.CK_ISNN_C, 0)
	return q
}
func (q *WhiteSameNameCQ) AddOrderBy_NextSchemaProductId_Asc() *WhiteSameNameCQ {
	q.BaseConditionQuery.RegOBA("nextSchemaProductId")
	return q
}
func (q *WhiteSameNameCQ) AddOrderBy_NextSchemaProductId_Desc() *WhiteSameNameCQ {
	q.BaseConditionQuery.RegOBD("nextSchemaProductId")
	return q
}
func (q *WhiteSameNameCQ) regNextSchemaProductId(key *df.ConditionKey, value interface{}) {
	if q.NextSchemaProductId == nil {
		q.NextSchemaProductId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.NextSchemaProductId, "nextSchemaProductId")
}

