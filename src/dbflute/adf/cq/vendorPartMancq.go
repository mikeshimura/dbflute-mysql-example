package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorPartManCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	PartManId *df.ConditionValue
	PartManName *df.ConditionValue
	PartManPoint *df.ConditionValue
	PartManDate *df.ConditionValue
}

func (q *VendorPartManCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *VendorPartManCQ) getCValuePartManId() *df.ConditionValue {
	if q.PartManId == nil {
		q.PartManId = new(df.ConditionValue)
	}
	return q.PartManId
}



func (q *VendorPartManCQ) SetPartManId_Equal(value int64) *VendorPartManCQ {
	q.regPartManId(df.CK_EQ_C, value)
	return q
}

func (q *VendorPartManCQ) SetPartManId_NotEqual(value int64) *VendorPartManCQ {
	q.regPartManId(df.CK_NE_C, value)
	return q
}

func (q *VendorPartManCQ) SetPartManId_GreaterThan(value int64) *VendorPartManCQ {
	q.regPartManId(df.CK_GT_C, value)
	return q
}

func (q *VendorPartManCQ) SetPartManId_LessThan(value int64) *VendorPartManCQ {
	q.regPartManId(df.CK_LT_C, value)
	return q
}

func (q *VendorPartManCQ) SetPartManId_GreaterEqual(value int64) *VendorPartManCQ {
	q.regPartManId(df.CK_GE_C, value)
	return q
}

func (q *VendorPartManCQ) SetPartManId_LessEqual(value int64) *VendorPartManCQ {
	q.regPartManId(df.CK_LE_C, value)
	return q
}
func (q *VendorPartManCQ) SetPartManId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValuePartManId(),"partManId",rangeOfOption)
}	


func (q *VendorPartManCQ) SetPartManId_IsNull() *VendorPartManCQ {
	q.regPartManId(df.CK_ISN_C, 0)
	return q
}
func (q *VendorPartManCQ) SetPartManId_IsNotNull() *VendorPartManCQ {
	q.regPartManId(df.CK_ISNN_C, 0)
	return q
}
func (q *VendorPartManCQ) AddOrderBy_PartManId_Asc() *VendorPartManCQ {
	q.BaseConditionQuery.RegOBA("partManId")
	return q
}
func (q *VendorPartManCQ) AddOrderBy_PartManId_Desc() *VendorPartManCQ {
	q.BaseConditionQuery.RegOBD("partManId")
	return q
}
func (q *VendorPartManCQ) regPartManId(key *df.ConditionKey, value interface{}) {
	if q.PartManId == nil {
		q.PartManId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PartManId, "partManId")
}

func (q *VendorPartManCQ) getCValuePartManName() *df.ConditionValue {
	if q.PartManName == nil {
		q.PartManName = new(df.ConditionValue)
	}
	return q.PartManName
}


func (q *VendorPartManCQ) SetPartManName_Equal(value string) *VendorPartManCQ {
	q.regPartManName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *VendorPartManCQ) SetPartManName_NotEqual(value string) *VendorPartManCQ {
	q.regPartManName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorPartManCQ) SetPartManName_GreaterThan(value string) *VendorPartManCQ {
	q.regPartManName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorPartManCQ) SetPartManName_LessThan(value string) *VendorPartManCQ {
	q.regPartManName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorPartManCQ) SetPartManName_GreaterEqualThan(value string) *VendorPartManCQ {
	q.regPartManName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *VendorPartManCQ) SetPartManName_LessEqualThan(value string) *VendorPartManCQ {
	q.regPartManName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorPartManCQ) SetPartManName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValuePartManName(), "partManName", option)
}

func (q *VendorPartManCQ) SetPartManName_PrefixSearch(value string) error {
	return q.SetPartManName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *VendorPartManCQ) SetPartManName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValuePartManName(), "partManName", option)
}



func (q *VendorPartManCQ) AddOrderBy_PartManName_Asc() *VendorPartManCQ {
	q.BaseConditionQuery.RegOBA("partManName")
	return q
}
func (q *VendorPartManCQ) AddOrderBy_PartManName_Desc() *VendorPartManCQ {
	q.BaseConditionQuery.RegOBD("partManName")
	return q
}
func (q *VendorPartManCQ) regPartManName(key *df.ConditionKey, value interface{}) {
	if q.PartManName == nil {
		q.PartManName = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PartManName, "partManName")
}

