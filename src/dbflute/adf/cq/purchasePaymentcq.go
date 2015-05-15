package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type PurchasePaymentCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	PurchasePaymentId *df.ConditionValue
	PurchaseId *df.ConditionValue
	PaymentAmount *df.ConditionValue
	PaymentDatetime *df.ConditionValue
	PaymentMethodCode *df.ConditionValue
	RegisterDatetime *df.ConditionValue
	RegisterUser *df.ConditionValue
	UpdateDatetime *df.ConditionValue
	UpdateUser *df.ConditionValue
    conditionQueryPurchase *PurchaseCQ
}

func (q *PurchasePaymentCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *PurchasePaymentCQ) getCValuePurchasePaymentId() *df.ConditionValue {
	if q.PurchasePaymentId == nil {
		q.PurchasePaymentId = new(df.ConditionValue)
	}
	return q.PurchasePaymentId
}



func (q *PurchasePaymentCQ) SetPurchasePaymentId_Equal(value int64) *PurchasePaymentCQ {
	q.regPurchasePaymentId(df.CK_EQ_C, value)
	return q
}
func (q *PurchasePaymentCQ) SetPurchasePaymentId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValuePurchasePaymentId(), "purchasePaymentId")
}
func (q *PurchasePaymentCQ) SetPurchasePaymentId_NotEqual(value int64) *PurchasePaymentCQ {
	q.regPurchasePaymentId(df.CK_NE_C, value)
	return q
}

func (q *PurchasePaymentCQ) SetPurchasePaymentId_GreaterThan(value int64) *PurchasePaymentCQ {
	q.regPurchasePaymentId(df.CK_GT_C, value)
	return q
}

func (q *PurchasePaymentCQ) SetPurchasePaymentId_LessThan(value int64) *PurchasePaymentCQ {
	q.regPurchasePaymentId(df.CK_LT_C, value)
	return q
}

func (q *PurchasePaymentCQ) SetPurchasePaymentId_GreaterEqual(value int64) *PurchasePaymentCQ {
	q.regPurchasePaymentId(df.CK_GE_C, value)
	return q
}

func (q *PurchasePaymentCQ) SetPurchasePaymentId_LessEqual(value int64) *PurchasePaymentCQ {
	q.regPurchasePaymentId(df.CK_LE_C, value)
	return q
}
func (q *PurchasePaymentCQ) SetPurchasePaymentId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValuePurchasePaymentId(),"purchasePaymentId",rangeOfOption)
}	


func (q *PurchasePaymentCQ) SetPurchasePaymentId_IsNull() *PurchasePaymentCQ {
	q.regPurchasePaymentId(df.CK_ISN_C, 0)
	return q
}
func (q *PurchasePaymentCQ) SetPurchasePaymentId_IsNotNull() *PurchasePaymentCQ {
	q.regPurchasePaymentId(df.CK_ISNN_C, 0)
	return q
}
func (q *PurchasePaymentCQ) AddOrderBy_PurchasePaymentId_Asc() *PurchasePaymentCQ {
	q.BaseConditionQuery.RegOBA("purchasePaymentId")
	return q
}
func (q *PurchasePaymentCQ) AddOrderBy_PurchasePaymentId_Desc() *PurchasePaymentCQ {
	q.BaseConditionQuery.RegOBD("purchasePaymentId")
	return q
}
func (q *PurchasePaymentCQ) regPurchasePaymentId(key *df.ConditionKey, value interface{}) {
	if q.PurchasePaymentId == nil {
		q.PurchasePaymentId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PurchasePaymentId, "purchasePaymentId")
}

func (q *PurchasePaymentCQ) getCValuePurchaseId() *df.ConditionValue {
	if q.PurchaseId == nil {
		q.PurchaseId = new(df.ConditionValue)
	}
	return q.PurchaseId
}



func (q *PurchasePaymentCQ) SetPurchaseId_Equal(value int64) *PurchasePaymentCQ {
	q.regPurchaseId(df.CK_EQ_C, value)
	return q
}
func (q *PurchasePaymentCQ) SetPurchaseId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValuePurchaseId(), "purchaseId")
}
func (q *PurchasePaymentCQ) SetPurchaseId_NotEqual(value int64) *PurchasePaymentCQ {
	q.regPurchaseId(df.CK_NE_C, value)
	return q
}

func (q *PurchasePaymentCQ) SetPurchaseId_GreaterThan(value int64) *PurchasePaymentCQ {
	q.regPurchaseId(df.CK_GT_C, value)
	return q
}

