package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type RegionBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *RegionBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.RegionDbm
	return &meta
}
func (l *RegionBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *RegionBhv) Update(region *entity.Region, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = region
	return l.BaseBehavior.DoUpdate(&entity, nil, tx, ctx)
}
func (l *RegionBhv) Insert(region *entity.Region, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = region
	return l.BaseBehavior.DoInsert(&entity, nil, tx, ctx)
}
func (l *RegionBhv) Delete(region *entity.Region, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = region
	return l.BaseBehavior.DoDelete(&entity, nil, tx, ctx)
}
func (l *RegionBhv) SelectList(cb *cb.RegionCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "Region", tx)
}
func (l *RegionBhv) SelectCount(cb *cb.RegionCB, tx *sql.Tx) (int64, error) {

	return l.BaseBehavior.DoSelectCount(cb, tx)
}
func (l *RegionBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *RegionBhv) ReadNextVal(tx *sql.Tx) int64{
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var RegionBhv_I *RegionBhv
