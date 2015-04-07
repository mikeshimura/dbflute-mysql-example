package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type MemberAddressBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *MemberAddressBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.MemberAddressDbm
	return &meta
}
func (l *MemberAddressBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *MemberAddressBhv) Update(memberAddress *entity.MemberAddress, tx *sql.Tx) (int64, error) {
	var entity df.Entity = memberAddress
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *MemberAddressBhv) Insert(memberAddress *entity.MemberAddress, tx *sql.Tx) (int64, error) {
	var entity df.Entity = memberAddress
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *MemberAddressBhv) Delete(memberAddress *entity.MemberAddress, tx *sql.Tx) (int64, error) {
	var entity df.Entity = memberAddress
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *MemberAddressBhv) SelectList(cb *cb.MemberAddressCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "MemberAddress", tx)
}
func (l *MemberAddressBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *MemberAddressBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var MemberAddressBhv_I *MemberAddressBhv
