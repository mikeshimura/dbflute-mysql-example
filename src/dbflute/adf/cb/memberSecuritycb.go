package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/cq"
	"dbflute/adf/meta"
)

type MemberSecurityCB struct {
	BaseConditionBean *df.BaseConditionBean
	Query             *cq.MemberSecurityCQ
}

func CreateMemberSecurityCB() *MemberSecurityCB {
	cb := new(MemberSecurityCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "MemberSecurity"
	cb.BaseConditionBean.Name = "MemberSecurityCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	//dm:=DBMetaProvider_I.TableDbNameInstanceMap["MemberSecurity"]
	var dmx df.DBMeta = meta.MemberSecurityDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	cb.Query = cb.createConditionQuery(nil, cb.BaseConditionBean.SqlClause, (*cb.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
	return cb
}
func (l *MemberSecurityCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *MemberSecurityCB) createConditionQuery(referrerQuery *df.ConditionQuery, sqlClause *df.SqlClause, aliasName string, nestlevel int8) *cq.MemberSecurityCQ {
	cq := new(cq.MemberSecurityCQ)
	cq.BaseConditionQuery = new(df.BaseConditionQuery)
	cq.BaseConditionQuery.TableDbName = l.BaseConditionBean.TableDbName
	cq.BaseConditionQuery.ReferrerQuery = referrerQuery
	cq.BaseConditionQuery.SqlClause = sqlClause
	cq.BaseConditionQuery.AliasName = aliasName
	cq.BaseConditionQuery.NestLevel = nestlevel
	cq.BaseConditionQuery.DBMetaProvider = l.BaseConditionBean.DBMetaProvider
	cq.BaseConditionQuery.CQ_PROPERTY = "Query"
	cq.BaseConditionQuery.ConditionQuery=cq
	return cq
}

func (l *MemberSecurityCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}
