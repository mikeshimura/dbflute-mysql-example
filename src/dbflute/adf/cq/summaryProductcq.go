package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type SummaryProductCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	ProductId *df.ConditionValue
	ProductName *df.ConditionValue
	ProductHandleCode *df.ConditionValue
	ProductStatusCode *df.ConditionValue
	LatestPurchaseDatetime *df.ConditionValue
}

func (q *SummaryProductCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *SummaryProductCQ) getCValueProductId() *df.ConditionValue {
	if q.ProductId == nil {
		q.ProductId = new(df.ConditionValue)
	}
	return q.ProductId
}



func (q *SummaryProductCQ) SetProductId_Equal(value int64) *SummaryProductCQ {
	q.regProductId(df.CK_EQ_C, value)
	return q
}

func (q *SummaryProductCQ) SetProductId_NotEqual(value int64) *SummaryProductCQ {
	q.regProductId(df.CK_NE_C, value)
	return q
}

func (q *SummaryProductCQ) SetProductId_GreaterThan(value int64) *SummaryProductCQ {
	q.regProductId(df.CK_GT_C, value)
	return q
}

func (q *SummaryProductCQ) SetProductId_LessThan(value int64) *SummaryProductCQ {
	q.regProductId(df.CK_LT_C, value)
	return q
}

func (q *SummaryProductCQ) SetProductId_GreaterEqual(value int64) *SummaryProductCQ {
	q.regProductId(df.CK_GE_C, value)
	return q
}

func (q *SummaryProductCQ) SetProductId_LessEqual(value int64) *SummaryProductCQ {
	q.regProductId(df.CK_LE_C, value)
	return q
}
func (q *SummaryProductCQ) SetProductId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueProductId(),"productId",rangeOfOption)
}	


func (q *SummaryProductCQ) AddOrderBy_ProductId_Asc() *SummaryProductCQ {
	q.BaseConditionQuery.RegOBA("productId")
	return q
}
func (q *SummaryProductCQ) AddOrderBy_ProductId_Desc() *SummaryProductCQ {
	q.BaseConditionQuery.RegOBD("productId")
	return q
}
func (q *SummaryProductCQ) regProductId(key *df.ConditionKey, value interface{}) {
	if q.ProductId == nil {
		q.ProductId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ProductId, "productId")
}

func (q *SummaryProductCQ) getCValueProductName() *df.ConditionValue {
	if q.ProductName == nil {
		q.ProductName = new(df.ConditionValue)
	}
	return q.ProductName
}


func (q *SummaryProductCQ) SetProductName_Equal(value string) *SummaryProductCQ {
	q.regProductName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *SummaryProductCQ) SetProductName_NotEqual(value string) *SummaryProductCQ {
	q.regProductName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryProductCQ) SetProductName_GreaterThan(value string) *SummaryProductCQ {
	q.regProductName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryProductCQ) SetProductName_LessThan(value string) *SummaryProductCQ {
	q.regProductName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryProductCQ) SetProductName_GreaterEqualThan(value string) *SummaryProductCQ {
	q.regProductName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SummaryProductCQ) SetProductName_LessEqualThan(value string) *SummaryProductCQ {
	q.regProductName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryProductCQ) SetProductName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueProductName(), "productName", option)
}

func (q *SummaryProductCQ) SetProductName_PrefixSearch(value string) error {
	return q.SetProductName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SummaryProductCQ) SetProductName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueProductName(), "productName", option)
}



func (q *SummaryProductCQ) AddOrderBy_ProductName_Asc() *SummaryProductCQ {
	q.BaseConditionQuery.RegOBA("productName")
	return q
}
func (q *SummaryProductCQ) AddOrderBy_ProductName_Desc() *SummaryProductCQ {
	q.BaseConditionQuery.RegOBD("productName")
	return q
}
func (q *SummaryProductCQ) regProductName(key *df.ConditionKey, value interface{}) {
	if q.ProductName == nil {
		q.ProductName = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ProductName, "productName")
}

func (q *SummaryProductCQ) getCValueProductHandleCode() *df.ConditionValue {
	if q.ProductHandleCode == nil {
		q.ProductHandleCode = new(df.ConditionValue)
	}
	return q.ProductHandleCode
}


