package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type MemberWithdrawalCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	MemberId *df.ConditionValue
	WithdrawalReasonCode *df.ConditionValue
	WithdrawalReasonInputText *df.ConditionValue
	WithdrawalDatetime *df.ConditionValue
	RegisterDatetime *df.ConditionValue
	RegisterUser *df.ConditionValue
	UpdateDatetime *df.ConditionValue
	UpdateUser *df.ConditionValue
}

func (q *MemberWithdrawalCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *MemberWithdrawalCQ) getCValueMemberId() *df.ConditionValue {
	if q.MemberId == nil {
		q.MemberId = new(df.ConditionValue)
	}
	return q.MemberId
}



func (q *MemberWithdrawalCQ) SetMemberId_Equal(value int64) *MemberWithdrawalCQ {
	q.regMemberId(df.CK_EQ_C, value)
	return q
}

func (q *MemberWithdrawalCQ) SetMemberId_NotEqual(value int64) *MemberWithdrawalCQ {
	q.regMemberId(df.CK_NE_C, value)
	return q
}

func (q *MemberWithdrawalCQ) SetMemberId_GreaterThan(value int64) *MemberWithdrawalCQ {
	q.regMemberId(df.CK_GT_C, value)
	return q
}

func (q *MemberWithdrawalCQ) SetMemberId_LessThan(value int64) *MemberWithdrawalCQ {
	q.regMemberId(df.CK_LT_C, value)
	return q
}

func (q *MemberWithdrawalCQ) SetMemberId_GreaterEqual(value int64) *MemberWithdrawalCQ {
	q.regMemberId(df.CK_GE_C, value)
	return q
}

func (q *MemberWithdrawalCQ) SetMemberId_LessEqual(value int64) *MemberWithdrawalCQ {
	q.regMemberId(df.CK_LE_C, value)
	return q
}
func (q *MemberWithdrawalCQ) SetMemberId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueMemberId(),"memberId",rangeOfOption)
}	


func (q *MemberWithdrawalCQ) SetMemberId_IsNull() *MemberWithdrawalCQ {
	q.regMemberId(df.CK_ISN_C, 0)
	return q
}
func (q *MemberWithdrawalCQ) SetMemberId_IsNotNull() *MemberWithdrawalCQ {
	q.regMemberId(df.CK_ISNN_C, 0)
	return q
}
func (q *MemberWithdrawalCQ) AddOrderBy_MemberId_Asc() *MemberWithdrawalCQ {
	q.BaseConditionQuery.RegOBA("memberId")
	return q
}
func (q *MemberWithdrawalCQ) AddOrderBy_MemberId_Desc() *MemberWithdrawalCQ {
	q.BaseConditionQuery.RegOBD("memberId")
	return q
}
func (q *MemberWithdrawalCQ) regMemberId(key *df.ConditionKey, value interface{}) {
	if q.MemberId == nil {
		q.MemberId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.MemberId, "memberId")
}

func (q *MemberWithdrawalCQ) getCValueWithdrawalReasonCode() *df.ConditionValue {
	if q.WithdrawalReasonCode == nil {
		q.WithdrawalReasonCode = new(df.ConditionValue)
	}
	return q.WithdrawalReasonCode
}


