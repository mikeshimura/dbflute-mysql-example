package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type MemberSecurityCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	MemberId *df.ConditionValue
	LoginPassword *df.ConditionValue
	ReminderQuestion *df.ConditionValue
	ReminderAnswer *df.ConditionValue
	ReminderUseCount *df.ConditionValue
	RegisterDatetime *df.ConditionValue
	RegisterUser *df.ConditionValue
	UpdateDatetime *df.ConditionValue
	UpdateUser *df.ConditionValue
	VersionNo *df.ConditionValue
    conditionQueryMember *MemberCQ
}

func (q *MemberSecurityCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *MemberSecurityCQ) getCValueMemberId() *df.ConditionValue {
	if q.MemberId == nil {
		q.MemberId = new(df.ConditionValue)
	}
	return q.MemberId
}



func (q *MemberSecurityCQ) SetMemberId_Equal(value int64) *MemberSecurityCQ {
	q.regMemberId(df.CK_EQ_C, value)
	return q
}
func (q *MemberSecurityCQ) SetMemberId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueMemberId(), "memberId")
}
func (q *MemberSecurityCQ) SetMemberId_NotEqual(value int64) *MemberSecurityCQ {
	q.regMemberId(df.CK_NE_C, value)
	return q
}

func (q *MemberSecurityCQ) SetMemberId_GreaterThan(value int64) *MemberSecurityCQ {
	q.regMemberId(df.CK_GT_C, value)
	return q
}

func (q *MemberSecurityCQ) SetMemberId_LessThan(value int64) *MemberSecurityCQ {
	q.regMemberId(df.CK_LT_C, value)
	return q
}

func (q *MemberSecurityCQ) SetMemberId_GreaterEqual(value int64) *MemberSecurityCQ {
	q.regMemberId(df.CK_GE_C, value)
	return q
}

func (q *MemberSecurityCQ) SetMemberId_LessEqual(value int64) *MemberSecurityCQ {
	q.regMemberId(df.CK_LE_C, value)
	return q
}
func (q *MemberSecurityCQ) SetMemberId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueMemberId(),"memberId",rangeOfOption)
}	


func (q *MemberSecurityCQ) SetMemberId_IsNull() *MemberSecurityCQ {
	q.regMemberId(df.CK_ISN_C, 0)
	return q
}
func (q *MemberSecurityCQ) SetMemberId_IsNotNull() *MemberSecurityCQ {
	q.regMemberId(df.CK_ISNN_C, 0)
	return q
}
func (q *MemberSecurityCQ) AddOrderBy_MemberId_Asc() *MemberSecurityCQ {
	q.BaseConditionQuery.RegOBA("memberId")
	return q
}
func (q *MemberSecurityCQ) AddOrderBy_MemberId_Desc() *MemberSecurityCQ {
	q.BaseConditionQuery.RegOBD("memberId")
	return q
}
func (q *MemberSecurityCQ) regMemberId(key *df.ConditionKey, value interface{}) {
	if q.MemberId == nil {
		q.MemberId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.MemberId, "memberId")
}

func (q *MemberSecurityCQ) getCValueLoginPassword() *df.ConditionValue {
	if q.LoginPassword == nil {
		q.LoginPassword = new(df.ConditionValue)
	}
	return q.LoginPassword
}


