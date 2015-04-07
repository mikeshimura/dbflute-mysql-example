package centity

import (
	"github.com/mikeshimura/dbflute/df"
)

func init() {
	SelectMember := func() *df.Entity {
		var te df.Entity = new(C_SelectMember)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("C_SelectMember", SelectMember)
}