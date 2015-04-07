package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type ServiceRankBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *ServiceRankBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.ServiceRankDbm
	return &meta
}
func (l *ServiceRankBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *ServiceRankBhv) Update(serviceRank *entity.ServiceRank, tx *sql.Tx) (int64, error) {
	var entity df.Entity = serviceRank
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *ServiceRankBhv) Insert(serviceRank *entity.ServiceRank, tx *sql.Tx) (int64, error) {
	var entity df.Entity = serviceRank
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *ServiceRankBhv) Delete(serviceRank *entity.ServiceRank, tx *sql.Tx) (int64, error) {
	var entity df.Entity = serviceRank
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *ServiceRankBhv) SelectList(cb *cb.ServiceRankCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "ServiceRank", tx)
}
func (l *ServiceRankBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *ServiceRankBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var ServiceRankBhv_I *ServiceRankBhv
