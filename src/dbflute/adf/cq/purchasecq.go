package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type PurchaseCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	PurchaseId *df.ConditionValue
	MemberId *df.ConditionValue
	ProductId *df.ConditionValue
	PurchaseDatetime *df.ConditionValue
	PurchaseCount *df.ConditionValue
	PurchasePrice *df.ConditionValue
	PaymentCompleteFlg *df.ConditionValue
	RegisterDatetime *df.ConditionValue
	RegisterUser *df.ConditionValue
	UpdateDatetime *df.ConditionValue
	UpdateUser *df.ConditionValue
	VersionNo *df.ConditionValue
    conditionQueryMember *MemberCQ
    conditionQueryProduct *ProductCQ
}

func (q *PurchaseCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *PurchaseCQ) getCValuePurchaseId() *df.ConditionValue {
	if q.PurchaseId == nil {
		q.PurchaseId = new(df.ConditionValue)
	}
	return q.PurchaseId
}



func (q *PurchaseCQ) SetPurchaseId_Equal(value int64) *PurchaseCQ {
	q.regPurchaseId(df.CK_EQ_C, value)
	return q
}
func (q *PurchaseCQ) SetPurchaseId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValuePurchaseId(), "purchaseId")
}
func (q *PurchaseCQ) SetPurchaseId_NotEqual(value int64) *PurchaseCQ {
	q.regPurchaseId(df.CK_NE_C, value)
	return q
}

func (q *PurchaseCQ) SetPurchaseId_GreaterThan(value int64) *PurchaseCQ {
	q.regPurchaseId(df.CK_GT_C, value)
	return q
}

func (q *PurchaseCQ) SetPurchaseId_LessThan(value int64) *PurchaseCQ {
	q.regPurchaseId(df.CK_LT_C, value)
	return q
}

func (q *PurchaseCQ) SetPurchaseId_GreaterEqual(value int64) *PurchaseCQ {
	q.regPurchaseId(df.CK_GE_C, value)
	return q
}

func (q *PurchaseCQ) SetPurchaseId_LessEqual(value int64) *PurchaseCQ {
	q.regPurchaseId(df.CK_LE_C, value)
	return q
}
func (q *PurchaseCQ) SetPurchaseId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValuePurchaseId(),"purchaseId",rangeOfOption)
}	


func (q *PurchaseCQ) SetPurchaseId_IsNull() *PurchaseCQ {
	q.regPurchaseId(df.CK_ISN_C, 0)
	return q
}
func (q *PurchaseCQ) SetPurchaseId_IsNotNull() *PurchaseCQ {
	q.regPurchaseId(df.CK_ISNN_C, 0)
	return q
}
func (q *PurchaseCQ) AddOrderBy_PurchaseId_Asc() *PurchaseCQ {
	q.BaseConditionQuery.RegOBA("purchaseId")
	return q
}
func (q *PurchaseCQ) AddOrderBy_PurchaseId_Desc() *PurchaseCQ {
	q.BaseConditionQuery.RegOBD("purchaseId")
	return q
}
func (q *PurchaseCQ) regPurchaseId(key *df.ConditionKey, value interface{}) {
	if q.PurchaseId == nil {
		q.PurchaseId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PurchaseId, "purchaseId")
}

func (q *PurchaseCQ) getCValueMemberId() *df.ConditionValue {
	if q.MemberId == nil {
		q.MemberId = new(df.ConditionValue)
	}
	return q.MemberId
}



func (q *PurchaseCQ) SetMemberId_Equal(value int64) *PurchaseCQ {
	q.regMemberId(df.CK_EQ_C, value)
	return q
}
func (q *PurchaseCQ) SetMemberId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueMemberId(), "memberId")
}
func (q *PurchaseCQ) SetMemberId_NotEqual(value int64) *PurchaseCQ {
	q.regMemberId(df.CK_NE_C, value)
	return q
}

func (q *PurchaseCQ) SetMemberId_GreaterThan(value int64) *PurchaseCQ {
	q.regMemberId(df.CK_GT_C, value)
	return q
}

