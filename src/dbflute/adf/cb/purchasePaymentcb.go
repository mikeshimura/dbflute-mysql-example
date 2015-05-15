package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/cq"
	"dbflute/adf/meta"
)

type PurchasePaymentCB struct {
	BaseConditionBean *df.BaseConditionBean
	query             *cq.PurchasePaymentCQ
    NssPurchase *PurchaseNss
}

func CreatePurchasePaymentCB() *PurchasePaymentCB {
	cb := new(PurchasePaymentCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "PurchasePayment"
	cb.BaseConditionBean.Name = "PurchasePaymentCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	var dmx df.DBMeta = meta.PurchasePaymentDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	return cb
}

func (l *PurchasePaymentCB) Query() *cq.PurchasePaymentCQ {
	if l.query == nil {
		l.query = cq.CreatePurchasePaymentCQ(nil, l.BaseConditionBean.SqlClause, (*l.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
		var cb df.ConditionBean = l
		l.query.BaseConditionQuery.BaseCB = &cb	
	}
	return l.query
}
func (l *PurchasePaymentCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *PurchasePaymentCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}

func (l *PurchasePaymentCB) SetupSelect_Purchase() *PurchaseNss {
	l.BaseConditionBean.DoSetupSelect(l.Query().GetBaseConditionQuery(),
		l.Query().QueryPurchase().GetBaseConditionQuery())
	if l.NssPurchase == nil || !l.NssPurchase.hasConditionQuery() {
		l.NssPurchase = new(PurchaseNss)
		l.NssPurchase.Query = l.Query().QueryPurchase()
	}
	return l.NssPurchase
}

func (l *PurchasePaymentCB) FetchFirst(fetchSize int){
	(*l.GetBaseConditionBean().SqlClause).FetchFirst(fetchSize)
}

func (l *PurchasePaymentCB) OrScopeQuery(fquery func(*PurchasePaymentCB)){
	(*l.BaseConditionBean.SqlClause).MakeOrScopeQueryEffective()
	fquery(l)
	(*l.BaseConditionBean.SqlClause).CloseOrScopeQuery()
}

type PurchasePaymentNss struct {
	Query *cq.PurchasePaymentCQ
    NssPurchase *PurchaseNss
}
func (p *PurchasePaymentNss) WithPurchase() *PurchaseNss{
	(*p.Query.BaseConditionQuery.BaseCB).GetBaseConditionBean().
	DoSetupSelect(p.Query.BaseConditionQuery,p.Query.QueryPurchase().GetBaseConditionQuery())
	if p.NssPurchase == nil || !p.NssPurchase.hasConditionQuery() {
		p.NssPurchase = new(PurchaseNss)
		p.NssPurchase.Query = p.Query.QueryPurchase()
	}
	return p.NssPurchase
}
func (p *PurchasePaymentNss) hasConditionQuery() bool {
	return p.Query != nil
}