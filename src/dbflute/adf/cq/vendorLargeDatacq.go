package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorLargeDataCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	LargeDataId *df.ConditionValue
	StringIndex *df.ConditionValue
	StringNoIndex *df.ConditionValue
	StringUniqueIndex *df.ConditionValue
	IntflgIndex *df.ConditionValue
	NumericIntegerIndex *df.ConditionValue
	NumericIntegerNoIndex *df.ConditionValue
}

func (q *VendorLargeDataCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *VendorLargeDataCQ) getCValueLargeDataId() *df.ConditionValue {
	if q.LargeDataId == nil {
		q.LargeDataId = new(df.ConditionValue)
	}
	return q.LargeDataId
}



func (q *VendorLargeDataCQ) SetLargeDataId_Equal(value int64) *VendorLargeDataCQ {
	q.regLargeDataId(df.CK_EQ_C, value)
	return q
}

func (q *VendorLargeDataCQ) SetLargeDataId_NotEqual(value int64) *VendorLargeDataCQ {
	q.regLargeDataId(df.CK_NE_C, value)
	return q
}

func (q *VendorLargeDataCQ) SetLargeDataId_GreaterThan(value int64) *VendorLargeDataCQ {
	q.regLargeDataId(df.CK_GT_C, value)
	return q
}

func (q *VendorLargeDataCQ) SetLargeDataId_LessThan(value int64) *VendorLargeDataCQ {
	q.regLargeDataId(df.CK_LT_C, value)
	return q
}

func (q *VendorLargeDataCQ) SetLargeDataId_GreaterEqual(value int64) *VendorLargeDataCQ {
	q.regLargeDataId(df.CK_GE_C, value)
	return q
}

func (q *VendorLargeDataCQ) SetLargeDataId_LessEqual(value int64) *VendorLargeDataCQ {
	q.regLargeDataId(df.CK_LE_C, value)
	return q
}
func (q *VendorLargeDataCQ) SetLargeDataId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueLargeDataId(),"largeDataId",rangeOfOption)
}	


func (q *VendorLargeDataCQ) SetLargeDataId_IsNull() *VendorLargeDataCQ {
	q.regLargeDataId(df.CK_ISN_C, 0)
	return q
}
func (q *VendorLargeDataCQ) SetLargeDataId_IsNotNull() *VendorLargeDataCQ {
	q.regLargeDataId(df.CK_ISNN_C, 0)
	return q
}
func (q *VendorLargeDataCQ) AddOrderBy_LargeDataId_Asc() *VendorLargeDataCQ {
	q.BaseConditionQuery.RegOBA("largeDataId")
	return q
}
func (q *VendorLargeDataCQ) AddOrderBy_LargeDataId_Desc() *VendorLargeDataCQ {
	q.BaseConditionQuery.RegOBD("largeDataId")
	return q
}
func (q *VendorLargeDataCQ) regLargeDataId(key *df.ConditionKey, value interface{}) {
	if q.LargeDataId == nil {
		q.LargeDataId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.LargeDataId, "largeDataId")
}

func (q *VendorLargeDataCQ) getCValueStringIndex() *df.ConditionValue {
	if q.StringIndex == nil {
		q.StringIndex = new(df.ConditionValue)
	}
	return q.StringIndex
}