func (q *PurchaseCQ) SetMemberId_LessThan(value int64) *PurchaseCQ {
	q.regMemberId(df.CK_LT_C, value)
	return q
}

func (q *PurchaseCQ) SetMemberId_GreaterEqual(value int64) *PurchaseCQ {
	q.regMemberId(df.CK_GE_C, value)
	return q
}

func (q *PurchaseCQ) SetMemberId_LessEqual(value int64) *PurchaseCQ {
	q.regMemberId(df.CK_LE_C, value)
	return q
}
func (q *PurchaseCQ) SetMemberId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueMemberId(),"memberId",rangeOfOption)
}	


func (q *PurchaseCQ) AddOrderBy_MemberId_Asc() *PurchaseCQ {
	q.BaseConditionQuery.RegOBA("memberId")
	return q
}
func (q *PurchaseCQ) AddOrderBy_MemberId_Desc() *PurchaseCQ {
	q.BaseConditionQuery.RegOBD("memberId")
	return q
}
func (q *PurchaseCQ) regMemberId(key *df.ConditionKey, value interface{}) {
	if q.MemberId == nil {
		q.MemberId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.MemberId, "memberId")
}

func (q *PurchaseCQ) getCValueProductId() *df.ConditionValue {
	if q.ProductId == nil {
		q.ProductId = new(df.ConditionValue)
	}
	return q.ProductId
}



func (q *PurchaseCQ) SetProductId_Equal(value int64) *PurchaseCQ {
	q.regProductId(df.CK_EQ_C, value)
	return q
}
func (q *PurchaseCQ) SetProductId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueProductId(), "productId")
}
func (q *PurchaseCQ) SetProductId_NotEqual(value int64) *PurchaseCQ {
	q.regProductId(df.CK_NE_C, value)
	return q
}

func (q *PurchaseCQ) SetProductId_GreaterThan(value int64) *PurchaseCQ {
	q.regProductId(df.CK_GT_C, value)
	return q
}

func (q *PurchaseCQ) SetProductId_LessThan(value int64) *PurchaseCQ {
	q.regProductId(df.CK_LT_C, value)
	return q
}

func (q *PurchaseCQ) SetProductId_GreaterEqual(value int64) *PurchaseCQ {
	q.regProductId(df.CK_GE_C, value)
	return q
}

func (q *PurchaseCQ) SetProductId_LessEqual(value int64) *PurchaseCQ {
	q.regProductId(df.CK_LE_C, value)
	return q
}
func (q *PurchaseCQ) SetProductId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueProductId(),"productId",rangeOfOption)
}	


func (q *PurchaseCQ) AddOrderBy_ProductId_Asc() *PurchaseCQ {
	q.BaseConditionQuery.RegOBA("productId")
	return q
}
func (q *PurchaseCQ) AddOrderBy_ProductId_Desc() *PurchaseCQ {
	q.BaseConditionQuery.RegOBD("productId")
	return q
}
func (q *PurchaseCQ) regProductId(key *df.ConditionKey, value interface{}) {
	if q.ProductId == nil {
		q.ProductId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ProductId, "productId")
}

func (q *PurchaseCQ) getCValuePurchaseDatetime() *df.ConditionValue {
	if q.PurchaseDatetime == nil {
		q.PurchaseDatetime = new(df.ConditionValue)
	}
	return q.PurchaseDatetime
}




func (q *PurchaseCQ) SetPurchaseDatetime_Equal(value df.MysqlTimestamp) *PurchaseCQ {
	q.regPurchaseDatetime(df.CK_EQ_C, value)
	return q
}


func (q *PurchaseCQ) SetPurchaseDatetime_GreaterThan(value df.MysqlTimestamp) *PurchaseCQ {
	q.regPurchaseDatetime(df.CK_GT_C, value)
	return q
}

func (q *PurchaseCQ) SetPurchaseDatetime_LessThan(value df.MysqlTimestamp) *PurchaseCQ {
	q.regPurchaseDatetime(df.CK_LT_C, value)
	return q
}

