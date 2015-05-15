package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/cq"
	"dbflute/adf/meta"
)

type MemberCB struct {
	BaseConditionBean *df.BaseConditionBean
	query             *cq.MemberCQ
    NssMemberStatus *MemberStatusNss
}

func CreateMemberCB() *MemberCB {
	cb := new(MemberCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "Member"
	cb.BaseConditionBean.Name = "MemberCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	var dmx df.DBMeta = meta.MemberDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	return cb
}

func (l *MemberCB) Query() *cq.MemberCQ {
	if l.query == nil {
		l.query = cq.CreateMemberCQ(nil, l.BaseConditionBean.SqlClause, (*l.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
		var cb df.ConditionBean = l
		l.query.BaseConditionQuery.BaseCB = &cb	
	}
	return l.query
}
func (l *MemberCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *MemberCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}

func (l *MemberCB) SetupSelect_MemberStatus() *MemberStatusNss {
	l.BaseConditionBean.DoSetupSelect(l.Query().GetBaseConditionQuery(),
		l.Query().QueryMemberStatus().GetBaseConditionQuery())
	if l.NssMemberStatus == nil || !l.NssMemberStatus.hasConditionQuery() {
		l.NssMemberStatus = new(MemberStatusNss)
		l.NssMemberStatus.Query = l.Query().QueryMemberStatus()
	}
	return l.NssMemberStatus
}

func (l *MemberCB) FetchFirst(fetchSize int){
	(*l.GetBaseConditionBean().SqlClause).FetchFirst(fetchSize)
}

func (l *MemberCB) OrScopeQuery(fquery func(*MemberCB)){
	(*l.BaseConditionBean.SqlClause).MakeOrScopeQueryEffective()
	fquery(l)
	(*l.BaseConditionBean.SqlClause).CloseOrScopeQuery()
}

type MemberNss struct {
	Query *cq.MemberCQ
    NssMemberStatus *MemberStatusNss
}
func (p *MemberNss) WithMemberStatus() *MemberStatusNss{
	(*p.Query.BaseConditionQuery.BaseCB).GetBaseConditionBean().
	DoSetupSelect(p.Query.BaseConditionQuery,p.Query.QueryMemberStatus().GetBaseConditionQuery())
	if p.NssMemberStatus == nil || !p.NssMemberStatus.hasConditionQuery() {
		p.NssMemberStatus = new(MemberStatusNss)
		p.NssMemberStatus.Query = p.Query.QueryMemberStatus()
	}
	return p.NssMemberStatus
}
func (p *MemberNss) hasConditionQuery() bool {
	return p.Query != nil
}