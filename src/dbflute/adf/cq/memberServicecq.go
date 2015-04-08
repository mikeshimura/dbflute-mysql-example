package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type MemberServiceCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	MemberServiceId *df.ConditionValue
	MemberId *df.ConditionValue
	ServicePointCount *df.ConditionValue
	ServiceRankCode *df.ConditionValue
	RegisterDatetime *df.ConditionValue
	RegisterUser *df.ConditionValue
	UpdateDatetime *df.ConditionValue
	UpdateUser *df.ConditionValue
	VersionNo *df.ConditionValue
}

func (q *MemberServiceCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *MemberServiceCQ) getCValueMemberServiceId() *df.ConditionValue {
	if q.MemberServiceId == nil {
		q.MemberServiceId = new(df.ConditionValue)
	}
	return q.MemberServiceId
}



func (q *MemberServiceCQ) SetMemberServiceId_Equal(value int64) *MemberServiceCQ {
	q.regMemberServiceId(df.CK_EQ_C, value)
	return q
}

func (q *MemberServiceCQ) SetMemberServiceId_NotEqual(value int64) *MemberServiceCQ {
	q.regMemberServiceId(df.CK_NE_C, value)
	return q
}

func (q *MemberServiceCQ) SetMemberServiceId_GreaterThan(value int64) *MemberServiceCQ {
	q.regMemberServiceId(df.CK_GT_C, value)
	return q
}

func (q *MemberServiceCQ) SetMemberServiceId_LessThan(value int64) *MemberServiceCQ {
	q.regMemberServiceId(df.CK_LT_C, value)
	return q
}

func (q *MemberServiceCQ) SetMemberServiceId_GreaterEqual(value int64) *MemberServiceCQ {
	q.regMemberServiceId(df.CK_GE_C, value)
	return q
}

func (q *MemberServiceCQ) SetMemberServiceId_LessEqual(value int64) *MemberServiceCQ {
	q.regMemberServiceId(df.CK_LE_C, value)
	return q
}
func (q *MemberServiceCQ) SetMemberServiceId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueMemberServiceId(),"memberServiceId",rangeOfOption)
}	


func (q *MemberServiceCQ) SetMemberServiceId_IsNull() *MemberServiceCQ {
	q.regMemberServiceId(df.CK_ISN_C, 0)
	return q
}
func (q *MemberServiceCQ) SetMemberServiceId_IsNotNull() *MemberServiceCQ {
	q.regMemberServiceId(df.CK_ISNN_C, 0)
	return q
}
func (q *MemberServiceCQ) AddOrderBy_MemberServiceId_Asc() *MemberServiceCQ {
	q.BaseConditionQuery.RegOBA("memberServiceId")
	return q
}
func (q *MemberServiceCQ) AddOrderBy_MemberServiceId_Desc() *MemberServiceCQ {
	q.BaseConditionQuery.RegOBD("memberServiceId")
	return q
}
func (q *MemberServiceCQ) regMemberServiceId(key *df.ConditionKey, value interface{}) {
	if q.MemberServiceId == nil {
		q.MemberServiceId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.MemberServiceId, "memberServiceId")
}

func (q *MemberServiceCQ) getCValueMemberId() *df.ConditionValue {
	if q.MemberId == nil {
		q.MemberId = new(df.ConditionValue)
	}
	return q.MemberId
}



func (q *MemberServiceCQ) SetMemberId_Equal(value int64) *MemberServiceCQ {
	q.regMemberId(df.CK_EQ_C, value)
	return q
}

func (q *MemberServiceCQ) SetMemberId_NotEqual(value int64) *MemberServiceCQ {
	q.regMemberId(df.CK_NE_C, value)
	return q
}

func (q *MemberServiceCQ) SetMemberId_GreaterThan(value int64) *MemberServiceCQ {
	q.regMemberId(df.CK_GT_C, value)
	return q
}

func (q *MemberServiceCQ) SetMemberId_LessThan(value int64) *MemberServiceCQ {
	q.regMemberId(df.CK_LT_C, value)
	return q
}

func (q *MemberServiceCQ) SetMemberId_GreaterEqual(value int64) *MemberServiceCQ {
	q.regMemberId(df.CK_GE_C, value)
	return q
}

func (q *MemberServiceCQ) SetMemberId_LessEqual(value int64) *MemberServiceCQ {
	q.regMemberId(df.CK_LE_C, value)
	return q
}
func (q *MemberServiceCQ) SetMemberId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueMemberId(),"memberId",rangeOfOption)
}	


