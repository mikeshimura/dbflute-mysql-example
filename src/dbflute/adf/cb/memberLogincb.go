package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/cq"
	"dbflute/adf/meta"
)

type MemberLoginCB struct {
	BaseConditionBean *df.BaseConditionBean
	query             *cq.MemberLoginCQ
    NssMemberStatus *MemberStatusNss
    NssMember *MemberNss
}

func CreateMemberLoginCB() *MemberLoginCB {
	cb := new(MemberLoginCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "MemberLogin"
	cb.BaseConditionBean.Name = "MemberLoginCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	var dmx df.DBMeta = meta.MemberLoginDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	return cb
}

func (l *MemberLoginCB) Query() *cq.MemberLoginCQ {
	if l.query == nil {
		l.query = cq.CreateMemberLoginCQ(nil, l.BaseConditionBean.SqlClause, (*l.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
		var cb df.ConditionBean = l
		l.query.BaseConditionQuery.BaseCB = &cb	
	}
	return l.query
}
func (l *MemberLoginCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *MemberLoginCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}

func (l *MemberLoginCB) SetupSelect_MemberStatus() *MemberStatusNss {
	l.BaseConditionBean.DoSetupSelect(l.Query().GetBaseConditionQuery(),
		l.Query().QueryMemberStatus().GetBaseConditionQuery())
	if l.NssMemberStatus == nil || !l.NssMemberStatus.hasConditionQuery() {
		l.NssMemberStatus = new(MemberStatusNss)
		l.NssMemberStatus.Query = l.Query().QueryMemberStatus()
	}
	return l.NssMemberStatus
}
func (l *MemberLoginCB) SetupSelect_Member() *MemberNss {
	l.BaseConditionBean.DoSetupSelect(l.Query().GetBaseConditionQuery(),
		l.Query().QueryMember().GetBaseConditionQuery())
	if l.NssMember == nil || !l.NssMember.hasConditionQuery() {
		l.NssMember = new(MemberNss)
		l.NssMember.Query = l.Query().QueryMember()
	}
	return l.NssMember
}

func (l *MemberLoginCB) FetchFirst(fetchSize int){
	(*l.GetBaseConditionBean().SqlClause).FetchFirst(fetchSize)
}

func (l *MemberLoginCB) OrScopeQuery(fquery func(*MemberLoginCB)){
	(*l.BaseConditionBean.SqlClause).MakeOrScopeQueryEffective()
	fquery(l)
	(*l.BaseConditionBean.SqlClause).CloseOrScopeQuery()
}

type MemberLoginNss struct {
	Query *cq.MemberLoginCQ
    NssMemberStatus *MemberStatusNss
    NssMember *MemberNss
}
func (p *MemberLoginNss) WithMemberStatus() *MemberStatusNss{
	(*p.Query.BaseConditionQuery.BaseCB).GetBaseConditionBean().
	DoSetupSelect(p.Query.BaseConditionQuery,p.Query.QueryMemberStatus().GetBaseConditionQuery())
	if p.NssMemberStatus == nil || !p.NssMemberStatus.hasConditionQuery() {
		p.NssMemberStatus = new(MemberStatusNss)
		p.NssMemberStatus.Query = p.Query.QueryMemberStatus()
	}
	return p.NssMemberStatus
}
func (p *MemberLoginNss) WithMember() *MemberNss{
	(*p.Query.BaseConditionQuery.BaseCB).GetBaseConditionBean().
	DoSetupSelect(p.Query.BaseConditionQuery,p.Query.QueryMember().GetBaseConditionQuery())
	if p.NssMember == nil || !p.NssMember.hasConditionQuery() {
		p.NssMember = new(MemberNss)
		p.NssMember.Query = p.Query.QueryMember()
	}
	return p.NssMember
}
func (p *MemberLoginNss) hasConditionQuery() bool {
	return p.Query != nil
}