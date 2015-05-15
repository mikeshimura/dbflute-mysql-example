package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/cq"
	"dbflute/adf/meta"
)

type MemberWithdrawalCB struct {
	BaseConditionBean *df.BaseConditionBean
	query             *cq.MemberWithdrawalCQ
    NssMember *MemberNss
    NssWithdrawalReason *WithdrawalReasonNss
}

func CreateMemberWithdrawalCB() *MemberWithdrawalCB {
	cb := new(MemberWithdrawalCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "MemberWithdrawal"
	cb.BaseConditionBean.Name = "MemberWithdrawalCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	var dmx df.DBMeta = meta.MemberWithdrawalDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	return cb
}

func (l *MemberWithdrawalCB) Query() *cq.MemberWithdrawalCQ {
	if l.query == nil {
		l.query = cq.CreateMemberWithdrawalCQ(nil, l.BaseConditionBean.SqlClause, (*l.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
		var cb df.ConditionBean = l
		l.query.BaseConditionQuery.BaseCB = &cb	
	}
	return l.query
}
func (l *MemberWithdrawalCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *MemberWithdrawalCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}

func (l *MemberWithdrawalCB) SetupSelect_Member() *MemberNss {
	l.BaseConditionBean.DoSetupSelect(l.Query().GetBaseConditionQuery(),
		l.Query().QueryMember().GetBaseConditionQuery())
	if l.NssMember == nil || !l.NssMember.hasConditionQuery() {
		l.NssMember = new(MemberNss)
		l.NssMember.Query = l.Query().QueryMember()
	}
	return l.NssMember
}
func (l *MemberWithdrawalCB) SetupSelect_WithdrawalReason() *WithdrawalReasonNss {
	l.BaseConditionBean.DoSetupSelect(l.Query().GetBaseConditionQuery(),
		l.Query().QueryWithdrawalReason().GetBaseConditionQuery())
	if l.NssWithdrawalReason == nil || !l.NssWithdrawalReason.hasConditionQuery() {
		l.NssWithdrawalReason = new(WithdrawalReasonNss)
		l.NssWithdrawalReason.Query = l.Query().QueryWithdrawalReason()
	}
	return l.NssWithdrawalReason
}

func (l *MemberWithdrawalCB) FetchFirst(fetchSize int){
	(*l.GetBaseConditionBean().SqlClause).FetchFirst(fetchSize)
}

func (l *MemberWithdrawalCB) OrScopeQuery(fquery func(*MemberWithdrawalCB)){
	(*l.BaseConditionBean.SqlClause).MakeOrScopeQueryEffective()
	fquery(l)
	(*l.BaseConditionBean.SqlClause).CloseOrScopeQuery()
}

type MemberWithdrawalNss struct {
	Query *cq.MemberWithdrawalCQ
    NssMember *MemberNss
    NssWithdrawalReason *WithdrawalReasonNss
}
func (p *MemberWithdrawalNss) WithMember() *MemberNss{
	(*p.Query.BaseConditionQuery.BaseCB).GetBaseConditionBean().
	DoSetupSelect(p.Query.BaseConditionQuery,p.Query.QueryMember().GetBaseConditionQuery())
	if p.NssMember == nil || !p.NssMember.hasConditionQuery() {
		p.NssMember = new(MemberNss)
		p.NssMember.Query = p.Query.QueryMember()
	}
	return p.NssMember
}
func (p *MemberWithdrawalNss) WithWithdrawalReason() *WithdrawalReasonNss{
	(*p.Query.BaseConditionQuery.BaseCB).GetBaseConditionBean().
	DoSetupSelect(p.Query.BaseConditionQuery,p.Query.QueryWithdrawalReason().GetBaseConditionQuery())
	if p.NssWithdrawalReason == nil || !p.NssWithdrawalReason.hasConditionQuery() {
		p.NssWithdrawalReason = new(WithdrawalReasonNss)
		p.NssWithdrawalReason.Query = p.Query.QueryWithdrawalReason()
	}
	return p.NssWithdrawalReason
}
func (p *MemberWithdrawalNss) hasConditionQuery() bool {
	return p.Query != nil
}