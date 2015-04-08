package pmb

import (
    "github.com/mikeshimura/dbflute/df"
    "database/sql"
	"dbflute/adf/bhv"
	_ "dbflute/adf/cmeta"
	_ "dbflute/adf/centity"
)

type C_UpdateMemberPmb struct {
    df.BasePmb
    Name interface{}
}
func (o *C_UpdateMemberPmb) GetName() interface{} {
	return o.Name
}
func (o *C_UpdateMemberPmb) SetName(value df.NullString) {
	o.Name = value	
}

func (o *C_UpdateMemberPmb) GetEntityType() string {
	return "C_UpdateMember"
}
func (o *C_UpdateMemberPmb) GetOutsideSqlPath() string {
	return "dbflute/adf/bhv/sql/MemberBhv_testUpdateMember.sql"
}

func (o *C_UpdateMemberPmb) Execute(tx *sql.Tx) (int64, error) {
	return bhv.MemberBhv_I.OutsideSql().Execute(o,tx)

}
