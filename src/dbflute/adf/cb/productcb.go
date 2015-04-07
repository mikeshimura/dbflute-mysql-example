package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/cq"
	"dbflute/adf/meta"
)

type ProductCB struct {
	BaseConditionBean *df.BaseConditionBean
	Query             *cq.ProductCQ
}

func CreateProductCB() *ProductCB {
	cb := new(ProductCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "Product"
	cb.BaseConditionBean.Name = "ProductCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	//dm:=DBMetaProvider_I.TableDbNameInstanceMap["Product"]
	var dmx df.DBMeta = meta.ProductDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	cb.Query = cb.createConditionQuery(nil, cb.BaseConditionBean.SqlClause, (*cb.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
	return cb
}
func (l *ProductCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *ProductCB) createConditionQuery(referrerQuery *df.ConditionQuery, sqlClause *df.SqlClause, aliasName string, nestlevel int8) *cq.ProductCQ {
	cq := new(cq.ProductCQ)
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

func (l *ProductCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}