func (q *MemberServiceCQ) AddOrderBy_MemberId_Asc() *MemberServiceCQ {
	q.BaseConditionQuery.RegOBA("memberId")
	return q
}
func (q *MemberServiceCQ) AddOrderBy_MemberId_Desc() *MemberServiceCQ {
	q.BaseConditionQuery.RegOBD("memberId")
	return q
}
func (q *MemberServiceCQ) regMemberId(key *df.ConditionKey, value interface{}) {
	if q.MemberId == nil {
		q.MemberId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.MemberId, "memberId")
}

func (q *MemberServiceCQ) getCValueServicePointCount() *df.ConditionValue {
	if q.ServicePointCount == nil {
		q.ServicePointCount = new(df.ConditionValue)
	}
	return q.ServicePointCount
}



func (q *MemberServiceCQ) SetServicePointCount_Equal(value int64) *MemberServiceCQ {
	q.regServicePointCount(df.CK_EQ_C, value)
	return q
}

func (q *MemberServiceCQ) SetServicePointCount_NotEqual(value int64) *MemberServiceCQ {
	q.regServicePointCount(df.CK_NE_C, value)
	return q
}

func (q *MemberServiceCQ) SetServicePointCount_GreaterThan(value int64) *MemberServiceCQ {
	q.regServicePointCount(df.CK_GT_C, value)
	return q
}

func (q *MemberServiceCQ) SetServicePointCount_LessThan(value int64) *MemberServiceCQ {
	q.regServicePointCount(df.CK_LT_C, value)
	return q
}

func (q *MemberServiceCQ) SetServicePointCount_GreaterEqual(value int64) *MemberServiceCQ {
	q.regServicePointCount(df.CK_GE_C, value)
	return q
}

func (q *MemberServiceCQ) SetServicePointCount_LessEqual(value int64) *MemberServiceCQ {
	q.regServicePointCount(df.CK_LE_C, value)
	return q
}
func (q *MemberServiceCQ) SetServicePointCount_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueServicePointCount(),"servicePointCount",rangeOfOption)
}	


func (q *MemberServiceCQ) AddOrderBy_ServicePointCount_Asc() *MemberServiceCQ {
	q.BaseConditionQuery.RegOBA("servicePointCount")
	return q
}
func (q *MemberServiceCQ) AddOrderBy_ServicePointCount_Desc() *MemberServiceCQ {
	q.BaseConditionQuery.RegOBD("servicePointCount")
	return q
}
func (q *MemberServiceCQ) regServicePointCount(key *df.ConditionKey, value interface{}) {
	if q.ServicePointCount == nil {
		q.ServicePointCount = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ServicePointCount, "servicePointCount")
}

func (q *MemberServiceCQ) getCValueServiceRankCode() *df.ConditionValue {
	if q.ServiceRankCode == nil {
		q.ServiceRankCode = new(df.ConditionValue)
	}
	return q.ServiceRankCode
}


