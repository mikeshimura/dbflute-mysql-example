package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorInheritNekoCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	InuId *df.ConditionValue
	InuName *df.ConditionValue
	InuDate *df.ConditionValue
	NekoId *df.ConditionValue
	NekoName *df.ConditionValue
	NekoDate *df.ConditionValue
}

func (q *VendorInheritNekoCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *VendorInheritNekoCQ) getCValueInuId() *df.ConditionValue {
	if q.InuId == nil {
		q.InuId = new(df.ConditionValue)
	}
	return q.InuId
}



func (q *VendorInheritNekoCQ) SetInuId_Equal(value int64) *VendorInheritNekoCQ {
	q.regInuId(df.CK_EQ_C, value)
	return q
}

func (q *VendorInheritNekoCQ) SetInuId_NotEqual(value int64) *VendorInheritNekoCQ {
	q.regInuId(df.CK_NE_C, value)
	return q
}

func (q *VendorInheritNekoCQ) SetInuId_GreaterThan(value int64) *VendorInheritNekoCQ {
	q.regInuId(df.CK_GT_C, value)
	return q
}

func (q *VendorInheritNekoCQ) SetInuId_LessThan(value int64) *VendorInheritNekoCQ {
	q.regInuId(df.CK_LT_C, value)
	return q
}

func (q *VendorInheritNekoCQ) SetInuId_GreaterEqual(value int64) *VendorInheritNekoCQ {
	q.regInuId(df.CK_GE_C, value)
	return q
}

func (q *VendorInheritNekoCQ) SetInuId_LessEqual(value int64) *VendorInheritNekoCQ {
	q.regInuId(df.CK_LE_C, value)
	return q
}
func (q *VendorInheritNekoCQ) SetInuId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueInuId(),"inuId",rangeOfOption)
}	


func (q *VendorInheritNekoCQ) AddOrderBy_InuId_Asc() *VendorInheritNekoCQ {
	q.BaseConditionQuery.RegOBA("inuId")
	return q
}
func (q *VendorInheritNekoCQ) AddOrderBy_InuId_Desc() *VendorInheritNekoCQ {
	q.BaseConditionQuery.RegOBD("inuId")
	return q
}
func (q *VendorInheritNekoCQ) regInuId(key *df.ConditionKey, value interface{}) {
	if q.InuId == nil {
		q.InuId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.InuId, "inuId")
}

func (q *VendorInheritNekoCQ) getCValueInuName() *df.ConditionValue {
	if q.InuName == nil {
		q.InuName = new(df.ConditionValue)
	}
	return q.InuName
}


func (q *VendorInheritNekoCQ) SetInuName_Equal(value string) *VendorInheritNekoCQ {
	q.regInuName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *VendorInheritNekoCQ) SetInuName_NotEqual(value string) *VendorInheritNekoCQ {
	q.regInuName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorInheritNekoCQ) SetInuName_GreaterThan(value string) *VendorInheritNekoCQ {
	q.regInuName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorInheritNekoCQ) SetInuName_LessThan(value string) *VendorInheritNekoCQ {
	q.regInuName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorInheritNekoCQ) SetInuName_GreaterEqualThan(value string) *VendorInheritNekoCQ {
	q.regInuName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *VendorInheritNekoCQ) SetInuName_LessEqualThan(value string) *VendorInheritNekoCQ {
	q.regInuName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorInheritNekoCQ) SetInuName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueInuName(), "inuName", option)
}

func (q *VendorInheritNekoCQ) SetInuName_PrefixSearch(value string) error {
	return q.SetInuName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *VendorInheritNekoCQ) SetInuName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueInuName(), "inuName", option)
}



func (q *VendorInheritNekoCQ) AddOrderBy_InuName_Asc() *VendorInheritNekoCQ {
	q.BaseConditionQuery.RegOBA("inuName")
	return q
}
func (q *VendorInheritNekoCQ) AddOrderBy_InuName_Desc() *VendorInheritNekoCQ {
	q.BaseConditionQuery.RegOBD("inuName")
	return q
}
func (q *VendorInheritNekoCQ) regInuName(key *df.ConditionKey, value interface{}) {
	if q.InuName == nil {
		q.InuName = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.InuName, "inuName")
}

func (q *VendorInheritNekoCQ) getCValueInuDate() *df.ConditionValue {
	if q.InuDate == nil {
		q.InuDate = new(df.ConditionValue)
	}
	return q.InuDate
}




