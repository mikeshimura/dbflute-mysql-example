package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type SummaryWithdrawalCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	MemberId *df.ConditionValue
	MemberName *df.ConditionValue
	WithdrawalReasonCode *df.ConditionValue
	WithdrawalReasonText *df.ConditionValue
	WithdrawalReasonInputText *df.ConditionValue
	WithdrawalDatetime *df.ConditionValue
	MemberStatusCode *df.ConditionValue
	MemberStatusName *df.ConditionValue
	MaxPurchasePrice *df.ConditionValue
}

func (q *SummaryWithdrawalCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *SummaryWithdrawalCQ) getCValueMemberId() *df.ConditionValue {
	if q.MemberId == nil {
		q.MemberId = new(df.ConditionValue)
	}
	return q.MemberId
}



func (q *SummaryWithdrawalCQ) SetMemberId_Equal(value int64) *SummaryWithdrawalCQ {
	q.regMemberId(df.CK_EQ_C, value)
	return q
}
func (q *SummaryWithdrawalCQ) SetMemberId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueMemberId(), "memberId")
}
func (q *SummaryWithdrawalCQ) SetMemberId_NotEqual(value int64) *SummaryWithdrawalCQ {
	q.regMemberId(df.CK_NE_C, value)
	return q
}

func (q *SummaryWithdrawalCQ) SetMemberId_GreaterThan(value int64) *SummaryWithdrawalCQ {
	q.regMemberId(df.CK_GT_C, value)
	return q
}

func (q *SummaryWithdrawalCQ) SetMemberId_LessThan(value int64) *SummaryWithdrawalCQ {
	q.regMemberId(df.CK_LT_C, value)
	return q
}

func (q *SummaryWithdrawalCQ) SetMemberId_GreaterEqual(value int64) *SummaryWithdrawalCQ {
	q.regMemberId(df.CK_GE_C, value)
	return q
}

func (q *SummaryWithdrawalCQ) SetMemberId_LessEqual(value int64) *SummaryWithdrawalCQ {
	q.regMemberId(df.CK_LE_C, value)
	return q
}
func (q *SummaryWithdrawalCQ) SetMemberId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueMemberId(),"memberId",rangeOfOption)
}	


func (q *SummaryWithdrawalCQ) AddOrderBy_MemberId_Asc() *SummaryWithdrawalCQ {
	q.BaseConditionQuery.RegOBA("memberId")
	return q
}
func (q *SummaryWithdrawalCQ) AddOrderBy_MemberId_Desc() *SummaryWithdrawalCQ {
	q.BaseConditionQuery.RegOBD("memberId")
	return q
}
func (q *SummaryWithdrawalCQ) regMemberId(key *df.ConditionKey, value interface{}) {
	if q.MemberId == nil {
		q.MemberId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.MemberId, "memberId")
}

func (q *SummaryWithdrawalCQ) getCValueMemberName() *df.ConditionValue {
	if q.MemberName == nil {
		q.MemberName = new(df.ConditionValue)
	}
	return q.MemberName
}


func (q *SummaryWithdrawalCQ) SetMemberName_Equal(value string) *SummaryWithdrawalCQ {
	q.regMemberName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *SummaryWithdrawalCQ) SetMemberName_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueMemberName(), "memberName")
}
func (q *SummaryWithdrawalCQ) SetMemberName_NotEqual(value string) *SummaryWithdrawalCQ {
	q.regMemberName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryWithdrawalCQ) SetMemberName_GreaterThan(value string) *SummaryWithdrawalCQ {
	q.regMemberName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryWithdrawalCQ) SetMemberName_LessThan(value string) *SummaryWithdrawalCQ {
	q.regMemberName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryWithdrawalCQ) SetMemberName_GreaterEqualThan(value string) *SummaryWithdrawalCQ {
	q.regMemberName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SummaryWithdrawalCQ) SetMemberName_LessEqualThan(value string) *SummaryWithdrawalCQ {
	q.regMemberName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryWithdrawalCQ) SetMemberName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueMemberName(), "memberName", option)
}

