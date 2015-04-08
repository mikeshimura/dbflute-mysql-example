package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type MemberCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	MemberId *df.ConditionValue
	MemberName *df.ConditionValue
	MemberAccount *df.ConditionValue
	MemberStatusCode *df.ConditionValue
	FormalizedDatetime *df.ConditionValue
	Birthdate *df.ConditionValue
	RegisterDatetime *df.ConditionValue
	RegisterUser *df.ConditionValue
	UpdateDatetime *df.ConditionValue
	UpdateUser *df.ConditionValue
	VersionNo *df.ConditionValue
}

func (q *MemberCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *MemberCQ) getCValueMemberId() *df.ConditionValue {
	if q.MemberId == nil {
		q.MemberId = new(df.ConditionValue)
	}
	return q.MemberId
}



func (q *MemberCQ) SetMemberId_Equal(value int64) *MemberCQ {
	q.regMemberId(df.CK_EQ_C, value)
	return q
}

func (q *MemberCQ) SetMemberId_NotEqual(value int64) *MemberCQ {
	q.regMemberId(df.CK_NE_C, value)
	return q
}

func (q *MemberCQ) SetMemberId_GreaterThan(value int64) *MemberCQ {
	q.regMemberId(df.CK_GT_C, value)
	return q
}

func (q *MemberCQ) SetMemberId_LessThan(value int64) *MemberCQ {
	q.regMemberId(df.CK_LT_C, value)
	return q
}

func (q *MemberCQ) SetMemberId_GreaterEqual(value int64) *MemberCQ {
	q.regMemberId(df.CK_GE_C, value)
	return q
}

func (q *MemberCQ) SetMemberId_LessEqual(value int64) *MemberCQ {
	q.regMemberId(df.CK_LE_C, value)
	return q
}
func (q *MemberCQ) SetMemberId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueMemberId(),"memberId",rangeOfOption)
}	


func (q *MemberCQ) SetMemberId_IsNull() *MemberCQ {
	q.regMemberId(df.CK_ISN_C, 0)
	return q
}
func (q *MemberCQ) SetMemberId_IsNotNull() *MemberCQ {
	q.regMemberId(df.CK_ISNN_C, 0)
	return q
}
func (q *MemberCQ) AddOrderBy_MemberId_Asc() *MemberCQ {
	q.BaseConditionQuery.RegOBA("memberId")
	return q
}
func (q *MemberCQ) AddOrderBy_MemberId_Desc() *MemberCQ {
	q.BaseConditionQuery.RegOBD("memberId")
	return q
}
func (q *MemberCQ) regMemberId(key *df.ConditionKey, value interface{}) {
	if q.MemberId == nil {
		q.MemberId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.MemberId, "memberId")
}

func (q *MemberCQ) getCValueMemberName() *df.ConditionValue {
	if q.MemberName == nil {
		q.MemberName = new(df.ConditionValue)
	}
	return q.MemberName
}


func (q *MemberCQ) SetMemberName_Equal(value string) *MemberCQ {
	q.regMemberName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *MemberCQ) SetMemberName_NotEqual(value string) *MemberCQ {
	q.regMemberName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberCQ) SetMemberName_GreaterThan(value string) *MemberCQ {
	q.regMemberName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberCQ) SetMemberName_LessThan(value string) *MemberCQ {
	q.regMemberName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberCQ) SetMemberName_GreaterEqualThan(value string) *MemberCQ {
	q.regMemberName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *MemberCQ) SetMemberName_LessEqualThan(value string) *MemberCQ {
	q.regMemberName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberCQ) SetMemberName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueMemberName(), "memberName", option)
}

func (q *MemberCQ) SetMemberName_PrefixSearch(value string) error {
	return q.SetMemberName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *MemberCQ) SetMemberName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueMemberName(), "memberName", option)
}



func (q *MemberCQ) AddOrderBy_MemberName_Asc() *MemberCQ {
	q.BaseConditionQuery.RegOBA("memberName")
	return q
}
func (q *MemberCQ) AddOrderBy_MemberName_Desc() *MemberCQ {
	q.BaseConditionQuery.RegOBD("memberName")
	return q
}
func (q *MemberCQ) regMemberName(key *df.ConditionKey, value interface{}) {
	if q.MemberName == nil {
		q.MemberName = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.MemberName, "memberName")
}