func (q *MemberSecurityCQ) SetLoginPassword_Equal(value string) *MemberSecurityCQ {
	q.regLoginPassword(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *MemberSecurityCQ) SetLoginPassword_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueLoginPassword(), "loginPassword")
}
func (q *MemberSecurityCQ) SetLoginPassword_NotEqual(value string) *MemberSecurityCQ {
	q.regLoginPassword(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberSecurityCQ) SetLoginPassword_GreaterThan(value string) *MemberSecurityCQ {
	q.regLoginPassword(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberSecurityCQ) SetLoginPassword_LessThan(value string) *MemberSecurityCQ {
	q.regLoginPassword(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberSecurityCQ) SetLoginPassword_GreaterEqualThan(value string) *MemberSecurityCQ {
	q.regLoginPassword(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *MemberSecurityCQ) SetLoginPassword_LessEqualThan(value string) *MemberSecurityCQ {
	q.regLoginPassword(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberSecurityCQ) SetLoginPassword_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueLoginPassword(), "loginPassword", option)
}

func (q *MemberSecurityCQ) SetLoginPassword_PrefixSearch(value string) error {
	return q.SetLoginPassword_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *MemberSecurityCQ) SetLoginPassword_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueLoginPassword(), "loginPassword", option)
}



func (q *MemberSecurityCQ) AddOrderBy_LoginPassword_Asc() *MemberSecurityCQ {
	q.BaseConditionQuery.RegOBA("loginPassword")
	return q
}
func (q *MemberSecurityCQ) AddOrderBy_LoginPassword_Desc() *MemberSecurityCQ {
	q.BaseConditionQuery.RegOBD("loginPassword")
	return q
}
func (q *MemberSecurityCQ) regLoginPassword(key *df.ConditionKey, value interface{}) {
	if q.LoginPassword == nil {
		q.LoginPassword = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.LoginPassword, "loginPassword")
}

func (q *MemberSecurityCQ) getCValueReminderQuestion() *df.ConditionValue {
	if q.ReminderQuestion == nil {
		q.ReminderQuestion = new(df.ConditionValue)
	}
	return q.ReminderQuestion
}


func (q *MemberSecurityCQ) SetReminderQuestion_Equal(value string) *MemberSecurityCQ {
	q.regReminderQuestion(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *MemberSecurityCQ) SetReminderQuestion_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueReminderQuestion(), "reminderQuestion")
}
func (q *MemberSecurityCQ) SetReminderQuestion_NotEqual(value string) *MemberSecurityCQ {
	q.regReminderQuestion(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberSecurityCQ) SetReminderQuestion_GreaterThan(value string) *MemberSecurityCQ {
	q.regReminderQuestion(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberSecurityCQ) SetReminderQuestion_LessThan(value string) *MemberSecurityCQ {
	q.regReminderQuestion(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberSecurityCQ) SetReminderQuestion_GreaterEqualThan(value string) *MemberSecurityCQ {
	q.regReminderQuestion(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *MemberSecurityCQ) SetReminderQuestion_LessEqualThan(value string) *MemberSecurityCQ {
	q.regReminderQuestion(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberSecurityCQ) SetReminderQuestion_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueReminderQuestion(), "reminderQuestion", option)
}

func (q *MemberSecurityCQ) SetReminderQuestion_PrefixSearch(value string) error {
	return q.SetReminderQuestion_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *MemberSecurityCQ) SetReminderQuestion_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueReminderQuestion(), "reminderQuestion", option)
}



func (q *MemberSecurityCQ) AddOrderBy_ReminderQuestion_Asc() *MemberSecurityCQ {
	q.BaseConditionQuery.RegOBA("reminderQuestion")
	return q
}
func (q *MemberSecurityCQ) AddOrderBy_ReminderQuestion_Desc() *MemberSecurityCQ {
	q.BaseConditionQuery.RegOBD("reminderQuestion")
	return q
}
func (q *MemberSecurityCQ) regReminderQuestion(key *df.ConditionKey, value interface{}) {
	if q.ReminderQuestion == nil {
		q.ReminderQuestion = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ReminderQuestion, "reminderQuestion")
}

func (q *MemberSecurityCQ) getCValueReminderAnswer() *df.ConditionValue {
	if q.ReminderAnswer == nil {
		q.ReminderAnswer = new(df.ConditionValue)
	}
	return q.ReminderAnswer
}


func (q *MemberSecurityCQ) SetReminderAnswer_Equal(value string) *MemberSecurityCQ {
	q.regReminderAnswer(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *MemberSecurityCQ) SetReminderAnswer_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueReminderAnswer(), "reminderAnswer")
}
func (q *MemberSecurityCQ) SetReminderAnswer_NotEqual(value string) *MemberSecurityCQ {
	q.regReminderAnswer(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberSecurityCQ) SetReminderAnswer_GreaterThan(value string) *MemberSecurityCQ {
	q.regReminderAnswer(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberSecurityCQ) SetReminderAnswer_LessThan(value string) *MemberSecurityCQ {
	q.regReminderAnswer(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberSecurityCQ) SetReminderAnswer_GreaterEqualThan(value string) *MemberSecurityCQ {
	q.regReminderAnswer(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *MemberSecurityCQ) SetReminderAnswer_LessEqualThan(value string) *MemberSecurityCQ {
	q.regReminderAnswer(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberSecurityCQ) SetReminderAnswer_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueReminderAnswer(), "reminderAnswer", option)
}

func (q *MemberSecurityCQ) SetReminderAnswer_PrefixSearch(value string) error {
	return q.SetReminderAnswer_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *MemberSecurityCQ) SetReminderAnswer_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueReminderAnswer(), "reminderAnswer", option)
}



func (q *MemberSecurityCQ) AddOrderBy_ReminderAnswer_Asc() *MemberSecurityCQ {
	q.BaseConditionQuery.RegOBA("reminderAnswer")
	return q
}
func (q *MemberSecurityCQ) AddOrderBy_ReminderAnswer_Desc() *MemberSecurityCQ {
	q.BaseConditionQuery.RegOBD("reminderAnswer")
	return q
}
func (q *MemberSecurityCQ) regReminderAnswer(key *df.ConditionKey, value interface{}) {
	if q.ReminderAnswer == nil {
		q.ReminderAnswer = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ReminderAnswer, "reminderAnswer")
}

func (q *MemberSecurityCQ) getCValueReminderUseCount() *df.ConditionValue {
	if q.ReminderUseCount == nil {
		q.ReminderUseCount = new(df.ConditionValue)
	}
	return q.ReminderUseCount
}



func (q *MemberSecurityCQ) SetReminderUseCount_Equal(value int64) *MemberSecurityCQ {
	q.regReminderUseCount(df.CK_EQ_C, value)
	return q
}
func (q *MemberSecurityCQ) SetReminderUseCount_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueReminderUseCount(), "reminderUseCount")
}
func (q *MemberSecurityCQ) SetReminderUseCount_NotEqual(value int64) *MemberSecurityCQ {
	q.regReminderUseCount(df.CK_NE_C, value)
	return q
}