func (q *VendorInheritNekoCQ) SetInuDate_Equal(value df.Date) *VendorInheritNekoCQ {
	q.regInuDate(df.CK_EQ_C, value)
	return q
}


func (q *VendorInheritNekoCQ) SetInuDate_GreaterThan(value df.Date) *VendorInheritNekoCQ {
	q.regInuDate(df.CK_GT_C, value)
	return q
}

func (q *VendorInheritNekoCQ) SetInuDate_LessThan(value df.Date) *VendorInheritNekoCQ {
	q.regInuDate(df.CK_LT_C, value)
	return q
}

func (q *VendorInheritNekoCQ) SetInuDate_GreaterEqual(value df.Date) *VendorInheritNekoCQ {
	q.regInuDate(df.CK_GE_C, value)
	return q
}

func (q *VendorInheritNekoCQ) SetInuDate_LessEqual(value df.Date) *VendorInheritNekoCQ {
	q.regInuDate(df.CK_LE_C, value)
	return q
}

func (q *VendorInheritNekoCQ) SetInuDate_IsNull() *VendorInheritNekoCQ {
	q.regInuDate(df.CK_ISN_C, 0)
	return q
}
func (q *VendorInheritNekoCQ) SetInuDate_IsNotNull() *VendorInheritNekoCQ {
	q.regInuDate(df.CK_ISNN_C, 0)
	return q
}
func (q *VendorInheritNekoCQ) AddOrderBy_InuDate_Asc() *VendorInheritNekoCQ {
	q.BaseConditionQuery.RegOBA("inuDate")
	return q
}
func (q *VendorInheritNekoCQ) AddOrderBy_InuDate_Desc() *VendorInheritNekoCQ {
	q.BaseConditionQuery.RegOBD("inuDate")
	return q
}
func (q *VendorInheritNekoCQ) regInuDate(key *df.ConditionKey, value interface{}) {
	if q.InuDate == nil {
		q.InuDate = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.InuDate, "inuDate")
}

func (q *VendorInheritNekoCQ) getCValueNekoId() *df.ConditionValue {
	if q.NekoId == nil {
		q.NekoId = new(df.ConditionValue)
	}
	return q.NekoId
}



func (q *VendorInheritNekoCQ) SetNekoId_Equal(value int64) *VendorInheritNekoCQ {
	q.regNekoId(df.CK_EQ_C, value)
	return q
}

func (q *VendorInheritNekoCQ) SetNekoId_NotEqual(value int64) *VendorInheritNekoCQ {
	q.regNekoId(df.CK_NE_C, value)
	return q
}

func (q *VendorInheritNekoCQ) SetNekoId_GreaterThan(value int64) *VendorInheritNekoCQ {
	q.regNekoId(df.CK_GT_C, value)
	return q
}

func (q *VendorInheritNekoCQ) SetNekoId_LessThan(value int64) *VendorInheritNekoCQ {
	q.regNekoId(df.CK_LT_C, value)
	return q
}

func (q *VendorInheritNekoCQ) SetNekoId_GreaterEqual(value int64) *VendorInheritNekoCQ {
	q.regNekoId(df.CK_GE_C, value)
	return q
}

func (q *VendorInheritNekoCQ) SetNekoId_LessEqual(value int64) *VendorInheritNekoCQ {
	q.regNekoId(df.CK_LE_C, value)
	return q
}
func (q *VendorInheritNekoCQ) SetNekoId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueNekoId(),"nekoId",rangeOfOption)
}	


func (q *VendorInheritNekoCQ) SetNekoId_IsNull() *VendorInheritNekoCQ {
	q.regNekoId(df.CK_ISN_C, 0)
	return q
}
func (q *VendorInheritNekoCQ) SetNekoId_IsNotNull() *VendorInheritNekoCQ {
	q.regNekoId(df.CK_ISNN_C, 0)
	return q
}
func (q *VendorInheritNekoCQ) AddOrderBy_NekoId_Asc() *VendorInheritNekoCQ {
	q.BaseConditionQuery.RegOBA("nekoId")
	return q
}
func (q *VendorInheritNekoCQ) AddOrderBy_NekoId_Desc() *VendorInheritNekoCQ {
	q.BaseConditionQuery.RegOBD("nekoId")
	return q
}
func (q *VendorInheritNekoCQ) regNekoId(key *df.ConditionKey, value interface{}) {
	if q.NekoId == nil {
		q.NekoId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.NekoId, "nekoId")
}