func (q *VendorPartManCQ) getCValuePartManPoint() *df.ConditionValue {
	if q.PartManPoint == nil {
		q.PartManPoint = new(df.ConditionValue)
	}
	return q.PartManPoint
}



func (q *VendorPartManCQ) SetPartManPoint_Equal(value int64) *VendorPartManCQ {
	q.regPartManPoint(df.CK_EQ_C, value)
	return q
}

func (q *VendorPartManCQ) SetPartManPoint_NotEqual(value int64) *VendorPartManCQ {
	q.regPartManPoint(df.CK_NE_C, value)
	return q
}

func (q *VendorPartManCQ) SetPartManPoint_GreaterThan(value int64) *VendorPartManCQ {
	q.regPartManPoint(df.CK_GT_C, value)
	return q
}

func (q *VendorPartManCQ) SetPartManPoint_LessThan(value int64) *VendorPartManCQ {
	q.regPartManPoint(df.CK_LT_C, value)
	return q
}

func (q *VendorPartManCQ) SetPartManPoint_GreaterEqual(value int64) *VendorPartManCQ {
	q.regPartManPoint(df.CK_GE_C, value)
	return q
}

func (q *VendorPartManCQ) SetPartManPoint_LessEqual(value int64) *VendorPartManCQ {
	q.regPartManPoint(df.CK_LE_C, value)
	return q
}
func (q *VendorPartManCQ) SetPartManPoint_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValuePartManPoint(),"partManPoint",rangeOfOption)
}	


func (q *VendorPartManCQ) AddOrderBy_PartManPoint_Asc() *VendorPartManCQ {
	q.BaseConditionQuery.RegOBA("partManPoint")
	return q
}
func (q *VendorPartManCQ) AddOrderBy_PartManPoint_Desc() *VendorPartManCQ {
	q.BaseConditionQuery.RegOBD("partManPoint")
	return q
}
func (q *VendorPartManCQ) regPartManPoint(key *df.ConditionKey, value interface{}) {
	if q.PartManPoint == nil {
		q.PartManPoint = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PartManPoint, "partManPoint")
}

func (q *VendorPartManCQ) getCValuePartManDate() *df.ConditionValue {
	if q.PartManDate == nil {
		q.PartManDate = new(df.ConditionValue)
	}
	return q.PartManDate
}




func (q *VendorPartManCQ) SetPartManDate_Equal(value df.Date) *VendorPartManCQ {
	q.regPartManDate(df.CK_EQ_C, value)
	return q
}


func (q *VendorPartManCQ) SetPartManDate_GreaterThan(value df.Date) *VendorPartManCQ {
	q.regPartManDate(df.CK_GT_C, value)
	return q
}

func (q *VendorPartManCQ) SetPartManDate_LessThan(value df.Date) *VendorPartManCQ {
	q.regPartManDate(df.CK_LT_C, value)
	return q
}

func (q *VendorPartManCQ) SetPartManDate_GreaterEqual(value df.Date) *VendorPartManCQ {
	q.regPartManDate(df.CK_GE_C, value)
	return q
}

func (q *VendorPartManCQ) SetPartManDate_LessEqual(value df.Date) *VendorPartManCQ {
	q.regPartManDate(df.CK_LE_C, value)
	return q
}

func (q *VendorPartManCQ) SetPartManDate_IsNull() *VendorPartManCQ {
	q.regPartManDate(df.CK_ISN_C, 0)
	return q
}
func (q *VendorPartManCQ) SetPartManDate_IsNotNull() *VendorPartManCQ {
	q.regPartManDate(df.CK_ISNN_C, 0)
	return q
}
func (q *VendorPartManCQ) AddOrderBy_PartManDate_Asc() *VendorPartManCQ {
	q.BaseConditionQuery.RegOBA("partManDate")
	return q
}
func (q *VendorPartManCQ) AddOrderBy_PartManDate_Desc() *VendorPartManCQ {
	q.BaseConditionQuery.RegOBD("partManDate")
	return q
}
func (q *VendorPartManCQ) regPartManDate(key *df.ConditionKey, value interface{}) {
	if q.PartManDate == nil {
		q.PartManDate = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PartManDate, "partManDate")
}