func (q *MemberSecurityCQ) SetReminderUseCount_GreaterThan(value int64) *MemberSecurityCQ {
	q.regReminderUseCount(df.CK_GT_C, value)
	return q
}

func (q *MemberSecurityCQ) SetReminderUseCount_LessThan(value int64) *MemberSecurityCQ {
	q.regReminderUseCount(df.CK_LT_C, value)
	return q
}

func (q *MemberSecurityCQ) SetReminderUseCount_GreaterEqual(value int64) *MemberSecurityCQ {
	q.regReminderUseCount(df.CK_GE_C, value)
	return q
}

func (q *MemberSecurityCQ) SetReminderUseCount_LessEqual(value int64) *MemberSecurityCQ {
	q.regReminderUseCount(df.CK_LE_C, value)
	return q
}
func (q *MemberSecurityCQ) SetReminderUseCount_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueReminderUseCount(),"reminderUseCount",rangeOfOption)
}	


func (q *MemberSecurityCQ) AddOrderBy_ReminderUseCount_Asc() *MemberSecurityCQ {
	q.BaseConditionQuery.RegOBA("reminderUseCount")
	return q
}
func (q *MemberSecurityCQ) AddOrderBy_ReminderUseCount_Desc() *MemberSecurityCQ {
	q.BaseConditionQuery.RegOBD("reminderUseCount")
	return q
}
func (q *MemberSecurityCQ) regReminderUseCount(key *df.ConditionKey, value interface{}) {
	if q.ReminderUseCount == nil {
		q.ReminderUseCount = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ReminderUseCount, "reminderUseCount")
}

func (q *MemberSecurityCQ) getCValueRegisterDatetime() *df.ConditionValue {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	return q.RegisterDatetime
}




func (q *MemberSecurityCQ) SetRegisterDatetime_Equal(value df.MysqlTimestamp) *MemberSecurityCQ {
	q.regRegisterDatetime(df.CK_EQ_C, value)
	return q
}


func (q *MemberSecurityCQ) SetRegisterDatetime_GreaterThan(value df.MysqlTimestamp) *MemberSecurityCQ {
	q.regRegisterDatetime(df.CK_GT_C, value)
	return q
}

func (q *MemberSecurityCQ) SetRegisterDatetime_LessThan(value df.MysqlTimestamp) *MemberSecurityCQ {
	q.regRegisterDatetime(df.CK_LT_C, value)
	return q
}

func (q *MemberSecurityCQ) SetRegisterDatetime_GreaterEqual(value df.MysqlTimestamp) *MemberSecurityCQ {
	q.regRegisterDatetime(df.CK_GE_C, value)
	return q
}

func (q *MemberSecurityCQ) SetRegisterDatetime_LessEqual(value df.MysqlTimestamp) *MemberSecurityCQ {
	q.regRegisterDatetime(df.CK_LE_C, value)
	return q
}

func (q *MemberSecurityCQ) AddOrderBy_RegisterDatetime_Asc() *MemberSecurityCQ {
	q.BaseConditionQuery.RegOBA("registerDatetime")
	return q
}
func (q *MemberSecurityCQ) AddOrderBy_RegisterDatetime_Desc() *MemberSecurityCQ {
	q.BaseConditionQuery.RegOBD("registerDatetime")
	return q
}
func (q *MemberSecurityCQ) regRegisterDatetime(key *df.ConditionKey, value interface{}) {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterDatetime, "registerDatetime")
}