func (q *PurchaseCQ) SetPurchaseDatetime_GreaterEqual(value df.MysqlTimestamp) *PurchaseCQ {
	q.regPurchaseDatetime(df.CK_GE_C, value)
	return q
}

func (q *PurchaseCQ) SetPurchaseDatetime_LessEqual(value df.MysqlTimestamp) *PurchaseCQ {
	q.regPurchaseDatetime(df.CK_LE_C, value)
	return q
}

func (q *PurchaseCQ) AddOrderBy_PurchaseDatetime_Asc() *PurchaseCQ {
	q.BaseConditionQuery.RegOBA("purchaseDatetime")
	return q
}
func (q *PurchaseCQ) AddOrderBy_PurchaseDatetime_Desc() *PurchaseCQ {
	q.BaseConditionQuery.RegOBD("purchaseDatetime")
	return q
}
func (q *PurchaseCQ) regPurchaseDatetime(key *df.ConditionKey, value interface{}) {
	if q.PurchaseDatetime == nil {
		q.PurchaseDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PurchaseDatetime, "purchaseDatetime")
}

func (q *PurchaseCQ) getCValuePurchaseCount() *df.ConditionValue {
	if q.PurchaseCount == nil {
		q.PurchaseCount = new(df.ConditionValue)
	}
	return q.PurchaseCount
}



func (q *PurchaseCQ) SetPurchaseCount_Equal(value int64) *PurchaseCQ {
	q.regPurchaseCount(df.CK_EQ_C, value)
	return q
}
func (q *PurchaseCQ) SetPurchaseCount_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValuePurchaseCount(), "purchaseCount")
}
func (q *PurchaseCQ) SetPurchaseCount_NotEqual(value int64) *PurchaseCQ {
	q.regPurchaseCount(df.CK_NE_C, value)
	return q
}

func (q *PurchaseCQ) SetPurchaseCount_GreaterThan(value int64) *PurchaseCQ {
	q.regPurchaseCount(df.CK_GT_C, value)
	return q
}

func (q *PurchaseCQ) SetPurchaseCount_LessThan(value int64) *PurchaseCQ {
	q.regPurchaseCount(df.CK_LT_C, value)
	return q
}

func (q *PurchaseCQ) SetPurchaseCount_GreaterEqual(value int64) *PurchaseCQ {
	q.regPurchaseCount(df.CK_GE_C, value)
	return q
}

func (q *PurchaseCQ) SetPurchaseCount_LessEqual(value int64) *PurchaseCQ {
	q.regPurchaseCount(df.CK_LE_C, value)
	return q
}
func (q *PurchaseCQ) SetPurchaseCount_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValuePurchaseCount(),"purchaseCount",rangeOfOption)
}	


func (q *PurchaseCQ) AddOrderBy_PurchaseCount_Asc() *PurchaseCQ {
	q.BaseConditionQuery.RegOBA("purchaseCount")
	return q
}
func (q *PurchaseCQ) AddOrderBy_PurchaseCount_Desc() *PurchaseCQ {
	q.BaseConditionQuery.RegOBD("purchaseCount")
	return q
}
func (q *PurchaseCQ) regPurchaseCount(key *df.ConditionKey, value interface{}) {
	if q.PurchaseCount == nil {
		q.PurchaseCount = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PurchaseCount, "purchaseCount")
}

func (q *PurchaseCQ) getCValuePurchasePrice() *df.ConditionValue {
	if q.PurchasePrice == nil {
		q.PurchasePrice = new(df.ConditionValue)
	}
	return q.PurchasePrice
}



func (q *PurchaseCQ) SetPurchasePrice_Equal(value int64) *PurchaseCQ {
	q.regPurchasePrice(df.CK_EQ_C, value)
	return q
}
func (q *PurchaseCQ) SetPurchasePrice_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValuePurchasePrice(), "purchasePrice")
}
func (q *PurchaseCQ) SetPurchasePrice_NotEqual(value int64) *PurchaseCQ {
	q.regPurchasePrice(df.CK_NE_C, value)
	return q
}

func (q *PurchaseCQ) SetPurchasePrice_GreaterThan(value int64) *PurchaseCQ {
	q.regPurchasePrice(df.CK_GT_C, value)
	return q
}

