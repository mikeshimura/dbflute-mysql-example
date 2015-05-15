package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/cq"
	"dbflute/adf/meta"
)

type MemberAddressCB struct {
	BaseConditionBean *df.BaseConditionBean
	query             *cq.MemberAddressCQ
    NssMember *MemberNss
    NssRegion *RegionNss
}

func CreateMemberAddressCB() *MemberAddressCB {
	cb := new(MemberAddressCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "MemberAddress"
	cb.BaseConditionBean.Name = "MemberAddressCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	var dmx df.DBMeta = meta.MemberAddressDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	return cb
}

func (l *MemberAddressCB) Query() *cq.MemberAddressCQ {
	if l.query == nil {
		l.query = cq.CreateMemberAddressCQ(nil, l.BaseConditionBean.SqlClause, (*l.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
		var cb df.ConditionBean = l
		l.query.BaseConditionQuery.BaseCB = &cb	
	}
	return l.query
}
func (l *MemberAddressCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *MemberAddressCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}

func (l *MemberAddressCB) SetupSelect_Member() *MemberNss {
	l.BaseConditionBean.DoSetupSelect(l.Query().GetBaseConditionQuery(),
		l.Query().QueryMember().GetBaseConditionQuery())
	if l.NssMember == nil || !l.NssMember.hasConditionQuery() {
		l.NssMember = new(MemberNss)
		l.NssMember.Query = l.Query().QueryMember()
	}
	return l.NssMember
}
func (l *MemberAddressCB) SetupSelect_Region() *RegionNss {
	l.BaseConditionBean.DoSetupSelect(l.Query().GetBaseConditionQuery(),
		l.Query().QueryRegion().GetBaseConditionQuery())
	if l.NssRegion == nil || !l.NssRegion.hasConditionQuery() {
		l.NssRegion = new(RegionNss)
		l.NssRegion.Query = l.Query().QueryRegion()
	}
	return l.NssRegion
}

func (l *MemberAddressCB) FetchFirst(fetchSize int){
	(*l.GetBaseConditionBean().SqlClause).FetchFirst(fetchSize)
}

func (l *MemberAddressCB) OrScopeQuery(fquery func(*MemberAddressCB)){
	(*l.BaseConditionBean.SqlClause).MakeOrScopeQueryEffective()
	fquery(l)
	(*l.BaseConditionBean.SqlClause).CloseOrScopeQuery()
}

type MemberAddressNss struct {
	Query *cq.MemberAddressCQ
    NssMember *MemberNss
    NssRegion *RegionNss
}
func (p *MemberAddressNss) WithMember() *MemberNss{
	(*p.Query.BaseConditionQuery.BaseCB).GetBaseConditionBean().
	DoSetupSelect(p.Query.BaseConditionQuery,p.Query.QueryMember().GetBaseConditionQuery())
	if p.NssMember == nil || !p.NssMember.hasConditionQuery() {
		p.NssMember = new(MemberNss)
		p.NssMember.Query = p.Query.QueryMember()
	}
	return p.NssMember
}
func (p *MemberAddressNss) WithRegion() *RegionNss{
	(*p.Query.BaseConditionQuery.BaseCB).GetBaseConditionBean().
	DoSetupSelect(p.Query.BaseConditionQuery,p.Query.QueryRegion().GetBaseConditionQuery())
	if p.NssRegion == nil || !p.NssRegion.hasConditionQuery() {
		p.NssRegion = new(RegionNss)
		p.NssRegion.Query = p.Query.QueryRegion()
	}
	return p.NssRegion
}
func (p *MemberAddressNss) hasConditionQuery() bool {
	return p.Query != nil
}