func (q *PurchasePaymentCQ) SetPurchaseId_LessThan(value int64) *PurchasePaymentCQ {
	q.regPurchaseId(df.CK_LT_C, value)
	return q
}

func (q *PurchasePaymentCQ) SetPurchaseId_GreaterEqual(value int64) *PurchasePaymentCQ {
	q.regPurchaseId(df.CK_GE_C, value)
	return q
}

func (q *PurchasePaymentCQ) SetPurchaseId_LessEqual(value int64) *PurchasePaymentCQ {
	q.regPurchaseId(df.CK_LE_C, value)
	return q
}
func (q *PurchasePaymentCQ) SetPurchaseId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValuePurchaseId(),"purchaseId",rangeOfOption)
}	


func (q *PurchasePaymentCQ) AddOrderBy_PurchaseId_Asc() *PurchasePaymentCQ {
	q.BaseConditionQuery.RegOBA("purchaseId")
	return q
}
func (q *PurchasePaymentCQ) AddOrderBy_PurchaseId_Desc() *PurchasePaymentCQ {
	q.BaseConditionQuery.RegOBD("purchaseId")
	return q
}
func (q *PurchasePaymentCQ) regPurchaseId(key *df.ConditionKey, value interface{}) {
	if q.PurchaseId == nil {
		q.PurchaseId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PurchaseId, "purchaseId")
}

func (q *PurchasePaymentCQ) getCValuePaymentAmount() *df.ConditionValue {
	if q.PaymentAmount == nil {
		q.PaymentAmount = new(df.ConditionValue)
	}
	return q.PaymentAmount
}



func (q *PurchasePaymentCQ) SetPaymentAmount_Equal(value df.Numeric) *PurchasePaymentCQ {
	q.regPaymentAmount(df.CK_EQ_C, value)
	return q
}
func (q *PurchasePaymentCQ) SetPaymentAmount_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValuePaymentAmount(), "paymentAmount")
}
func (q *PurchasePaymentCQ) SetPaymentAmount_NotEqual(value df.Numeric) *PurchasePaymentCQ {
	q.regPaymentAmount(df.CK_NE_C, value)
	return q
}

func (q *PurchasePaymentCQ) SetPaymentAmount_GreaterThan(value df.Numeric) *PurchasePaymentCQ {
	q.regPaymentAmount(df.CK_GT_C, value)
	return q
}

func (q *PurchasePaymentCQ) SetPaymentAmount_LessThan(value df.Numeric) *PurchasePaymentCQ {
	q.regPaymentAmount(df.CK_LT_C, value)
	return q
}

func (q *PurchasePaymentCQ) SetPaymentAmount_GreaterEqual(value df.Numeric) *PurchasePaymentCQ {
	q.regPaymentAmount(df.CK_GE_C, value)
	return q
}

func (q *PurchasePaymentCQ) SetPaymentAmount_LessEqual(value df.Numeric) *PurchasePaymentCQ {
	q.regPaymentAmount(df.CK_LE_C, value)
	return q
}
func (q *PurchasePaymentCQ) SetPaymentAmount_RangeOf(minNumber df.Numeric, maxNumber df.Numeric, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValuePaymentAmount(),"paymentAmount",rangeOfOption)
}	


func (q *PurchasePaymentCQ) AddOrderBy_PaymentAmount_Asc() *PurchasePaymentCQ {
	q.BaseConditionQuery.RegOBA("paymentAmount")
	return q
}
func (q *PurchasePaymentCQ) AddOrderBy_PaymentAmount_Desc() *PurchasePaymentCQ {
	q.BaseConditionQuery.RegOBD("paymentAmount")
	return q
}
func (q *PurchasePaymentCQ) regPaymentAmount(key *df.ConditionKey, value interface{}) {
	if q.PaymentAmount == nil {
		q.PaymentAmount = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PaymentAmount, "paymentAmount")
}

func (q *PurchasePaymentCQ) getCValuePaymentDatetime() *df.ConditionValue {
	if q.PaymentDatetime == nil {
		q.PaymentDatetime = new(df.ConditionValue)
	}
	return q.PaymentDatetime
}




func (q *PurchasePaymentCQ) SetPaymentDatetime_Equal(value df.MysqlTimestamp) *PurchasePaymentCQ {
	q.regPaymentDatetime(df.CK_EQ_C, value)
	return q
}