func (q *PurchaseCQ) SetPurchasePrice_LessThan(value int64) *PurchaseCQ {
	q.regPurchasePrice(df.CK_LT_C, value)
	return q
}

func (q *PurchaseCQ) SetPurchasePrice_GreaterEqual(value int64) *PurchaseCQ {
	q.regPurchasePrice(df.CK_GE_C, value)
	return q
}

func (q *PurchaseCQ) SetPurchasePrice_LessEqual(value int64) *PurchaseCQ {
	q.regPurchasePrice(df.CK_LE_C, value)
	return q
}
func (q *PurchaseCQ) SetPurchasePrice_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValuePurchasePrice(),"purchasePrice",rangeOfOption)
}	


func (q *PurchaseCQ) AddOrderBy_PurchasePrice_Asc() *PurchaseCQ {
	q.BaseConditionQuery.RegOBA("purchasePrice")
	return q
}
func (q *PurchaseCQ) AddOrderBy_PurchasePrice_Desc() *PurchaseCQ {
	q.BaseConditionQuery.RegOBD("purchasePrice")
	return q
}
func (q *PurchaseCQ) regPurchasePrice(key *df.ConditionKey, value interface{}) {
	if q.PurchasePrice == nil {
		q.PurchasePrice = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PurchasePrice, "purchasePrice")
}

func (q *PurchaseCQ) getCValuePaymentCompleteFlg() *df.ConditionValue {
	if q.PaymentCompleteFlg == nil {
		q.PaymentCompleteFlg = new(df.ConditionValue)
	}
	return q.PaymentCompleteFlg
}



func (q *PurchaseCQ) SetPaymentCompleteFlg_Equal(value int64) *PurchaseCQ {
	q.regPaymentCompleteFlg(df.CK_EQ_C, value)
	return q
}
func (q *PurchaseCQ) SetPaymentCompleteFlg_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValuePaymentCompleteFlg(), "paymentCompleteFlg")
}
func (q *PurchaseCQ) SetPaymentCompleteFlg_NotEqual(value int64) *PurchaseCQ {
	q.regPaymentCompleteFlg(df.CK_NE_C, value)
	return q
}

func (q *PurchaseCQ) SetPaymentCompleteFlg_GreaterThan(value int64) *PurchaseCQ {
	q.regPaymentCompleteFlg(df.CK_GT_C, value)
	return q
}

func (q *PurchaseCQ) SetPaymentCompleteFlg_LessThan(value int64) *PurchaseCQ {
	q.regPaymentCompleteFlg(df.CK_LT_C, value)
	return q
}

func (q *PurchaseCQ) SetPaymentCompleteFlg_GreaterEqual(value int64) *PurchaseCQ {
	q.regPaymentCompleteFlg(df.CK_GE_C, value)
	return q
}

func (q *PurchaseCQ) SetPaymentCompleteFlg_LessEqual(value int64) *PurchaseCQ {
	q.regPaymentCompleteFlg(df.CK_LE_C, value)
	return q
}
func (q *PurchaseCQ) SetPaymentCompleteFlg_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValuePaymentCompleteFlg(),"paymentCompleteFlg",rangeOfOption)
}	


func (q *PurchaseCQ) AddOrderBy_PaymentCompleteFlg_Asc() *PurchaseCQ {
	q.BaseConditionQuery.RegOBA("paymentCompleteFlg")
	return q
}
func (q *PurchaseCQ) AddOrderBy_PaymentCompleteFlg_Desc() *PurchaseCQ {
	q.BaseConditionQuery.RegOBD("paymentCompleteFlg")
	return q
}
func (q *PurchaseCQ) regPaymentCompleteFlg(key *df.ConditionKey, value interface{}) {
	if q.PaymentCompleteFlg == nil {
		q.PaymentCompleteFlg = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PaymentCompleteFlg, "paymentCompleteFlg")
}

func (q *PurchaseCQ) getCValueRegisterDatetime() *df.ConditionValue {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	return q.RegisterDatetime
}




