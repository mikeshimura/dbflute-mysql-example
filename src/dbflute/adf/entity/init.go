package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

func EntityInit() {
	Member := func() *df.Entity {
		var te df.Entity = new(Member)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("Member", Member)
	MemberAddress := func() *df.Entity {
		var te df.Entity = new(MemberAddress)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("MemberAddress", MemberAddress)
	MemberLogin := func() *df.Entity {
		var te df.Entity = new(MemberLogin)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("MemberLogin", MemberLogin)
	MemberSecurity := func() *df.Entity {
		var te df.Entity = new(MemberSecurity)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("MemberSecurity", MemberSecurity)
	MemberService := func() *df.Entity {
		var te df.Entity = new(MemberService)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("MemberService", MemberService)
	MemberStatus := func() *df.Entity {
		var te df.Entity = new(MemberStatus)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("MemberStatus", MemberStatus)
	MemberWithdrawal := func() *df.Entity {
		var te df.Entity = new(MemberWithdrawal)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("MemberWithdrawal", MemberWithdrawal)
	Product := func() *df.Entity {
		var te df.Entity = new(Product)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("Product", Product)
	ProductCategory := func() *df.Entity {
		var te df.Entity = new(ProductCategory)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("ProductCategory", ProductCategory)
	ProductStatus := func() *df.Entity {
		var te df.Entity = new(ProductStatus)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("ProductStatus", ProductStatus)
	Purchase := func() *df.Entity {
		var te df.Entity = new(Purchase)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("Purchase", Purchase)
	PurchasePayment := func() *df.Entity {
		var te df.Entity = new(PurchasePayment)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("PurchasePayment", PurchasePayment)
	Region := func() *df.Entity {
		var te df.Entity = new(Region)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("Region", Region)
	ServiceRank := func() *df.Entity {
		var te df.Entity = new(ServiceRank)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("ServiceRank", ServiceRank)
	SummaryProduct := func() *df.Entity {
		var te df.Entity = new(SummaryProduct)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("SummaryProduct", SummaryProduct)
	VendorDateFk := func() *df.Entity {
		var te df.Entity = new(VendorDateFk)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("VendorDateFk", VendorDateFk)
	VendorDatePk := func() *df.Entity {
		var te df.Entity = new(VendorDatePk)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("VendorDatePk", VendorDatePk)
	VendorInheritInu := func() *df.Entity {
		var te df.Entity = new(VendorInheritInu)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("VendorInheritInu", VendorInheritInu)
	VendorInheritNeko := func() *df.Entity {
		var te df.Entity = new(VendorInheritNeko)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("VendorInheritNeko", VendorInheritNeko)
	VendorLargeData := func() *df.Entity {
		var te df.Entity = new(VendorLargeData)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("VendorLargeData", VendorLargeData)
	VendorLargeDataRef := func() *df.Entity {
		var te df.Entity = new(VendorLargeDataRef)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("VendorLargeDataRef", VendorLargeDataRef)
	VendorPartMan := func() *df.Entity {
		var te df.Entity = new(VendorPartMan)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("VendorPartMan", VendorPartMan)
	VendorPartManHigh := func() *df.Entity {
		var te df.Entity = new(VendorPartManHigh)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("VendorPartManHigh", VendorPartManHigh)
	WhiteCompoundPk := func() *df.Entity {
		var te df.Entity = new(WhiteCompoundPk)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("WhiteCompoundPk", WhiteCompoundPk)
	WhiteCompoundPkRef := func() *df.Entity {
		var te df.Entity = new(WhiteCompoundPkRef)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("WhiteCompoundPkRef", WhiteCompoundPkRef)
	WhiteCompoundPkWrongOrder := func() *df.Entity {
		var te df.Entity = new(WhiteCompoundPkWrongOrder)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("WhiteCompoundPkWrongOrder", WhiteCompoundPkWrongOrder)
	WhiteNotPk := func() *df.Entity {
		var te df.Entity = new(WhiteNotPk)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("WhiteNotPk", WhiteNotPk)
	WhiteSameName := func() *df.Entity {
		var te df.Entity = new(WhiteSameName)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("WhiteSameName", WhiteSameName)
	WhiteSameNameRef := func() *df.Entity {
		var te df.Entity = new(WhiteSameNameRef)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("WhiteSameNameRef", WhiteSameNameRef)
	WhiteXlsMan := func() *df.Entity {
		var te df.Entity = new(WhiteXlsMan)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("WhiteXlsMan", WhiteXlsMan)
	WithdrawalReason := func() *df.Entity {
		var te df.Entity = new(WithdrawalReason)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("WithdrawalReason", WithdrawalReason)
}