func (q *PurchasePaymentCQ) SetPaymentDatetime_GreaterThan(value df.MysqlTimestamp) *PurchasePaymentCQ {
	q.regPaymentDatetime(df.CK_GT_C, value)
	return q
}

func (q *PurchasePaymentCQ) SetPaymentDatetime_LessThan(value df.MysqlTimestamp) *PurchasePaymentCQ {
	q.regPaymentDatetime(df.CK_LT_C, value)
	return q
}

func (q *PurchasePaymentCQ) SetPaymentDatetime_GreaterEqual(value df.MysqlTimestamp) *PurchasePaymentCQ {
	q.regPaymentDatetime(df.CK_GE_C, value)
	return q
}

func (q *PurchasePaymentCQ) SetPaymentDatetime_LessEqual(value df.MysqlTimestamp) *PurchasePaymentCQ {
	q.regPaymentDatetime(df.CK_LE_C, value)
	return q
}

func (q *PurchasePaymentCQ) AddOrderBy_PaymentDatetime_Asc() *PurchasePaymentCQ {
	q.BaseConditionQuery.RegOBA("paymentDatetime")
	return q
}
func (q *PurchasePaymentCQ) AddOrderBy_PaymentDatetime_Desc() *PurchasePaymentCQ {
	q.BaseConditionQuery.RegOBD("paymentDatetime")
	return q
}
func (q *PurchasePaymentCQ) regPaymentDatetime(key *df.ConditionKey, value interface{}) {
	if q.PaymentDatetime == nil {
		q.PaymentDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PaymentDatetime, "paymentDatetime")
}

func (q *PurchasePaymentCQ) getCValuePaymentMethodCode() *df.ConditionValue {
	if q.PaymentMethodCode == nil {
		q.PaymentMethodCode = new(df.ConditionValue)
	}
	return q.PaymentMethodCode
}


func (q *PurchasePaymentCQ) SetPaymentMethodCode_Equal(value string) *PurchasePaymentCQ {
	q.regPaymentMethodCode(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *PurchasePaymentCQ) SetPaymentMethodCode_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValuePaymentMethodCode(), "paymentMethodCode")
}
func (q *PurchasePaymentCQ) SetPaymentMethodCode_NotEqual(value string) *PurchasePaymentCQ {
	q.regPaymentMethodCode(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *PurchasePaymentCQ) SetPaymentMethodCode_GreaterThan(value string) *PurchasePaymentCQ {
	q.regPaymentMethodCode(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *PurchasePaymentCQ) SetPaymentMethodCode_LessThan(value string) *PurchasePaymentCQ {
	q.regPaymentMethodCode(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *PurchasePaymentCQ) SetPaymentMethodCode_GreaterEqualThan(value string) *PurchasePaymentCQ {
	q.regPaymentMethodCode(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *PurchasePaymentCQ) SetPaymentMethodCode_LessEqualThan(value string) *PurchasePaymentCQ {
	q.regPaymentMethodCode(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *PurchasePaymentCQ) SetPaymentMethodCode_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValuePaymentMethodCode(), "paymentMethodCode", option)
}

func (q *PurchasePaymentCQ) SetPaymentMethodCode_PrefixSearch(value string) error {
	return q.SetPaymentMethodCode_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *PurchasePaymentCQ) SetPaymentMethodCode_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValuePaymentMethodCode(), "paymentMethodCode", option)
}



func (q *PurchasePaymentCQ) AddOrderBy_PaymentMethodCode_Asc() *PurchasePaymentCQ {
	q.BaseConditionQuery.RegOBA("paymentMethodCode")
	return q
}
func (q *PurchasePaymentCQ) AddOrderBy_PaymentMethodCode_Desc() *PurchasePaymentCQ {
	q.BaseConditionQuery.RegOBD("paymentMethodCode")
	return q
}
func (q *PurchasePaymentCQ) regPaymentMethodCode(key *df.ConditionKey, value interface{}) {
	if q.PaymentMethodCode == nil {
		q.PaymentMethodCode = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PaymentMethodCode, "paymentMethodCode")
}

func (q *PurchasePaymentCQ) getCValueRegisterDatetime() *df.ConditionValue {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	return q.RegisterDatetime
}




func (q *PurchasePaymentCQ) SetRegisterDatetime_Equal(value df.MysqlTimestamp) *PurchasePaymentCQ {
	q.regRegisterDatetime(df.CK_EQ_C, value)
	return q
}


func (q *PurchasePaymentCQ) SetRegisterDatetime_GreaterThan(value df.MysqlTimestamp) *PurchasePaymentCQ {
	q.regRegisterDatetime(df.CK_GT_C, value)
	return q
}

func (q *PurchasePaymentCQ) SetRegisterDatetime_LessThan(value df.MysqlTimestamp) *PurchasePaymentCQ {
	q.regRegisterDatetime(df.CK_LT_C, value)
	return q
}

func (q *PurchasePaymentCQ) SetRegisterDatetime_GreaterEqual(value df.MysqlTimestamp) *PurchasePaymentCQ {
	q.regRegisterDatetime(df.CK_GE_C, value)
	return q
}

func (q *PurchasePaymentCQ) SetRegisterDatetime_LessEqual(value df.MysqlTimestamp) *PurchasePaymentCQ {
	q.regRegisterDatetime(df.CK_LE_C, value)
	return q
}

func (q *PurchasePaymentCQ) AddOrderBy_RegisterDatetime_Asc() *PurchasePaymentCQ {
	q.BaseConditionQuery.RegOBA("registerDatetime")
	return q
}
func (q *PurchasePaymentCQ) AddOrderBy_RegisterDatetime_Desc() *PurchasePaymentCQ {
	q.BaseConditionQuery.RegOBD("registerDatetime")
	return q
}
func (q *PurchasePaymentCQ) regRegisterDatetime(key *df.ConditionKey, value interface{}) {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterDatetime, "registerDatetime")
}

