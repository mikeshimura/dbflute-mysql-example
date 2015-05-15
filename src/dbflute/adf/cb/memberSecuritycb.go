package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/cq"
	"dbflute/adf/meta"
)

type MemberSecurityCB struct {
	BaseConditionBean *df.BaseConditionBean
	query             *cq.MemberSecurityCQ
    NssMember *MemberNss
}

func CreateMemberSecurityCB() *MemberSecurityCB {
	cb := new(MemberSecurityCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "MemberSecurity"
	cb.BaseConditionBean.Name = "MemberSecurityCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	var dmx df.DBMeta = meta.MemberSecurityDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	return cb
}

func (l *MemberSecurityCB) Query() *cq.MemberSecurityCQ {
	if l.query == nil {
		l.query = cq.CreateMemberSecurityCQ(nil, l.BaseConditionBean.SqlClause, (*l.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
		var cb df.ConditionBean = l
		l.query.BaseConditionQuery.BaseCB = &cb	
	}
	return l.query
}
func (l *MemberSecurityCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *MemberSecurityCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}

func (l *MemberSecurityCB) SetupSelect_Member() *MemberNss {
	l.BaseConditionBean.DoSetupSelect(l.Query().GetBaseConditionQuery(),
		l.Query().QueryMember().GetBaseConditionQuery())
	if l.NssMember == nil || !l.NssMember.hasConditionQuery() {
		l.NssMember = new(MemberNss)
		l.NssMember.Query = l.Query().QueryMember()
	}
	return l.NssMember
}

func (l *MemberSecurityCB) FetchFirst(fetchSize int){
	(*l.GetBaseConditionBean().SqlClause).FetchFirst(fetchSize)
}

func (l *MemberSecurityCB) OrScopeQuery(fquery func(*MemberSecurityCB)){
	(*l.BaseConditionBean.SqlClause).MakeOrScopeQueryEffective()
	fquery(l)
	(*l.BaseConditionBean.SqlClause).CloseOrScopeQuery()
}

type MemberSecurityNss struct {
	Query *cq.MemberSecurityCQ
    NssMember *MemberNss
}
func (p *MemberSecurityNss) WithMember() *MemberNss{
	(*p.Query.BaseConditionQuery.BaseCB).GetBaseConditionBean().
	DoSetupSelect(p.Query.BaseConditionQuery,p.Query.QueryMember().GetBaseConditionQuery())
	if p.NssMember == nil || !p.NssMember.hasConditionQuery() {
		p.NssMember = new(MemberNss)
		p.NssMember.Query = p.Query.QueryMember()
	}
	return p.NssMember
}
func (p *MemberSecurityNss) hasConditionQuery() bool {
	return p.Query != nil
}