func (q *PurchaseCQ) SetRegisterDatetime_Equal(value df.MysqlTimestamp) *PurchaseCQ {
	q.regRegisterDatetime(df.CK_EQ_C, value)
	return q
}


func (q *PurchaseCQ) SetRegisterDatetime_GreaterThan(value df.MysqlTimestamp) *PurchaseCQ {
	q.regRegisterDatetime(df.CK_GT_C, value)
	return q
}

func (q *PurchaseCQ) SetRegisterDatetime_LessThan(value df.MysqlTimestamp) *PurchaseCQ {
	q.regRegisterDatetime(df.CK_LT_C, value)
	return q
}

func (q *PurchaseCQ) SetRegisterDatetime_GreaterEqual(value df.MysqlTimestamp) *PurchaseCQ {
	q.regRegisterDatetime(df.CK_GE_C, value)
	return q
}

func (q *PurchaseCQ) SetRegisterDatetime_LessEqual(value df.MysqlTimestamp) *PurchaseCQ {
	q.regRegisterDatetime(df.CK_LE_C, value)
	return q
}

func (q *PurchaseCQ) AddOrderBy_RegisterDatetime_Asc() *PurchaseCQ {
	q.BaseConditionQuery.RegOBA("registerDatetime")
	return q
}
func (q *PurchaseCQ) AddOrderBy_RegisterDatetime_Desc() *PurchaseCQ {
	q.BaseConditionQuery.RegOBD("registerDatetime")
	return q
}
func (q *PurchaseCQ) regRegisterDatetime(key *df.ConditionKey, value interface{}) {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterDatetime, "registerDatetime")
}

func (q *PurchaseCQ) getCValueRegisterUser() *df.ConditionValue {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	return q.RegisterUser
}


func (q *PurchaseCQ) SetRegisterUser_Equal(value string) *PurchaseCQ {
	q.regRegisterUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *PurchaseCQ) SetRegisterUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegisterUser(), "registerUser")
}
func (q *PurchaseCQ) SetRegisterUser_NotEqual(value string) *PurchaseCQ {
	q.regRegisterUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *PurchaseCQ) SetRegisterUser_GreaterThan(value string) *PurchaseCQ {
	q.regRegisterUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *PurchaseCQ) SetRegisterUser_LessThan(value string) *PurchaseCQ {
	q.regRegisterUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *PurchaseCQ) SetRegisterUser_GreaterEqualThan(value string) *PurchaseCQ {
	q.regRegisterUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *PurchaseCQ) SetRegisterUser_LessEqualThan(value string) *PurchaseCQ {
	q.regRegisterUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *PurchaseCQ) SetRegisterUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}

func (q *PurchaseCQ) SetRegisterUser_PrefixSearch(value string) error {
	return q.SetRegisterUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *PurchaseCQ) SetRegisterUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}



func (q *PurchaseCQ) AddOrderBy_RegisterUser_Asc() *PurchaseCQ {
	q.BaseConditionQuery.RegOBA("registerUser")
	return q
}
func (q *PurchaseCQ) AddOrderBy_RegisterUser_Desc() *PurchaseCQ {
	q.BaseConditionQuery.RegOBD("registerUser")
	return q
}
func (q *PurchaseCQ) regRegisterUser(key *df.ConditionKey, value interface{}) {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterUser, "registerUser")
}

func (q *PurchaseCQ) getCValueUpdateDatetime() *df.ConditionValue {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	return q.UpdateDatetime
}




func (q *PurchaseCQ) SetUpdateDatetime_Equal(value df.MysqlTimestamp) *PurchaseCQ {
	q.regUpdateDatetime(df.CK_EQ_C, value)
	return q
}


func (q *PurchaseCQ) SetUpdateDatetime_GreaterThan(value df.MysqlTimestamp) *PurchaseCQ {
	q.regUpdateDatetime(df.CK_GT_C, value)
	return q
}

func (q *PurchaseCQ) SetUpdateDatetime_LessThan(value df.MysqlTimestamp) *PurchaseCQ {
	q.regUpdateDatetime(df.CK_LT_C, value)
	return q
}

