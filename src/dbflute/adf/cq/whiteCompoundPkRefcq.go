package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type WhiteCompoundPkRefCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	MultipleFirstId *df.ConditionValue
	MultipleSecondId *df.ConditionValue
	RefFirstId *df.ConditionValue
	RefSecondId *df.ConditionValue
}

func (q *WhiteCompoundPkRefCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *WhiteCompoundPkRefCQ) getCValueMultipleFirstId() *df.ConditionValue {
	if q.MultipleFirstId == nil {
		q.MultipleFirstId = new(df.ConditionValue)
	}
	return q.MultipleFirstId
}



func (q *WhiteCompoundPkRefCQ) SetMultipleFirstId_Equal(value int64) *WhiteCompoundPkRefCQ {
	q.regMultipleFirstId(df.CK_EQ_C, value)
	return q
}

func (q *WhiteCompoundPkRefCQ) SetMultipleFirstId_NotEqual(value int64) *WhiteCompoundPkRefCQ {
	q.regMultipleFirstId(df.CK_NE_C, value)
	return q
}

func (q *WhiteCompoundPkRefCQ) SetMultipleFirstId_GreaterThan(value int64) *WhiteCompoundPkRefCQ {
	q.regMultipleFirstId(df.CK_GT_C, value)
	return q
}

func (q *WhiteCompoundPkRefCQ) SetMultipleFirstId_LessThan(value int64) *WhiteCompoundPkRefCQ {
	q.regMultipleFirstId(df.CK_LT_C, value)
	return q
}

func (q *WhiteCompoundPkRefCQ) SetMultipleFirstId_GreaterEqual(value int64) *WhiteCompoundPkRefCQ {
	q.regMultipleFirstId(df.CK_GE_C, value)
	return q
}

func (q *WhiteCompoundPkRefCQ) SetMultipleFirstId_LessEqual(value int64) *WhiteCompoundPkRefCQ {
	q.regMultipleFirstId(df.CK_LE_C, value)
	return q
}
func (q *WhiteCompoundPkRefCQ) SetMultipleFirstId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueMultipleFirstId(),"multipleFirstId",rangeOfOption)
}	


func (q *WhiteCompoundPkRefCQ) SetMultipleFirstId_IsNull() *WhiteCompoundPkRefCQ {
	q.regMultipleFirstId(df.CK_ISN_C, 0)
	return q
}
func (q *WhiteCompoundPkRefCQ) SetMultipleFirstId_IsNotNull() *WhiteCompoundPkRefCQ {
	q.regMultipleFirstId(df.CK_ISNN_C, 0)
	return q
}
func (q *WhiteCompoundPkRefCQ) AddOrderBy_MultipleFirstId_Asc() *WhiteCompoundPkRefCQ {
	q.BaseConditionQuery.RegOBA("multipleFirstId")
	return q
}
func (q *WhiteCompoundPkRefCQ) AddOrderBy_MultipleFirstId_Desc() *WhiteCompoundPkRefCQ {
	q.BaseConditionQuery.RegOBD("multipleFirstId")
	return q
}
func (q *WhiteCompoundPkRefCQ) regMultipleFirstId(key *df.ConditionKey, value interface{}) {
	if q.MultipleFirstId == nil {
		q.MultipleFirstId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.MultipleFirstId, "multipleFirstId")
}

func (q *WhiteCompoundPkRefCQ) getCValueMultipleSecondId() *df.ConditionValue {
	if q.MultipleSecondId == nil {
		q.MultipleSecondId = new(df.ConditionValue)
	}
	return q.MultipleSecondId
}



func (q *WhiteCompoundPkRefCQ) SetMultipleSecondId_Equal(value int64) *WhiteCompoundPkRefCQ {
	q.regMultipleSecondId(df.CK_EQ_C, value)
	return q
}

func (q *WhiteCompoundPkRefCQ) SetMultipleSecondId_NotEqual(value int64) *WhiteCompoundPkRefCQ {
	q.regMultipleSecondId(df.CK_NE_C, value)
	return q
}

func (q *WhiteCompoundPkRefCQ) SetMultipleSecondId_GreaterThan(value int64) *WhiteCompoundPkRefCQ {
	q.regMultipleSecondId(df.CK_GT_C, value)
	return q
}

func (q *WhiteCompoundPkRefCQ) SetMultipleSecondId_LessThan(value int64) *WhiteCompoundPkRefCQ {
	q.regMultipleSecondId(df.CK_LT_C, value)
	return q
}

func (q *WhiteCompoundPkRefCQ) SetMultipleSecondId_GreaterEqual(value int64) *WhiteCompoundPkRefCQ {
	q.regMultipleSecondId(df.CK_GE_C, value)
	return q
}

func (q *WhiteCompoundPkRefCQ) SetMultipleSecondId_LessEqual(value int64) *WhiteCompoundPkRefCQ {
	q.regMultipleSecondId(df.CK_LE_C, value)
	return q
}
func (q *WhiteCompoundPkRefCQ) SetMultipleSecondId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueMultipleSecondId(),"multipleSecondId",rangeOfOption)
}	


