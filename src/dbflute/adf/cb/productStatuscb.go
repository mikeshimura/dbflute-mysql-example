package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/cq"
	"dbflute/adf/meta"
)

type ProductStatusCB struct {
	BaseConditionBean *df.BaseConditionBean
	query             *cq.ProductStatusCQ
}

func CreateProductStatusCB() *ProductStatusCB {
	cb := new(ProductStatusCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "ProductStatus"
	cb.BaseConditionBean.Name = "ProductStatusCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	var dmx df.DBMeta = meta.ProductStatusDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	return cb
}

func (l *ProductStatusCB) Query() *cq.ProductStatusCQ {
	if l.query == nil {
		l.query = cq.CreateProductStatusCQ(nil, l.BaseConditionBean.SqlClause, (*l.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
		var cb df.ConditionBean = l
		l.query.BaseConditionQuery.BaseCB = &cb	
	}
	return l.query
}
func (l *ProductStatusCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *ProductStatusCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}


func (l *ProductStatusCB) FetchFirst(fetchSize int){
	(*l.GetBaseConditionBean().SqlClause).FetchFirst(fetchSize)
}

func (l *ProductStatusCB) OrScopeQuery(fquery func(*ProductStatusCB)){
	(*l.BaseConditionBean.SqlClause).MakeOrScopeQueryEffective()
	fquery(l)
	(*l.BaseConditionBean.SqlClause).CloseOrScopeQuery()
}

type ProductStatusNss struct {
	Query *cq.ProductStatusCQ
}
func (p *ProductStatusNss) hasConditionQuery() bool {
	return p.Query != nil
}