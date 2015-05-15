package adf

import (
	"github.com/mikeshimura/dbflute/df"
)

type CommonColumnAutoSetupperImpl struct {
	df.BaseCommonColumnAutoSetupper
}

func (p *CommonColumnAutoSetupperImpl) DoHandleCommonColumnOfInsertIfNeeds(
	entity *df.Entity, context *df.Context) {

}
func (p *CommonColumnAutoSetupperImpl) DoHandleCommonColumnOfUpdateIfNeeds(
	entity *df.Entity, context *df.Context) {
}