func (q *MemberCQ) getCValueMemberAccount() *df.ConditionValue {
	if q.MemberAccount == nil {
		q.MemberAccount = new(df.ConditionValue)
	}
	return q.MemberAccount
}


func (q *MemberCQ) SetMemberAccount_Equal(value string) *MemberCQ {
	q.regMemberAccount(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *MemberCQ) SetMemberAccount_NotEqual(value string) *MemberCQ {
	q.regMemberAccount(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberCQ) SetMemberAccount_GreaterThan(value string) *MemberCQ {
	q.regMemberAccount(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberCQ) SetMemberAccount_LessThan(value string) *MemberCQ {
	q.regMemberAccount(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberCQ) SetMemberAccount_GreaterEqualThan(value string) *MemberCQ {
	q.regMemberAccount(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *MemberCQ) SetMemberAccount_LessEqualThan(value string) *MemberCQ {
	q.regMemberAccount(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberCQ) SetMemberAccount_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueMemberAccount(), "memberAccount", option)
}

func (q *MemberCQ) SetMemberAccount_PrefixSearch(value string) error {
	return q.SetMemberAccount_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *MemberCQ) SetMemberAccount_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueMemberAccount(), "memberAccount", option)
}



func (q *MemberCQ) AddOrderBy_MemberAccount_Asc() *MemberCQ {
	q.BaseConditionQuery.RegOBA("memberAccount")
	return q
}
func (q *MemberCQ) AddOrderBy_MemberAccount_Desc() *MemberCQ {
	q.BaseConditionQuery.RegOBD("memberAccount")
	return q
}
func (q *MemberCQ) regMemberAccount(key *df.ConditionKey, value interface{}) {
	if q.MemberAccount == nil {
		q.MemberAccount = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.MemberAccount, "memberAccount")
}

func (q *MemberCQ) getCValueMemberStatusCode() *df.ConditionValue {
	if q.MemberStatusCode == nil {
		q.MemberStatusCode = new(df.ConditionValue)
	}
	return q.MemberStatusCode
}


func (q *MemberCQ) SetMemberStatusCode_Equal(value string) *MemberCQ {
	q.regMemberStatusCode(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *MemberCQ) SetMemberStatusCode_NotEqual(value string) *MemberCQ {
	q.regMemberStatusCode(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberCQ) SetMemberStatusCode_GreaterThan(value string) *MemberCQ {
	q.regMemberStatusCode(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberCQ) SetMemberStatusCode_LessThan(value string) *MemberCQ {
	q.regMemberStatusCode(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberCQ) SetMemberStatusCode_GreaterEqualThan(value string) *MemberCQ {
	q.regMemberStatusCode(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *MemberCQ) SetMemberStatusCode_LessEqualThan(value string) *MemberCQ {
	q.regMemberStatusCode(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberCQ) SetMemberStatusCode_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueMemberStatusCode(), "memberStatusCode", option)
}

func (q *MemberCQ) SetMemberStatusCode_PrefixSearch(value string) error {
	return q.SetMemberStatusCode_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *MemberCQ) SetMemberStatusCode_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueMemberStatusCode(), "memberStatusCode", option)
}



func (q *MemberCQ) AddOrderBy_MemberStatusCode_Asc() *MemberCQ {
	q.BaseConditionQuery.RegOBA("memberStatusCode")
	return q
}
func (q *MemberCQ) AddOrderBy_MemberStatusCode_Desc() *MemberCQ {
	q.BaseConditionQuery.RegOBD("memberStatusCode")
	return q
}
func (q *MemberCQ) regMemberStatusCode(key *df.ConditionKey, value interface{}) {
	if q.MemberStatusCode == nil {
		q.MemberStatusCode = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.MemberStatusCode, "memberStatusCode")
}

func (q *MemberCQ) getCValueFormalizedDatetime() *df.ConditionValue {
	if q.FormalizedDatetime == nil {
		q.FormalizedDatetime = new(df.ConditionValue)
	}
	return q.FormalizedDatetime
}




func (q *MemberCQ) SetFormalizedDatetime_Equal(value df.MysqlTimestamp) *MemberCQ {
	q.regFormalizedDatetime(df.CK_EQ_C, value)
	return q
}


func (q *MemberCQ) SetFormalizedDatetime_GreaterThan(value df.MysqlTimestamp) *MemberCQ {
	q.regFormalizedDatetime(df.CK_GT_C, value)
	return q
}

func (q *MemberCQ) SetFormalizedDatetime_LessThan(value df.MysqlTimestamp) *MemberCQ {
	q.regFormalizedDatetime(df.CK_LT_C, value)
	return q
}

func (q *MemberCQ) SetFormalizedDatetime_GreaterEqual(value df.MysqlTimestamp) *MemberCQ {
	q.regFormalizedDatetime(df.CK_GE_C, value)
	return q
}

func (q *MemberCQ) SetFormalizedDatetime_LessEqual(value df.MysqlTimestamp) *MemberCQ {
	q.regFormalizedDatetime(df.CK_LE_C, value)
	return q
}

func (q *MemberCQ) SetFormalizedDatetime_IsNull() *MemberCQ {
	q.regFormalizedDatetime(df.CK_ISN_C, 0)
	return q
}
func (q *MemberCQ) SetFormalizedDatetime_IsNotNull() *MemberCQ {
	q.regFormalizedDatetime(df.CK_ISNN_C, 0)
	return q
}
func (q *MemberCQ) AddOrderBy_FormalizedDatetime_Asc() *MemberCQ {
	q.BaseConditionQuery.RegOBA("formalizedDatetime")
	return q
}
func (q *MemberCQ) AddOrderBy_FormalizedDatetime_Desc() *MemberCQ {
	q.BaseConditionQuery.RegOBD("formalizedDatetime")
	return q
}
func (q *MemberCQ) regFormalizedDatetime(key *df.ConditionKey, value interface{}) {
	if q.FormalizedDatetime == nil {
		q.FormalizedDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.FormalizedDatetime, "formalizedDatetime")
}

func (q *MemberCQ) getCValueBirthdate() *df.ConditionValue {
	if q.Birthdate == nil {
		q.Birthdate = new(df.ConditionValue)
	}
	return q.Birthdate
}




func (q *MemberCQ) SetBirthdate_Equal(value df.MysqlDate) *MemberCQ {
	q.regBirthdate(df.CK_EQ_C, value)
	return q
}


func (q *MemberCQ) SetBirthdate_GreaterThan(value df.MysqlDate) *MemberCQ {
	q.regBirthdate(df.CK_GT_C, value)
	return q
}

func (q *MemberCQ) SetBirthdate_LessThan(value df.MysqlDate) *MemberCQ {
	q.regBirthdate(df.CK_LT_C, value)
	return q
}

func (q *MemberCQ) SetBirthdate_GreaterEqual(value df.MysqlDate) *MemberCQ {
	q.regBirthdate(df.CK_GE_C, value)
	return q
}

func (q *MemberCQ) SetBirthdate_LessEqual(value df.MysqlDate) *MemberCQ {
	q.regBirthdate(df.CK_LE_C, value)
	return q
}

func (q *MemberCQ) SetBirthdate_IsNull() *MemberCQ {
	q.regBirthdate(df.CK_ISN_C, 0)
	return q
}
func (q *MemberCQ) SetBirthdate_IsNotNull() *MemberCQ {
	q.regBirthdate(df.CK_ISNN_C, 0)
	return q
}
func (q *MemberCQ) AddOrderBy_Birthdate_Asc() *MemberCQ {
	q.BaseConditionQuery.RegOBA("birthdate")
	return q
}
func (q *MemberCQ) AddOrderBy_Birthdate_Desc() *MemberCQ {
	q.BaseConditionQuery.RegOBD("birthdate")
	return q
}
func (q *MemberCQ) regBirthdate(key *df.ConditionKey, value interface{}) {
	if q.Birthdate == nil {
		q.Birthdate = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Birthdate, "birthdate")
}

func (q *MemberCQ) getCValueRegisterDatetime() *df.ConditionValue {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	return q.RegisterDatetime
}




func (q *MemberCQ) SetRegisterDatetime_Equal(value df.MysqlTimestamp) *MemberCQ {
	q.regRegisterDatetime(df.CK_EQ_C, value)
	return q
}


func (q *MemberCQ) SetRegisterDatetime_GreaterThan(value df.MysqlTimestamp) *MemberCQ {
	q.regRegisterDatetime(df.CK_GT_C, value)
	return q
}

func (q *MemberCQ) SetRegisterDatetime_LessThan(value df.MysqlTimestamp) *MemberCQ {
	q.regRegisterDatetime(df.CK_LT_C, value)
	return q
}

func (q *MemberCQ) SetRegisterDatetime_GreaterEqual(value df.MysqlTimestamp) *MemberCQ {
	q.regRegisterDatetime(df.CK_GE_C, value)
	return q
}

func (q *MemberCQ) SetRegisterDatetime_LessEqual(value df.MysqlTimestamp) *MemberCQ {
	q.regRegisterDatetime(df.CK_LE_C, value)
	return q
}

func (q *MemberCQ) AddOrderBy_RegisterDatetime_Asc() *MemberCQ {
	q.BaseConditionQuery.RegOBA("registerDatetime")
	return q
}
func (q *MemberCQ) AddOrderBy_RegisterDatetime_Desc() *MemberCQ {
	q.BaseConditionQuery.RegOBD("registerDatetime")
	return q
}
func (q *MemberCQ) regRegisterDatetime(key *df.ConditionKey, value interface{}) {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterDatetime, "registerDatetime")
}

func (q *MemberCQ) getCValueRegisterUser() *df.ConditionValue {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	return q.RegisterUser
}


func (q *MemberCQ) SetRegisterUser_Equal(value string) *MemberCQ {
	q.regRegisterUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *MemberCQ) SetRegisterUser_NotEqual(value string) *MemberCQ {
	q.regRegisterUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberCQ) SetRegisterUser_GreaterThan(value string) *MemberCQ {
	q.regRegisterUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberCQ) SetRegisterUser_LessThan(value string) *MemberCQ {
	q.regRegisterUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberCQ) SetRegisterUser_GreaterEqualThan(value string) *MemberCQ {
	q.regRegisterUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *MemberCQ) SetRegisterUser_LessEqualThan(value string) *MemberCQ {
	q.regRegisterUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberCQ) SetRegisterUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}

func (q *MemberCQ) SetRegisterUser_PrefixSearch(value string) error {
	return q.SetRegisterUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *MemberCQ) SetRegisterUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}



func (q *MemberCQ) AddOrderBy_RegisterUser_Asc() *MemberCQ {
	q.BaseConditionQuery.RegOBA("registerUser")
	return q
}
func (q *MemberCQ) AddOrderBy_RegisterUser_Desc() *MemberCQ {
	q.BaseConditionQuery.RegOBD("registerUser")
	return q
}
func (q *MemberCQ) regRegisterUser(key *df.ConditionKey, value interface{}) {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterUser, "registerUser")
}