func (q *SummaryProductCQ) SetProductHandleCode_Equal(value string) *SummaryProductCQ {
	q.regProductHandleCode(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *SummaryProductCQ) SetProductHandleCode_NotEqual(value string) *SummaryProductCQ {
	q.regProductHandleCode(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryProductCQ) SetProductHandleCode_GreaterThan(value string) *SummaryProductCQ {
	q.regProductHandleCode(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryProductCQ) SetProductHandleCode_LessThan(value string) *SummaryProductCQ {
	q.regProductHandleCode(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryProductCQ) SetProductHandleCode_GreaterEqualThan(value string) *SummaryProductCQ {
	q.regProductHandleCode(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SummaryProductCQ) SetProductHandleCode_LessEqualThan(value string) *SummaryProductCQ {
	q.regProductHandleCode(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryProductCQ) SetProductHandleCode_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueProductHandleCode(), "productHandleCode", option)
}

func (q *SummaryProductCQ) SetProductHandleCode_PrefixSearch(value string) error {
	return q.SetProductHandleCode_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SummaryProductCQ) SetProductHandleCode_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueProductHandleCode(), "productHandleCode", option)
}



func (q *SummaryProductCQ) AddOrderBy_ProductHandleCode_Asc() *SummaryProductCQ {
	q.BaseConditionQuery.RegOBA("productHandleCode")
	return q
}
func (q *SummaryProductCQ) AddOrderBy_ProductHandleCode_Desc() *SummaryProductCQ {
	q.BaseConditionQuery.RegOBD("productHandleCode")
	return q
}
func (q *SummaryProductCQ) regProductHandleCode(key *df.ConditionKey, value interface{}) {
	if q.ProductHandleCode == nil {
		q.ProductHandleCode = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ProductHandleCode, "productHandleCode")
}

func (q *SummaryProductCQ) getCValueProductStatusCode() *df.ConditionValue {
	if q.ProductStatusCode == nil {
		q.ProductStatusCode = new(df.ConditionValue)
	}
	return q.ProductStatusCode
}


func (q *SummaryProductCQ) SetProductStatusCode_Equal(value string) *SummaryProductCQ {
	q.regProductStatusCode(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *SummaryProductCQ) SetProductStatusCode_NotEqual(value string) *SummaryProductCQ {
	q.regProductStatusCode(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryProductCQ) SetProductStatusCode_GreaterThan(value string) *SummaryProductCQ {
	q.regProductStatusCode(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryProductCQ) SetProductStatusCode_LessThan(value string) *SummaryProductCQ {
	q.regProductStatusCode(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryProductCQ) SetProductStatusCode_GreaterEqualThan(value string) *SummaryProductCQ {
	q.regProductStatusCode(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SummaryProductCQ) SetProductStatusCode_LessEqualThan(value string) *SummaryProductCQ {
	q.regProductStatusCode(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryProductCQ) SetProductStatusCode_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueProductStatusCode(), "productStatusCode", option)
}

func (q *SummaryProductCQ) SetProductStatusCode_PrefixSearch(value string) error {
	return q.SetProductStatusCode_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SummaryProductCQ) SetProductStatusCode_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueProductStatusCode(), "productStatusCode", option)
}



func (q *SummaryProductCQ) AddOrderBy_ProductStatusCode_Asc() *SummaryProductCQ {
	q.BaseConditionQuery.RegOBA("productStatusCode")
	return q
}
func (q *SummaryProductCQ) AddOrderBy_ProductStatusCode_Desc() *SummaryProductCQ {
	q.BaseConditionQuery.RegOBD("productStatusCode")
	return q
}
func (q *SummaryProductCQ) regProductStatusCode(key *df.ConditionKey, value interface{}) {
	if q.ProductStatusCode == nil {
		q.ProductStatusCode = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ProductStatusCode, "productStatusCode")
}

func (q *SummaryProductCQ) getCValueLatestPurchaseDatetime() *df.ConditionValue {
	if q.LatestPurchaseDatetime == nil {
		q.LatestPurchaseDatetime = new(df.ConditionValue)
	}
	return q.LatestPurchaseDatetime
}




func (q *SummaryProductCQ) SetLatestPurchaseDatetime_Equal(value df.MysqlTimestamp) *SummaryProductCQ {
	q.regLatestPurchaseDatetime(df.CK_EQ_C, value)
	return q
}


func (q *SummaryProductCQ) SetLatestPurchaseDatetime_GreaterThan(value df.MysqlTimestamp) *SummaryProductCQ {
	q.regLatestPurchaseDatetime(df.CK_GT_C, value)
	return q
}

func (q *SummaryProductCQ) SetLatestPurchaseDatetime_LessThan(value df.MysqlTimestamp) *SummaryProductCQ {
	q.regLatestPurchaseDatetime(df.CK_LT_C, value)
	return q
}

func (q *SummaryProductCQ) SetLatestPurchaseDatetime_GreaterEqual(value df.MysqlTimestamp) *SummaryProductCQ {
	q.regLatestPurchaseDatetime(df.CK_GE_C, value)
	return q
}

func (q *SummaryProductCQ) SetLatestPurchaseDatetime_LessEqual(value df.MysqlTimestamp) *SummaryProductCQ {
	q.regLatestPurchaseDatetime(df.CK_LE_C, value)
	return q
}

func (q *SummaryProductCQ) SetLatestPurchaseDatetime_IsNull() *SummaryProductCQ {
	q.regLatestPurchaseDatetime(df.CK_ISN_C, 0)
	return q
}
func (q *SummaryProductCQ) SetLatestPurchaseDatetime_IsNotNull() *SummaryProductCQ {
	q.regLatestPurchaseDatetime(df.CK_ISNN_C, 0)
	return q
}
func (q *SummaryProductCQ) AddOrderBy_LatestPurchaseDatetime_Asc() *SummaryProductCQ {
	q.BaseConditionQuery.RegOBA("latestPurchaseDatetime")
	return q
}
func (q *SummaryProductCQ) AddOrderBy_LatestPurchaseDatetime_Desc() *SummaryProductCQ {
	q.BaseConditionQuery.RegOBD("latestPurchaseDatetime")
	return q
}
func (q *SummaryProductCQ) regLatestPurchaseDatetime(key *df.ConditionKey, value interface{}) {
	if q.LatestPurchaseDatetime == nil {
		q.LatestPurchaseDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.LatestPurchaseDatetime, "latestPurchaseDatetime")
}

