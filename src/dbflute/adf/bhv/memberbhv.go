package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type MemberBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *MemberBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.MemberDbm
	return &meta
}
func (l *MemberBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *MemberBhv) Update(member *entity.Member, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = member
	return l.BaseBehavior.DoUpdate(&entity, nil, tx, ctx)
}
func (l *MemberBhv) Insert(member *entity.Member, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = member
	return l.BaseBehavior.DoInsert(&entity, nil, tx, ctx)
}
func (l *MemberBhv) Delete(member *entity.Member, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = member
	return l.BaseBehavior.DoDelete(&entity, nil, tx, ctx)
}
func (l *MemberBhv) SelectList(cb *cb.MemberCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "Member", tx)
}
func (l *MemberBhv) SelectCount(cb *cb.MemberCB, tx *sql.Tx) (int64, error) {

	return l.BaseBehavior.DoSelectCount(cb, tx)
}
func (l *MemberBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *MemberBhv) ReadNextVal(tx *sql.Tx) int64{
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var MemberBhv_I *MemberBhv
