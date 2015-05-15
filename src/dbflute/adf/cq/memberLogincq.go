package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type MemberLoginCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	MemberLoginId *df.ConditionValue
	MemberId *df.ConditionValue
	LoginDatetime *df.ConditionValue
	MobileLoginFlg *df.ConditionValue
	LoginMemberStatusCode *df.ConditionValue
    conditionQueryMemberStatus *MemberStatusCQ
    conditionQueryMember *MemberCQ
}

func (q *MemberLoginCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *MemberLoginCQ) getCValueMemberLoginId() *df.ConditionValue {
	if q.MemberLoginId == nil {
		q.MemberLoginId = new(df.ConditionValue)
	}
	return q.MemberLoginId
}



func (q *MemberLoginCQ) SetMemberLoginId_Equal(value int64) *MemberLoginCQ {
	q.regMemberLoginId(df.CK_EQ_C, value)
	return q
}
func (q *MemberLoginCQ) SetMemberLoginId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueMemberLoginId(), "memberLoginId")
}
func (q *MemberLoginCQ) SetMemberLoginId_NotEqual(value int64) *MemberLoginCQ {
	q.regMemberLoginId(df.CK_NE_C, value)
	return q
}

func (q *MemberLoginCQ) SetMemberLoginId_GreaterThan(value int64) *MemberLoginCQ {
	q.regMemberLoginId(df.CK_GT_C, value)
	return q
}

func (q *MemberLoginCQ) SetMemberLoginId_LessThan(value int64) *MemberLoginCQ {
	q.regMemberLoginId(df.CK_LT_C, value)
	return q
}

func (q *MemberLoginCQ) SetMemberLoginId_GreaterEqual(value int64) *MemberLoginCQ {
	q.regMemberLoginId(df.CK_GE_C, value)
	return q
}

func (q *MemberLoginCQ) SetMemberLoginId_LessEqual(value int64) *MemberLoginCQ {
	q.regMemberLoginId(df.CK_LE_C, value)
	return q
}
func (q *MemberLoginCQ) SetMemberLoginId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueMemberLoginId(),"memberLoginId",rangeOfOption)
}	


func (q *MemberLoginCQ) SetMemberLoginId_IsNull() *MemberLoginCQ {
	q.regMemberLoginId(df.CK_ISN_C, 0)
	return q
}
func (q *MemberLoginCQ) SetMemberLoginId_IsNotNull() *MemberLoginCQ {
	q.regMemberLoginId(df.CK_ISNN_C, 0)
	return q
}
func (q *MemberLoginCQ) AddOrderBy_MemberLoginId_Asc() *MemberLoginCQ {
	q.BaseConditionQuery.RegOBA("memberLoginId")
	return q
}
func (q *MemberLoginCQ) AddOrderBy_MemberLoginId_Desc() *MemberLoginCQ {
	q.BaseConditionQuery.RegOBD("memberLoginId")
	return q
}
func (q *MemberLoginCQ) regMemberLoginId(key *df.ConditionKey, value interface{}) {
	if q.MemberLoginId == nil {
		q.MemberLoginId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.MemberLoginId, "memberLoginId")
}

func (q *MemberLoginCQ) getCValueMemberId() *df.ConditionValue {
	if q.MemberId == nil {
		q.MemberId = new(df.ConditionValue)
	}
	return q.MemberId
}



func (q *MemberLoginCQ) SetMemberId_Equal(value int64) *MemberLoginCQ {
	q.regMemberId(df.CK_EQ_C, value)
	return q
}
func (q *MemberLoginCQ) SetMemberId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueMemberId(), "memberId")
}
func (q *MemberLoginCQ) SetMemberId_NotEqual(value int64) *MemberLoginCQ {
	q.regMemberId(df.CK_NE_C, value)
	return q
}

func (q *MemberLoginCQ) SetMemberId_GreaterThan(value int64) *MemberLoginCQ {
	q.regMemberId(df.CK_GT_C, value)
	return q
}

func (q *MemberLoginCQ) SetMemberId_LessThan(value int64) *MemberLoginCQ {
	q.regMemberId(df.CK_LT_C, value)
	return q
}

func (q *MemberLoginCQ) SetMemberId_GreaterEqual(value int64) *MemberLoginCQ {
	q.regMemberId(df.CK_GE_C, value)
	return q
}

func (q *MemberLoginCQ) SetMemberId_LessEqual(value int64) *MemberLoginCQ {
	q.regMemberId(df.CK_LE_C, value)
	return q
}
func (q *MemberLoginCQ) SetMemberId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueMemberId(),"memberId",rangeOfOption)
}	


func (q *MemberLoginCQ) AddOrderBy_MemberId_Asc() *MemberLoginCQ {
	q.BaseConditionQuery.RegOBA("memberId")
	return q
}
func (q *MemberLoginCQ) AddOrderBy_MemberId_Desc() *MemberLoginCQ {
	q.BaseConditionQuery.RegOBD("memberId")
	return q
}
func (q *MemberLoginCQ) regMemberId(key *df.ConditionKey, value interface{}) {
	if q.MemberId == nil {
		q.MemberId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.MemberId, "memberId")
}