func (q *PurchaseCQ) SetUpdateDatetime_GreaterEqual(value df.MysqlTimestamp) *PurchaseCQ {
	q.regUpdateDatetime(df.CK_GE_C, value)
	return q
}

func (q *PurchaseCQ) SetUpdateDatetime_LessEqual(value df.MysqlTimestamp) *PurchaseCQ {
	q.regUpdateDatetime(df.CK_LE_C, value)
	return q
}

func (q *PurchaseCQ) AddOrderBy_UpdateDatetime_Asc() *PurchaseCQ {
	q.BaseConditionQuery.RegOBA("updateDatetime")
	return q
}
func (q *PurchaseCQ) AddOrderBy_UpdateDatetime_Desc() *PurchaseCQ {
	q.BaseConditionQuery.RegOBD("updateDatetime")
	return q
}
func (q *PurchaseCQ) regUpdateDatetime(key *df.ConditionKey, value interface{}) {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateDatetime, "updateDatetime")
}

func (q *PurchaseCQ) getCValueUpdateUser() *df.ConditionValue {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	return q.UpdateUser
}


func (q *PurchaseCQ) SetUpdateUser_Equal(value string) *PurchaseCQ {
	q.regUpdateUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *PurchaseCQ) SetUpdateUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUpdateUser(), "updateUser")
}
func (q *PurchaseCQ) SetUpdateUser_NotEqual(value string) *PurchaseCQ {
	q.regUpdateUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *PurchaseCQ) SetUpdateUser_GreaterThan(value string) *PurchaseCQ {
	q.regUpdateUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *PurchaseCQ) SetUpdateUser_LessThan(value string) *PurchaseCQ {
	q.regUpdateUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *PurchaseCQ) SetUpdateUser_GreaterEqualThan(value string) *PurchaseCQ {
	q.regUpdateUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *PurchaseCQ) SetUpdateUser_LessEqualThan(value string) *PurchaseCQ {
	q.regUpdateUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *PurchaseCQ) SetUpdateUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}

func (q *PurchaseCQ) SetUpdateUser_PrefixSearch(value string) error {
	return q.SetUpdateUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *PurchaseCQ) SetUpdateUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}



func (q *PurchaseCQ) AddOrderBy_UpdateUser_Asc() *PurchaseCQ {
	q.BaseConditionQuery.RegOBA("updateUser")
	return q
}
func (q *PurchaseCQ) AddOrderBy_UpdateUser_Desc() *PurchaseCQ {
	q.BaseConditionQuery.RegOBD("updateUser")
	return q
}
func (q *PurchaseCQ) regUpdateUser(key *df.ConditionKey, value interface{}) {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateUser, "updateUser")
}

func (q *PurchaseCQ) getCValueVersionNo() *df.ConditionValue {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	return q.VersionNo
}



func (q *PurchaseCQ) SetVersionNo_Equal(value int64) *PurchaseCQ {
	q.regVersionNo(df.CK_EQ_C, value)
	return q
}
func (q *PurchaseCQ) SetVersionNo_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueVersionNo(), "versionNo")
}
func (q *PurchaseCQ) SetVersionNo_NotEqual(value int64) *PurchaseCQ {
	q.regVersionNo(df.CK_NE_C, value)
	return q
}

func (q *PurchaseCQ) SetVersionNo_GreaterThan(value int64) *PurchaseCQ {
	q.regVersionNo(df.CK_GT_C, value)
	return q
}

func (q *PurchaseCQ) SetVersionNo_LessThan(value int64) *PurchaseCQ {
	q.regVersionNo(df.CK_LT_C, value)
	return q
}

func (q *PurchaseCQ) SetVersionNo_GreaterEqual(value int64) *PurchaseCQ {
	q.regVersionNo(df.CK_GE_C, value)
	return q
}

func (q *PurchaseCQ) SetVersionNo_LessEqual(value int64) *PurchaseCQ {
	q.regVersionNo(df.CK_LE_C, value)
	return q
}
func (q *PurchaseCQ) SetVersionNo_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueVersionNo(),"versionNo",rangeOfOption)
}	