func (q *MemberCQ) getCValueUpdateDatetime() *df.ConditionValue {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	return q.UpdateDatetime
}




func (q *MemberCQ) SetUpdateDatetime_Equal(value df.MysqlTimestamp) *MemberCQ {
	q.regUpdateDatetime(df.CK_EQ_C, value)
	return q
}


func (q *MemberCQ) SetUpdateDatetime_GreaterThan(value df.MysqlTimestamp) *MemberCQ {
	q.regUpdateDatetime(df.CK_GT_C, value)
	return q
}

func (q *MemberCQ) SetUpdateDatetime_LessThan(value df.MysqlTimestamp) *MemberCQ {
	q.regUpdateDatetime(df.CK_LT_C, value)
	return q
}

func (q *MemberCQ) SetUpdateDatetime_GreaterEqual(value df.MysqlTimestamp) *MemberCQ {
	q.regUpdateDatetime(df.CK_GE_C, value)
	return q
}

func (q *MemberCQ) SetUpdateDatetime_LessEqual(value df.MysqlTimestamp) *MemberCQ {
	q.regUpdateDatetime(df.CK_LE_C, value)
	return q
}

func (q *MemberCQ) AddOrderBy_UpdateDatetime_Asc() *MemberCQ {
	q.BaseConditionQuery.RegOBA("updateDatetime")
	return q
}
func (q *MemberCQ) AddOrderBy_UpdateDatetime_Desc() *MemberCQ {
	q.BaseConditionQuery.RegOBD("updateDatetime")
	return q
}
func (q *MemberCQ) regUpdateDatetime(key *df.ConditionKey, value interface{}) {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateDatetime, "updateDatetime")
}

