package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type WhiteXlsManCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	XlsManId *df.ConditionValue
	StringConverted *df.ConditionValue
	TimestampZeroDefaultMillis *df.ConditionValue
	TimestampZeroPrefixMillis *df.ConditionValue
	TimestampZeroSuffixMillis *df.ConditionValue
}

func (q *WhiteXlsManCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *WhiteXlsManCQ) getCValueXlsManId() *df.ConditionValue {
	if q.XlsManId == nil {
		q.XlsManId = new(df.ConditionValue)
	}
	return q.XlsManId
}



func (q *WhiteXlsManCQ) SetXlsManId_Equal(value int64) *WhiteXlsManCQ {
	q.regXlsManId(df.CK_EQ_C, value)
	return q
}

func (q *WhiteXlsManCQ) SetXlsManId_NotEqual(value int64) *WhiteXlsManCQ {
	q.regXlsManId(df.CK_NE_C, value)
	return q
}

func (q *WhiteXlsManCQ) SetXlsManId_GreaterThan(value int64) *WhiteXlsManCQ {
	q.regXlsManId(df.CK_GT_C, value)
	return q
}

func (q *WhiteXlsManCQ) SetXlsManId_LessThan(value int64) *WhiteXlsManCQ {
	q.regXlsManId(df.CK_LT_C, value)
	return q
}

func (q *WhiteXlsManCQ) SetXlsManId_GreaterEqual(value int64) *WhiteXlsManCQ {
	q.regXlsManId(df.CK_GE_C, value)
	return q
}

func (q *WhiteXlsManCQ) SetXlsManId_LessEqual(value int64) *WhiteXlsManCQ {
	q.regXlsManId(df.CK_LE_C, value)
	return q
}
func (q *WhiteXlsManCQ) SetXlsManId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueXlsManId(),"xlsManId",rangeOfOption)
}	


func (q *WhiteXlsManCQ) SetXlsManId_IsNull() *WhiteXlsManCQ {
	q.regXlsManId(df.CK_ISN_C, 0)
	return q
}
func (q *WhiteXlsManCQ) SetXlsManId_IsNotNull() *WhiteXlsManCQ {
	q.regXlsManId(df.CK_ISNN_C, 0)
	return q
}
func (q *WhiteXlsManCQ) AddOrderBy_XlsManId_Asc() *WhiteXlsManCQ {
	q.BaseConditionQuery.RegOBA("xlsManId")
	return q
}
func (q *WhiteXlsManCQ) AddOrderBy_XlsManId_Desc() *WhiteXlsManCQ {
	q.BaseConditionQuery.RegOBD("xlsManId")
	return q
}
func (q *WhiteXlsManCQ) regXlsManId(key *df.ConditionKey, value interface{}) {
	if q.XlsManId == nil {
		q.XlsManId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.XlsManId, "xlsManId")
}

func (q *WhiteXlsManCQ) getCValueStringConverted() *df.ConditionValue {
	if q.StringConverted == nil {
		q.StringConverted = new(df.ConditionValue)
	}
	return q.StringConverted
}