func (q *VendorLargeDataCQ) SetStringIndex_Equal(value string) *VendorLargeDataCQ {
	q.regStringIndex(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *VendorLargeDataCQ) SetStringIndex_NotEqual(value string) *VendorLargeDataCQ {
	q.regStringIndex(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorLargeDataCQ) SetStringIndex_GreaterThan(value string) *VendorLargeDataCQ {
	q.regStringIndex(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorLargeDataCQ) SetStringIndex_LessThan(value string) *VendorLargeDataCQ {
	q.regStringIndex(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorLargeDataCQ) SetStringIndex_GreaterEqualThan(value string) *VendorLargeDataCQ {
	q.regStringIndex(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *VendorLargeDataCQ) SetStringIndex_LessEqualThan(value string) *VendorLargeDataCQ {
	q.regStringIndex(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorLargeDataCQ) SetStringIndex_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueStringIndex(), "stringIndex", option)
}

func (q *VendorLargeDataCQ) SetStringIndex_PrefixSearch(value string) error {
	return q.SetStringIndex_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *VendorLargeDataCQ) SetStringIndex_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueStringIndex(), "stringIndex", option)
}



func (q *VendorLargeDataCQ) AddOrderBy_StringIndex_Asc() *VendorLargeDataCQ {
	q.BaseConditionQuery.RegOBA("stringIndex")
	return q
}
func (q *VendorLargeDataCQ) AddOrderBy_StringIndex_Desc() *VendorLargeDataCQ {
	q.BaseConditionQuery.RegOBD("stringIndex")
	return q
}
func (q *VendorLargeDataCQ) regStringIndex(key *df.ConditionKey, value interface{}) {
	if q.StringIndex == nil {
		q.StringIndex = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.StringIndex, "stringIndex")
}

func (q *VendorLargeDataCQ) getCValueStringNoIndex() *df.ConditionValue {
	if q.StringNoIndex == nil {
		q.StringNoIndex = new(df.ConditionValue)
	}
	return q.StringNoIndex
}


func (q *VendorLargeDataCQ) SetStringNoIndex_Equal(value string) *VendorLargeDataCQ {
	q.regStringNoIndex(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *VendorLargeDataCQ) SetStringNoIndex_NotEqual(value string) *VendorLargeDataCQ {
	q.regStringNoIndex(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorLargeDataCQ) SetStringNoIndex_GreaterThan(value string) *VendorLargeDataCQ {
	q.regStringNoIndex(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorLargeDataCQ) SetStringNoIndex_LessThan(value string) *VendorLargeDataCQ {
	q.regStringNoIndex(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorLargeDataCQ) SetStringNoIndex_GreaterEqualThan(value string) *VendorLargeDataCQ {
	q.regStringNoIndex(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *VendorLargeDataCQ) SetStringNoIndex_LessEqualThan(value string) *VendorLargeDataCQ {
	q.regStringNoIndex(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorLargeDataCQ) SetStringNoIndex_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueStringNoIndex(), "stringNoIndex", option)
}

func (q *VendorLargeDataCQ) SetStringNoIndex_PrefixSearch(value string) error {
	return q.SetStringNoIndex_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *VendorLargeDataCQ) SetStringNoIndex_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueStringNoIndex(), "stringNoIndex", option)
}



func (q *VendorLargeDataCQ) AddOrderBy_StringNoIndex_Asc() *VendorLargeDataCQ {
	q.BaseConditionQuery.RegOBA("stringNoIndex")
	return q
}
func (q *VendorLargeDataCQ) AddOrderBy_StringNoIndex_Desc() *VendorLargeDataCQ {
	q.BaseConditionQuery.RegOBD("stringNoIndex")
	return q
}
func (q *VendorLargeDataCQ) regStringNoIndex(key *df.ConditionKey, value interface{}) {
	if q.StringNoIndex == nil {
		q.StringNoIndex = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.StringNoIndex, "stringNoIndex")
}

func (q *VendorLargeDataCQ) getCValueStringUniqueIndex() *df.ConditionValue {
	if q.StringUniqueIndex == nil {
		q.StringUniqueIndex = new(df.ConditionValue)
	}
	return q.StringUniqueIndex
}


func (q *VendorLargeDataCQ) SetStringUniqueIndex_Equal(value string) *VendorLargeDataCQ {
	q.regStringUniqueIndex(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *VendorLargeDataCQ) SetStringUniqueIndex_NotEqual(value string) *VendorLargeDataCQ {
	q.regStringUniqueIndex(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorLargeDataCQ) SetStringUniqueIndex_GreaterThan(value string) *VendorLargeDataCQ {
	q.regStringUniqueIndex(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorLargeDataCQ) SetStringUniqueIndex_LessThan(value string) *VendorLargeDataCQ {
	q.regStringUniqueIndex(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorLargeDataCQ) SetStringUniqueIndex_GreaterEqualThan(value string) *VendorLargeDataCQ {
	q.regStringUniqueIndex(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *VendorLargeDataCQ) SetStringUniqueIndex_LessEqualThan(value string) *VendorLargeDataCQ {
	q.regStringUniqueIndex(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorLargeDataCQ) SetStringUniqueIndex_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueStringUniqueIndex(), "stringUniqueIndex", option)
}

func (q *VendorLargeDataCQ) SetStringUniqueIndex_PrefixSearch(value string) error {
	return q.SetStringUniqueIndex_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *VendorLargeDataCQ) SetStringUniqueIndex_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueStringUniqueIndex(), "stringUniqueIndex", option)
}



func (q *VendorLargeDataCQ) AddOrderBy_StringUniqueIndex_Asc() *VendorLargeDataCQ {
	q.BaseConditionQuery.RegOBA("stringUniqueIndex")
	return q
}
func (q *VendorLargeDataCQ) AddOrderBy_StringUniqueIndex_Desc() *VendorLargeDataCQ {
	q.BaseConditionQuery.RegOBD("stringUniqueIndex")
	return q
}
func (q *VendorLargeDataCQ) regStringUniqueIndex(key *df.ConditionKey, value interface{}) {
	if q.StringUniqueIndex == nil {
		q.StringUniqueIndex = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.StringUniqueIndex, "stringUniqueIndex")
}

func (q *VendorLargeDataCQ) getCValueIntflgIndex() *df.ConditionValue {
	if q.IntflgIndex == nil {
		q.IntflgIndex = new(df.ConditionValue)
	}
	return q.IntflgIndex
}



func (q *VendorLargeDataCQ) SetIntflgIndex_Equal(value int64) *VendorLargeDataCQ {
	q.regIntflgIndex(df.CK_EQ_C, value)
	return q
}

func (q *VendorLargeDataCQ) SetIntflgIndex_NotEqual(value int64) *VendorLargeDataCQ {
	q.regIntflgIndex(df.CK_NE_C, value)
	return q
}

func (q *VendorLargeDataCQ) SetIntflgIndex_GreaterThan(value int64) *VendorLargeDataCQ {
	q.regIntflgIndex(df.CK_GT_C, value)
	return q
}

func (q *VendorLargeDataCQ) SetIntflgIndex_LessThan(value int64) *VendorLargeDataCQ {
	q.regIntflgIndex(df.CK_LT_C, value)
	return q
}

func (q *VendorLargeDataCQ) SetIntflgIndex_GreaterEqual(value int64) *VendorLargeDataCQ {
	q.regIntflgIndex(df.CK_GE_C, value)
	return q
}

func (q *VendorLargeDataCQ) SetIntflgIndex_LessEqual(value int64) *VendorLargeDataCQ {
	q.regIntflgIndex(df.CK_LE_C, value)
	return q
}
func (q *VendorLargeDataCQ) SetIntflgIndex_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueIntflgIndex(),"intflgIndex",rangeOfOption)
}	


func (q *VendorLargeDataCQ) AddOrderBy_IntflgIndex_Asc() *VendorLargeDataCQ {
	q.BaseConditionQuery.RegOBA("intflgIndex")
	return q
}
func (q *VendorLargeDataCQ) AddOrderBy_IntflgIndex_Desc() *VendorLargeDataCQ {
	q.BaseConditionQuery.RegOBD("intflgIndex")
	return q
}
func (q *VendorLargeDataCQ) regIntflgIndex(key *df.ConditionKey, value interface{}) {
	if q.IntflgIndex == nil {
		q.IntflgIndex = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.IntflgIndex, "intflgIndex")
}

func (q *VendorLargeDataCQ) getCValueNumericIntegerIndex() *df.ConditionValue {
	if q.NumericIntegerIndex == nil {
		q.NumericIntegerIndex = new(df.ConditionValue)
	}
	return q.NumericIntegerIndex
}



func (q *VendorLargeDataCQ) SetNumericIntegerIndex_Equal(value df.Numeric) *VendorLargeDataCQ {
	q.regNumericIntegerIndex(df.CK_EQ_C, value)
	return q
}

func (q *VendorLargeDataCQ) SetNumericIntegerIndex_NotEqual(value df.Numeric) *VendorLargeDataCQ {
	q.regNumericIntegerIndex(df.CK_NE_C, value)
	return q
}

func (q *VendorLargeDataCQ) SetNumericIntegerIndex_GreaterThan(value df.Numeric) *VendorLargeDataCQ {
	q.regNumericIntegerIndex(df.CK_GT_C, value)
	return q
}

func (q *VendorLargeDataCQ) SetNumericIntegerIndex_LessThan(value df.Numeric) *VendorLargeDataCQ {
	q.regNumericIntegerIndex(df.CK_LT_C, value)
	return q
}

func (q *VendorLargeDataCQ) SetNumericIntegerIndex_GreaterEqual(value df.Numeric) *VendorLargeDataCQ {
	q.regNumericIntegerIndex(df.CK_GE_C, value)
	return q
}

func (q *VendorLargeDataCQ) SetNumericIntegerIndex_LessEqual(value df.Numeric) *VendorLargeDataCQ {
	q.regNumericIntegerIndex(df.CK_LE_C, value)
	return q
}
func (q *VendorLargeDataCQ) SetNumericIntegerIndex_RangeOf(minNumber df.Numeric, maxNumber df.Numeric, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueNumericIntegerIndex(),"numericIntegerIndex",rangeOfOption)
}	


func (q *VendorLargeDataCQ) AddOrderBy_NumericIntegerIndex_Asc() *VendorLargeDataCQ {
	q.BaseConditionQuery.RegOBA("numericIntegerIndex")
	return q
}
func (q *VendorLargeDataCQ) AddOrderBy_NumericIntegerIndex_Desc() *VendorLargeDataCQ {
	q.BaseConditionQuery.RegOBD("numericIntegerIndex")
	return q
}
func (q *VendorLargeDataCQ) regNumericIntegerIndex(key *df.ConditionKey, value interface{}) {
	if q.NumericIntegerIndex == nil {
		q.NumericIntegerIndex = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.NumericIntegerIndex, "numericIntegerIndex")
}

func (q *VendorLargeDataCQ) getCValueNumericIntegerNoIndex() *df.ConditionValue {
	if q.NumericIntegerNoIndex == nil {
		q.NumericIntegerNoIndex = new(df.ConditionValue)
	}
	return q.NumericIntegerNoIndex
}



func (q *VendorLargeDataCQ) SetNumericIntegerNoIndex_Equal(value df.Numeric) *VendorLargeDataCQ {
	q.regNumericIntegerNoIndex(df.CK_EQ_C, value)
	return q
}

func (q *VendorLargeDataCQ) SetNumericIntegerNoIndex_NotEqual(value df.Numeric) *VendorLargeDataCQ {
	q.regNumericIntegerNoIndex(df.CK_NE_C, value)
	return q
}

func (q *VendorLargeDataCQ) SetNumericIntegerNoIndex_GreaterThan(value df.Numeric) *VendorLargeDataCQ {
	q.regNumericIntegerNoIndex(df.CK_GT_C, value)
	return q
}

func (q *VendorLargeDataCQ) SetNumericIntegerNoIndex_LessThan(value df.Numeric) *VendorLargeDataCQ {
	q.regNumericIntegerNoIndex(df.CK_LT_C, value)
	return q
}

func (q *VendorLargeDataCQ) SetNumericIntegerNoIndex_GreaterEqual(value df.Numeric) *VendorLargeDataCQ {
	q.regNumericIntegerNoIndex(df.CK_GE_C, value)
	return q
}

func (q *VendorLargeDataCQ) SetNumericIntegerNoIndex_LessEqual(value df.Numeric) *VendorLargeDataCQ {
	q.regNumericIntegerNoIndex(df.CK_LE_C, value)
	return q
}
func (q *VendorLargeDataCQ) SetNumericIntegerNoIndex_RangeOf(minNumber df.Numeric, maxNumber df.Numeric, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueNumericIntegerNoIndex(),"numericIntegerNoIndex",rangeOfOption)
}	


func (q *VendorLargeDataCQ) AddOrderBy_NumericIntegerNoIndex_Asc() *VendorLargeDataCQ {
	q.BaseConditionQuery.RegOBA("numericIntegerNoIndex")
	return q
}
func (q *VendorLargeDataCQ) AddOrderBy_NumericIntegerNoIndex_Desc() *VendorLargeDataCQ {
	q.BaseConditionQuery.RegOBD("numericIntegerNoIndex")
	return q
}
func (q *VendorLargeDataCQ) regNumericIntegerNoIndex(key *df.ConditionKey, value interface{}) {
	if q.NumericIntegerNoIndex == nil {
		q.NumericIntegerNoIndex = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.NumericIntegerNoIndex, "numericIntegerNoIndex")
}