func (q *SummaryWithdrawalCQ) SetMemberName_PrefixSearch(value string) error {
	return q.SetMemberName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SummaryWithdrawalCQ) SetMemberName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueMemberName(), "memberName", option)
}



func (q *SummaryWithdrawalCQ) SetMemberName_IsNull() *SummaryWithdrawalCQ {
	q.regMemberName(df.CK_ISN_C, 0)
	return q
}
func (q *SummaryWithdrawalCQ) SetMemberName_IsNullOrEmpty() *SummaryWithdrawalCQ {
	q.regMemberName(df.CK_ISNOE_C, 0)
	return q
}
func (q *SummaryWithdrawalCQ) SetMemberName_IsNotNull() *SummaryWithdrawalCQ {
	q.regMemberName(df.CK_ISNN_C, 0)
	return q
}
func (q *SummaryWithdrawalCQ) AddOrderBy_MemberName_Asc() *SummaryWithdrawalCQ {
	q.BaseConditionQuery.RegOBA("memberName")
	return q
}
func (q *SummaryWithdrawalCQ) AddOrderBy_MemberName_Desc() *SummaryWithdrawalCQ {
	q.BaseConditionQuery.RegOBD("memberName")
	return q
}
func (q *SummaryWithdrawalCQ) regMemberName(key *df.ConditionKey, value interface{}) {
	if q.MemberName == nil {
		q.MemberName = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.MemberName, "memberName")
}

func (q *SummaryWithdrawalCQ) getCValueWithdrawalReasonCode() *df.ConditionValue {
	if q.WithdrawalReasonCode == nil {
		q.WithdrawalReasonCode = new(df.ConditionValue)
	}
	return q.WithdrawalReasonCode
}