func (q *MemberCQ) getCValueUpdateUser() *df.ConditionValue {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	return q.UpdateUser
}


func (q *MemberCQ) SetUpdateUser_Equal(value string) *MemberCQ {
	q.regUpdateUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *MemberCQ) SetUpdateUser_NotEqual(value string) *MemberCQ {
	q.regUpdateUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberCQ) SetUpdateUser_GreaterThan(value string) *MemberCQ {
	q.regUpdateUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberCQ) SetUpdateUser_LessThan(value string) *MemberCQ {
	q.regUpdateUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberCQ) SetUpdateUser_GreaterEqualThan(value string) *MemberCQ {
	q.regUpdateUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *MemberCQ) SetUpdateUser_LessEqualThan(value string) *MemberCQ {
	q.regUpdateUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberCQ) SetUpdateUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}

func (q *MemberCQ) SetUpdateUser_PrefixSearch(value string) error {
	return q.SetUpdateUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *MemberCQ) SetUpdateUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}



func (q *MemberCQ) AddOrderBy_UpdateUser_Asc() *MemberCQ {
	q.BaseConditionQuery.RegOBA("updateUser")
	return q
}
func (q *MemberCQ) AddOrderBy_UpdateUser_Desc() *MemberCQ {
	q.BaseConditionQuery.RegOBD("updateUser")
	return q
}
func (q *MemberCQ) regUpdateUser(key *df.ConditionKey, value interface{}) {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateUser, "updateUser")
}

