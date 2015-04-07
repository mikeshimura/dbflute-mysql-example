package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorDateFkCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	BarId *df.ConditionValue
	BarDate *df.ConditionValue
}

func (q *VendorDateFkCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *VendorDateFkCQ) getCValueBarId() *df.ConditionValue {
	if q.BarId == nil {
		q.BarId = new(df.ConditionValue)
	}
	return q.BarId
}



func (q *VendorDateFkCQ) SetBarId_Equal(value int64) *VendorDateFkCQ {
	q.regBarId(df.CK_EQ_C, value)
	return q
}

func (q *VendorDateFkCQ) SetBarId_NotEqual(value int64) *VendorDateFkCQ {
	q.regBarId(df.CK_NE_C, value)
	return q
}

func (q *VendorDateFkCQ) SetBarId_GreaterThan(value int64) *VendorDateFkCQ {
	q.regBarId(df.CK_GT_C, value)
	return q
}

func (q *VendorDateFkCQ) SetBarId_LessThan(value int64) *VendorDateFkCQ {
	q.regBarId(df.CK_LT_C, value)
	return q
}

func (q *VendorDateFkCQ) SetBarId_GreaterEqual(value int64) *VendorDateFkCQ {
	q.regBarId(df.CK_GE_C, value)
	return q
}

func (q *VendorDateFkCQ) SetBarId_LessEqual(value int64) *VendorDateFkCQ {
	q.regBarId(df.CK_LE_C, value)
	return q
}
func (q *VendorDateFkCQ) SetBarId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueBarId(),"barId",rangeOfOption)
}	


func (q *VendorDateFkCQ) SetBarId_IsNull() *VendorDateFkCQ {
	q.regBarId(df.CK_ISN_C, 0)
	return q
}
func (q *VendorDateFkCQ) SetBarId_IsNotNull() *VendorDateFkCQ {
	q.regBarId(df.CK_ISNN_C, 0)
	return q
}
func (q *VendorDateFkCQ) AddOrderBy_BarId_Asc() *VendorDateFkCQ {
	q.BaseConditionQuery.RegOBA("barId")
	return q
}
func (q *VendorDateFkCQ) AddOrderBy_BarId_Desc() *VendorDateFkCQ {
	q.BaseConditionQuery.RegOBD("barId")
	return q
}
func (q *VendorDateFkCQ) regBarId(key *df.ConditionKey, value interface{}) {
	if q.BarId == nil {
		q.BarId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.BarId, "barId")
}

func (q *VendorDateFkCQ) getCValueBarDate() *df.ConditionValue {
	if q.BarDate == nil {
		q.BarDate = new(df.ConditionValue)
	}
	return q.BarDate
}




func (q *VendorDateFkCQ) SetBarDate_Equal(value df.Date) *VendorDateFkCQ {
	q.regBarDate(df.CK_EQ_C, value)
	return q
}


func (q *VendorDateFkCQ) SetBarDate_GreaterThan(value df.Date) *VendorDateFkCQ {
	q.regBarDate(df.CK_GT_C, value)
	return q
}

func (q *VendorDateFkCQ) SetBarDate_LessThan(value df.Date) *VendorDateFkCQ {
	q.regBarDate(df.CK_LT_C, value)
	return q
}

func (q *VendorDateFkCQ) SetBarDate_GreaterEqual(value df.Date) *VendorDateFkCQ {
	q.regBarDate(df.CK_GE_C, value)
	return q
}

func (q *VendorDateFkCQ) SetBarDate_LessEqual(value df.Date) *VendorDateFkCQ {
	q.regBarDate(df.CK_LE_C, value)
	return q
}

func (q *VendorDateFkCQ) AddOrderBy_BarDate_Asc() *VendorDateFkCQ {
	q.BaseConditionQuery.RegOBA("barDate")
	return q
}
func (q *VendorDateFkCQ) AddOrderBy_BarDate_Desc() *VendorDateFkCQ {
	q.BaseConditionQuery.RegOBD("barDate")
	return q
}
func (q *VendorDateFkCQ) regBarDate(key *df.ConditionKey, value interface{}) {
	if q.BarDate == nil {
		q.BarDate = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.BarDate, "barDate")
}