func (q *MemberWithdrawalCQ) SetWithdrawalReasonCode_Equal(value string) *MemberWithdrawalCQ {
	q.regWithdrawalReasonCode(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *MemberWithdrawalCQ) SetWithdrawalReasonCode_NotEqual(value string) *MemberWithdrawalCQ {
	q.regWithdrawalReasonCode(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberWithdrawalCQ) SetWithdrawalReasonCode_GreaterThan(value string) *MemberWithdrawalCQ {
	q.regWithdrawalReasonCode(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberWithdrawalCQ) SetWithdrawalReasonCode_LessThan(value string) *MemberWithdrawalCQ {
	q.regWithdrawalReasonCode(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberWithdrawalCQ) SetWithdrawalReasonCode_GreaterEqualThan(value string) *MemberWithdrawalCQ {
	q.regWithdrawalReasonCode(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *MemberWithdrawalCQ) SetWithdrawalReasonCode_LessEqualThan(value string) *MemberWithdrawalCQ {
	q.regWithdrawalReasonCode(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberWithdrawalCQ) SetWithdrawalReasonCode_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueWithdrawalReasonCode(), "withdrawalReasonCode", option)
}

func (q *MemberWithdrawalCQ) SetWithdrawalReasonCode_PrefixSearch(value string) error {
	return q.SetWithdrawalReasonCode_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *MemberWithdrawalCQ) SetWithdrawalReasonCode_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueWithdrawalReasonCode(), "withdrawalReasonCode", option)
}



func (q *MemberWithdrawalCQ) SetWithdrawalReasonCode_IsNull() *MemberWithdrawalCQ {
	q.regWithdrawalReasonCode(df.CK_ISN_C, 0)
	return q
}
func (q *MemberWithdrawalCQ) SetWithdrawalReasonCode_IsNullOrEmpty() *MemberWithdrawalCQ {
	q.regWithdrawalReasonCode(df.CK_ISNOE_C, 0)
	return q
}
func (q *MemberWithdrawalCQ) SetWithdrawalReasonCode_IsNotNull() *MemberWithdrawalCQ {
	q.regWithdrawalReasonCode(df.CK_ISNN_C, 0)
	return q
}
func (q *MemberWithdrawalCQ) AddOrderBy_WithdrawalReasonCode_Asc() *MemberWithdrawalCQ {
	q.BaseConditionQuery.RegOBA("withdrawalReasonCode")
	return q
}
func (q *MemberWithdrawalCQ) AddOrderBy_WithdrawalReasonCode_Desc() *MemberWithdrawalCQ {
	q.BaseConditionQuery.RegOBD("withdrawalReasonCode")
	return q
}
func (q *MemberWithdrawalCQ) regWithdrawalReasonCode(key *df.ConditionKey, value interface{}) {
	if q.WithdrawalReasonCode == nil {
		q.WithdrawalReasonCode = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.WithdrawalReasonCode, "withdrawalReasonCode")
}

func (q *MemberWithdrawalCQ) getCValueWithdrawalReasonInputText() *df.ConditionValue {
	if q.WithdrawalReasonInputText == nil {
		q.WithdrawalReasonInputText = new(df.ConditionValue)
	}
	return q.WithdrawalReasonInputText
}


func (q *MemberWithdrawalCQ) SetWithdrawalReasonInputText_Equal(value string) *MemberWithdrawalCQ {
	q.regWithdrawalReasonInputText(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *MemberWithdrawalCQ) SetWithdrawalReasonInputText_NotEqual(value string) *MemberWithdrawalCQ {
	q.regWithdrawalReasonInputText(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberWithdrawalCQ) SetWithdrawalReasonInputText_GreaterThan(value string) *MemberWithdrawalCQ {
	q.regWithdrawalReasonInputText(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberWithdrawalCQ) SetWithdrawalReasonInputText_LessThan(value string) *MemberWithdrawalCQ {
	q.regWithdrawalReasonInputText(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberWithdrawalCQ) SetWithdrawalReasonInputText_GreaterEqualThan(value string) *MemberWithdrawalCQ {
	q.regWithdrawalReasonInputText(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *MemberWithdrawalCQ) SetWithdrawalReasonInputText_LessEqualThan(value string) *MemberWithdrawalCQ {
	q.regWithdrawalReasonInputText(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberWithdrawalCQ) SetWithdrawalReasonInputText_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueWithdrawalReasonInputText(), "withdrawalReasonInputText", option)
}

func (q *MemberWithdrawalCQ) SetWithdrawalReasonInputText_PrefixSearch(value string) error {
	return q.SetWithdrawalReasonInputText_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *MemberWithdrawalCQ) SetWithdrawalReasonInputText_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueWithdrawalReasonInputText(), "withdrawalReasonInputText", option)
}



func (q *MemberWithdrawalCQ) SetWithdrawalReasonInputText_IsNull() *MemberWithdrawalCQ {
	q.regWithdrawalReasonInputText(df.CK_ISN_C, 0)
	return q
}
func (q *MemberWithdrawalCQ) SetWithdrawalReasonInputText_IsNullOrEmpty() *MemberWithdrawalCQ {
	q.regWithdrawalReasonInputText(df.CK_ISNOE_C, 0)
	return q
}
func (q *MemberWithdrawalCQ) SetWithdrawalReasonInputText_IsNotNull() *MemberWithdrawalCQ {
	q.regWithdrawalReasonInputText(df.CK_ISNN_C, 0)
	return q
}
func (q *MemberWithdrawalCQ) AddOrderBy_WithdrawalReasonInputText_Asc() *MemberWithdrawalCQ {
	q.BaseConditionQuery.RegOBA("withdrawalReasonInputText")
	return q
}
func (q *MemberWithdrawalCQ) AddOrderBy_WithdrawalReasonInputText_Desc() *MemberWithdrawalCQ {
	q.BaseConditionQuery.RegOBD("withdrawalReasonInputText")
	return q
}
func (q *MemberWithdrawalCQ) regWithdrawalReasonInputText(key *df.ConditionKey, value interface{}) {
	if q.WithdrawalReasonInputText == nil {
		q.WithdrawalReasonInputText = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.WithdrawalReasonInputText, "withdrawalReasonInputText")
}

func (q *MemberWithdrawalCQ) getCValueWithdrawalDatetime() *df.ConditionValue {
	if q.WithdrawalDatetime == nil {
		q.WithdrawalDatetime = new(df.ConditionValue)
	}
	return q.WithdrawalDatetime
}




func (q *MemberWithdrawalCQ) SetWithdrawalDatetime_Equal(value df.MysqlTimestamp) *MemberWithdrawalCQ {
	q.regWithdrawalDatetime(df.CK_EQ_C, value)
	return q
}


func (q *MemberWithdrawalCQ) SetWithdrawalDatetime_GreaterThan(value df.MysqlTimestamp) *MemberWithdrawalCQ {
	q.regWithdrawalDatetime(df.CK_GT_C, value)
	return q
}

func (q *MemberWithdrawalCQ) SetWithdrawalDatetime_LessThan(value df.MysqlTimestamp) *MemberWithdrawalCQ {
	q.regWithdrawalDatetime(df.CK_LT_C, value)
	return q
}

func (q *MemberWithdrawalCQ) SetWithdrawalDatetime_GreaterEqual(value df.MysqlTimestamp) *MemberWithdrawalCQ {
	q.regWithdrawalDatetime(df.CK_GE_C, value)
	return q
}

func (q *MemberWithdrawalCQ) SetWithdrawalDatetime_LessEqual(value df.MysqlTimestamp) *MemberWithdrawalCQ {
	q.regWithdrawalDatetime(df.CK_LE_C, value)
	return q
}

func (q *MemberWithdrawalCQ) AddOrderBy_WithdrawalDatetime_Asc() *MemberWithdrawalCQ {
	q.BaseConditionQuery.RegOBA("withdrawalDatetime")
	return q
}
func (q *MemberWithdrawalCQ) AddOrderBy_WithdrawalDatetime_Desc() *MemberWithdrawalCQ {
	q.BaseConditionQuery.RegOBD("withdrawalDatetime")
	return q
}
func (q *MemberWithdrawalCQ) regWithdrawalDatetime(key *df.ConditionKey, value interface{}) {
	if q.WithdrawalDatetime == nil {
		q.WithdrawalDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.WithdrawalDatetime, "withdrawalDatetime")
}

func (q *MemberWithdrawalCQ) getCValueRegisterDatetime() *df.ConditionValue {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	return q.RegisterDatetime
}




func (q *MemberWithdrawalCQ) SetRegisterDatetime_Equal(value df.MysqlTimestamp) *MemberWithdrawalCQ {
	q.regRegisterDatetime(df.CK_EQ_C, value)
	return q
}


func (q *MemberWithdrawalCQ) SetRegisterDatetime_GreaterThan(value df.MysqlTimestamp) *MemberWithdrawalCQ {
	q.regRegisterDatetime(df.CK_GT_C, value)
	return q
}

func (q *MemberWithdrawalCQ) SetRegisterDatetime_LessThan(value df.MysqlTimestamp) *MemberWithdrawalCQ {
	q.regRegisterDatetime(df.CK_LT_C, value)
	return q
}

func (q *MemberWithdrawalCQ) SetRegisterDatetime_GreaterEqual(value df.MysqlTimestamp) *MemberWithdrawalCQ {
	q.regRegisterDatetime(df.CK_GE_C, value)
	return q
}

func (q *MemberWithdrawalCQ) SetRegisterDatetime_LessEqual(value df.MysqlTimestamp) *MemberWithdrawalCQ {
	q.regRegisterDatetime(df.CK_LE_C, value)
	return q
}

func (q *MemberWithdrawalCQ) AddOrderBy_RegisterDatetime_Asc() *MemberWithdrawalCQ {
	q.BaseConditionQuery.RegOBA("registerDatetime")
	return q
}
func (q *MemberWithdrawalCQ) AddOrderBy_RegisterDatetime_Desc() *MemberWithdrawalCQ {
	q.BaseConditionQuery.RegOBD("registerDatetime")
	return q
}
func (q *MemberWithdrawalCQ) regRegisterDatetime(key *df.ConditionKey, value interface{}) {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterDatetime, "registerDatetime")
}

func (q *MemberWithdrawalCQ) getCValueRegisterUser() *df.ConditionValue {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	return q.RegisterUser
}


func (q *MemberWithdrawalCQ) SetRegisterUser_Equal(value string) *MemberWithdrawalCQ {
	q.regRegisterUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *MemberWithdrawalCQ) SetRegisterUser_NotEqual(value string) *MemberWithdrawalCQ {
	q.regRegisterUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberWithdrawalCQ) SetRegisterUser_GreaterThan(value string) *MemberWithdrawalCQ {
	q.regRegisterUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberWithdrawalCQ) SetRegisterUser_LessThan(value string) *MemberWithdrawalCQ {
	q.regRegisterUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberWithdrawalCQ) SetRegisterUser_GreaterEqualThan(value string) *MemberWithdrawalCQ {
	q.regRegisterUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *MemberWithdrawalCQ) SetRegisterUser_LessEqualThan(value string) *MemberWithdrawalCQ {
	q.regRegisterUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberWithdrawalCQ) SetRegisterUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}

func (q *MemberWithdrawalCQ) SetRegisterUser_PrefixSearch(value string) error {
	return q.SetRegisterUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *MemberWithdrawalCQ) SetRegisterUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}



func (q *MemberWithdrawalCQ) AddOrderBy_RegisterUser_Asc() *MemberWithdrawalCQ {
	q.BaseConditionQuery.RegOBA("registerUser")
	return q
}
func (q *MemberWithdrawalCQ) AddOrderBy_RegisterUser_Desc() *MemberWithdrawalCQ {
	q.BaseConditionQuery.RegOBD("registerUser")
	return q
}
func (q *MemberWithdrawalCQ) regRegisterUser(key *df.ConditionKey, value interface{}) {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterUser, "registerUser")
}