func (q *MemberSecurityCQ) getCValueRegisterUser() *df.ConditionValue {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	return q.RegisterUser
}


func (q *MemberSecurityCQ) SetRegisterUser_Equal(value string) *MemberSecurityCQ {
	q.regRegisterUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *MemberSecurityCQ) SetRegisterUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegisterUser(), "registerUser")
}
func (q *MemberSecurityCQ) SetRegisterUser_NotEqual(value string) *MemberSecurityCQ {
	q.regRegisterUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberSecurityCQ) SetRegisterUser_GreaterThan(value string) *MemberSecurityCQ {
	q.regRegisterUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberSecurityCQ) SetRegisterUser_LessThan(value string) *MemberSecurityCQ {
	q.regRegisterUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberSecurityCQ) SetRegisterUser_GreaterEqualThan(value string) *MemberSecurityCQ {
	q.regRegisterUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *MemberSecurityCQ) SetRegisterUser_LessEqualThan(value string) *MemberSecurityCQ {
	q.regRegisterUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberSecurityCQ) SetRegisterUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}

func (q *MemberSecurityCQ) SetRegisterUser_PrefixSearch(value string) error {
	return q.SetRegisterUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *MemberSecurityCQ) SetRegisterUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}



func (q *MemberSecurityCQ) AddOrderBy_RegisterUser_Asc() *MemberSecurityCQ {
	q.BaseConditionQuery.RegOBA("registerUser")
	return q
}
func (q *MemberSecurityCQ) AddOrderBy_RegisterUser_Desc() *MemberSecurityCQ {
	q.BaseConditionQuery.RegOBD("registerUser")
	return q
}
func (q *MemberSecurityCQ) regRegisterUser(key *df.ConditionKey, value interface{}) {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterUser, "registerUser")
}

func (q *MemberSecurityCQ) getCValueUpdateDatetime() *df.ConditionValue {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	return q.UpdateDatetime
}




func (q *MemberSecurityCQ) SetUpdateDatetime_Equal(value df.MysqlTimestamp) *MemberSecurityCQ {
	q.regUpdateDatetime(df.CK_EQ_C, value)
	return q
}


func (q *MemberSecurityCQ) SetUpdateDatetime_GreaterThan(value df.MysqlTimestamp) *MemberSecurityCQ {
	q.regUpdateDatetime(df.CK_GT_C, value)
	return q
}

func (q *MemberSecurityCQ) SetUpdateDatetime_LessThan(value df.MysqlTimestamp) *MemberSecurityCQ {
	q.regUpdateDatetime(df.CK_LT_C, value)
	return q
}

func (q *MemberSecurityCQ) SetUpdateDatetime_GreaterEqual(value df.MysqlTimestamp) *MemberSecurityCQ {
	q.regUpdateDatetime(df.CK_GE_C, value)
	return q
}

func (q *MemberSecurityCQ) SetUpdateDatetime_LessEqual(value df.MysqlTimestamp) *MemberSecurityCQ {
	q.regUpdateDatetime(df.CK_LE_C, value)
	return q
}

func (q *MemberSecurityCQ) AddOrderBy_UpdateDatetime_Asc() *MemberSecurityCQ {
	q.BaseConditionQuery.RegOBA("updateDatetime")
	return q
}
func (q *MemberSecurityCQ) AddOrderBy_UpdateDatetime_Desc() *MemberSecurityCQ {
	q.BaseConditionQuery.RegOBD("updateDatetime")
	return q
}
func (q *MemberSecurityCQ) regUpdateDatetime(key *df.ConditionKey, value interface{}) {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateDatetime, "updateDatetime")
}

func (q *MemberSecurityCQ) getCValueUpdateUser() *df.ConditionValue {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	return q.UpdateUser
}


func (q *MemberSecurityCQ) SetUpdateUser_Equal(value string) *MemberSecurityCQ {
	q.regUpdateUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *MemberSecurityCQ) SetUpdateUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUpdateUser(), "updateUser")
}
func (q *MemberSecurityCQ) SetUpdateUser_NotEqual(value string) *MemberSecurityCQ {
	q.regUpdateUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberSecurityCQ) SetUpdateUser_GreaterThan(value string) *MemberSecurityCQ {
	q.regUpdateUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberSecurityCQ) SetUpdateUser_LessThan(value string) *MemberSecurityCQ {
	q.regUpdateUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberSecurityCQ) SetUpdateUser_GreaterEqualThan(value string) *MemberSecurityCQ {
	q.regUpdateUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *MemberSecurityCQ) SetUpdateUser_LessEqualThan(value string) *MemberSecurityCQ {
	q.regUpdateUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberSecurityCQ) SetUpdateUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}