func (q *SummaryWithdrawalCQ) SetWithdrawalReasonCode_Equal(value string) *SummaryWithdrawalCQ {
	q.regWithdrawalReasonCode(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *SummaryWithdrawalCQ) SetWithdrawalReasonCode_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueWithdrawalReasonCode(), "withdrawalReasonCode")
}
func (q *SummaryWithdrawalCQ) SetWithdrawalReasonCode_NotEqual(value string) *SummaryWithdrawalCQ {
	q.regWithdrawalReasonCode(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryWithdrawalCQ) SetWithdrawalReasonCode_GreaterThan(value string) *SummaryWithdrawalCQ {
	q.regWithdrawalReasonCode(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryWithdrawalCQ) SetWithdrawalReasonCode_LessThan(value string) *SummaryWithdrawalCQ {
	q.regWithdrawalReasonCode(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryWithdrawalCQ) SetWithdrawalReasonCode_GreaterEqualThan(value string) *SummaryWithdrawalCQ {
	q.regWithdrawalReasonCode(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SummaryWithdrawalCQ) SetWithdrawalReasonCode_LessEqualThan(value string) *SummaryWithdrawalCQ {
	q.regWithdrawalReasonCode(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryWithdrawalCQ) SetWithdrawalReasonCode_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueWithdrawalReasonCode(), "withdrawalReasonCode", option)
}

func (q *SummaryWithdrawalCQ) SetWithdrawalReasonCode_PrefixSearch(value string) error {
	return q.SetWithdrawalReasonCode_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SummaryWithdrawalCQ) SetWithdrawalReasonCode_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueWithdrawalReasonCode(), "withdrawalReasonCode", option)
}



func (q *SummaryWithdrawalCQ) SetWithdrawalReasonCode_IsNull() *SummaryWithdrawalCQ {
	q.regWithdrawalReasonCode(df.CK_ISN_C, 0)
	return q
}
func (q *SummaryWithdrawalCQ) SetWithdrawalReasonCode_IsNullOrEmpty() *SummaryWithdrawalCQ {
	q.regWithdrawalReasonCode(df.CK_ISNOE_C, 0)
	return q
}
func (q *SummaryWithdrawalCQ) SetWithdrawalReasonCode_IsNotNull() *SummaryWithdrawalCQ {
	q.regWithdrawalReasonCode(df.CK_ISNN_C, 0)
	return q
}
func (q *SummaryWithdrawalCQ) AddOrderBy_WithdrawalReasonCode_Asc() *SummaryWithdrawalCQ {
	q.BaseConditionQuery.RegOBA("withdrawalReasonCode")
	return q
}
func (q *SummaryWithdrawalCQ) AddOrderBy_WithdrawalReasonCode_Desc() *SummaryWithdrawalCQ {
	q.BaseConditionQuery.RegOBD("withdrawalReasonCode")
	return q
}
func (q *SummaryWithdrawalCQ) regWithdrawalReasonCode(key *df.ConditionKey, value interface{}) {
	if q.WithdrawalReasonCode == nil {
		q.WithdrawalReasonCode = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.WithdrawalReasonCode, "withdrawalReasonCode")
}

func (q *SummaryWithdrawalCQ) getCValueWithdrawalReasonText() *df.ConditionValue {
	if q.WithdrawalReasonText == nil {
		q.WithdrawalReasonText = new(df.ConditionValue)
	}
	return q.WithdrawalReasonText
}


func (q *SummaryWithdrawalCQ) SetWithdrawalReasonText_Equal(value string) *SummaryWithdrawalCQ {
	q.regWithdrawalReasonText(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *SummaryWithdrawalCQ) SetWithdrawalReasonText_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueWithdrawalReasonText(), "withdrawalReasonText")
}
func (q *SummaryWithdrawalCQ) SetWithdrawalReasonText_NotEqual(value string) *SummaryWithdrawalCQ {
	q.regWithdrawalReasonText(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryWithdrawalCQ) SetWithdrawalReasonText_GreaterThan(value string) *SummaryWithdrawalCQ {
	q.regWithdrawalReasonText(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryWithdrawalCQ) SetWithdrawalReasonText_LessThan(value string) *SummaryWithdrawalCQ {
	q.regWithdrawalReasonText(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryWithdrawalCQ) SetWithdrawalReasonText_GreaterEqualThan(value string) *SummaryWithdrawalCQ {
	q.regWithdrawalReasonText(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SummaryWithdrawalCQ) SetWithdrawalReasonText_LessEqualThan(value string) *SummaryWithdrawalCQ {
	q.regWithdrawalReasonText(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryWithdrawalCQ) SetWithdrawalReasonText_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueWithdrawalReasonText(), "withdrawalReasonText", option)
}

func (q *SummaryWithdrawalCQ) SetWithdrawalReasonText_PrefixSearch(value string) error {
	return q.SetWithdrawalReasonText_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SummaryWithdrawalCQ) SetWithdrawalReasonText_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueWithdrawalReasonText(), "withdrawalReasonText", option)
}



func (q *SummaryWithdrawalCQ) SetWithdrawalReasonText_IsNull() *SummaryWithdrawalCQ {
	q.regWithdrawalReasonText(df.CK_ISN_C, 0)
	return q
}
func (q *SummaryWithdrawalCQ) SetWithdrawalReasonText_IsNullOrEmpty() *SummaryWithdrawalCQ {
	q.regWithdrawalReasonText(df.CK_ISNOE_C, 0)
	return q
}
func (q *SummaryWithdrawalCQ) SetWithdrawalReasonText_IsNotNull() *SummaryWithdrawalCQ {
	q.regWithdrawalReasonText(df.CK_ISNN_C, 0)
	return q
}
func (q *SummaryWithdrawalCQ) AddOrderBy_WithdrawalReasonText_Asc() *SummaryWithdrawalCQ {
	q.BaseConditionQuery.RegOBA("withdrawalReasonText")
	return q
}
func (q *SummaryWithdrawalCQ) AddOrderBy_WithdrawalReasonText_Desc() *SummaryWithdrawalCQ {
	q.BaseConditionQuery.RegOBD("withdrawalReasonText")
	return q
}
func (q *SummaryWithdrawalCQ) regWithdrawalReasonText(key *df.ConditionKey, value interface{}) {
	if q.WithdrawalReasonText == nil {
		q.WithdrawalReasonText = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.WithdrawalReasonText, "withdrawalReasonText")
}

func (q *SummaryWithdrawalCQ) getCValueWithdrawalReasonInputText() *df.ConditionValue {
	if q.WithdrawalReasonInputText == nil {
		q.WithdrawalReasonInputText = new(df.ConditionValue)
	}
	return q.WithdrawalReasonInputText
}


func (q *SummaryWithdrawalCQ) SetWithdrawalReasonInputText_Equal(value string) *SummaryWithdrawalCQ {
	q.regWithdrawalReasonInputText(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *SummaryWithdrawalCQ) SetWithdrawalReasonInputText_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueWithdrawalReasonInputText(), "withdrawalReasonInputText")
}
func (q *SummaryWithdrawalCQ) SetWithdrawalReasonInputText_NotEqual(value string) *SummaryWithdrawalCQ {
	q.regWithdrawalReasonInputText(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryWithdrawalCQ) SetWithdrawalReasonInputText_GreaterThan(value string) *SummaryWithdrawalCQ {
	q.regWithdrawalReasonInputText(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryWithdrawalCQ) SetWithdrawalReasonInputText_LessThan(value string) *SummaryWithdrawalCQ {
	q.regWithdrawalReasonInputText(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryWithdrawalCQ) SetWithdrawalReasonInputText_GreaterEqualThan(value string) *SummaryWithdrawalCQ {
	q.regWithdrawalReasonInputText(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SummaryWithdrawalCQ) SetWithdrawalReasonInputText_LessEqualThan(value string) *SummaryWithdrawalCQ {
	q.regWithdrawalReasonInputText(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryWithdrawalCQ) SetWithdrawalReasonInputText_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueWithdrawalReasonInputText(), "withdrawalReasonInputText", option)
}

func (q *SummaryWithdrawalCQ) SetWithdrawalReasonInputText_PrefixSearch(value string) error {
	return q.SetWithdrawalReasonInputText_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SummaryWithdrawalCQ) SetWithdrawalReasonInputText_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueWithdrawalReasonInputText(), "withdrawalReasonInputText", option)
}



func (q *SummaryWithdrawalCQ) SetWithdrawalReasonInputText_IsNull() *SummaryWithdrawalCQ {
	q.regWithdrawalReasonInputText(df.CK_ISN_C, 0)
	return q
}
func (q *SummaryWithdrawalCQ) SetWithdrawalReasonInputText_IsNullOrEmpty() *SummaryWithdrawalCQ {
	q.regWithdrawalReasonInputText(df.CK_ISNOE_C, 0)
	return q
}
func (q *SummaryWithdrawalCQ) SetWithdrawalReasonInputText_IsNotNull() *SummaryWithdrawalCQ {
	q.regWithdrawalReasonInputText(df.CK_ISNN_C, 0)
	return q
}
func (q *SummaryWithdrawalCQ) AddOrderBy_WithdrawalReasonInputText_Asc() *SummaryWithdrawalCQ {
	q.BaseConditionQuery.RegOBA("withdrawalReasonInputText")
	return q
}
func (q *SummaryWithdrawalCQ) AddOrderBy_WithdrawalReasonInputText_Desc() *SummaryWithdrawalCQ {
	q.BaseConditionQuery.RegOBD("withdrawalReasonInputText")
	return q
}
func (q *SummaryWithdrawalCQ) regWithdrawalReasonInputText(key *df.ConditionKey, value interface{}) {
	if q.WithdrawalReasonInputText == nil {
		q.WithdrawalReasonInputText = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.WithdrawalReasonInputText, "withdrawalReasonInputText")
}

func (q *SummaryWithdrawalCQ) getCValueWithdrawalDatetime() *df.ConditionValue {
	if q.WithdrawalDatetime == nil {
		q.WithdrawalDatetime = new(df.ConditionValue)
	}
	return q.WithdrawalDatetime
}




func (q *SummaryWithdrawalCQ) SetWithdrawalDatetime_Equal(value df.MysqlTimestamp) *SummaryWithdrawalCQ {
	q.regWithdrawalDatetime(df.CK_EQ_C, value)
	return q
}


func (q *SummaryWithdrawalCQ) SetWithdrawalDatetime_GreaterThan(value df.MysqlTimestamp) *SummaryWithdrawalCQ {
	q.regWithdrawalDatetime(df.CK_GT_C, value)
	return q
}

func (q *SummaryWithdrawalCQ) SetWithdrawalDatetime_LessThan(value df.MysqlTimestamp) *SummaryWithdrawalCQ {
	q.regWithdrawalDatetime(df.CK_LT_C, value)
	return q
}

func (q *SummaryWithdrawalCQ) SetWithdrawalDatetime_GreaterEqual(value df.MysqlTimestamp) *SummaryWithdrawalCQ {
	q.regWithdrawalDatetime(df.CK_GE_C, value)
	return q
}

func (q *SummaryWithdrawalCQ) SetWithdrawalDatetime_LessEqual(value df.MysqlTimestamp) *SummaryWithdrawalCQ {
	q.regWithdrawalDatetime(df.CK_LE_C, value)
	return q
}

func (q *SummaryWithdrawalCQ) AddOrderBy_WithdrawalDatetime_Asc() *SummaryWithdrawalCQ {
	q.BaseConditionQuery.RegOBA("withdrawalDatetime")
	return q
}
func (q *SummaryWithdrawalCQ) AddOrderBy_WithdrawalDatetime_Desc() *SummaryWithdrawalCQ {
	q.BaseConditionQuery.RegOBD("withdrawalDatetime")
	return q
}
func (q *SummaryWithdrawalCQ) regWithdrawalDatetime(key *df.ConditionKey, value interface{}) {
	if q.WithdrawalDatetime == nil {
		q.WithdrawalDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.WithdrawalDatetime, "withdrawalDatetime")
}

func (q *SummaryWithdrawalCQ) getCValueMemberStatusCode() *df.ConditionValue {
	if q.MemberStatusCode == nil {
		q.MemberStatusCode = new(df.ConditionValue)
	}
	return q.MemberStatusCode
}


func (q *SummaryWithdrawalCQ) SetMemberStatusCode_Equal(value string) *SummaryWithdrawalCQ {
	q.regMemberStatusCode(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *SummaryWithdrawalCQ) SetMemberStatusCode_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueMemberStatusCode(), "memberStatusCode")
}
func (q *SummaryWithdrawalCQ) SetMemberStatusCode_NotEqual(value string) *SummaryWithdrawalCQ {
	q.regMemberStatusCode(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryWithdrawalCQ) SetMemberStatusCode_GreaterThan(value string) *SummaryWithdrawalCQ {
	q.regMemberStatusCode(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryWithdrawalCQ) SetMemberStatusCode_LessThan(value string) *SummaryWithdrawalCQ {
	q.regMemberStatusCode(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryWithdrawalCQ) SetMemberStatusCode_GreaterEqualThan(value string) *SummaryWithdrawalCQ {
	q.regMemberStatusCode(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SummaryWithdrawalCQ) SetMemberStatusCode_LessEqualThan(value string) *SummaryWithdrawalCQ {
	q.regMemberStatusCode(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryWithdrawalCQ) SetMemberStatusCode_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueMemberStatusCode(), "memberStatusCode", option)
}

func (q *SummaryWithdrawalCQ) SetMemberStatusCode_PrefixSearch(value string) error {
	return q.SetMemberStatusCode_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SummaryWithdrawalCQ) SetMemberStatusCode_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueMemberStatusCode(), "memberStatusCode", option)
}



func (q *SummaryWithdrawalCQ) SetMemberStatusCode_IsNull() *SummaryWithdrawalCQ {
	q.regMemberStatusCode(df.CK_ISN_C, 0)
	return q
}
func (q *SummaryWithdrawalCQ) SetMemberStatusCode_IsNullOrEmpty() *SummaryWithdrawalCQ {
	q.regMemberStatusCode(df.CK_ISNOE_C, 0)
	return q
}
func (q *SummaryWithdrawalCQ) SetMemberStatusCode_IsNotNull() *SummaryWithdrawalCQ {
	q.regMemberStatusCode(df.CK_ISNN_C, 0)
	return q
}
func (q *SummaryWithdrawalCQ) AddOrderBy_MemberStatusCode_Asc() *SummaryWithdrawalCQ {
	q.BaseConditionQuery.RegOBA("memberStatusCode")
	return q
}
func (q *SummaryWithdrawalCQ) AddOrderBy_MemberStatusCode_Desc() *SummaryWithdrawalCQ {
	q.BaseConditionQuery.RegOBD("memberStatusCode")
	return q
}
func (q *SummaryWithdrawalCQ) regMemberStatusCode(key *df.ConditionKey, value interface{}) {
	if q.MemberStatusCode == nil {
		q.MemberStatusCode = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.MemberStatusCode, "memberStatusCode")
}

func (q *SummaryWithdrawalCQ) getCValueMemberStatusName() *df.ConditionValue {
	if q.MemberStatusName == nil {
		q.MemberStatusName = new(df.ConditionValue)
	}
	return q.MemberStatusName
}


func (q *SummaryWithdrawalCQ) SetMemberStatusName_Equal(value string) *SummaryWithdrawalCQ {
	q.regMemberStatusName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *SummaryWithdrawalCQ) SetMemberStatusName_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueMemberStatusName(), "memberStatusName")
}
func (q *SummaryWithdrawalCQ) SetMemberStatusName_NotEqual(value string) *SummaryWithdrawalCQ {
	q.regMemberStatusName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryWithdrawalCQ) SetMemberStatusName_GreaterThan(value string) *SummaryWithdrawalCQ {
	q.regMemberStatusName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryWithdrawalCQ) SetMemberStatusName_LessThan(value string) *SummaryWithdrawalCQ {
	q.regMemberStatusName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryWithdrawalCQ) SetMemberStatusName_GreaterEqualThan(value string) *SummaryWithdrawalCQ {
	q.regMemberStatusName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *SummaryWithdrawalCQ) SetMemberStatusName_LessEqualThan(value string) *SummaryWithdrawalCQ {
	q.regMemberStatusName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *SummaryWithdrawalCQ) SetMemberStatusName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueMemberStatusName(), "memberStatusName", option)
}