func (q *MemberWithdrawalCQ) getCValueUpdateDatetime() *df.ConditionValue {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	return q.UpdateDatetime
}




func (q *MemberWithdrawalCQ) SetUpdateDatetime_Equal(value df.MysqlTimestamp) *MemberWithdrawalCQ {
	q.regUpdateDatetime(df.CK_EQ_C, value)
	return q
}


func (q *MemberWithdrawalCQ) SetUpdateDatetime_GreaterThan(value df.MysqlTimestamp) *MemberWithdrawalCQ {
	q.regUpdateDatetime(df.CK_GT_C, value)
	return q
}

func (q *MemberWithdrawalCQ) SetUpdateDatetime_LessThan(value df.MysqlTimestamp) *MemberWithdrawalCQ {
	q.regUpdateDatetime(df.CK_LT_C, value)
	return q
}

func (q *MemberWithdrawalCQ) SetUpdateDatetime_GreaterEqual(value df.MysqlTimestamp) *MemberWithdrawalCQ {
	q.regUpdateDatetime(df.CK_GE_C, value)
	return q
}

func (q *MemberWithdrawalCQ) SetUpdateDatetime_LessEqual(value df.MysqlTimestamp) *MemberWithdrawalCQ {
	q.regUpdateDatetime(df.CK_LE_C, value)
	return q
}