func (q *VendorInheritNekoCQ) getCValueNekoName() *df.ConditionValue {
	if q.NekoName == nil {
		q.NekoName = new(df.ConditionValue)
	}
	return q.NekoName
}


func (q *VendorInheritNekoCQ) SetNekoName_Equal(value string) *VendorInheritNekoCQ {
	q.regNekoName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *VendorInheritNekoCQ) SetNekoName_NotEqual(value string) *VendorInheritNekoCQ {
	q.regNekoName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorInheritNekoCQ) SetNekoName_GreaterThan(value string) *VendorInheritNekoCQ {
	q.regNekoName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorInheritNekoCQ) SetNekoName_LessThan(value string) *VendorInheritNekoCQ {
	q.regNekoName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorInheritNekoCQ) SetNekoName_GreaterEqualThan(value string) *VendorInheritNekoCQ {
	q.regNekoName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *VendorInheritNekoCQ) SetNekoName_LessEqualThan(value string) *VendorInheritNekoCQ {
	q.regNekoName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorInheritNekoCQ) SetNekoName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueNekoName(), "nekoName", option)
}

func (q *VendorInheritNekoCQ) SetNekoName_PrefixSearch(value string) error {
	return q.SetNekoName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *VendorInheritNekoCQ) SetNekoName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueNekoName(), "nekoName", option)
}



func (q *VendorInheritNekoCQ) AddOrderBy_NekoName_Asc() *VendorInheritNekoCQ {
	q.BaseConditionQuery.RegOBA("nekoName")
	return q
}
func (q *VendorInheritNekoCQ) AddOrderBy_NekoName_Desc() *VendorInheritNekoCQ {
	q.BaseConditionQuery.RegOBD("nekoName")
	return q
}
func (q *VendorInheritNekoCQ) regNekoName(key *df.ConditionKey, value interface{}) {
	if q.NekoName == nil {
		q.NekoName = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.NekoName, "nekoName")
}

func (q *VendorInheritNekoCQ) getCValueNekoDate() *df.ConditionValue {
	if q.NekoDate == nil {
		q.NekoDate = new(df.ConditionValue)
	}
	return q.NekoDate
}




func (q *VendorInheritNekoCQ) SetNekoDate_Equal(value df.Date) *VendorInheritNekoCQ {
	q.regNekoDate(df.CK_EQ_C, value)
	return q
}


func (q *VendorInheritNekoCQ) SetNekoDate_GreaterThan(value df.Date) *VendorInheritNekoCQ {
	q.regNekoDate(df.CK_GT_C, value)
	return q
}

func (q *VendorInheritNekoCQ) SetNekoDate_LessThan(value df.Date) *VendorInheritNekoCQ {
	q.regNekoDate(df.CK_LT_C, value)
	return q
}

func (q *VendorInheritNekoCQ) SetNekoDate_GreaterEqual(value df.Date) *VendorInheritNekoCQ {
	q.regNekoDate(df.CK_GE_C, value)
	return q
}

func (q *VendorInheritNekoCQ) SetNekoDate_LessEqual(value df.Date) *VendorInheritNekoCQ {
	q.regNekoDate(df.CK_LE_C, value)
	return q
}

func (q *VendorInheritNekoCQ) SetNekoDate_IsNull() *VendorInheritNekoCQ {
	q.regNekoDate(df.CK_ISN_C, 0)
	return q
}
func (q *VendorInheritNekoCQ) SetNekoDate_IsNotNull() *VendorInheritNekoCQ {
	q.regNekoDate(df.CK_ISNN_C, 0)
	return q
}
func (q *VendorInheritNekoCQ) AddOrderBy_NekoDate_Asc() *VendorInheritNekoCQ {
	q.BaseConditionQuery.RegOBA("nekoDate")
	return q
}
func (q *VendorInheritNekoCQ) AddOrderBy_NekoDate_Desc() *VendorInheritNekoCQ {
	q.BaseConditionQuery.RegOBD("nekoDate")
	return q
}
func (q *VendorInheritNekoCQ) regNekoDate(key *df.ConditionKey, value interface{}) {
	if q.NekoDate == nil {
		q.NekoDate = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.NekoDate, "nekoDate")
}

