package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type MemberServiceBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *MemberServiceBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.MemberServiceDbm
	return &meta
}
func (l *MemberServiceBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *MemberServiceBhv) Update(memberService *entity.MemberService, tx *sql.Tx) (int64, error) {
	var entity df.Entity = memberService
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *MemberServiceBhv) Insert(memberService *entity.MemberService, tx *sql.Tx) (int64, error) {
	var entity df.Entity = memberService
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *MemberServiceBhv) Delete(memberService *entity.MemberService, tx *sql.Tx) (int64, error) {
	var entity df.Entity = memberService
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *MemberServiceBhv) SelectList(cb *cb.MemberServiceCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "MemberService", tx)
}
func (l *MemberServiceBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *MemberServiceBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var MemberServiceBhv_I *MemberServiceBhv
