package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type MemberLoginBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *MemberLoginBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.MemberLoginDbm
	return &meta
}
func (l *MemberLoginBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *MemberLoginBhv) Update(memberLogin *entity.MemberLogin, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = memberLogin
	return l.BaseBehavior.DoUpdate(&entity, nil, tx, ctx)
}
func (l *MemberLoginBhv) Insert(memberLogin *entity.MemberLogin, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = memberLogin
	return l.BaseBehavior.DoInsert(&entity, nil, tx, ctx)
}
func (l *MemberLoginBhv) Delete(memberLogin *entity.MemberLogin, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = memberLogin
	return l.BaseBehavior.DoDelete(&entity, nil, tx, ctx)
}
func (l *MemberLoginBhv) SelectList(cb *cb.MemberLoginCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "MemberLogin", tx)
}
func (l *MemberLoginBhv) SelectCount(cb *cb.MemberLoginCB, tx *sql.Tx) (int64, error) {

	return l.BaseBehavior.DoSelectCount(cb, tx)
}
func (l *MemberLoginBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *MemberLoginBhv) ReadNextVal(tx *sql.Tx) int64{
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var MemberLoginBhv_I *MemberLoginBhv
