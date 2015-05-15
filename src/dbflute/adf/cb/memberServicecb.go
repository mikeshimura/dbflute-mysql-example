package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/cq"
	"dbflute/adf/meta"
)

type MemberServiceCB struct {
	BaseConditionBean *df.BaseConditionBean
	query             *cq.MemberServiceCQ
    NssMember *MemberNss
    NssServiceRank *ServiceRankNss
}

func CreateMemberServiceCB() *MemberServiceCB {
	cb := new(MemberServiceCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "MemberService"
	cb.BaseConditionBean.Name = "MemberServiceCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	var dmx df.DBMeta = meta.MemberServiceDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	return cb
}

func (l *MemberServiceCB) Query() *cq.MemberServiceCQ {
	if l.query == nil {
		l.query = cq.CreateMemberServiceCQ(nil, l.BaseConditionBean.SqlClause, (*l.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
		var cb df.ConditionBean = l
		l.query.BaseConditionQuery.BaseCB = &cb	
	}
	return l.query
}
func (l *MemberServiceCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *MemberServiceCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}

func (l *MemberServiceCB) SetupSelect_Member() *MemberNss {
	l.BaseConditionBean.DoSetupSelect(l.Query().GetBaseConditionQuery(),
		l.Query().QueryMember().GetBaseConditionQuery())
	if l.NssMember == nil || !l.NssMember.hasConditionQuery() {
		l.NssMember = new(MemberNss)
		l.NssMember.Query = l.Query().QueryMember()
	}
	return l.NssMember
}
func (l *MemberServiceCB) SetupSelect_ServiceRank() *ServiceRankNss {
	l.BaseConditionBean.DoSetupSelect(l.Query().GetBaseConditionQuery(),
		l.Query().QueryServiceRank().GetBaseConditionQuery())
	if l.NssServiceRank == nil || !l.NssServiceRank.hasConditionQuery() {
		l.NssServiceRank = new(ServiceRankNss)
		l.NssServiceRank.Query = l.Query().QueryServiceRank()
	}
	return l.NssServiceRank
}

func (l *MemberServiceCB) FetchFirst(fetchSize int){
	(*l.GetBaseConditionBean().SqlClause).FetchFirst(fetchSize)
}

func (l *MemberServiceCB) OrScopeQuery(fquery func(*MemberServiceCB)){
	(*l.BaseConditionBean.SqlClause).MakeOrScopeQueryEffective()
	fquery(l)
	(*l.BaseConditionBean.SqlClause).CloseOrScopeQuery()
}

type MemberServiceNss struct {
	Query *cq.MemberServiceCQ
    NssMember *MemberNss
    NssServiceRank *ServiceRankNss
}
func (p *MemberServiceNss) WithMember() *MemberNss{
	(*p.Query.BaseConditionQuery.BaseCB).GetBaseConditionBean().
	DoSetupSelect(p.Query.BaseConditionQuery,p.Query.QueryMember().GetBaseConditionQuery())
	if p.NssMember == nil || !p.NssMember.hasConditionQuery() {
		p.NssMember = new(MemberNss)
		p.NssMember.Query = p.Query.QueryMember()
	}
	return p.NssMember
}
func (p *MemberServiceNss) WithServiceRank() *ServiceRankNss{
	(*p.Query.BaseConditionQuery.BaseCB).GetBaseConditionBean().
	DoSetupSelect(p.Query.BaseConditionQuery,p.Query.QueryServiceRank().GetBaseConditionQuery())
	if p.NssServiceRank == nil || !p.NssServiceRank.hasConditionQuery() {
		p.NssServiceRank = new(ServiceRankNss)
		p.NssServiceRank.Query = p.Query.QueryServiceRank()
	}
	return p.NssServiceRank
}
func (p *MemberServiceNss) hasConditionQuery() bool {
	return p.Query != nil
}