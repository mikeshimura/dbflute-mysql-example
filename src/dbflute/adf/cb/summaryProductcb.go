package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/cq"
	"dbflute/adf/meta"
)

type SummaryProductCB struct {
	BaseConditionBean *df.BaseConditionBean
	Query             *cq.SummaryProductCQ
}

func CreateSummaryProductCB() *SummaryProductCB {
	cb := new(SummaryProductCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "SummaryProduct"
	cb.BaseConditionBean.Name = "SummaryProductCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	//dm:=DBMetaProvider_I.TableDbNameInstanceMap["SummaryProduct"]
	var dmx df.DBMeta = meta.SummaryProductDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	cb.Query = cb.createConditionQuery(nil, cb.BaseConditionBean.SqlClause, (*cb.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
	return cb
}
func (l *SummaryProductCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *SummaryProductCB) createConditionQuery(referrerQuery *df.ConditionQuery, sqlClause *df.SqlClause, aliasName string, nestlevel int8) *cq.SummaryProductCQ {
	cq := new(cq.SummaryProductCQ)
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

func (l *SummaryProductCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}
