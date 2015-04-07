package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type ProductStatusCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	ProductStatusCode *df.ConditionValue
	ProductStatusName *df.ConditionValue
	DisplayOrder *df.ConditionValue
}

func (q *ProductStatusCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *ProductStatusCQ) getCValueProductStatusCode() *df.ConditionValue {
	if q.ProductStatusCode == nil {
		q.ProductStatusCode = new(df.ConditionValue)
	}
	return q.ProductStatusCode
}


func (q *ProductStatusCQ) SetProductStatusCode_Equal(value string) *ProductStatusCQ {
	q.regProductStatusCode(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *ProductStatusCQ) SetProductStatusCode_NotEqual(value string) *ProductStatusCQ {
	q.regProductStatusCode(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductStatusCQ) SetProductStatusCode_GreaterThan(value string) *ProductStatusCQ {
	q.regProductStatusCode(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductStatusCQ) SetProductStatusCode_LessThan(value string) *ProductStatusCQ {
	q.regProductStatusCode(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductStatusCQ) SetProductStatusCode_GreaterEqualThan(value string) *ProductStatusCQ {
	q.regProductStatusCode(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *ProductStatusCQ) SetProductStatusCode_LessEqualThan(value string) *ProductStatusCQ {
	q.regProductStatusCode(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductStatusCQ) SetProductStatusCode_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueProductStatusCode(), "productStatusCode", option)
}

func (q *ProductStatusCQ) SetProductStatusCode_PrefixSearch(value string) error {
	return q.SetProductStatusCode_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *ProductStatusCQ) SetProductStatusCode_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueProductStatusCode(), "productStatusCode", option)
}



func (q *ProductStatusCQ) SetProductStatusCode_IsNull() *ProductStatusCQ {
	q.regProductStatusCode(df.CK_ISN_C, 0)
	return q
}
func (q *ProductStatusCQ) SetProductStatusCode_IsNotNull() *ProductStatusCQ {
	q.regProductStatusCode(df.CK_ISNN_C, 0)
	return q
}
func (q *ProductStatusCQ) AddOrderBy_ProductStatusCode_Asc() *ProductStatusCQ {
	q.BaseConditionQuery.RegOBA("productStatusCode")
	return q
}
func (q *ProductStatusCQ) AddOrderBy_ProductStatusCode_Desc() *ProductStatusCQ {
	q.BaseConditionQuery.RegOBD("productStatusCode")
	return q
}
func (q *ProductStatusCQ) regProductStatusCode(key *df.ConditionKey, value interface{}) {
	if q.ProductStatusCode == nil {
		q.ProductStatusCode = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ProductStatusCode, "productStatusCode")
}

func (q *ProductStatusCQ) getCValueProductStatusName() *df.ConditionValue {
	if q.ProductStatusName == nil {
		q.ProductStatusName = new(df.ConditionValue)
	}
	return q.ProductStatusName
}


func (q *ProductStatusCQ) SetProductStatusName_Equal(value string) *ProductStatusCQ {
	q.regProductStatusName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *ProductStatusCQ) SetProductStatusName_NotEqual(value string) *ProductStatusCQ {
	q.regProductStatusName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductStatusCQ) SetProductStatusName_GreaterThan(value string) *ProductStatusCQ {
	q.regProductStatusName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductStatusCQ) SetProductStatusName_LessThan(value string) *ProductStatusCQ {
	q.regProductStatusName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductStatusCQ) SetProductStatusName_GreaterEqualThan(value string) *ProductStatusCQ {
	q.regProductStatusName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *ProductStatusCQ) SetProductStatusName_LessEqualThan(value string) *ProductStatusCQ {
	q.regProductStatusName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductStatusCQ) SetProductStatusName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueProductStatusName(), "productStatusName", option)
}

func (q *ProductStatusCQ) SetProductStatusName_PrefixSearch(value string) error {
	return q.SetProductStatusName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *ProductStatusCQ) SetProductStatusName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueProductStatusName(), "productStatusName", option)
}



func (q *ProductStatusCQ) AddOrderBy_ProductStatusName_Asc() *ProductStatusCQ {
	q.BaseConditionQuery.RegOBA("productStatusName")
	return q
}
func (q *ProductStatusCQ) AddOrderBy_ProductStatusName_Desc() *ProductStatusCQ {
	q.BaseConditionQuery.RegOBD("productStatusName")
	return q
}
func (q *ProductStatusCQ) regProductStatusName(key *df.ConditionKey, value interface{}) {
	if q.ProductStatusName == nil {
		q.ProductStatusName = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ProductStatusName, "productStatusName")
}

func (q *ProductStatusCQ) getCValueDisplayOrder() *df.ConditionValue {
	if q.DisplayOrder == nil {
		q.DisplayOrder = new(df.ConditionValue)
	}
	return q.DisplayOrder
}



func (q *ProductStatusCQ) SetDisplayOrder_Equal(value int64) *ProductStatusCQ {
	q.regDisplayOrder(df.CK_EQ_C, value)
	return q
}

func (q *ProductStatusCQ) SetDisplayOrder_NotEqual(value int64) *ProductStatusCQ {
	q.regDisplayOrder(df.CK_NE_C, value)
	return q
}

func (q *ProductStatusCQ) SetDisplayOrder_GreaterThan(value int64) *ProductStatusCQ {
	q.regDisplayOrder(df.CK_GT_C, value)
	return q
}

func (q *ProductStatusCQ) SetDisplayOrder_LessThan(value int64) *ProductStatusCQ {
	q.regDisplayOrder(df.CK_LT_C, value)
	return q
}

func (q *ProductStatusCQ) SetDisplayOrder_GreaterEqual(value int64) *ProductStatusCQ {
	q.regDisplayOrder(df.CK_GE_C, value)
	return q
}

func (q *ProductStatusCQ) SetDisplayOrder_LessEqual(value int64) *ProductStatusCQ {
	q.regDisplayOrder(df.CK_LE_C, value)
	return q
}
func (q *ProductStatusCQ) SetDisplayOrder_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueDisplayOrder(),"displayOrder",rangeOfOption)
}	


func (q *ProductStatusCQ) AddOrderBy_DisplayOrder_Asc() *ProductStatusCQ {
	q.BaseConditionQuery.RegOBA("displayOrder")
	return q
}
func (q *ProductStatusCQ) AddOrderBy_DisplayOrder_Desc() *ProductStatusCQ {
	q.BaseConditionQuery.RegOBD("displayOrder")
	return q
}
func (q *ProductStatusCQ) regDisplayOrder(key *df.ConditionKey, value interface{}) {
	if q.DisplayOrder == nil {
		q.DisplayOrder = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.DisplayOrder, "displayOrder")
}