func (q *WhiteXlsManCQ) SetStringConverted_Equal(value string) *WhiteXlsManCQ {
	q.regStringConverted(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *WhiteXlsManCQ) SetStringConverted_NotEqual(value string) *WhiteXlsManCQ {
	q.regStringConverted(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WhiteXlsManCQ) SetStringConverted_GreaterThan(value string) *WhiteXlsManCQ {
	q.regStringConverted(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WhiteXlsManCQ) SetStringConverted_LessThan(value string) *WhiteXlsManCQ {
	q.regStringConverted(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WhiteXlsManCQ) SetStringConverted_GreaterEqualThan(value string) *WhiteXlsManCQ {
	q.regStringConverted(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *WhiteXlsManCQ) SetStringConverted_LessEqualThan(value string) *WhiteXlsManCQ {
	q.regStringConverted(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WhiteXlsManCQ) SetStringConverted_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueStringConverted(), "stringConverted", option)
}

func (q *WhiteXlsManCQ) SetStringConverted_PrefixSearch(value string) error {
	return q.SetStringConverted_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *WhiteXlsManCQ) SetStringConverted_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueStringConverted(), "stringConverted", option)
}



func (q *WhiteXlsManCQ) SetStringConverted_IsNull() *WhiteXlsManCQ {
	q.regStringConverted(df.CK_ISN_C, 0)
	return q
}
func (q *WhiteXlsManCQ) SetStringConverted_IsNullOrEmpty() *WhiteXlsManCQ {
	q.regStringConverted(df.CK_ISNOE_C, 0)
	return q
}
func (q *WhiteXlsManCQ) SetStringConverted_IsNotNull() *WhiteXlsManCQ {
	q.regStringConverted(df.CK_ISNN_C, 0)
	return q
}
func (q *WhiteXlsManCQ) AddOrderBy_StringConverted_Asc() *WhiteXlsManCQ {
	q.BaseConditionQuery.RegOBA("stringConverted")
	return q
}
func (q *WhiteXlsManCQ) AddOrderBy_StringConverted_Desc() *WhiteXlsManCQ {
	q.BaseConditionQuery.RegOBD("stringConverted")
	return q
}
func (q *WhiteXlsManCQ) regStringConverted(key *df.ConditionKey, value interface{}) {
	if q.StringConverted == nil {
		q.StringConverted = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.StringConverted, "stringConverted")
}

func (q *WhiteXlsManCQ) getCValueTimestampZeroDefaultMillis() *df.ConditionValue {
	if q.TimestampZeroDefaultMillis == nil {
		q.TimestampZeroDefaultMillis = new(df.ConditionValue)
	}
	return q.TimestampZeroDefaultMillis
}




func (q *WhiteXlsManCQ) SetTimestampZeroDefaultMillis_Equal(value df.Timestamp) *WhiteXlsManCQ {
	q.regTimestampZeroDefaultMillis(df.CK_EQ_C, value)
	return q
}


func (q *WhiteXlsManCQ) SetTimestampZeroDefaultMillis_GreaterThan(value df.Timestamp) *WhiteXlsManCQ {
	q.regTimestampZeroDefaultMillis(df.CK_GT_C, value)
	return q
}

func (q *WhiteXlsManCQ) SetTimestampZeroDefaultMillis_LessThan(value df.Timestamp) *WhiteXlsManCQ {
	q.regTimestampZeroDefaultMillis(df.CK_LT_C, value)
	return q
}

func (q *WhiteXlsManCQ) SetTimestampZeroDefaultMillis_GreaterEqual(value df.Timestamp) *WhiteXlsManCQ {
	q.regTimestampZeroDefaultMillis(df.CK_GE_C, value)
	return q
}

func (q *WhiteXlsManCQ) SetTimestampZeroDefaultMillis_LessEqual(value df.Timestamp) *WhiteXlsManCQ {
	q.regTimestampZeroDefaultMillis(df.CK_LE_C, value)
	return q
}

func (q *WhiteXlsManCQ) SetTimestampZeroDefaultMillis_IsNull() *WhiteXlsManCQ {
	q.regTimestampZeroDefaultMillis(df.CK_ISN_C, 0)
	return q
}
func (q *WhiteXlsManCQ) SetTimestampZeroDefaultMillis_IsNotNull() *WhiteXlsManCQ {
	q.regTimestampZeroDefaultMillis(df.CK_ISNN_C, 0)
	return q
}
func (q *WhiteXlsManCQ) AddOrderBy_TimestampZeroDefaultMillis_Asc() *WhiteXlsManCQ {
	q.BaseConditionQuery.RegOBA("timestampZeroDefaultMillis")
	return q
}
func (q *WhiteXlsManCQ) AddOrderBy_TimestampZeroDefaultMillis_Desc() *WhiteXlsManCQ {
	q.BaseConditionQuery.RegOBD("timestampZeroDefaultMillis")
	return q
}
func (q *WhiteXlsManCQ) regTimestampZeroDefaultMillis(key *df.ConditionKey, value interface{}) {
	if q.TimestampZeroDefaultMillis == nil {
		q.TimestampZeroDefaultMillis = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.TimestampZeroDefaultMillis, "timestampZeroDefaultMillis")
}

func (q *WhiteXlsManCQ) getCValueTimestampZeroPrefixMillis() *df.ConditionValue {
	if q.TimestampZeroPrefixMillis == nil {
		q.TimestampZeroPrefixMillis = new(df.ConditionValue)
	}
	return q.TimestampZeroPrefixMillis
}




func (q *WhiteXlsManCQ) SetTimestampZeroPrefixMillis_Equal(value df.Timestamp) *WhiteXlsManCQ {
	q.regTimestampZeroPrefixMillis(df.CK_EQ_C, value)
	return q
}


func (q *WhiteXlsManCQ) SetTimestampZeroPrefixMillis_GreaterThan(value df.Timestamp) *WhiteXlsManCQ {
	q.regTimestampZeroPrefixMillis(df.CK_GT_C, value)
	return q
}

func (q *WhiteXlsManCQ) SetTimestampZeroPrefixMillis_LessThan(value df.Timestamp) *WhiteXlsManCQ {
	q.regTimestampZeroPrefixMillis(df.CK_LT_C, value)
	return q
}

func (q *WhiteXlsManCQ) SetTimestampZeroPrefixMillis_GreaterEqual(value df.Timestamp) *WhiteXlsManCQ {
	q.regTimestampZeroPrefixMillis(df.CK_GE_C, value)
	return q
}

func (q *WhiteXlsManCQ) SetTimestampZeroPrefixMillis_LessEqual(value df.Timestamp) *WhiteXlsManCQ {
	q.regTimestampZeroPrefixMillis(df.CK_LE_C, value)
	return q
}

func (q *WhiteXlsManCQ) SetTimestampZeroPrefixMillis_IsNull() *WhiteXlsManCQ {
	q.regTimestampZeroPrefixMillis(df.CK_ISN_C, 0)
	return q
}
func (q *WhiteXlsManCQ) SetTimestampZeroPrefixMillis_IsNotNull() *WhiteXlsManCQ {
	q.regTimestampZeroPrefixMillis(df.CK_ISNN_C, 0)
	return q
}
func (q *WhiteXlsManCQ) AddOrderBy_TimestampZeroPrefixMillis_Asc() *WhiteXlsManCQ {
	q.BaseConditionQuery.RegOBA("timestampZeroPrefixMillis")
	return q
}
func (q *WhiteXlsManCQ) AddOrderBy_TimestampZeroPrefixMillis_Desc() *WhiteXlsManCQ {
	q.BaseConditionQuery.RegOBD("timestampZeroPrefixMillis")
	return q
}
func (q *WhiteXlsManCQ) regTimestampZeroPrefixMillis(key *df.ConditionKey, value interface{}) {
	if q.TimestampZeroPrefixMillis == nil {
		q.TimestampZeroPrefixMillis = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.TimestampZeroPrefixMillis, "timestampZeroPrefixMillis")
}

func (q *WhiteXlsManCQ) getCValueTimestampZeroSuffixMillis() *df.ConditionValue {
	if q.TimestampZeroSuffixMillis == nil {
		q.TimestampZeroSuffixMillis = new(df.ConditionValue)
	}
	return q.TimestampZeroSuffixMillis
}




func (q *WhiteXlsManCQ) SetTimestampZeroSuffixMillis_Equal(value df.Timestamp) *WhiteXlsManCQ {
	q.regTimestampZeroSuffixMillis(df.CK_EQ_C, value)
	return q
}


func (q *WhiteXlsManCQ) SetTimestampZeroSuffixMillis_GreaterThan(value df.Timestamp) *WhiteXlsManCQ {
	q.regTimestampZeroSuffixMillis(df.CK_GT_C, value)
	return q
}

func (q *WhiteXlsManCQ) SetTimestampZeroSuffixMillis_LessThan(value df.Timestamp) *WhiteXlsManCQ {
	q.regTimestampZeroSuffixMillis(df.CK_LT_C, value)
	return q
}

func (q *WhiteXlsManCQ) SetTimestampZeroSuffixMillis_GreaterEqual(value df.Timestamp) *WhiteXlsManCQ {
	q.regTimestampZeroSuffixMillis(df.CK_GE_C, value)
	return q
}

func (q *WhiteXlsManCQ) SetTimestampZeroSuffixMillis_LessEqual(value df.Timestamp) *WhiteXlsManCQ {
	q.regTimestampZeroSuffixMillis(df.CK_LE_C, value)
	return q
}

func (q *WhiteXlsManCQ) SetTimestampZeroSuffixMillis_IsNull() *WhiteXlsManCQ {
	q.regTimestampZeroSuffixMillis(df.CK_ISN_C, 0)
	return q
}
func (q *WhiteXlsManCQ) SetTimestampZeroSuffixMillis_IsNotNull() *WhiteXlsManCQ {
	q.regTimestampZeroSuffixMillis(df.CK_ISNN_C, 0)
	return q
}
func (q *WhiteXlsManCQ) AddOrderBy_TimestampZeroSuffixMillis_Asc() *WhiteXlsManCQ {
	q.BaseConditionQuery.RegOBA("timestampZeroSuffixMillis")
	return q
}
func (q *WhiteXlsManCQ) AddOrderBy_TimestampZeroSuffixMillis_Desc() *WhiteXlsManCQ {
	q.BaseConditionQuery.RegOBD("timestampZeroSuffixMillis")
	return q
}
func (q *WhiteXlsManCQ) regTimestampZeroSuffixMillis(key *df.ConditionKey, value interface{}) {
	if q.TimestampZeroSuffixMillis == nil {
		q.TimestampZeroSuffixMillis = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.TimestampZeroSuffixMillis, "timestampZeroSuffixMillis")
}

