package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type MemberStatusBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *MemberStatusBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.MemberStatusDbm
	return &meta
}
func (l *MemberStatusBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *MemberStatusBhv) Update(memberStatus *entity.MemberStatus, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = memberStatus
	return l.BaseBehavior.DoUpdate(&entity, nil, tx, ctx)
}
func (l *MemberStatusBhv) Insert(memberStatus *entity.MemberStatus, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = memberStatus
	return l.BaseBehavior.DoInsert(&entity, nil, tx, ctx)
}
func (l *MemberStatusBhv) Delete(memberStatus *entity.MemberStatus, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = memberStatus
	return l.BaseBehavior.DoDelete(&entity, nil, tx, ctx)
}
func (l *MemberStatusBhv) SelectList(cb *cb.MemberStatusCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "MemberStatus", tx)
}
func (l *MemberStatusBhv) SelectCount(cb *cb.MemberStatusCB, tx *sql.Tx) (int64, error) {

	return l.BaseBehavior.DoSelectCount(cb, tx)
}
func (l *MemberStatusBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *MemberStatusBhv) ReadNextVal(tx *sql.Tx) int64{
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var MemberStatusBhv_I *MemberStatusBhv
