package cmeta

import (
	"github.com/mikeshimura/dbflute/df"
)

func init() {
	Create_C_SelectMemberDbm()
	var SelectMember_meta df.DBMeta = C_SelectMemberDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["C_SelectMember"] = &SelectMember_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("C_SelectMember", "C_SelectMember")
}