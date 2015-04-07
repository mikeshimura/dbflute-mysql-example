package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorPartManHighCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	PartManId *df.ConditionValue
	PartManName *df.ConditionValue
	PartManPoint *df.ConditionValue
	PartManDate *df.ConditionValue
}

func (q *VendorPartManHighCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *VendorPartManHighCQ) getCValuePartManId() *df.ConditionValue {
	if q.PartManId == nil {
		q.PartManId = new(df.ConditionValue)
	}
	return q.PartManId
}



func (q *VendorPartManHighCQ) SetPartManId_Equal(value int64) *VendorPartManHighCQ {
	q.regPartManId(df.CK_EQ_C, value)
	return q
}

func (q *VendorPartManHighCQ) SetPartManId_NotEqual(value int64) *VendorPartManHighCQ {
	q.regPartManId(df.CK_NE_C, value)
	return q
}

func (q *VendorPartManHighCQ) SetPartManId_GreaterThan(value int64) *VendorPartManHighCQ {
	q.regPartManId(df.CK_GT_C, value)
	return q
}

func (q *VendorPartManHighCQ) SetPartManId_LessThan(value int64) *VendorPartManHighCQ {
	q.regPartManId(df.CK_LT_C, value)
	return q
}

func (q *VendorPartManHighCQ) SetPartManId_GreaterEqual(value int64) *VendorPartManHighCQ {
	q.regPartManId(df.CK_GE_C, value)
	return q
}

func (q *VendorPartManHighCQ) SetPartManId_LessEqual(value int64) *VendorPartManHighCQ {
	q.regPartManId(df.CK_LE_C, value)
	return q
}
func (q *VendorPartManHighCQ) SetPartManId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValuePartManId(),"partManId",rangeOfOption)
}	


func (q *VendorPartManHighCQ) SetPartManId_IsNull() *VendorPartManHighCQ {
	q.regPartManId(df.CK_ISN_C, 0)
	return q
}
func (q *VendorPartManHighCQ) SetPartManId_IsNotNull() *VendorPartManHighCQ {
	q.regPartManId(df.CK_ISNN_C, 0)
	return q
}
func (q *VendorPartManHighCQ) AddOrderBy_PartManId_Asc() *VendorPartManHighCQ {
	q.BaseConditionQuery.RegOBA("partManId")
	return q
}
func (q *VendorPartManHighCQ) AddOrderBy_PartManId_Desc() *VendorPartManHighCQ {
	q.BaseConditionQuery.RegOBD("partManId")
	return q
}
func (q *VendorPartManHighCQ) regPartManId(key *df.ConditionKey, value interface{}) {
	if q.PartManId == nil {
		q.PartManId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PartManId, "partManId")
}

func (q *VendorPartManHighCQ) getCValuePartManName() *df.ConditionValue {
	if q.PartManName == nil {
		q.PartManName = new(df.ConditionValue)
	}
	return q.PartManName
}


func (q *VendorPartManHighCQ) SetPartManName_Equal(value string) *VendorPartManHighCQ {
	q.regPartManName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *VendorPartManHighCQ) SetPartManName_NotEqual(value string) *VendorPartManHighCQ {
	q.regPartManName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorPartManHighCQ) SetPartManName_GreaterThan(value string) *VendorPartManHighCQ {
	q.regPartManName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorPartManHighCQ) SetPartManName_LessThan(value string) *VendorPartManHighCQ {
	q.regPartManName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorPartManHighCQ) SetPartManName_GreaterEqualThan(value string) *VendorPartManHighCQ {
	q.regPartManName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *VendorPartManHighCQ) SetPartManName_LessEqualThan(value string) *VendorPartManHighCQ {
	q.regPartManName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorPartManHighCQ) SetPartManName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValuePartManName(), "partManName", option)
}

func (q *VendorPartManHighCQ) SetPartManName_PrefixSearch(value string) error {
	return q.SetPartManName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *VendorPartManHighCQ) SetPartManName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValuePartManName(), "partManName", option)
}



func (q *VendorPartManHighCQ) AddOrderBy_PartManName_Asc() *VendorPartManHighCQ {
	q.BaseConditionQuery.RegOBA("partManName")
	return q
}
func (q *VendorPartManHighCQ) AddOrderBy_PartManName_Desc() *VendorPartManHighCQ {
	q.BaseConditionQuery.RegOBD("partManName")
	return q
}
func (q *VendorPartManHighCQ) regPartManName(key *df.ConditionKey, value interface{}) {
	if q.PartManName == nil {
		q.PartManName = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PartManName, "partManName")
}

