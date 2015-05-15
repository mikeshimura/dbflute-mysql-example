package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/cq"
	"dbflute/adf/meta"
)

type ProductCB struct {
	BaseConditionBean *df.BaseConditionBean
	query             *cq.ProductCQ
    NssProductCategory *ProductCategoryNss
    NssProductStatus *ProductStatusNss
}

func CreateProductCB() *ProductCB {
	cb := new(ProductCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "Product"
	cb.BaseConditionBean.Name = "ProductCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	var dmx df.DBMeta = meta.ProductDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	return cb
}

func (l *ProductCB) Query() *cq.ProductCQ {
	if l.query == nil {
		l.query = cq.CreateProductCQ(nil, l.BaseConditionBean.SqlClause, (*l.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
		var cb df.ConditionBean = l
		l.query.BaseConditionQuery.BaseCB = &cb	
	}
	return l.query
}
func (l *ProductCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *ProductCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}

func (l *ProductCB) SetupSelect_ProductCategory() *ProductCategoryNss {
	l.BaseConditionBean.DoSetupSelect(l.Query().GetBaseConditionQuery(),
		l.Query().QueryProductCategory().GetBaseConditionQuery())
	if l.NssProductCategory == nil || !l.NssProductCategory.hasConditionQuery() {
		l.NssProductCategory = new(ProductCategoryNss)
		l.NssProductCategory.Query = l.Query().QueryProductCategory()
	}
	return l.NssProductCategory
}
func (l *ProductCB) SetupSelect_ProductStatus() *ProductStatusNss {
	l.BaseConditionBean.DoSetupSelect(l.Query().GetBaseConditionQuery(),
		l.Query().QueryProductStatus().GetBaseConditionQuery())
	if l.NssProductStatus == nil || !l.NssProductStatus.hasConditionQuery() {
		l.NssProductStatus = new(ProductStatusNss)
		l.NssProductStatus.Query = l.Query().QueryProductStatus()
	}
	return l.NssProductStatus
}

func (l *ProductCB) FetchFirst(fetchSize int){
	(*l.GetBaseConditionBean().SqlClause).FetchFirst(fetchSize)
}

func (l *ProductCB) OrScopeQuery(fquery func(*ProductCB)){
	(*l.BaseConditionBean.SqlClause).MakeOrScopeQueryEffective()
	fquery(l)
	(*l.BaseConditionBean.SqlClause).CloseOrScopeQuery()
}

type ProductNss struct {
	Query *cq.ProductCQ
    NssProductCategory *ProductCategoryNss
    NssProductStatus *ProductStatusNss
}
func (p *ProductNss) WithProductCategory() *ProductCategoryNss{
	(*p.Query.BaseConditionQuery.BaseCB).GetBaseConditionBean().
	DoSetupSelect(p.Query.BaseConditionQuery,p.Query.QueryProductCategory().GetBaseConditionQuery())
	if p.NssProductCategory == nil || !p.NssProductCategory.hasConditionQuery() {
		p.NssProductCategory = new(ProductCategoryNss)
		p.NssProductCategory.Query = p.Query.QueryProductCategory()
	}
	return p.NssProductCategory
}
func (p *ProductNss) WithProductStatus() *ProductStatusNss{
	(*p.Query.BaseConditionQuery.BaseCB).GetBaseConditionBean().
	DoSetupSelect(p.Query.BaseConditionQuery,p.Query.QueryProductStatus().GetBaseConditionQuery())
	if p.NssProductStatus == nil || !p.NssProductStatus.hasConditionQuery() {
		p.NssProductStatus = new(ProductStatusNss)
		p.NssProductStatus.Query = p.Query.QueryProductStatus()
	}
	return p.NssProductStatus
}
func (p *ProductNss) hasConditionQuery() bool {
	return p.Query != nil
}