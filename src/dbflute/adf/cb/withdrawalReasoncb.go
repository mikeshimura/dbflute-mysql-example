package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/cq"
	"dbflute/adf/meta"
)

type WithdrawalReasonCB struct {
	BaseConditionBean *df.BaseConditionBean
	query             *cq.WithdrawalReasonCQ
}

func CreateWithdrawalReasonCB() *WithdrawalReasonCB {
	cb := new(WithdrawalReasonCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "WithdrawalReason"
	cb.BaseConditionBean.Name = "WithdrawalReasonCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	var dmx df.DBMeta = meta.WithdrawalReasonDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	return cb
}

func (l *WithdrawalReasonCB) Query() *cq.WithdrawalReasonCQ {
	if l.query == nil {
		l.query = cq.CreateWithdrawalReasonCQ(nil, l.BaseConditionBean.SqlClause, (*l.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
		var cb df.ConditionBean = l
		l.query.BaseConditionQuery.BaseCB = &cb	
	}
	return l.query
}
func (l *WithdrawalReasonCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *WithdrawalReasonCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}


func (l *WithdrawalReasonCB) FetchFirst(fetchSize int){
	(*l.GetBaseConditionBean().SqlClause).FetchFirst(fetchSize)
}

func (l *WithdrawalReasonCB) OrScopeQuery(fquery func(*WithdrawalReasonCB)){
	(*l.BaseConditionBean.SqlClause).MakeOrScopeQueryEffective()
	fquery(l)
	(*l.BaseConditionBean.SqlClause).CloseOrScopeQuery()
}

type WithdrawalReasonNss struct {
	Query *cq.WithdrawalReasonCQ
}
func (p *WithdrawalReasonNss) hasConditionQuery() bool {
	return p.Query != nil
}