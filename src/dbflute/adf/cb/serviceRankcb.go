package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/cq"
	"dbflute/adf/meta"
)

type ServiceRankCB struct {
	BaseConditionBean *df.BaseConditionBean
	query             *cq.ServiceRankCQ
}

func CreateServiceRankCB() *ServiceRankCB {
	cb := new(ServiceRankCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "ServiceRank"
	cb.BaseConditionBean.Name = "ServiceRankCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	var dmx df.DBMeta = meta.ServiceRankDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	return cb
}

func (l *ServiceRankCB) Query() *cq.ServiceRankCQ {
	if l.query == nil {
		l.query = cq.CreateServiceRankCQ(nil, l.BaseConditionBean.SqlClause, (*l.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
		var cb df.ConditionBean = l
		l.query.BaseConditionQuery.BaseCB = &cb	
	}
	return l.query
}
func (l *ServiceRankCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *ServiceRankCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}


func (l *ServiceRankCB) FetchFirst(fetchSize int){
	(*l.GetBaseConditionBean().SqlClause).FetchFirst(fetchSize)
}

func (l *ServiceRankCB) OrScopeQuery(fquery func(*ServiceRankCB)){
	(*l.BaseConditionBean.SqlClause).MakeOrScopeQueryEffective()
	fquery(l)
	(*l.BaseConditionBean.SqlClause).CloseOrScopeQuery()
}

type ServiceRankNss struct {
	Query *cq.ServiceRankCQ
}
func (p *ServiceRankNss) hasConditionQuery() bool {
	return p.Query != nil
}