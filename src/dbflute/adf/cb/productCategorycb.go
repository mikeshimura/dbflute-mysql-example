package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/cq"
	"dbflute/adf/meta"
)

type ProductCategoryCB struct {
	BaseConditionBean *df.BaseConditionBean
	query             *cq.ProductCategoryCQ
    NssProductCategory *ProductCategoryNss
}

func CreateProductCategoryCB() *ProductCategoryCB {
	cb := new(ProductCategoryCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "ProductCategory"
	cb.BaseConditionBean.Name = "ProductCategoryCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	var dmx df.DBMeta = meta.ProductCategoryDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	return cb
}

func (l *ProductCategoryCB) Query() *cq.ProductCategoryCQ {
	if l.query == nil {
		l.query = cq.CreateProductCategoryCQ(nil, l.BaseConditionBean.SqlClause, (*l.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
		var cb df.ConditionBean = l
		l.query.BaseConditionQuery.BaseCB = &cb	
	}
	return l.query
}
func (l *ProductCategoryCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *ProductCategoryCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}

func (l *ProductCategoryCB) SetupSelect_ProductCategorySelf() *ProductCategoryNss {
	l.BaseConditionBean.DoSetupSelect(l.Query().GetBaseConditionQuery(),
		l.Query().QueryProductCategorySelf().GetBaseConditionQuery())
	if l.NssProductCategory == nil || !l.NssProductCategory.hasConditionQuery() {
		l.NssProductCategory = new(ProductCategoryNss)
		l.NssProductCategory.Query = l.Query().QueryProductCategorySelf()
	}
	return l.NssProductCategory
}

func (l *ProductCategoryCB) FetchFirst(fetchSize int){
	(*l.GetBaseConditionBean().SqlClause).FetchFirst(fetchSize)
}

func (l *ProductCategoryCB) OrScopeQuery(fquery func(*ProductCategoryCB)){
	(*l.BaseConditionBean.SqlClause).MakeOrScopeQueryEffective()
	fquery(l)
	(*l.BaseConditionBean.SqlClause).CloseOrScopeQuery()
}

type ProductCategoryNss struct {
	Query *cq.ProductCategoryCQ
    NssProductCategory *ProductCategoryNss
}
func (p *ProductCategoryNss) WithProductCategory() *ProductCategoryNss{
	(*p.Query.BaseConditionQuery.BaseCB).GetBaseConditionBean().
	DoSetupSelect(p.Query.BaseConditionQuery,p.Query.QueryProductCategorySelf().GetBaseConditionQuery())
	if p.NssProductCategory == nil || !p.NssProductCategory.hasConditionQuery() {
		p.NssProductCategory = new(ProductCategoryNss)
		p.NssProductCategory.Query = p.Query.QueryProductCategorySelf()
	}
	return p.NssProductCategory
}
func (p *ProductCategoryNss) hasConditionQuery() bool {
	return p.Query != nil
}