package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/cq"
	"dbflute/adf/meta"
)

type SummaryWithdrawalCB struct {
	BaseConditionBean *df.BaseConditionBean
	query             *cq.SummaryWithdrawalCQ
}

func CreateSummaryWithdrawalCB() *SummaryWithdrawalCB {
	cb := new(SummaryWithdrawalCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "SummaryWithdrawal"
	cb.BaseConditionBean.Name = "SummaryWithdrawalCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	var dmx df.DBMeta = meta.SummaryWithdrawalDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	return cb
}

func (l *SummaryWithdrawalCB) Query() *cq.SummaryWithdrawalCQ {
	if l.query == nil {
		l.query = cq.CreateSummaryWithdrawalCQ(nil, l.BaseConditionBean.SqlClause, (*l.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
		var cb df.ConditionBean = l
		l.query.BaseConditionQuery.BaseCB = &cb	
	}
	return l.query
}
func (l *SummaryWithdrawalCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *SummaryWithdrawalCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}


func (l *SummaryWithdrawalCB) FetchFirst(fetchSize int){
	(*l.GetBaseConditionBean().SqlClause).FetchFirst(fetchSize)
}

func (l *SummaryWithdrawalCB) OrScopeQuery(fquery func(*SummaryWithdrawalCB)){
	(*l.BaseConditionBean.SqlClause).MakeOrScopeQueryEffective()
	fquery(l)
	(*l.BaseConditionBean.SqlClause).CloseOrScopeQuery()
}

type SummaryWithdrawalNss struct {
	Query *cq.SummaryWithdrawalCQ
}
func (p *SummaryWithdrawalNss) hasConditionQuery() bool {
	return p.Query != nil
}