func (q *SummaryWithdrawalCQ) SetMemberStatusName_PrefixSearch(value string) error {
	return q.SetMemberStatusName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *SummaryWithdrawalCQ) SetMemberStatusName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueMemberStatusName(), "memberStatusName", option)
}



func (q *SummaryWithdrawalCQ) SetMemberStatusName_IsNull() *SummaryWithdrawalCQ {
	q.regMemberStatusName(df.CK_ISN_C, 0)
	return q
}
func (q *SummaryWithdrawalCQ) SetMemberStatusName_IsNullOrEmpty() *SummaryWithdrawalCQ {
	q.regMemberStatusName(df.CK_ISNOE_C, 0)
	return q
}
func (q *SummaryWithdrawalCQ) SetMemberStatusName_IsNotNull() *SummaryWithdrawalCQ {
	q.regMemberStatusName(df.CK_ISNN_C, 0)
	return q
}
func (q *SummaryWithdrawalCQ) AddOrderBy_MemberStatusName_Asc() *SummaryWithdrawalCQ {
	q.BaseConditionQuery.RegOBA("memberStatusName")
	return q
}
func (q *SummaryWithdrawalCQ) AddOrderBy_MemberStatusName_Desc() *SummaryWithdrawalCQ {
	q.BaseConditionQuery.RegOBD("memberStatusName")
	return q
}
func (q *SummaryWithdrawalCQ) regMemberStatusName(key *df.ConditionKey, value interface{}) {
	if q.MemberStatusName == nil {
		q.MemberStatusName = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.MemberStatusName, "memberStatusName")
}