func (q *MemberSecurityCQ) SetUpdateUser_PrefixSearch(value string) error {
	return q.SetUpdateUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *MemberSecurityCQ) SetUpdateUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}



func (q *MemberSecurityCQ) AddOrderBy_UpdateUser_Asc() *MemberSecurityCQ {
	q.BaseConditionQuery.RegOBA("updateUser")
	return q
}
func (q *MemberSecurityCQ) AddOrderBy_UpdateUser_Desc() *MemberSecurityCQ {
	q.BaseConditionQuery.RegOBD("updateUser")
	return q
}
func (q *MemberSecurityCQ) regUpdateUser(key *df.ConditionKey, value interface{}) {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateUser, "updateUser")
}

func (q *MemberSecurityCQ) getCValueVersionNo() *df.ConditionValue {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	return q.VersionNo
}



func (q *MemberSecurityCQ) SetVersionNo_Equal(value int64) *MemberSecurityCQ {
	q.regVersionNo(df.CK_EQ_C, value)
	return q
}
func (q *MemberSecurityCQ) SetVersionNo_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueVersionNo(), "versionNo")
}
func (q *MemberSecurityCQ) SetVersionNo_NotEqual(value int64) *MemberSecurityCQ {
	q.regVersionNo(df.CK_NE_C, value)
	return q
}

func (q *MemberSecurityCQ) SetVersionNo_GreaterThan(value int64) *MemberSecurityCQ {
	q.regVersionNo(df.CK_GT_C, value)
	return q
}

func (q *MemberSecurityCQ) SetVersionNo_LessThan(value int64) *MemberSecurityCQ {
	q.regVersionNo(df.CK_LT_C, value)
	return q
}

func (q *MemberSecurityCQ) SetVersionNo_GreaterEqual(value int64) *MemberSecurityCQ {
	q.regVersionNo(df.CK_GE_C, value)
	return q
}

func (q *MemberSecurityCQ) SetVersionNo_LessEqual(value int64) *MemberSecurityCQ {
	q.regVersionNo(df.CK_LE_C, value)
	return q
}
func (q *MemberSecurityCQ) SetVersionNo_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueVersionNo(),"versionNo",rangeOfOption)
}	


func (q *MemberSecurityCQ) AddOrderBy_VersionNo_Asc() *MemberSecurityCQ {
	q.BaseConditionQuery.RegOBA("versionNo")
	return q
}
func (q *MemberSecurityCQ) AddOrderBy_VersionNo_Desc() *MemberSecurityCQ {
	q.BaseConditionQuery.RegOBD("versionNo")
	return q
}
func (q *MemberSecurityCQ) regVersionNo(key *df.ConditionKey, value interface{}) {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.VersionNo, "versionNo")
}


func (q *MemberSecurityCQ) QueryMember() *MemberCQ {
	if q.conditionQueryMember == nil {
		q.conditionQueryMember = q.xcreateQueryMember()
		q.xsetupOuterJoinMember()
	}
	return q.conditionQueryMember
}

func (q *MemberSecurityCQ) xcreateQueryMember() *MemberCQ {
	nrp := q.BaseConditionQuery.ResolveNextRelationPath("MemberSecurity", "Member")
	jan := q.BaseConditionQuery.ResolveJoinAliasName(nrp)
	var basecq df.ConditionQuery = q
	cq := CreateMemberCQ(&basecq, q.BaseConditionQuery.SqlClause, jan, q.BaseConditionQuery.NestLevel+1)
	cq.BaseConditionQuery.BaseCB = q.BaseConditionQuery.BaseCB
	cq.BaseConditionQuery.ForeignPropertyName = "Member"
	cq.BaseConditionQuery.RelationPath = nrp
	return cq
}
func (q *MemberSecurityCQ) xsetupOuterJoinMember() {
	    cq := q.QueryMember()
        joinOnMap := make(map[string]string)
        joinOnMap["memberId"]="memberId"
        q.BaseConditionQuery.RegisterOuterJoin(
        	cq.BaseConditionQuery.ConditionQuery, joinOnMap, "Member");
}	
	
func CreateMemberSecurityCQ(referrerQuery *df.ConditionQuery, sqlClause *df.SqlClause, aliasName string, nestlevel int8) *MemberSecurityCQ {
	cq := new(MemberSecurityCQ)
	cq.BaseConditionQuery = new(df.BaseConditionQuery)
	cq.BaseConditionQuery.TableDbName = "MemberSecurity"
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