func (q *WhiteCompoundPkRefCQ) SetMultipleSecondId_IsNull() *WhiteCompoundPkRefCQ {
	q.regMultipleSecondId(df.CK_ISN_C, 0)
	return q
}
func (q *WhiteCompoundPkRefCQ) SetMultipleSecondId_IsNotNull() *WhiteCompoundPkRefCQ {
	q.regMultipleSecondId(df.CK_ISNN_C, 0)
	return q
}
func (q *WhiteCompoundPkRefCQ) AddOrderBy_MultipleSecondId_Asc() *WhiteCompoundPkRefCQ {
	q.BaseConditionQuery.RegOBA("multipleSecondId")
	return q
}
func (q *WhiteCompoundPkRefCQ) AddOrderBy_MultipleSecondId_Desc() *WhiteCompoundPkRefCQ {
	q.BaseConditionQuery.RegOBD("multipleSecondId")
	return q
}
func (q *WhiteCompoundPkRefCQ) regMultipleSecondId(key *df.ConditionKey, value interface{}) {
	if q.MultipleSecondId == nil {
		q.MultipleSecondId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.MultipleSecondId, "multipleSecondId")
}

func (q *WhiteCompoundPkRefCQ) getCValueRefFirstId() *df.ConditionValue {
	if q.RefFirstId == nil {
		q.RefFirstId = new(df.ConditionValue)
	}
	return q.RefFirstId
}



func (q *WhiteCompoundPkRefCQ) SetRefFirstId_Equal(value int64) *WhiteCompoundPkRefCQ {
	q.regRefFirstId(df.CK_EQ_C, value)
	return q
}

func (q *WhiteCompoundPkRefCQ) SetRefFirstId_NotEqual(value int64) *WhiteCompoundPkRefCQ {
	q.regRefFirstId(df.CK_NE_C, value)
	return q
}

func (q *WhiteCompoundPkRefCQ) SetRefFirstId_GreaterThan(value int64) *WhiteCompoundPkRefCQ {
	q.regRefFirstId(df.CK_GT_C, value)
	return q
}

func (q *WhiteCompoundPkRefCQ) SetRefFirstId_LessThan(value int64) *WhiteCompoundPkRefCQ {
	q.regRefFirstId(df.CK_LT_C, value)
	return q
}

func (q *WhiteCompoundPkRefCQ) SetRefFirstId_GreaterEqual(value int64) *WhiteCompoundPkRefCQ {
	q.regRefFirstId(df.CK_GE_C, value)
	return q
}

func (q *WhiteCompoundPkRefCQ) SetRefFirstId_LessEqual(value int64) *WhiteCompoundPkRefCQ {
	q.regRefFirstId(df.CK_LE_C, value)
	return q
}
func (q *WhiteCompoundPkRefCQ) SetRefFirstId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueRefFirstId(),"refFirstId",rangeOfOption)
}	


func (q *WhiteCompoundPkRefCQ) AddOrderBy_RefFirstId_Asc() *WhiteCompoundPkRefCQ {
	q.BaseConditionQuery.RegOBA("refFirstId")
	return q
}
func (q *WhiteCompoundPkRefCQ) AddOrderBy_RefFirstId_Desc() *WhiteCompoundPkRefCQ {
	q.BaseConditionQuery.RegOBD("refFirstId")
	return q
}
func (q *WhiteCompoundPkRefCQ) regRefFirstId(key *df.ConditionKey, value interface{}) {
	if q.RefFirstId == nil {
		q.RefFirstId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RefFirstId, "refFirstId")
}

func (q *WhiteCompoundPkRefCQ) getCValueRefSecondId() *df.ConditionValue {
	if q.RefSecondId == nil {
		q.RefSecondId = new(df.ConditionValue)
	}
	return q.RefSecondId
}



func (q *WhiteCompoundPkRefCQ) SetRefSecondId_Equal(value int64) *WhiteCompoundPkRefCQ {
	q.regRefSecondId(df.CK_EQ_C, value)
	return q
}

func (q *WhiteCompoundPkRefCQ) SetRefSecondId_NotEqual(value int64) *WhiteCompoundPkRefCQ {
	q.regRefSecondId(df.CK_NE_C, value)
	return q
}

func (q *WhiteCompoundPkRefCQ) SetRefSecondId_GreaterThan(value int64) *WhiteCompoundPkRefCQ {
	q.regRefSecondId(df.CK_GT_C, value)
	return q
}

func (q *WhiteCompoundPkRefCQ) SetRefSecondId_LessThan(value int64) *WhiteCompoundPkRefCQ {
	q.regRefSecondId(df.CK_LT_C, value)
	return q
}

func (q *WhiteCompoundPkRefCQ) SetRefSecondId_GreaterEqual(value int64) *WhiteCompoundPkRefCQ {
	q.regRefSecondId(df.CK_GE_C, value)
	return q
}

func (q *WhiteCompoundPkRefCQ) SetRefSecondId_LessEqual(value int64) *WhiteCompoundPkRefCQ {
	q.regRefSecondId(df.CK_LE_C, value)
	return q
}
func (q *WhiteCompoundPkRefCQ) SetRefSecondId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueRefSecondId(),"refSecondId",rangeOfOption)
}	


func (q *WhiteCompoundPkRefCQ) AddOrderBy_RefSecondId_Asc() *WhiteCompoundPkRefCQ {
	q.BaseConditionQuery.RegOBA("refSecondId")
	return q
}
func (q *WhiteCompoundPkRefCQ) AddOrderBy_RefSecondId_Desc() *WhiteCompoundPkRefCQ {
	q.BaseConditionQuery.RegOBD("refSecondId")
	return q
}
func (q *WhiteCompoundPkRefCQ) regRefSecondId(key *df.ConditionKey, value interface{}) {
	if q.RefSecondId == nil {
		q.RefSecondId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RefSecondId, "refSecondId")
}

