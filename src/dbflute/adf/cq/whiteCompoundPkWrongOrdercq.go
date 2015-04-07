package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type WhiteCompoundPkWrongOrderCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	FirstId *df.ConditionValue
	SecondId *df.ConditionValue
	ThirdId *df.ConditionValue
	WrongName *df.ConditionValue
}

func (q *WhiteCompoundPkWrongOrderCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *WhiteCompoundPkWrongOrderCQ) getCValueFirstId() *df.ConditionValue {
	if q.FirstId == nil {
		q.FirstId = new(df.ConditionValue)
	}
	return q.FirstId
}



func (q *WhiteCompoundPkWrongOrderCQ) SetFirstId_Equal(value int64) *WhiteCompoundPkWrongOrderCQ {
	q.regFirstId(df.CK_EQ_C, value)
	return q
}

func (q *WhiteCompoundPkWrongOrderCQ) SetFirstId_NotEqual(value int64) *WhiteCompoundPkWrongOrderCQ {
	q.regFirstId(df.CK_NE_C, value)
	return q
}

func (q *WhiteCompoundPkWrongOrderCQ) SetFirstId_GreaterThan(value int64) *WhiteCompoundPkWrongOrderCQ {
	q.regFirstId(df.CK_GT_C, value)
	return q
}

func (q *WhiteCompoundPkWrongOrderCQ) SetFirstId_LessThan(value int64) *WhiteCompoundPkWrongOrderCQ {
	q.regFirstId(df.CK_LT_C, value)
	return q
}

func (q *WhiteCompoundPkWrongOrderCQ) SetFirstId_GreaterEqual(value int64) *WhiteCompoundPkWrongOrderCQ {
	q.regFirstId(df.CK_GE_C, value)
	return q
}

func (q *WhiteCompoundPkWrongOrderCQ) SetFirstId_LessEqual(value int64) *WhiteCompoundPkWrongOrderCQ {
	q.regFirstId(df.CK_LE_C, value)
	return q
}
func (q *WhiteCompoundPkWrongOrderCQ) SetFirstId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueFirstId(),"firstId",rangeOfOption)
}	


func (q *WhiteCompoundPkWrongOrderCQ) SetFirstId_IsNull() *WhiteCompoundPkWrongOrderCQ {
	q.regFirstId(df.CK_ISN_C, 0)
	return q
}
func (q *WhiteCompoundPkWrongOrderCQ) SetFirstId_IsNotNull() *WhiteCompoundPkWrongOrderCQ {
	q.regFirstId(df.CK_ISNN_C, 0)
	return q
}
func (q *WhiteCompoundPkWrongOrderCQ) AddOrderBy_FirstId_Asc() *WhiteCompoundPkWrongOrderCQ {
	q.BaseConditionQuery.RegOBA("firstId")
	return q
}
func (q *WhiteCompoundPkWrongOrderCQ) AddOrderBy_FirstId_Desc() *WhiteCompoundPkWrongOrderCQ {
	q.BaseConditionQuery.RegOBD("firstId")
	return q
}
func (q *WhiteCompoundPkWrongOrderCQ) regFirstId(key *df.ConditionKey, value interface{}) {
	if q.FirstId == nil {
		q.FirstId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.FirstId, "firstId")
}

func (q *WhiteCompoundPkWrongOrderCQ) getCValueSecondId() *df.ConditionValue {
	if q.SecondId == nil {
		q.SecondId = new(df.ConditionValue)
	}
	return q.SecondId
}



func (q *WhiteCompoundPkWrongOrderCQ) SetSecondId_Equal(value int64) *WhiteCompoundPkWrongOrderCQ {
	q.regSecondId(df.CK_EQ_C, value)
	return q
}

func (q *WhiteCompoundPkWrongOrderCQ) SetSecondId_NotEqual(value int64) *WhiteCompoundPkWrongOrderCQ {
	q.regSecondId(df.CK_NE_C, value)
	return q
}