func (q *PurchasePaymentCQ) getCValueRegisterUser() *df.ConditionValue {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	return q.RegisterUser
}


func (q *PurchasePaymentCQ) SetRegisterUser_Equal(value string) *PurchasePaymentCQ {
	q.regRegisterUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *PurchasePaymentCQ) SetRegisterUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegisterUser(), "registerUser")
}
func (q *PurchasePaymentCQ) SetRegisterUser_NotEqual(value string) *PurchasePaymentCQ {
	q.regRegisterUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *PurchasePaymentCQ) SetRegisterUser_GreaterThan(value string) *PurchasePaymentCQ {
	q.regRegisterUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *PurchasePaymentCQ) SetRegisterUser_LessThan(value string) *PurchasePaymentCQ {
	q.regRegisterUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *PurchasePaymentCQ) SetRegisterUser_GreaterEqualThan(value string) *PurchasePaymentCQ {
	q.regRegisterUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *PurchasePaymentCQ) SetRegisterUser_LessEqualThan(value string) *PurchasePaymentCQ {
	q.regRegisterUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *PurchasePaymentCQ) SetRegisterUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}

func (q *PurchasePaymentCQ) SetRegisterUser_PrefixSearch(value string) error {
	return q.SetRegisterUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *PurchasePaymentCQ) SetRegisterUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}



func (q *PurchasePaymentCQ) AddOrderBy_RegisterUser_Asc() *PurchasePaymentCQ {
	q.BaseConditionQuery.RegOBA("registerUser")
	return q
}
func (q *PurchasePaymentCQ) AddOrderBy_RegisterUser_Desc() *PurchasePaymentCQ {
	q.BaseConditionQuery.RegOBD("registerUser")
	return q
}
func (q *PurchasePaymentCQ) regRegisterUser(key *df.ConditionKey, value interface{}) {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterUser, "registerUser")
}

func (q *PurchasePaymentCQ) getCValueUpdateDatetime() *df.ConditionValue {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	return q.UpdateDatetime
}




func (q *PurchasePaymentCQ) SetUpdateDatetime_Equal(value df.MysqlTimestamp) *PurchasePaymentCQ {
	q.regUpdateDatetime(df.CK_EQ_C, value)
	return q
}


func (q *PurchasePaymentCQ) SetUpdateDatetime_GreaterThan(value df.MysqlTimestamp) *PurchasePaymentCQ {
	q.regUpdateDatetime(df.CK_GT_C, value)
	return q
}

func (q *PurchasePaymentCQ) SetUpdateDatetime_LessThan(value df.MysqlTimestamp) *PurchasePaymentCQ {
	q.regUpdateDatetime(df.CK_LT_C, value)
	return q
}

func (q *PurchasePaymentCQ) SetUpdateDatetime_GreaterEqual(value df.MysqlTimestamp) *PurchasePaymentCQ {
	q.regUpdateDatetime(df.CK_GE_C, value)
	return q
}

