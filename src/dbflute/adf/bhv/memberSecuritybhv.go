package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type MemberSecurityBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *MemberSecurityBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.MemberSecurityDbm
	return &meta
}
func (l *MemberSecurityBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *MemberSecurityBhv) Update(memberSecurity *entity.MemberSecurity, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = memberSecurity
	return l.BaseBehavior.DoUpdate(&entity, nil, tx, ctx)
}
func (l *MemberSecurityBhv) Insert(memberSecurity *entity.MemberSecurity, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = memberSecurity
	return l.BaseBehavior.DoInsert(&entity, nil, tx, ctx)
}
func (l *MemberSecurityBhv) Delete(memberSecurity *entity.MemberSecurity, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = memberSecurity
	return l.BaseBehavior.DoDelete(&entity, nil, tx, ctx)
}
func (l *MemberSecurityBhv) SelectList(cb *cb.MemberSecurityCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "MemberSecurity", tx)
}
func (l *MemberSecurityBhv) SelectCount(cb *cb.MemberSecurityCB, tx *sql.Tx) (int64, error) {

	return l.BaseBehavior.DoSelectCount(cb, tx)
}
func (l *MemberSecurityBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *MemberSecurityBhv) ReadNextVal(tx *sql.Tx) int64{
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var MemberSecurityBhv_I *MemberSecurityBhv