func (q *SummaryWithdrawalCQ) getCValueMaxPurchasePrice() *df.ConditionValue {
	if q.MaxPurchasePrice == nil {
		q.MaxPurchasePrice = new(df.ConditionValue)
	}
	return q.MaxPurchasePrice
}



func (q *SummaryWithdrawalCQ) SetMaxPurchasePrice_Equal(value int64) *SummaryWithdrawalCQ {
	q.regMaxPurchasePrice(df.CK_EQ_C, value)
	return q
}
func (q *SummaryWithdrawalCQ) SetMaxPurchasePrice_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueMaxPurchasePrice(), "maxPurchasePrice")
}
func (q *SummaryWithdrawalCQ) SetMaxPurchasePrice_NotEqual(value int64) *SummaryWithdrawalCQ {
	q.regMaxPurchasePrice(df.CK_NE_C, value)
	return q
}

func (q *SummaryWithdrawalCQ) SetMaxPurchasePrice_GreaterThan(value int64) *SummaryWithdrawalCQ {
	q.regMaxPurchasePrice(df.CK_GT_C, value)
	return q
}

func (q *SummaryWithdrawalCQ) SetMaxPurchasePrice_LessThan(value int64) *SummaryWithdrawalCQ {
	q.regMaxPurchasePrice(df.CK_LT_C, value)
	return q
}

