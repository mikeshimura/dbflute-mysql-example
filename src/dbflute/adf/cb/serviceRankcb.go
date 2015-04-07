package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/cq"
	"dbflute/adf/meta"
)

type ServiceRankCB struct {
	BaseConditionBean *df.BaseConditionBean
	Query             *cq.ServiceRankCQ
}

func CreateServiceRankCB() *ServiceRankCB {
	cb := new(ServiceRankCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "ServiceRank"
	cb.BaseConditionBean.Name = "ServiceRankCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	//dm:=DBMetaProvider_I.TableDbNameInstanceMap["ServiceRank"]
	var dmx df.DBMeta = meta.ServiceRankDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	cb.Query = cb.createConditionQuery(nil, cb.BaseConditionBean.SqlClause, (*cb.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
	return cb
}
func (l *ServiceRankCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *ServiceRankCB) createConditionQuery(referrerQuery *df.ConditionQuery, sqlClause *df.SqlClause, aliasName string, nestlevel int8) *cq.ServiceRankCQ {
	cq := new(cq.ServiceRankCQ)
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

func (l *ServiceRankCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}