func (q *VendorPartManHighCQ) getCValuePartManPoint() *df.ConditionValue {
	if q.PartManPoint == nil {
		q.PartManPoint = new(df.ConditionValue)
	}
	return q.PartManPoint
}



func (q *VendorPartManHighCQ) SetPartManPoint_Equal(value int64) *VendorPartManHighCQ {
	q.regPartManPoint(df.CK_EQ_C, value)
	return q
}

func (q *VendorPartManHighCQ) SetPartManPoint_NotEqual(value int64) *VendorPartManHighCQ {
	q.regPartManPoint(df.CK_NE_C, value)
	return q
}

func (q *VendorPartManHighCQ) SetPartManPoint_GreaterThan(value int64) *VendorPartManHighCQ {
	q.regPartManPoint(df.CK_GT_C, value)
	return q
}

func (q *VendorPartManHighCQ) SetPartManPoint_LessThan(value int64) *VendorPartManHighCQ {
	q.regPartManPoint(df.CK_LT_C, value)
	return q
}

func (q *VendorPartManHighCQ) SetPartManPoint_GreaterEqual(value int64) *VendorPartManHighCQ {
	q.regPartManPoint(df.CK_GE_C, value)
	return q
}

func (q *VendorPartManHighCQ) SetPartManPoint_LessEqual(value int64) *VendorPartManHighCQ {
	q.regPartManPoint(df.CK_LE_C, value)
	return q
}
func (q *VendorPartManHighCQ) SetPartManPoint_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValuePartManPoint(),"partManPoint",rangeOfOption)
}	


func (q *VendorPartManHighCQ) AddOrderBy_PartManPoint_Asc() *VendorPartManHighCQ {
	q.BaseConditionQuery.RegOBA("partManPoint")
	return q
}
func (q *VendorPartManHighCQ) AddOrderBy_PartManPoint_Desc() *VendorPartManHighCQ {
	q.BaseConditionQuery.RegOBD("partManPoint")
	return q
}
func (q *VendorPartManHighCQ) regPartManPoint(key *df.ConditionKey, value interface{}) {
	if q.PartManPoint == nil {
		q.PartManPoint = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PartManPoint, "partManPoint")
}

func (q *VendorPartManHighCQ) getCValuePartManDate() *df.ConditionValue {
	if q.PartManDate == nil {
		q.PartManDate = new(df.ConditionValue)
	}
	return q.PartManDate
}




func (q *VendorPartManHighCQ) SetPartManDate_Equal(value df.Date) *VendorPartManHighCQ {
	q.regPartManDate(df.CK_EQ_C, value)
	return q
}


func (q *VendorPartManHighCQ) SetPartManDate_GreaterThan(value df.Date) *VendorPartManHighCQ {
	q.regPartManDate(df.CK_GT_C, value)
	return q
}

func (q *VendorPartManHighCQ) SetPartManDate_LessThan(value df.Date) *VendorPartManHighCQ {
	q.regPartManDate(df.CK_LT_C, value)
	return q
}

func (q *VendorPartManHighCQ) SetPartManDate_GreaterEqual(value df.Date) *VendorPartManHighCQ {
	q.regPartManDate(df.CK_GE_C, value)
	return q
}

func (q *VendorPartManHighCQ) SetPartManDate_LessEqual(value df.Date) *VendorPartManHighCQ {
	q.regPartManDate(df.CK_LE_C, value)
	return q
}

func (q *VendorPartManHighCQ) SetPartManDate_IsNull() *VendorPartManHighCQ {
	q.regPartManDate(df.CK_ISN_C, 0)
	return q
}
func (q *VendorPartManHighCQ) SetPartManDate_IsNotNull() *VendorPartManHighCQ {
	q.regPartManDate(df.CK_ISNN_C, 0)
	return q
}
func (q *VendorPartManHighCQ) AddOrderBy_PartManDate_Asc() *VendorPartManHighCQ {
	q.BaseConditionQuery.RegOBA("partManDate")
	return q
}
func (q *VendorPartManHighCQ) AddOrderBy_PartManDate_Desc() *VendorPartManHighCQ {
	q.BaseConditionQuery.RegOBD("partManDate")
	return q
}
func (q *VendorPartManHighCQ) regPartManDate(key *df.ConditionKey, value interface{}) {
	if q.PartManDate == nil {
		q.PartManDate = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PartManDate, "partManDate")
}

