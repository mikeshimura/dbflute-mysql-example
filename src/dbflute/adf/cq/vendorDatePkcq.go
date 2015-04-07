package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorDatePkCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	FooDate *df.ConditionValue
	FooName *df.ConditionValue
}

func (q *VendorDatePkCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *VendorDatePkCQ) getCValueFooDate() *df.ConditionValue {
	if q.FooDate == nil {
		q.FooDate = new(df.ConditionValue)
	}
	return q.FooDate
}




func (q *VendorDatePkCQ) SetFooDate_Equal(value df.Date) *VendorDatePkCQ {
	q.regFooDate(df.CK_EQ_C, value)
	return q
}


func (q *VendorDatePkCQ) SetFooDate_GreaterThan(value df.Date) *VendorDatePkCQ {
	q.regFooDate(df.CK_GT_C, value)
	return q
}

func (q *VendorDatePkCQ) SetFooDate_LessThan(value df.Date) *VendorDatePkCQ {
	q.regFooDate(df.CK_LT_C, value)
	return q
}

func (q *VendorDatePkCQ) SetFooDate_GreaterEqual(value df.Date) *VendorDatePkCQ {
	q.regFooDate(df.CK_GE_C, value)
	return q
}

func (q *VendorDatePkCQ) SetFooDate_LessEqual(value df.Date) *VendorDatePkCQ {
	q.regFooDate(df.CK_LE_C, value)
	return q
}

func (q *VendorDatePkCQ) SetFooDate_IsNull() *VendorDatePkCQ {
	q.regFooDate(df.CK_ISN_C, 0)
	return q
}
func (q *VendorDatePkCQ) SetFooDate_IsNotNull() *VendorDatePkCQ {
	q.regFooDate(df.CK_ISNN_C, 0)
	return q
}
func (q *VendorDatePkCQ) AddOrderBy_FooDate_Asc() *VendorDatePkCQ {
	q.BaseConditionQuery.RegOBA("fooDate")
	return q
}
func (q *VendorDatePkCQ) AddOrderBy_FooDate_Desc() *VendorDatePkCQ {
	q.BaseConditionQuery.RegOBD("fooDate")
	return q
}
func (q *VendorDatePkCQ) regFooDate(key *df.ConditionKey, value interface{}) {
	if q.FooDate == nil {
		q.FooDate = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.FooDate, "fooDate")
}

func (q *VendorDatePkCQ) getCValueFooName() *df.ConditionValue {
	if q.FooName == nil {
		q.FooName = new(df.ConditionValue)
	}
	return q.FooName
}


func (q *VendorDatePkCQ) SetFooName_Equal(value string) *VendorDatePkCQ {
	q.regFooName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *VendorDatePkCQ) SetFooName_NotEqual(value string) *VendorDatePkCQ {
	q.regFooName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorDatePkCQ) SetFooName_GreaterThan(value string) *VendorDatePkCQ {
	q.regFooName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorDatePkCQ) SetFooName_LessThan(value string) *VendorDatePkCQ {
	q.regFooName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorDatePkCQ) SetFooName_GreaterEqualThan(value string) *VendorDatePkCQ {
	q.regFooName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *VendorDatePkCQ) SetFooName_LessEqualThan(value string) *VendorDatePkCQ {
	q.regFooName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorDatePkCQ) SetFooName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueFooName(), "fooName", option)
}

func (q *VendorDatePkCQ) SetFooName_PrefixSearch(value string) error {
	return q.SetFooName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *VendorDatePkCQ) SetFooName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueFooName(), "fooName", option)
}



func (q *VendorDatePkCQ) AddOrderBy_FooName_Asc() *VendorDatePkCQ {
	q.BaseConditionQuery.RegOBA("fooName")
	return q
}
func (q *VendorDatePkCQ) AddOrderBy_FooName_Desc() *VendorDatePkCQ {
	q.BaseConditionQuery.RegOBD("fooName")
	return q
}
func (q *VendorDatePkCQ) regFooName(key *df.ConditionKey, value interface{}) {
	if q.FooName == nil {
		q.FooName = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.FooName, "fooName")
}