func (q *WhiteCompoundPkWrongOrderCQ) SetSecondId_GreaterThan(value int64) *WhiteCompoundPkWrongOrderCQ {
	q.regSecondId(df.CK_GT_C, value)
	return q
}

func (q *WhiteCompoundPkWrongOrderCQ) SetSecondId_LessThan(value int64) *WhiteCompoundPkWrongOrderCQ {
	q.regSecondId(df.CK_LT_C, value)
	return q
}

func (q *WhiteCompoundPkWrongOrderCQ) SetSecondId_GreaterEqual(value int64) *WhiteCompoundPkWrongOrderCQ {
	q.regSecondId(df.CK_GE_C, value)
	return q
}

func (q *WhiteCompoundPkWrongOrderCQ) SetSecondId_LessEqual(value int64) *WhiteCompoundPkWrongOrderCQ {
	q.regSecondId(df.CK_LE_C, value)
	return q
}
func (q *WhiteCompoundPkWrongOrderCQ) SetSecondId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueSecondId(),"secondId",rangeOfOption)
}	


func (q *WhiteCompoundPkWrongOrderCQ) SetSecondId_IsNull() *WhiteCompoundPkWrongOrderCQ {
	q.regSecondId(df.CK_ISN_C, 0)
	return q
}
func (q *WhiteCompoundPkWrongOrderCQ) SetSecondId_IsNotNull() *WhiteCompoundPkWrongOrderCQ {
	q.regSecondId(df.CK_ISNN_C, 0)
	return q
}
func (q *WhiteCompoundPkWrongOrderCQ) AddOrderBy_SecondId_Asc() *WhiteCompoundPkWrongOrderCQ {
	q.BaseConditionQuery.RegOBA("secondId")
	return q
}
func (q *WhiteCompoundPkWrongOrderCQ) AddOrderBy_SecondId_Desc() *WhiteCompoundPkWrongOrderCQ {
	q.BaseConditionQuery.RegOBD("secondId")
	return q
}
func (q *WhiteCompoundPkWrongOrderCQ) regSecondId(key *df.ConditionKey, value interface{}) {
	if q.SecondId == nil {
		q.SecondId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.SecondId, "secondId")
}

func (q *WhiteCompoundPkWrongOrderCQ) getCValueThirdId() *df.ConditionValue {
	if q.ThirdId == nil {
		q.ThirdId = new(df.ConditionValue)
	}
	return q.ThirdId
}



func (q *WhiteCompoundPkWrongOrderCQ) SetThirdId_Equal(value int64) *WhiteCompoundPkWrongOrderCQ {
	q.regThirdId(df.CK_EQ_C, value)
	return q
}

func (q *WhiteCompoundPkWrongOrderCQ) SetThirdId_NotEqual(value int64) *WhiteCompoundPkWrongOrderCQ {
	q.regThirdId(df.CK_NE_C, value)
	return q
}

func (q *WhiteCompoundPkWrongOrderCQ) SetThirdId_GreaterThan(value int64) *WhiteCompoundPkWrongOrderCQ {
	q.regThirdId(df.CK_GT_C, value)
	return q
}

func (q *WhiteCompoundPkWrongOrderCQ) SetThirdId_LessThan(value int64) *WhiteCompoundPkWrongOrderCQ {
	q.regThirdId(df.CK_LT_C, value)
	return q
}

func (q *WhiteCompoundPkWrongOrderCQ) SetThirdId_GreaterEqual(value int64) *WhiteCompoundPkWrongOrderCQ {
	q.regThirdId(df.CK_GE_C, value)
	return q
}

func (q *WhiteCompoundPkWrongOrderCQ) SetThirdId_LessEqual(value int64) *WhiteCompoundPkWrongOrderCQ {
	q.regThirdId(df.CK_LE_C, value)
	return q
}
func (q *WhiteCompoundPkWrongOrderCQ) SetThirdId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueThirdId(),"thirdId",rangeOfOption)
}	