func (q *PurchaseCQ) AddOrderBy_VersionNo_Asc() *PurchaseCQ {
	q.BaseConditionQuery.RegOBA("versionNo")
	return q
}
func (q *PurchaseCQ) AddOrderBy_VersionNo_Desc() *PurchaseCQ {
	q.BaseConditionQuery.RegOBD("versionNo")
	return q
}
func (q *PurchaseCQ) regVersionNo(key *df.ConditionKey, value interface{}) {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.VersionNo, "versionNo")
}


func (q *PurchaseCQ) QueryMember() *MemberCQ {
	if q.conditionQueryMember == nil {
		q.conditionQueryMember = q.xcreateQueryMember()
		q.xsetupOuterJoinMember()
	}
	return q.conditionQueryMember
}

func (q *PurchaseCQ) xcreateQueryMember() *MemberCQ {
	nrp := q.BaseConditionQuery.ResolveNextRelationPath("Purchase", "Member")
	jan := q.BaseConditionQuery.ResolveJoinAliasName(nrp)
	var basecq df.ConditionQuery = q
	cq := CreateMemberCQ(&basecq, q.BaseConditionQuery.SqlClause, jan, q.BaseConditionQuery.NestLevel+1)
	cq.BaseConditionQuery.BaseCB = q.BaseConditionQuery.BaseCB
	cq.BaseConditionQuery.ForeignPropertyName = "Member"
	cq.BaseConditionQuery.RelationPath = nrp
	return cq
}
func (q *PurchaseCQ) xsetupOuterJoinMember() {
	    cq := q.QueryMember()
        joinOnMap := make(map[string]string)
        joinOnMap["memberId"]="memberId"
        q.BaseConditionQuery.RegisterOuterJoin(
        	cq.BaseConditionQuery.ConditionQuery, joinOnMap, "Member");
}	
	
func (q *PurchaseCQ) QueryProduct() *ProductCQ {
	if q.conditionQueryProduct == nil {
		q.conditionQueryProduct = q.xcreateQueryProduct()
		q.xsetupOuterJoinProduct()
	}
	return q.conditionQueryProduct
}

func (q *PurchaseCQ) xcreateQueryProduct() *ProductCQ {
	nrp := q.BaseConditionQuery.ResolveNextRelationPath("Purchase", "Product")
	jan := q.BaseConditionQuery.ResolveJoinAliasName(nrp)
	var basecq df.ConditionQuery = q
	cq := CreateProductCQ(&basecq, q.BaseConditionQuery.SqlClause, jan, q.BaseConditionQuery.NestLevel+1)
	cq.BaseConditionQuery.BaseCB = q.BaseConditionQuery.BaseCB
	cq.BaseConditionQuery.ForeignPropertyName = "Product"
	cq.BaseConditionQuery.RelationPath = nrp
	return cq
}
func (q *PurchaseCQ) xsetupOuterJoinProduct() {
	    cq := q.QueryProduct()
        joinOnMap := make(map[string]string)
        joinOnMap["productId"]="productId"
        q.BaseConditionQuery.RegisterOuterJoin(
        	cq.BaseConditionQuery.ConditionQuery, joinOnMap, "Product");
}	
	
func CreatePurchaseCQ(referrerQuery *df.ConditionQuery, sqlClause *df.SqlClause, aliasName string, nestlevel int8) *PurchaseCQ {
	cq := new(PurchaseCQ)
	cq.BaseConditionQuery = new(df.BaseConditionQuery)
	cq.BaseConditionQuery.TableDbName = "Purchase"
	cq.BaseConditionQuery.ReferrerQuery = referrerQuery
	cq.BaseConditionQuery.SqlClause = sqlClause
	cq.BaseConditionQuery.AliasName = aliasName
	cq.BaseConditionQuery.NestLevel = nestlevel
	cq.BaseConditionQuery.DBMetaProvider = df.DBMetaProvider_I
	cq.BaseConditionQuery.CQ_PROPERTY = "Query"
	var cqi df.ConditionQuery = cq
	cq.BaseConditionQuery.ConditionQuery=&cqi
	return cq
}	