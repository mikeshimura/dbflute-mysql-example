package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/cq"
	"dbflute/adf/meta"
)

type PurchaseCB struct {
	BaseConditionBean *df.BaseConditionBean
	query             *cq.PurchaseCQ
    NssMember *MemberNss
    NssProduct *ProductNss
}

func CreatePurchaseCB() *PurchaseCB {
	cb := new(PurchaseCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "Purchase"
	cb.BaseConditionBean.Name = "PurchaseCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	var dmx df.DBMeta = meta.PurchaseDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	return cb
}

func (l *PurchaseCB) Query() *cq.PurchaseCQ {
	if l.query == nil {
		l.query = cq.CreatePurchaseCQ(nil, l.BaseConditionBean.SqlClause, (*l.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
		var cb df.ConditionBean = l
		l.query.BaseConditionQuery.BaseCB = &cb	
	}
	return l.query
}
func (l *PurchaseCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *PurchaseCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}

func (l *PurchaseCB) SetupSelect_Member() *MemberNss {
	l.BaseConditionBean.DoSetupSelect(l.Query().GetBaseConditionQuery(),
		l.Query().QueryMember().GetBaseConditionQuery())
	if l.NssMember == nil || !l.NssMember.hasConditionQuery() {
		l.NssMember = new(MemberNss)
		l.NssMember.Query = l.Query().QueryMember()
	}
	return l.NssMember
}
func (l *PurchaseCB) SetupSelect_Product() *ProductNss {
	l.BaseConditionBean.DoSetupSelect(l.Query().GetBaseConditionQuery(),
		l.Query().QueryProduct().GetBaseConditionQuery())
	if l.NssProduct == nil || !l.NssProduct.hasConditionQuery() {
		l.NssProduct = new(ProductNss)
		l.NssProduct.Query = l.Query().QueryProduct()
	}
	return l.NssProduct
}

func (l *PurchaseCB) FetchFirst(fetchSize int){
	(*l.GetBaseConditionBean().SqlClause).FetchFirst(fetchSize)
}

func (l *PurchaseCB) OrScopeQuery(fquery func(*PurchaseCB)){
	(*l.BaseConditionBean.SqlClause).MakeOrScopeQueryEffective()
	fquery(l)
	(*l.BaseConditionBean.SqlClause).CloseOrScopeQuery()
}

type PurchaseNss struct {
	Query *cq.PurchaseCQ
    NssMember *MemberNss
    NssProduct *ProductNss
}
func (p *PurchaseNss) WithMember() *MemberNss{
	(*p.Query.BaseConditionQuery.BaseCB).GetBaseConditionBean().
	DoSetupSelect(p.Query.BaseConditionQuery,p.Query.QueryMember().GetBaseConditionQuery())
	if p.NssMember == nil || !p.NssMember.hasConditionQuery() {
		p.NssMember = new(MemberNss)
		p.NssMember.Query = p.Query.QueryMember()
	}
	return p.NssMember
}
func (p *PurchaseNss) WithProduct() *ProductNss{
	(*p.Query.BaseConditionQuery.BaseCB).GetBaseConditionBean().
	DoSetupSelect(p.Query.BaseConditionQuery,p.Query.QueryProduct().GetBaseConditionQuery())
	if p.NssProduct == nil || !p.NssProduct.hasConditionQuery() {
		p.NssProduct = new(ProductNss)
		p.NssProduct.Query = p.Query.QueryProduct()
	}
	return p.NssProduct
}
func (p *PurchaseNss) hasConditionQuery() bool {
	return p.Query != nil
}