func (q *PurchasePaymentCQ) SetUpdateDatetime_LessEqual(value df.MysqlTimestamp) *PurchasePaymentCQ {
	q.regUpdateDatetime(df.CK_LE_C, value)
	return q
}

func (q *PurchasePaymentCQ) AddOrderBy_UpdateDatetime_Asc() *PurchasePaymentCQ {
	q.BaseConditionQuery.RegOBA("updateDatetime")
	return q
}
func (q *PurchasePaymentCQ) AddOrderBy_UpdateDatetime_Desc() *PurchasePaymentCQ {
	q.BaseConditionQuery.RegOBD("updateDatetime")
	return q
}
func (q *PurchasePaymentCQ) regUpdateDatetime(key *df.ConditionKey, value interface{}) {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateDatetime, "updateDatetime")
}

func (q *PurchasePaymentCQ) getCValueUpdateUser() *df.ConditionValue {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	return q.UpdateUser
}


func (q *PurchasePaymentCQ) SetUpdateUser_Equal(value string) *PurchasePaymentCQ {
	q.regUpdateUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *PurchasePaymentCQ) SetUpdateUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUpdateUser(), "updateUser")
}
func (q *PurchasePaymentCQ) SetUpdateUser_NotEqual(value string) *PurchasePaymentCQ {
	q.regUpdateUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *PurchasePaymentCQ) SetUpdateUser_GreaterThan(value string) *PurchasePaymentCQ {
	q.regUpdateUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *PurchasePaymentCQ) SetUpdateUser_LessThan(value string) *PurchasePaymentCQ {
	q.regUpdateUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *PurchasePaymentCQ) SetUpdateUser_GreaterEqualThan(value string) *PurchasePaymentCQ {
	q.regUpdateUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *PurchasePaymentCQ) SetUpdateUser_LessEqualThan(value string) *PurchasePaymentCQ {
	q.regUpdateUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *PurchasePaymentCQ) SetUpdateUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}

func (q *PurchasePaymentCQ) SetUpdateUser_PrefixSearch(value string) error {
	return q.SetUpdateUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *PurchasePaymentCQ) SetUpdateUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}



func (q *PurchasePaymentCQ) AddOrderBy_UpdateUser_Asc() *PurchasePaymentCQ {
	q.BaseConditionQuery.RegOBA("updateUser")
	return q
}
func (q *PurchasePaymentCQ) AddOrderBy_UpdateUser_Desc() *PurchasePaymentCQ {
	q.BaseConditionQuery.RegOBD("updateUser")
	return q
}
func (q *PurchasePaymentCQ) regUpdateUser(key *df.ConditionKey, value interface{}) {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateUser, "updateUser")
}


func (q *PurchasePaymentCQ) QueryPurchase() *PurchaseCQ {
	if q.conditionQueryPurchase == nil {
		q.conditionQueryPurchase = q.xcreateQueryPurchase()
		q.xsetupOuterJoinPurchase()
	}
	return q.conditionQueryPurchase
}

func (q *PurchasePaymentCQ) xcreateQueryPurchase() *PurchaseCQ {
	nrp := q.BaseConditionQuery.ResolveNextRelationPath("PurchasePayment", "Purchase")
	jan := q.BaseConditionQuery.ResolveJoinAliasName(nrp)
	var basecq df.ConditionQuery = q
	cq := CreatePurchaseCQ(&basecq, q.BaseConditionQuery.SqlClause, jan, q.BaseConditionQuery.NestLevel+1)
	cq.BaseConditionQuery.BaseCB = q.BaseConditionQuery.BaseCB
	cq.BaseConditionQuery.ForeignPropertyName = "Purchase"
	cq.BaseConditionQuery.RelationPath = nrp
	return cq
}
func (q *PurchasePaymentCQ) xsetupOuterJoinPurchase() {
	    cq := q.QueryPurchase()
        joinOnMap := make(map[string]string)
        joinOnMap["purchaseId"]="purchaseId"
        q.BaseConditionQuery.RegisterOuterJoin(
        	cq.BaseConditionQuery.ConditionQuery, joinOnMap, "Purchase");
}	
	
func CreatePurchasePaymentCQ(referrerQuery *df.ConditionQuery, sqlClause *df.SqlClause, aliasName string, nestlevel int8) *PurchasePaymentCQ {
	cq := new(PurchasePaymentCQ)
	cq.BaseConditionQuery = new(df.BaseConditionQuery)
	cq.BaseConditionQuery.TableDbName = "PurchasePayment"
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