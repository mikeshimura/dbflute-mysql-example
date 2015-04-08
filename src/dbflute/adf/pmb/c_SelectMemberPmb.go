package pmb

import (
    "github.com/mikeshimura/dbflute/df"
    "database/sql"
	"dbflute/adf/bhv"
	_ "dbflute/adf/cmeta"
	_ "dbflute/adf/centity"
)

type C_SelectMemberPmb struct {
    df.BasePmb
    Name interface{}
}
func (o *C_SelectMemberPmb) GetName() interface{} {
	return o.Name
}
func (o *C_SelectMemberPmb) SetName(value df.NullString) {
	o.Name = value	
}

func (o *C_SelectMemberPmb) GetEntityType() string {
	return "C_SelectMember"
}
func (o *C_SelectMemberPmb) GetOutsideSqlPath() string {
	return "dbflute/adf/bhv/sql/MemberBhv_testSelectMember.sql"
}

func (o *C_SelectMemberPmb) SelectList(tx *sql.Tx) (*df.ListResultBean, error) {
	return bhv.MemberBhv_I.OutsideSql().SelectList(o,tx)

}