func (q *MemberServiceCQ) SetServiceRankCode_Equal(value string) *MemberServiceCQ {
	q.regServiceRankCode(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *MemberServiceCQ) SetServiceRankCode_NotEqual(value string) *MemberServiceCQ {
	q.regServiceRankCode(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberServiceCQ) SetServiceRankCode_GreaterThan(value string) *MemberServiceCQ {
	q.regServiceRankCode(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberServiceCQ) SetServiceRankCode_LessThan(value string) *MemberServiceCQ {
	q.regServiceRankCode(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberServiceCQ) SetServiceRankCode_GreaterEqualThan(value string) *MemberServiceCQ {
	q.regServiceRankCode(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *MemberServiceCQ) SetServiceRankCode_LessEqualThan(value string) *MemberServiceCQ {
	q.regServiceRankCode(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberServiceCQ) SetServiceRankCode_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueServiceRankCode(), "serviceRankCode", option)
}

func (q *MemberServiceCQ) SetServiceRankCode_PrefixSearch(value string) error {
	return q.SetServiceRankCode_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *MemberServiceCQ) SetServiceRankCode_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueServiceRankCode(), "serviceRankCode", option)
}



func (q *MemberServiceCQ) AddOrderBy_ServiceRankCode_Asc() *MemberServiceCQ {
	q.BaseConditionQuery.RegOBA("serviceRankCode")
	return q
}
func (q *MemberServiceCQ) AddOrderBy_ServiceRankCode_Desc() *MemberServiceCQ {
	q.BaseConditionQuery.RegOBD("serviceRankCode")
	return q
}
func (q *MemberServiceCQ) regServiceRankCode(key *df.ConditionKey, value interface{}) {
	if q.ServiceRankCode == nil {
		q.ServiceRankCode = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ServiceRankCode, "serviceRankCode")
}

func (q *MemberServiceCQ) getCValueRegisterDatetime() *df.ConditionValue {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	return q.RegisterDatetime
}




func (q *MemberServiceCQ) SetRegisterDatetime_Equal(value df.MysqlTimestamp) *MemberServiceCQ {
	q.regRegisterDatetime(df.CK_EQ_C, value)
	return q
}


func (q *MemberServiceCQ) SetRegisterDatetime_GreaterThan(value df.MysqlTimestamp) *MemberServiceCQ {
	q.regRegisterDatetime(df.CK_GT_C, value)
	return q
}

func (q *MemberServiceCQ) SetRegisterDatetime_LessThan(value df.MysqlTimestamp) *MemberServiceCQ {
	q.regRegisterDatetime(df.CK_LT_C, value)
	return q
}

func (q *MemberServiceCQ) SetRegisterDatetime_GreaterEqual(value df.MysqlTimestamp) *MemberServiceCQ {
	q.regRegisterDatetime(df.CK_GE_C, value)
	return q
}

func (q *MemberServiceCQ) SetRegisterDatetime_LessEqual(value df.MysqlTimestamp) *MemberServiceCQ {
	q.regRegisterDatetime(df.CK_LE_C, value)
	return q
}

func (q *MemberServiceCQ) AddOrderBy_RegisterDatetime_Asc() *MemberServiceCQ {
	q.BaseConditionQuery.RegOBA("registerDatetime")
	return q
}
func (q *MemberServiceCQ) AddOrderBy_RegisterDatetime_Desc() *MemberServiceCQ {
	q.BaseConditionQuery.RegOBD("registerDatetime")
	return q
}
func (q *MemberServiceCQ) regRegisterDatetime(key *df.ConditionKey, value interface{}) {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterDatetime, "registerDatetime")
}

func (q *MemberServiceCQ) getCValueRegisterUser() *df.ConditionValue {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	return q.RegisterUser
}


func (q *MemberServiceCQ) SetRegisterUser_Equal(value string) *MemberServiceCQ {
	q.regRegisterUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *MemberServiceCQ) SetRegisterUser_NotEqual(value string) *MemberServiceCQ {
	q.regRegisterUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberServiceCQ) SetRegisterUser_GreaterThan(value string) *MemberServiceCQ {
	q.regRegisterUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberServiceCQ) SetRegisterUser_LessThan(value string) *MemberServiceCQ {
	q.regRegisterUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberServiceCQ) SetRegisterUser_GreaterEqualThan(value string) *MemberServiceCQ {
	q.regRegisterUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *MemberServiceCQ) SetRegisterUser_LessEqualThan(value string) *MemberServiceCQ {
	q.regRegisterUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberServiceCQ) SetRegisterUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}

func (q *MemberServiceCQ) SetRegisterUser_PrefixSearch(value string) error {
	return q.SetRegisterUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *MemberServiceCQ) SetRegisterUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}



func (q *MemberServiceCQ) AddOrderBy_RegisterUser_Asc() *MemberServiceCQ {
	q.BaseConditionQuery.RegOBA("registerUser")
	return q
}
func (q *MemberServiceCQ) AddOrderBy_RegisterUser_Desc() *MemberServiceCQ {
	q.BaseConditionQuery.RegOBD("registerUser")
	return q
}
func (q *MemberServiceCQ) regRegisterUser(key *df.ConditionKey, value interface{}) {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterUser, "registerUser")
}

func (q *MemberServiceCQ) getCValueUpdateDatetime() *df.ConditionValue {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	return q.UpdateDatetime
}




func (q *MemberServiceCQ) SetUpdateDatetime_Equal(value df.MysqlTimestamp) *MemberServiceCQ {
	q.regUpdateDatetime(df.CK_EQ_C, value)
	return q
}


func (q *MemberServiceCQ) SetUpdateDatetime_GreaterThan(value df.MysqlTimestamp) *MemberServiceCQ {
	q.regUpdateDatetime(df.CK_GT_C, value)
	return q
}

func (q *MemberServiceCQ) SetUpdateDatetime_LessThan(value df.MysqlTimestamp) *MemberServiceCQ {
	q.regUpdateDatetime(df.CK_LT_C, value)
	return q
}