func (q *MemberLoginCQ) getCValueLoginDatetime() *df.ConditionValue {
	if q.LoginDatetime == nil {
		q.LoginDatetime = new(df.ConditionValue)
	}
	return q.LoginDatetime
}




func (q *MemberLoginCQ) SetLoginDatetime_Equal(value df.MysqlTimestamp) *MemberLoginCQ {
	q.regLoginDatetime(df.CK_EQ_C, value)
	return q
}


func (q *MemberLoginCQ) SetLoginDatetime_GreaterThan(value df.MysqlTimestamp) *MemberLoginCQ {
	q.regLoginDatetime(df.CK_GT_C, value)
	return q
}

func (q *MemberLoginCQ) SetLoginDatetime_LessThan(value df.MysqlTimestamp) *MemberLoginCQ {
	q.regLoginDatetime(df.CK_LT_C, value)
	return q
}

func (q *MemberLoginCQ) SetLoginDatetime_GreaterEqual(value df.MysqlTimestamp) *MemberLoginCQ {
	q.regLoginDatetime(df.CK_GE_C, value)
	return q
}

func (q *MemberLoginCQ) SetLoginDatetime_LessEqual(value df.MysqlTimestamp) *MemberLoginCQ {
	q.regLoginDatetime(df.CK_LE_C, value)
	return q
}

func (q *MemberLoginCQ) AddOrderBy_LoginDatetime_Asc() *MemberLoginCQ {
	q.BaseConditionQuery.RegOBA("loginDatetime")
	return q
}
func (q *MemberLoginCQ) AddOrderBy_LoginDatetime_Desc() *MemberLoginCQ {
	q.BaseConditionQuery.RegOBD("loginDatetime")
	return q
}
func (q *MemberLoginCQ) regLoginDatetime(key *df.ConditionKey, value interface{}) {
	if q.LoginDatetime == nil {
		q.LoginDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.LoginDatetime, "loginDatetime")
}

func (q *MemberLoginCQ) getCValueMobileLoginFlg() *df.ConditionValue {
	if q.MobileLoginFlg == nil {
		q.MobileLoginFlg = new(df.ConditionValue)
	}
	return q.MobileLoginFlg
}



func (q *MemberLoginCQ) SetMobileLoginFlg_Equal(value int64) *MemberLoginCQ {
	q.regMobileLoginFlg(df.CK_EQ_C, value)
	return q
}
func (q *MemberLoginCQ) SetMobileLoginFlg_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueMobileLoginFlg(), "mobileLoginFlg")
}
func (q *MemberLoginCQ) SetMobileLoginFlg_NotEqual(value int64) *MemberLoginCQ {
	q.regMobileLoginFlg(df.CK_NE_C, value)
	return q
}

func (q *MemberLoginCQ) SetMobileLoginFlg_GreaterThan(value int64) *MemberLoginCQ {
	q.regMobileLoginFlg(df.CK_GT_C, value)
	return q
}

func (q *MemberLoginCQ) SetMobileLoginFlg_LessThan(value int64) *MemberLoginCQ {
	q.regMobileLoginFlg(df.CK_LT_C, value)
	return q
}

func (q *MemberLoginCQ) SetMobileLoginFlg_GreaterEqual(value int64) *MemberLoginCQ {
	q.regMobileLoginFlg(df.CK_GE_C, value)
	return q
}

func (q *MemberLoginCQ) SetMobileLoginFlg_LessEqual(value int64) *MemberLoginCQ {
	q.regMobileLoginFlg(df.CK_LE_C, value)
	return q
}
func (q *MemberLoginCQ) SetMobileLoginFlg_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueMobileLoginFlg(),"mobileLoginFlg",rangeOfOption)
}	


func (q *MemberLoginCQ) AddOrderBy_MobileLoginFlg_Asc() *MemberLoginCQ {
	q.BaseConditionQuery.RegOBA("mobileLoginFlg")
	return q
}
func (q *MemberLoginCQ) AddOrderBy_MobileLoginFlg_Desc() *MemberLoginCQ {
	q.BaseConditionQuery.RegOBD("mobileLoginFlg")
	return q
}
func (q *MemberLoginCQ) regMobileLoginFlg(key *df.ConditionKey, value interface{}) {
	if q.MobileLoginFlg == nil {
		q.MobileLoginFlg = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.MobileLoginFlg, "mobileLoginFlg")
}

func (q *MemberLoginCQ) getCValueLoginMemberStatusCode() *df.ConditionValue {
	if q.LoginMemberStatusCode == nil {
		q.LoginMemberStatusCode = new(df.ConditionValue)
	}
	return q.LoginMemberStatusCode
}


