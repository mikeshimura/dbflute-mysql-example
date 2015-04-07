package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorLargeDataRefCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	LargeDataRefId *df.ConditionValue
	LargeDataId *df.ConditionValue
	DateIndex *df.ConditionValue
	DateNoIndex *df.ConditionValue
	TimestampIndex *df.ConditionValue
	TimestampNoIndex *df.ConditionValue
	NullableDecimalIndex *df.ConditionValue
	NullableDecimalNoIndex *df.ConditionValue
	SelfParentId *df.ConditionValue
}

func (q *VendorLargeDataRefCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *VendorLargeDataRefCQ) getCValueLargeDataRefId() *df.ConditionValue {
	if q.LargeDataRefId == nil {
		q.LargeDataRefId = new(df.ConditionValue)
	}
	return q.LargeDataRefId
}



func (q *VendorLargeDataRefCQ) SetLargeDataRefId_Equal(value int64) *VendorLargeDataRefCQ {
	q.regLargeDataRefId(df.CK_EQ_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetLargeDataRefId_NotEqual(value int64) *VendorLargeDataRefCQ {
	q.regLargeDataRefId(df.CK_NE_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetLargeDataRefId_GreaterThan(value int64) *VendorLargeDataRefCQ {
	q.regLargeDataRefId(df.CK_GT_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetLargeDataRefId_LessThan(value int64) *VendorLargeDataRefCQ {
	q.regLargeDataRefId(df.CK_LT_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetLargeDataRefId_GreaterEqual(value int64) *VendorLargeDataRefCQ {
	q.regLargeDataRefId(df.CK_GE_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetLargeDataRefId_LessEqual(value int64) *VendorLargeDataRefCQ {
	q.regLargeDataRefId(df.CK_LE_C, value)
	return q
}
func (q *VendorLargeDataRefCQ) SetLargeDataRefId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueLargeDataRefId(),"largeDataRefId",rangeOfOption)
}	


func (q *VendorLargeDataRefCQ) SetLargeDataRefId_IsNull() *VendorLargeDataRefCQ {
	q.regLargeDataRefId(df.CK_ISN_C, 0)
	return q
}
func (q *VendorLargeDataRefCQ) SetLargeDataRefId_IsNotNull() *VendorLargeDataRefCQ {
	q.regLargeDataRefId(df.CK_ISNN_C, 0)
	return q
}
func (q *VendorLargeDataRefCQ) AddOrderBy_LargeDataRefId_Asc() *VendorLargeDataRefCQ {
	q.BaseConditionQuery.RegOBA("largeDataRefId")
	return q
}
func (q *VendorLargeDataRefCQ) AddOrderBy_LargeDataRefId_Desc() *VendorLargeDataRefCQ {
	q.BaseConditionQuery.RegOBD("largeDataRefId")
	return q
}
func (q *VendorLargeDataRefCQ) regLargeDataRefId(key *df.ConditionKey, value interface{}) {
	if q.LargeDataRefId == nil {
		q.LargeDataRefId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.LargeDataRefId, "largeDataRefId")
}

func (q *VendorLargeDataRefCQ) getCValueLargeDataId() *df.ConditionValue {
	if q.LargeDataId == nil {
		q.LargeDataId = new(df.ConditionValue)
	}
	return q.LargeDataId
}



func (q *VendorLargeDataRefCQ) SetLargeDataId_Equal(value int64) *VendorLargeDataRefCQ {
	q.regLargeDataId(df.CK_EQ_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetLargeDataId_NotEqual(value int64) *VendorLargeDataRefCQ {
	q.regLargeDataId(df.CK_NE_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetLargeDataId_GreaterThan(value int64) *VendorLargeDataRefCQ {
	q.regLargeDataId(df.CK_GT_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetLargeDataId_LessThan(value int64) *VendorLargeDataRefCQ {
	q.regLargeDataId(df.CK_LT_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetLargeDataId_GreaterEqual(value int64) *VendorLargeDataRefCQ {
	q.regLargeDataId(df.CK_GE_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetLargeDataId_LessEqual(value int64) *VendorLargeDataRefCQ {
	q.regLargeDataId(df.CK_LE_C, value)
	return q
}
func (q *VendorLargeDataRefCQ) SetLargeDataId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueLargeDataId(),"largeDataId",rangeOfOption)
}	


func (q *VendorLargeDataRefCQ) AddOrderBy_LargeDataId_Asc() *VendorLargeDataRefCQ {
	q.BaseConditionQuery.RegOBA("largeDataId")
	return q
}
func (q *VendorLargeDataRefCQ) AddOrderBy_LargeDataId_Desc() *VendorLargeDataRefCQ {
	q.BaseConditionQuery.RegOBD("largeDataId")
	return q
}
func (q *VendorLargeDataRefCQ) regLargeDataId(key *df.ConditionKey, value interface{}) {
	if q.LargeDataId == nil {
		q.LargeDataId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.LargeDataId, "largeDataId")
}

func (q *VendorLargeDataRefCQ) getCValueDateIndex() *df.ConditionValue {
	if q.DateIndex == nil {
		q.DateIndex = new(df.ConditionValue)
	}
	return q.DateIndex
}




func (q *VendorLargeDataRefCQ) SetDateIndex_Equal(value df.Date) *VendorLargeDataRefCQ {
	q.regDateIndex(df.CK_EQ_C, value)
	return q
}


func (q *VendorLargeDataRefCQ) SetDateIndex_GreaterThan(value df.Date) *VendorLargeDataRefCQ {
	q.regDateIndex(df.CK_GT_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetDateIndex_LessThan(value df.Date) *VendorLargeDataRefCQ {
	q.regDateIndex(df.CK_LT_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetDateIndex_GreaterEqual(value df.Date) *VendorLargeDataRefCQ {
	q.regDateIndex(df.CK_GE_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetDateIndex_LessEqual(value df.Date) *VendorLargeDataRefCQ {
	q.regDateIndex(df.CK_LE_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) AddOrderBy_DateIndex_Asc() *VendorLargeDataRefCQ {
	q.BaseConditionQuery.RegOBA("dateIndex")
	return q
}
func (q *VendorLargeDataRefCQ) AddOrderBy_DateIndex_Desc() *VendorLargeDataRefCQ {
	q.BaseConditionQuery.RegOBD("dateIndex")
	return q
}
func (q *VendorLargeDataRefCQ) regDateIndex(key *df.ConditionKey, value interface{}) {
	if q.DateIndex == nil {
		q.DateIndex = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.DateIndex, "dateIndex")
}

func (q *VendorLargeDataRefCQ) getCValueDateNoIndex() *df.ConditionValue {
	if q.DateNoIndex == nil {
		q.DateNoIndex = new(df.ConditionValue)
	}
	return q.DateNoIndex
}




func (q *VendorLargeDataRefCQ) SetDateNoIndex_Equal(value df.Date) *VendorLargeDataRefCQ {
	q.regDateNoIndex(df.CK_EQ_C, value)
	return q
}


func (q *VendorLargeDataRefCQ) SetDateNoIndex_GreaterThan(value df.Date) *VendorLargeDataRefCQ {
	q.regDateNoIndex(df.CK_GT_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetDateNoIndex_LessThan(value df.Date) *VendorLargeDataRefCQ {
	q.regDateNoIndex(df.CK_LT_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetDateNoIndex_GreaterEqual(value df.Date) *VendorLargeDataRefCQ {
	q.regDateNoIndex(df.CK_GE_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetDateNoIndex_LessEqual(value df.Date) *VendorLargeDataRefCQ {
	q.regDateNoIndex(df.CK_LE_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) AddOrderBy_DateNoIndex_Asc() *VendorLargeDataRefCQ {
	q.BaseConditionQuery.RegOBA("dateNoIndex")
	return q
}
func (q *VendorLargeDataRefCQ) AddOrderBy_DateNoIndex_Desc() *VendorLargeDataRefCQ {
	q.BaseConditionQuery.RegOBD("dateNoIndex")
	return q
}
func (q *VendorLargeDataRefCQ) regDateNoIndex(key *df.ConditionKey, value interface{}) {
	if q.DateNoIndex == nil {
		q.DateNoIndex = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.DateNoIndex, "dateNoIndex")
}

func (q *VendorLargeDataRefCQ) getCValueTimestampIndex() *df.ConditionValue {
	if q.TimestampIndex == nil {
		q.TimestampIndex = new(df.ConditionValue)
	}
	return q.TimestampIndex
}




func (q *VendorLargeDataRefCQ) SetTimestampIndex_Equal(value df.Timestamp) *VendorLargeDataRefCQ {
	q.regTimestampIndex(df.CK_EQ_C, value)
	return q
}


func (q *VendorLargeDataRefCQ) SetTimestampIndex_GreaterThan(value df.Timestamp) *VendorLargeDataRefCQ {
	q.regTimestampIndex(df.CK_GT_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetTimestampIndex_LessThan(value df.Timestamp) *VendorLargeDataRefCQ {
	q.regTimestampIndex(df.CK_LT_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetTimestampIndex_GreaterEqual(value df.Timestamp) *VendorLargeDataRefCQ {
	q.regTimestampIndex(df.CK_GE_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetTimestampIndex_LessEqual(value df.Timestamp) *VendorLargeDataRefCQ {
	q.regTimestampIndex(df.CK_LE_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) AddOrderBy_TimestampIndex_Asc() *VendorLargeDataRefCQ {
	q.BaseConditionQuery.RegOBA("timestampIndex")
	return q
}
func (q *VendorLargeDataRefCQ) AddOrderBy_TimestampIndex_Desc() *VendorLargeDataRefCQ {
	q.BaseConditionQuery.RegOBD("timestampIndex")
	return q
}
func (q *VendorLargeDataRefCQ) regTimestampIndex(key *df.ConditionKey, value interface{}) {
	if q.TimestampIndex == nil {
		q.TimestampIndex = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.TimestampIndex, "timestampIndex")
}

func (q *VendorLargeDataRefCQ) getCValueTimestampNoIndex() *df.ConditionValue {
	if q.TimestampNoIndex == nil {
		q.TimestampNoIndex = new(df.ConditionValue)
	}
	return q.TimestampNoIndex
}




func (q *VendorLargeDataRefCQ) SetTimestampNoIndex_Equal(value df.Timestamp) *VendorLargeDataRefCQ {
	q.regTimestampNoIndex(df.CK_EQ_C, value)
	return q
}


func (q *VendorLargeDataRefCQ) SetTimestampNoIndex_GreaterThan(value df.Timestamp) *VendorLargeDataRefCQ {
	q.regTimestampNoIndex(df.CK_GT_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetTimestampNoIndex_LessThan(value df.Timestamp) *VendorLargeDataRefCQ {
	q.regTimestampNoIndex(df.CK_LT_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetTimestampNoIndex_GreaterEqual(value df.Timestamp) *VendorLargeDataRefCQ {
	q.regTimestampNoIndex(df.CK_GE_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetTimestampNoIndex_LessEqual(value df.Timestamp) *VendorLargeDataRefCQ {
	q.regTimestampNoIndex(df.CK_LE_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) AddOrderBy_TimestampNoIndex_Asc() *VendorLargeDataRefCQ {
	q.BaseConditionQuery.RegOBA("timestampNoIndex")
	return q
}
func (q *VendorLargeDataRefCQ) AddOrderBy_TimestampNoIndex_Desc() *VendorLargeDataRefCQ {
	q.BaseConditionQuery.RegOBD("timestampNoIndex")
	return q
}
func (q *VendorLargeDataRefCQ) regTimestampNoIndex(key *df.ConditionKey, value interface{}) {
	if q.TimestampNoIndex == nil {
		q.TimestampNoIndex = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.TimestampNoIndex, "timestampNoIndex")
}

func (q *VendorLargeDataRefCQ) getCValueNullableDecimalIndex() *df.ConditionValue {
	if q.NullableDecimalIndex == nil {
		q.NullableDecimalIndex = new(df.ConditionValue)
	}
	return q.NullableDecimalIndex
}



func (q *VendorLargeDataRefCQ) SetNullableDecimalIndex_Equal(value df.Numeric) *VendorLargeDataRefCQ {
	q.regNullableDecimalIndex(df.CK_EQ_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetNullableDecimalIndex_NotEqual(value df.Numeric) *VendorLargeDataRefCQ {
	q.regNullableDecimalIndex(df.CK_NE_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetNullableDecimalIndex_GreaterThan(value df.Numeric) *VendorLargeDataRefCQ {
	q.regNullableDecimalIndex(df.CK_GT_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetNullableDecimalIndex_LessThan(value df.Numeric) *VendorLargeDataRefCQ {
	q.regNullableDecimalIndex(df.CK_LT_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetNullableDecimalIndex_GreaterEqual(value df.Numeric) *VendorLargeDataRefCQ {
	q.regNullableDecimalIndex(df.CK_GE_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetNullableDecimalIndex_LessEqual(value df.Numeric) *VendorLargeDataRefCQ {
	q.regNullableDecimalIndex(df.CK_LE_C, value)
	return q
}
func (q *VendorLargeDataRefCQ) SetNullableDecimalIndex_RangeOf(minNumber df.Numeric, maxNumber df.Numeric, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueNullableDecimalIndex(),"nullableDecimalIndex",rangeOfOption)
}	


func (q *VendorLargeDataRefCQ) SetNullableDecimalIndex_IsNull() *VendorLargeDataRefCQ {
	q.regNullableDecimalIndex(df.CK_ISN_C, 0)
	return q
}
func (q *VendorLargeDataRefCQ) SetNullableDecimalIndex_IsNotNull() *VendorLargeDataRefCQ {
	q.regNullableDecimalIndex(df.CK_ISNN_C, 0)
	return q
}
func (q *VendorLargeDataRefCQ) AddOrderBy_NullableDecimalIndex_Asc() *VendorLargeDataRefCQ {
	q.BaseConditionQuery.RegOBA("nullableDecimalIndex")
	return q
}
func (q *VendorLargeDataRefCQ) AddOrderBy_NullableDecimalIndex_Desc() *VendorLargeDataRefCQ {
	q.BaseConditionQuery.RegOBD("nullableDecimalIndex")
	return q
}
func (q *VendorLargeDataRefCQ) regNullableDecimalIndex(key *df.ConditionKey, value interface{}) {
	if q.NullableDecimalIndex == nil {
		q.NullableDecimalIndex = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.NullableDecimalIndex, "nullableDecimalIndex")
}

func (q *VendorLargeDataRefCQ) getCValueNullableDecimalNoIndex() *df.ConditionValue {
	if q.NullableDecimalNoIndex == nil {
		q.NullableDecimalNoIndex = new(df.ConditionValue)
	}
	return q.NullableDecimalNoIndex
}



func (q *VendorLargeDataRefCQ) SetNullableDecimalNoIndex_Equal(value df.Numeric) *VendorLargeDataRefCQ {
	q.regNullableDecimalNoIndex(df.CK_EQ_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetNullableDecimalNoIndex_NotEqual(value df.Numeric) *VendorLargeDataRefCQ {
	q.regNullableDecimalNoIndex(df.CK_NE_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetNullableDecimalNoIndex_GreaterThan(value df.Numeric) *VendorLargeDataRefCQ {
	q.regNullableDecimalNoIndex(df.CK_GT_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetNullableDecimalNoIndex_LessThan(value df.Numeric) *VendorLargeDataRefCQ {
	q.regNullableDecimalNoIndex(df.CK_LT_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetNullableDecimalNoIndex_GreaterEqual(value df.Numeric) *VendorLargeDataRefCQ {
	q.regNullableDecimalNoIndex(df.CK_GE_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetNullableDecimalNoIndex_LessEqual(value df.Numeric) *VendorLargeDataRefCQ {
	q.regNullableDecimalNoIndex(df.CK_LE_C, value)
	return q
}
func (q *VendorLargeDataRefCQ) SetNullableDecimalNoIndex_RangeOf(minNumber df.Numeric, maxNumber df.Numeric, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueNullableDecimalNoIndex(),"nullableDecimalNoIndex",rangeOfOption)
}	


func (q *VendorLargeDataRefCQ) SetNullableDecimalNoIndex_IsNull() *VendorLargeDataRefCQ {
	q.regNullableDecimalNoIndex(df.CK_ISN_C, 0)
	return q
}
func (q *VendorLargeDataRefCQ) SetNullableDecimalNoIndex_IsNotNull() *VendorLargeDataRefCQ {
	q.regNullableDecimalNoIndex(df.CK_ISNN_C, 0)
	return q
}
func (q *VendorLargeDataRefCQ) AddOrderBy_NullableDecimalNoIndex_Asc() *VendorLargeDataRefCQ {
	q.BaseConditionQuery.RegOBA("nullableDecimalNoIndex")
	return q
}
func (q *VendorLargeDataRefCQ) AddOrderBy_NullableDecimalNoIndex_Desc() *VendorLargeDataRefCQ {
	q.BaseConditionQuery.RegOBD("nullableDecimalNoIndex")
	return q
}
func (q *VendorLargeDataRefCQ) regNullableDecimalNoIndex(key *df.ConditionKey, value interface{}) {
	if q.NullableDecimalNoIndex == nil {
		q.NullableDecimalNoIndex = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.NullableDecimalNoIndex, "nullableDecimalNoIndex")
}

func (q *VendorLargeDataRefCQ) getCValueSelfParentId() *df.ConditionValue {
	if q.SelfParentId == nil {
		q.SelfParentId = new(df.ConditionValue)
	}
	return q.SelfParentId
}



func (q *VendorLargeDataRefCQ) SetSelfParentId_Equal(value int64) *VendorLargeDataRefCQ {
	q.regSelfParentId(df.CK_EQ_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetSelfParentId_NotEqual(value int64) *VendorLargeDataRefCQ {
	q.regSelfParentId(df.CK_NE_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetSelfParentId_GreaterThan(value int64) *VendorLargeDataRefCQ {
	q.regSelfParentId(df.CK_GT_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetSelfParentId_LessThan(value int64) *VendorLargeDataRefCQ {
	q.regSelfParentId(df.CK_LT_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetSelfParentId_GreaterEqual(value int64) *VendorLargeDataRefCQ {
	q.regSelfParentId(df.CK_GE_C, value)
	return q
}

func (q *VendorLargeDataRefCQ) SetSelfParentId_LessEqual(value int64) *VendorLargeDataRefCQ {
	q.regSelfParentId(df.CK_LE_C, value)
	return q
}
func (q *VendorLargeDataRefCQ) SetSelfParentId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueSelfParentId(),"selfParentId",rangeOfOption)
}	


func (q *VendorLargeDataRefCQ) SetSelfParentId_IsNull() *VendorLargeDataRefCQ {
	q.regSelfParentId(df.CK_ISN_C, 0)
	return q
}
func (q *VendorLargeDataRefCQ) SetSelfParentId_IsNotNull() *VendorLargeDataRefCQ {
	q.regSelfParentId(df.CK_ISNN_C, 0)
	return q
}
func (q *VendorLargeDataRefCQ) AddOrderBy_SelfParentId_Asc() *VendorLargeDataRefCQ {
	q.BaseConditionQuery.RegOBA("selfParentId")
	return q
}
func (q *VendorLargeDataRefCQ) AddOrderBy_SelfParentId_Desc() *VendorLargeDataRefCQ {
	q.BaseConditionQuery.RegOBD("selfParentId")
	return q
}
func (q *VendorLargeDataRefCQ) regSelfParentId(key *df.ConditionKey, value interface{}) {
	if q.SelfParentId == nil {
		q.SelfParentId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.SelfParentId, "selfParentId")
}