func (q *MemberServiceCQ) SetUpdateDatetime_GreaterEqual(value df.MysqlTimestamp) *MemberServiceCQ {
	q.regUpdateDatetime(df.CK_GE_C, value)
	return q
}

func (q *MemberServiceCQ) SetUpdateDatetime_LessEqual(value df.MysqlTimestamp) *MemberServiceCQ {
	q.regUpdateDatetime(df.CK_LE_C, value)
	return q
}

func (q *MemberServiceCQ) AddOrderBy_UpdateDatetime_Asc() *MemberServiceCQ {
	q.BaseConditionQuery.RegOBA("updateDatetime")
	return q
}
func (q *MemberServiceCQ) AddOrderBy_UpdateDatetime_Desc() *MemberServiceCQ {
	q.BaseConditionQuery.RegOBD("updateDatetime")
	return q
}
func (q *MemberServiceCQ) regUpdateDatetime(key *df.ConditionKey, value interface{}) {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateDatetime, "updateDatetime")
}

func (q *MemberServiceCQ) getCValueUpdateUser() *df.ConditionValue {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	return q.UpdateUser
}


func (q *MemberServiceCQ) SetUpdateUser_Equal(value string) *MemberServiceCQ {
	q.regUpdateUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *MemberServiceCQ) SetUpdateUser_NotEqual(value string) *MemberServiceCQ {
	q.regUpdateUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberServiceCQ) SetUpdateUser_GreaterThan(value string) *MemberServiceCQ {
	q.regUpdateUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberServiceCQ) SetUpdateUser_LessThan(value string) *MemberServiceCQ {
	q.regUpdateUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberServiceCQ) SetUpdateUser_GreaterEqualThan(value string) *MemberServiceCQ {
	q.regUpdateUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *MemberServiceCQ) SetUpdateUser_LessEqualThan(value string) *MemberServiceCQ {
	q.regUpdateUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberServiceCQ) SetUpdateUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}

func (q *MemberServiceCQ) SetUpdateUser_PrefixSearch(value string) error {
	return q.SetUpdateUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *MemberServiceCQ) SetUpdateUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}



func (q *MemberServiceCQ) AddOrderBy_UpdateUser_Asc() *MemberServiceCQ {
	q.BaseConditionQuery.RegOBA("updateUser")
	return q
}
func (q *MemberServiceCQ) AddOrderBy_UpdateUser_Desc() *MemberServiceCQ {
	q.BaseConditionQuery.RegOBD("updateUser")
	return q
}
func (q *MemberServiceCQ) regUpdateUser(key *df.ConditionKey, value interface{}) {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateUser, "updateUser")
}

func (q *MemberServiceCQ) getCValueVersionNo() *df.ConditionValue {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	return q.VersionNo
}



func (q *MemberServiceCQ) SetVersionNo_Equal(value int64) *MemberServiceCQ {
	q.regVersionNo(df.CK_EQ_C, value)
	return q
}

func (q *MemberServiceCQ) SetVersionNo_NotEqual(value int64) *MemberServiceCQ {
	q.regVersionNo(df.CK_NE_C, value)
	return q
}

func (q *MemberServiceCQ) SetVersionNo_GreaterThan(value int64) *MemberServiceCQ {
	q.regVersionNo(df.CK_GT_C, value)
	return q
}

func (q *MemberServiceCQ) SetVersionNo_LessThan(value int64) *MemberServiceCQ {
	q.regVersionNo(df.CK_LT_C, value)
	return q
}

func (q *MemberServiceCQ) SetVersionNo_GreaterEqual(value int64) *MemberServiceCQ {
	q.regVersionNo(df.CK_GE_C, value)
	return q
}

func (q *MemberServiceCQ) SetVersionNo_LessEqual(value int64) *MemberServiceCQ {
	q.regVersionNo(df.CK_LE_C, value)
	return q
}
func (q *MemberServiceCQ) SetVersionNo_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueVersionNo(),"versionNo",rangeOfOption)
}	


func (q *MemberServiceCQ) AddOrderBy_VersionNo_Asc() *MemberServiceCQ {
	q.BaseConditionQuery.RegOBA("versionNo")
	return q
}
func (q *MemberServiceCQ) AddOrderBy_VersionNo_Desc() *MemberServiceCQ {
	q.BaseConditionQuery.RegOBD("versionNo")
	return q
}
func (q *MemberServiceCQ) regVersionNo(key *df.ConditionKey, value interface{}) {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.VersionNo, "versionNo")
}