func (q *MemberCQ) getCValueVersionNo() *df.ConditionValue {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	return q.VersionNo
}



func (q *MemberCQ) SetVersionNo_Equal(value int64) *MemberCQ {
	q.regVersionNo(df.CK_EQ_C, value)
	return q
}

func (q *MemberCQ) SetVersionNo_NotEqual(value int64) *MemberCQ {
	q.regVersionNo(df.CK_NE_C, value)
	return q
}

func (q *MemberCQ) SetVersionNo_GreaterThan(value int64) *MemberCQ {
	q.regVersionNo(df.CK_GT_C, value)
	return q
}

func (q *MemberCQ) SetVersionNo_LessThan(value int64) *MemberCQ {
	q.regVersionNo(df.CK_LT_C, value)
	return q
}

func (q *MemberCQ) SetVersionNo_GreaterEqual(value int64) *MemberCQ {
	q.regVersionNo(df.CK_GE_C, value)
	return q
}

func (q *MemberCQ) SetVersionNo_LessEqual(value int64) *MemberCQ {
	q.regVersionNo(df.CK_LE_C, value)
	return q
}
func (q *MemberCQ) SetVersionNo_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueVersionNo(),"versionNo",rangeOfOption)
}	


func (q *MemberCQ) AddOrderBy_VersionNo_Asc() *MemberCQ {
	q.BaseConditionQuery.RegOBA("versionNo")
	return q
}
func (q *MemberCQ) AddOrderBy_VersionNo_Desc() *MemberCQ {
	q.BaseConditionQuery.RegOBD("versionNo")
	return q
}
func (q *MemberCQ) regVersionNo(key *df.ConditionKey, value interface{}) {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.VersionNo, "versionNo")
}