func (q *MemberWithdrawalCQ) AddOrderBy_UpdateDatetime_Asc() *MemberWithdrawalCQ {
	q.BaseConditionQuery.RegOBA("updateDatetime")
	return q
}
func (q *MemberWithdrawalCQ) AddOrderBy_UpdateDatetime_Desc() *MemberWithdrawalCQ {
	q.BaseConditionQuery.RegOBD("updateDatetime")
	return q
}
func (q *MemberWithdrawalCQ) regUpdateDatetime(key *df.ConditionKey, value interface{}) {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateDatetime, "updateDatetime")
}

func (q *MemberWithdrawalCQ) getCValueUpdateUser() *df.ConditionValue {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	return q.UpdateUser
}


func (q *MemberWithdrawalCQ) SetUpdateUser_Equal(value string) *MemberWithdrawalCQ {
	q.regUpdateUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *MemberWithdrawalCQ) SetUpdateUser_NotEqual(value string) *MemberWithdrawalCQ {
	q.regUpdateUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberWithdrawalCQ) SetUpdateUser_GreaterThan(value string) *MemberWithdrawalCQ {
	q.regUpdateUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberWithdrawalCQ) SetUpdateUser_LessThan(value string) *MemberWithdrawalCQ {
	q.regUpdateUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberWithdrawalCQ) SetUpdateUser_GreaterEqualThan(value string) *MemberWithdrawalCQ {
	q.regUpdateUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *MemberWithdrawalCQ) SetUpdateUser_LessEqualThan(value string) *MemberWithdrawalCQ {
	q.regUpdateUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberWithdrawalCQ) SetUpdateUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}

func (q *MemberWithdrawalCQ) SetUpdateUser_PrefixSearch(value string) error {
	return q.SetUpdateUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *MemberWithdrawalCQ) SetUpdateUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}



func (q *MemberWithdrawalCQ) AddOrderBy_UpdateUser_Asc() *MemberWithdrawalCQ {
	q.BaseConditionQuery.RegOBA("updateUser")
	return q
}
func (q *MemberWithdrawalCQ) AddOrderBy_UpdateUser_Desc() *MemberWithdrawalCQ {
	q.BaseConditionQuery.RegOBD("updateUser")
	return q
}
func (q *MemberWithdrawalCQ) regUpdateUser(key *df.ConditionKey, value interface{}) {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateUser, "updateUser")
}