func (q *WhiteCompoundPkWrongOrderCQ) SetThirdId_IsNull() *WhiteCompoundPkWrongOrderCQ {
	q.regThirdId(df.CK_ISN_C, 0)
	return q
}
func (q *WhiteCompoundPkWrongOrderCQ) SetThirdId_IsNotNull() *WhiteCompoundPkWrongOrderCQ {
	q.regThirdId(df.CK_ISNN_C, 0)
	return q
}
func (q *WhiteCompoundPkWrongOrderCQ) AddOrderBy_ThirdId_Asc() *WhiteCompoundPkWrongOrderCQ {
	q.BaseConditionQuery.RegOBA("thirdId")
	return q
}
func (q *WhiteCompoundPkWrongOrderCQ) AddOrderBy_ThirdId_Desc() *WhiteCompoundPkWrongOrderCQ {
	q.BaseConditionQuery.RegOBD("thirdId")
	return q
}
func (q *WhiteCompoundPkWrongOrderCQ) regThirdId(key *df.ConditionKey, value interface{}) {
	if q.ThirdId == nil {
		q.ThirdId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ThirdId, "thirdId")
}

func (q *WhiteCompoundPkWrongOrderCQ) getCValueWrongName() *df.ConditionValue {
	if q.WrongName == nil {
		q.WrongName = new(df.ConditionValue)
	}
	return q.WrongName
}


func (q *WhiteCompoundPkWrongOrderCQ) SetWrongName_Equal(value string) *WhiteCompoundPkWrongOrderCQ {
	q.regWrongName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *WhiteCompoundPkWrongOrderCQ) SetWrongName_NotEqual(value string) *WhiteCompoundPkWrongOrderCQ {
	q.regWrongName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WhiteCompoundPkWrongOrderCQ) SetWrongName_GreaterThan(value string) *WhiteCompoundPkWrongOrderCQ {
	q.regWrongName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WhiteCompoundPkWrongOrderCQ) SetWrongName_LessThan(value string) *WhiteCompoundPkWrongOrderCQ {
	q.regWrongName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WhiteCompoundPkWrongOrderCQ) SetWrongName_GreaterEqualThan(value string) *WhiteCompoundPkWrongOrderCQ {
	q.regWrongName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *WhiteCompoundPkWrongOrderCQ) SetWrongName_LessEqualThan(value string) *WhiteCompoundPkWrongOrderCQ {
	q.regWrongName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WhiteCompoundPkWrongOrderCQ) SetWrongName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueWrongName(), "wrongName", option)
}

func (q *WhiteCompoundPkWrongOrderCQ) SetWrongName_PrefixSearch(value string) error {
	return q.SetWrongName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *WhiteCompoundPkWrongOrderCQ) SetWrongName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueWrongName(), "wrongName", option)
}



func (q *WhiteCompoundPkWrongOrderCQ) SetWrongName_IsNull() *WhiteCompoundPkWrongOrderCQ {
	q.regWrongName(df.CK_ISN_C, 0)
	return q
}
func (q *WhiteCompoundPkWrongOrderCQ) SetWrongName_IsNullOrEmpty() *WhiteCompoundPkWrongOrderCQ {
	q.regWrongName(df.CK_ISNOE_C, 0)
	return q
}
func (q *WhiteCompoundPkWrongOrderCQ) SetWrongName_IsNotNull() *WhiteCompoundPkWrongOrderCQ {
	q.regWrongName(df.CK_ISNN_C, 0)
	return q
}
func (q *WhiteCompoundPkWrongOrderCQ) AddOrderBy_WrongName_Asc() *WhiteCompoundPkWrongOrderCQ {
	q.BaseConditionQuery.RegOBA("wrongName")
	return q
}
func (q *WhiteCompoundPkWrongOrderCQ) AddOrderBy_WrongName_Desc() *WhiteCompoundPkWrongOrderCQ {
	q.BaseConditionQuery.RegOBD("wrongName")
	return q
}
func (q *WhiteCompoundPkWrongOrderCQ) regWrongName(key *df.ConditionKey, value interface{}) {
	if q.WrongName == nil {
		q.WrongName = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.WrongName, "wrongName")
}