func (q *SummaryWithdrawalCQ) SetMaxPurchasePrice_GreaterEqual(value int64) *SummaryWithdrawalCQ {
	q.regMaxPurchasePrice(df.CK_GE_C, value)
	return q
}

func (q *SummaryWithdrawalCQ) SetMaxPurchasePrice_LessEqual(value int64) *SummaryWithdrawalCQ {
	q.regMaxPurchasePrice(df.CK_LE_C, value)
	return q
}
func (q *SummaryWithdrawalCQ) SetMaxPurchasePrice_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueMaxPurchasePrice(),"maxPurchasePrice",rangeOfOption)
}	


func (q *SummaryWithdrawalCQ) SetMaxPurchasePrice_IsNull() *SummaryWithdrawalCQ {
	q.regMaxPurchasePrice(df.CK_ISN_C, 0)
	return q
}
func (q *SummaryWithdrawalCQ) SetMaxPurchasePrice_IsNotNull() *SummaryWithdrawalCQ {
	q.regMaxPurchasePrice(df.CK_ISNN_C, 0)
	return q
}
func (q *SummaryWithdrawalCQ) AddOrderBy_MaxPurchasePrice_Asc() *SummaryWithdrawalCQ {
	q.BaseConditionQuery.RegOBA("maxPurchasePrice")
	return q
}
func (q *SummaryWithdrawalCQ) AddOrderBy_MaxPurchasePrice_Desc() *SummaryWithdrawalCQ {
	q.BaseConditionQuery.RegOBD("maxPurchasePrice")
	return q
}
func (q *SummaryWithdrawalCQ) regMaxPurchasePrice(key *df.ConditionKey, value interface{}) {
	if q.MaxPurchasePrice == nil {
		q.MaxPurchasePrice = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.MaxPurchasePrice, "maxPurchasePrice")
}


func CreateSummaryWithdrawalCQ(referrerQuery *df.ConditionQuery, sqlClause *df.SqlClause, aliasName string, nestlevel int8) *SummaryWithdrawalCQ {
	cq := new(SummaryWithdrawalCQ)
	cq.BaseConditionQuery = new(df.BaseConditionQuery)
	cq.BaseConditionQuery.TableDbName = "SummaryWithdrawal"
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