func (q *MemberLoginCQ) SetLoginMemberStatusCode_Equal(value string) *MemberLoginCQ {
	q.regLoginMemberStatusCode(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *MemberLoginCQ) SetLoginMemberStatusCode_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueLoginMemberStatusCode(), "loginMemberStatusCode")
}
func (q *MemberLoginCQ) SetLoginMemberStatusCode_NotEqual(value string) *MemberLoginCQ {
	q.regLoginMemberStatusCode(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberLoginCQ) SetLoginMemberStatusCode_GreaterThan(value string) *MemberLoginCQ {
	q.regLoginMemberStatusCode(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberLoginCQ) SetLoginMemberStatusCode_LessThan(value string) *MemberLoginCQ {
	q.regLoginMemberStatusCode(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberLoginCQ) SetLoginMemberStatusCode_GreaterEqualThan(value string) *MemberLoginCQ {
	q.regLoginMemberStatusCode(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *MemberLoginCQ) SetLoginMemberStatusCode_LessEqualThan(value string) *MemberLoginCQ {
	q.regLoginMemberStatusCode(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberLoginCQ) SetLoginMemberStatusCode_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueLoginMemberStatusCode(), "loginMemberStatusCode", option)
}

func (q *MemberLoginCQ) SetLoginMemberStatusCode_PrefixSearch(value string) error {
	return q.SetLoginMemberStatusCode_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *MemberLoginCQ) SetLoginMemberStatusCode_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueLoginMemberStatusCode(), "loginMemberStatusCode", option)
}



func (q *MemberLoginCQ) AddOrderBy_LoginMemberStatusCode_Asc() *MemberLoginCQ {
	q.BaseConditionQuery.RegOBA("loginMemberStatusCode")
	return q
}
func (q *MemberLoginCQ) AddOrderBy_LoginMemberStatusCode_Desc() *MemberLoginCQ {
	q.BaseConditionQuery.RegOBD("loginMemberStatusCode")
	return q
}
func (q *MemberLoginCQ) regLoginMemberStatusCode(key *df.ConditionKey, value interface{}) {
	if q.LoginMemberStatusCode == nil {
		q.LoginMemberStatusCode = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.LoginMemberStatusCode, "loginMemberStatusCode")
}


func (q *MemberLoginCQ) QueryMemberStatus() *MemberStatusCQ {
	if q.conditionQueryMemberStatus == nil {
		q.conditionQueryMemberStatus = q.xcreateQueryMemberStatus()
		q.xsetupOuterJoinMemberStatus()
	}
	return q.conditionQueryMemberStatus
}

func (q *MemberLoginCQ) xcreateQueryMemberStatus() *MemberStatusCQ {
	nrp := q.BaseConditionQuery.ResolveNextRelationPath("MemberLogin", "MemberStatus")
	jan := q.BaseConditionQuery.ResolveJoinAliasName(nrp)
	var basecq df.ConditionQuery = q
	cq := CreateMemberStatusCQ(&basecq, q.BaseConditionQuery.SqlClause, jan, q.BaseConditionQuery.NestLevel+1)
	cq.BaseConditionQuery.BaseCB = q.BaseConditionQuery.BaseCB
	cq.BaseConditionQuery.ForeignPropertyName = "MemberStatus"
	cq.BaseConditionQuery.RelationPath = nrp
	return cq
}
func (q *MemberLoginCQ) xsetupOuterJoinMemberStatus() {
	    cq := q.QueryMemberStatus()
        joinOnMap := make(map[string]string)
        joinOnMap["loginMemberStatusCode"]="memberStatusCode"
        q.BaseConditionQuery.RegisterOuterJoin(
        	cq.BaseConditionQuery.ConditionQuery, joinOnMap, "MemberStatus");
}	
	
func (q *MemberLoginCQ) QueryMember() *MemberCQ {
	if q.conditionQueryMember == nil {
		q.conditionQueryMember = q.xcreateQueryMember()
		q.xsetupOuterJoinMember()
	}
	return q.conditionQueryMember
}

func (q *MemberLoginCQ) xcreateQueryMember() *MemberCQ {
	nrp := q.BaseConditionQuery.ResolveNextRelationPath("MemberLogin", "Member")
	jan := q.BaseConditionQuery.ResolveJoinAliasName(nrp)
	var basecq df.ConditionQuery = q
	cq := CreateMemberCQ(&basecq, q.BaseConditionQuery.SqlClause, jan, q.BaseConditionQuery.NestLevel+1)
	cq.BaseConditionQuery.BaseCB = q.BaseConditionQuery.BaseCB
	cq.BaseConditionQuery.ForeignPropertyName = "Member"
	cq.BaseConditionQuery.RelationPath = nrp
	return cq
}
func (q *MemberLoginCQ) xsetupOuterJoinMember() {
	    cq := q.QueryMember()
        joinOnMap := make(map[string]string)
        joinOnMap["memberId"]="memberId"
        q.BaseConditionQuery.RegisterOuterJoin(
        	cq.BaseConditionQuery.ConditionQuery, joinOnMap, "Member");
}	
	
func CreateMemberLoginCQ(referrerQuery *df.ConditionQuery, sqlClause *df.SqlClause, aliasName string, nestlevel int8) *MemberLoginCQ {
	cq := new(MemberLoginCQ)
	cq.BaseConditionQuery = new(df.BaseConditionQuery)
	cq.BaseConditionQuery.TableDbName = "MemberLogin"
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