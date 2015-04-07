package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorInheritInuCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	InuId *df.ConditionValue
	InuName *df.ConditionValue
	InuDate *df.ConditionValue
}

func (q *VendorInheritInuCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *VendorInheritInuCQ) getCValueInuId() *df.ConditionValue {
	if q.InuId == nil {
		q.InuId = new(df.ConditionValue)
	}
	return q.InuId
}



func (q *VendorInheritInuCQ) SetInuId_Equal(value int64) *VendorInheritInuCQ {
	q.regInuId(df.CK_EQ_C, value)
	return q
}

func (q *VendorInheritInuCQ) SetInuId_NotEqual(value int64) *VendorInheritInuCQ {
	q.regInuId(df.CK_NE_C, value)
	return q
}

func (q *VendorInheritInuCQ) SetInuId_GreaterThan(value int64) *VendorInheritInuCQ {
	q.regInuId(df.CK_GT_C, value)
	return q
}

func (q *VendorInheritInuCQ) SetInuId_LessThan(value int64) *VendorInheritInuCQ {
	q.regInuId(df.CK_LT_C, value)
	return q
}

func (q *VendorInheritInuCQ) SetInuId_GreaterEqual(value int64) *VendorInheritInuCQ {
	q.regInuId(df.CK_GE_C, value)
	return q
}

func (q *VendorInheritInuCQ) SetInuId_LessEqual(value int64) *VendorInheritInuCQ {
	q.regInuId(df.CK_LE_C, value)
	return q
}
func (q *VendorInheritInuCQ) SetInuId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueInuId(),"inuId",rangeOfOption)
}	


func (q *VendorInheritInuCQ) SetInuId_IsNull() *VendorInheritInuCQ {
	q.regInuId(df.CK_ISN_C, 0)
	return q
}
func (q *VendorInheritInuCQ) SetInuId_IsNotNull() *VendorInheritInuCQ {
	q.regInuId(df.CK_ISNN_C, 0)
	return q
}
func (q *VendorInheritInuCQ) AddOrderBy_InuId_Asc() *VendorInheritInuCQ {
	q.BaseConditionQuery.RegOBA("inuId")
	return q
}
func (q *VendorInheritInuCQ) AddOrderBy_InuId_Desc() *VendorInheritInuCQ {
	q.BaseConditionQuery.RegOBD("inuId")
	return q
}
func (q *VendorInheritInuCQ) regInuId(key *df.ConditionKey, value interface{}) {
	if q.InuId == nil {
		q.InuId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.InuId, "inuId")
}

func (q *VendorInheritInuCQ) getCValueInuName() *df.ConditionValue {
	if q.InuName == nil {
		q.InuName = new(df.ConditionValue)
	}
	return q.InuName
}


func (q *VendorInheritInuCQ) SetInuName_Equal(value string) *VendorInheritInuCQ {
	q.regInuName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *VendorInheritInuCQ) SetInuName_NotEqual(value string) *VendorInheritInuCQ {
	q.regInuName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorInheritInuCQ) SetInuName_GreaterThan(value string) *VendorInheritInuCQ {
	q.regInuName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorInheritInuCQ) SetInuName_LessThan(value string) *VendorInheritInuCQ {
	q.regInuName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorInheritInuCQ) SetInuName_GreaterEqualThan(value string) *VendorInheritInuCQ {
	q.regInuName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *VendorInheritInuCQ) SetInuName_LessEqualThan(value string) *VendorInheritInuCQ {
	q.regInuName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorInheritInuCQ) SetInuName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueInuName(), "inuName", option)
}

func (q *VendorInheritInuCQ) SetInuName_PrefixSearch(value string) error {
	return q.SetInuName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *VendorInheritInuCQ) SetInuName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueInuName(), "inuName", option)
}



func (q *VendorInheritInuCQ) AddOrderBy_InuName_Asc() *VendorInheritInuCQ {
	q.BaseConditionQuery.RegOBA("inuName")
	return q
}
func (q *VendorInheritInuCQ) AddOrderBy_InuName_Desc() *VendorInheritInuCQ {
	q.BaseConditionQuery.RegOBD("inuName")
	return q
}
func (q *VendorInheritInuCQ) regInuName(key *df.ConditionKey, value interface{}) {
	if q.InuName == nil {
		q.InuName = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.InuName, "inuName")
}

func (q *VendorInheritInuCQ) getCValueInuDate() *df.ConditionValue {
	if q.InuDate == nil {
		q.InuDate = new(df.ConditionValue)
	}
	return q.InuDate
}




func (q *VendorInheritInuCQ) SetInuDate_Equal(value df.Date) *VendorInheritInuCQ {
	q.regInuDate(df.CK_EQ_C, value)
	return q
}


func (q *VendorInheritInuCQ) SetInuDate_GreaterThan(value df.Date) *VendorInheritInuCQ {
	q.regInuDate(df.CK_GT_C, value)
	return q
}

func (q *VendorInheritInuCQ) SetInuDate_LessThan(value df.Date) *VendorInheritInuCQ {
	q.regInuDate(df.CK_LT_C, value)
	return q
}

func (q *VendorInheritInuCQ) SetInuDate_GreaterEqual(value df.Date) *VendorInheritInuCQ {
	q.regInuDate(df.CK_GE_C, value)
	return q
}

func (q *VendorInheritInuCQ) SetInuDate_LessEqual(value df.Date) *VendorInheritInuCQ {
	q.regInuDate(df.CK_LE_C, value)
	return q
}

func (q *VendorInheritInuCQ) SetInuDate_IsNull() *VendorInheritInuCQ {
	q.regInuDate(df.CK_ISN_C, 0)
	return q
}
func (q *VendorInheritInuCQ) SetInuDate_IsNotNull() *VendorInheritInuCQ {
	q.regInuDate(df.CK_ISNN_C, 0)
	return q
}
func (q *VendorInheritInuCQ) AddOrderBy_InuDate_Asc() *VendorInheritInuCQ {
	q.BaseConditionQuery.RegOBA("inuDate")
	return q
}
func (q *VendorInheritInuCQ) AddOrderBy_InuDate_Desc() *VendorInheritInuCQ {
	q.BaseConditionQuery.RegOBD("inuDate")
	return q
}
func (q *VendorInheritInuCQ) regInuDate(key *df.ConditionKey, value interface{}) {
	if q.InuDate == nil {
		q.InuDate = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.InuDate, "inuDate")
}

