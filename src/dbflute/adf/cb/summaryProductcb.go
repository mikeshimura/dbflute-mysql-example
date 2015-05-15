package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/cq"
	"dbflute/adf/meta"
)

type SummaryProductCB struct {
	BaseConditionBean *df.BaseConditionBean
	query             *cq.SummaryProductCQ
}

func CreateSummaryProductCB() *SummaryProductCB {
	cb := new(SummaryProductCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "SummaryProduct"
	cb.BaseConditionBean.Name = "SummaryProductCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	var dmx df.DBMeta = meta.SummaryProductDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	return cb
}

func (l *SummaryProductCB) Query() *cq.SummaryProductCQ {
	if l.query == nil {
		l.query = cq.CreateSummaryProductCQ(nil, l.BaseConditionBean.SqlClause, (*l.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
		var cb df.ConditionBean = l
		l.query.BaseConditionQuery.BaseCB = &cb	
	}
	return l.query
}
func (l *SummaryProductCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *SummaryProductCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}


func (l *SummaryProductCB) FetchFirst(fetchSize int){
	(*l.GetBaseConditionBean().SqlClause).FetchFirst(fetchSize)
}

func (l *SummaryProductCB) OrScopeQuery(fquery func(*SummaryProductCB)){
	(*l.BaseConditionBean.SqlClause).MakeOrScopeQueryEffective()
	fquery(l)
	(*l.BaseConditionBean.SqlClause).CloseOrScopeQuery()
}

type SummaryProductNss struct {
	Query *cq.SummaryProductCQ
}
func (p *SummaryProductNss) hasConditionQuery() bool {
	return p.Query != nil
}