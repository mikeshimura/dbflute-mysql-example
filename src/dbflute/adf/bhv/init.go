package bhv

import (
	"dbflute/adf"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/meta"
	_	"github.com/lib/pq"  //Don't eraze require to use this package
)
//To fix init sequence, all init routines are called from here.
func init() {
	adf.AdfInit()
	entity.EntityInit()
	meta.MetaInit()
	MemberBhv_I = new(MemberBhv)
	MemberBhv_I.BaseBehavior = new(df.BaseBehavior)
	MemberBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	MemberBhv_I.BaseBehavior.TableDbName = "Member"
	var member df.Behavior =MemberBhv_I
	MemberBhv_I.BaseBehavior.Behavior=&member
	MemberBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	MemberAddressBhv_I = new(MemberAddressBhv)
	MemberAddressBhv_I.BaseBehavior = new(df.BaseBehavior)
	MemberAddressBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	MemberAddressBhv_I.BaseBehavior.TableDbName = "MemberAddress"
	var memberAddress df.Behavior =MemberAddressBhv_I
	MemberAddressBhv_I.BaseBehavior.Behavior=&memberAddress
	MemberAddressBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	MemberLoginBhv_I = new(MemberLoginBhv)
	MemberLoginBhv_I.BaseBehavior = new(df.BaseBehavior)
	MemberLoginBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	MemberLoginBhv_I.BaseBehavior.TableDbName = "MemberLogin"
	var memberLogin df.Behavior =MemberLoginBhv_I
	MemberLoginBhv_I.BaseBehavior.Behavior=&memberLogin
	MemberLoginBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	MemberSecurityBhv_I = new(MemberSecurityBhv)
	MemberSecurityBhv_I.BaseBehavior = new(df.BaseBehavior)
	MemberSecurityBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	MemberSecurityBhv_I.BaseBehavior.TableDbName = "MemberSecurity"
	var memberSecurity df.Behavior =MemberSecurityBhv_I
	MemberSecurityBhv_I.BaseBehavior.Behavior=&memberSecurity
	MemberSecurityBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	MemberServiceBhv_I = new(MemberServiceBhv)
	MemberServiceBhv_I.BaseBehavior = new(df.BaseBehavior)
	MemberServiceBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	MemberServiceBhv_I.BaseBehavior.TableDbName = "MemberService"
	var memberService df.Behavior =MemberServiceBhv_I
	MemberServiceBhv_I.BaseBehavior.Behavior=&memberService
	MemberServiceBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	MemberStatusBhv_I = new(MemberStatusBhv)
	MemberStatusBhv_I.BaseBehavior = new(df.BaseBehavior)
	MemberStatusBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	MemberStatusBhv_I.BaseBehavior.TableDbName = "MemberStatus"
	var memberStatus df.Behavior =MemberStatusBhv_I
	MemberStatusBhv_I.BaseBehavior.Behavior=&memberStatus
	MemberStatusBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	MemberWithdrawalBhv_I = new(MemberWithdrawalBhv)
	MemberWithdrawalBhv_I.BaseBehavior = new(df.BaseBehavior)
	MemberWithdrawalBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	MemberWithdrawalBhv_I.BaseBehavior.TableDbName = "MemberWithdrawal"
	var memberWithdrawal df.Behavior =MemberWithdrawalBhv_I
	MemberWithdrawalBhv_I.BaseBehavior.Behavior=&memberWithdrawal
	MemberWithdrawalBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	ProductBhv_I = new(ProductBhv)
	ProductBhv_I.BaseBehavior = new(df.BaseBehavior)
	ProductBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	ProductBhv_I.BaseBehavior.TableDbName = "Product"
	var product df.Behavior =ProductBhv_I
	ProductBhv_I.BaseBehavior.Behavior=&product
	ProductBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	ProductCategoryBhv_I = new(ProductCategoryBhv)
	ProductCategoryBhv_I.BaseBehavior = new(df.BaseBehavior)
	ProductCategoryBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	ProductCategoryBhv_I.BaseBehavior.TableDbName = "ProductCategory"
	var productCategory df.Behavior =ProductCategoryBhv_I
	ProductCategoryBhv_I.BaseBehavior.Behavior=&productCategory
	ProductCategoryBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	ProductStatusBhv_I = new(ProductStatusBhv)
	ProductStatusBhv_I.BaseBehavior = new(df.BaseBehavior)
	ProductStatusBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	ProductStatusBhv_I.BaseBehavior.TableDbName = "ProductStatus"
	var productStatus df.Behavior =ProductStatusBhv_I
	ProductStatusBhv_I.BaseBehavior.Behavior=&productStatus
	ProductStatusBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	PurchaseBhv_I = new(PurchaseBhv)
	PurchaseBhv_I.BaseBehavior = new(df.BaseBehavior)
	PurchaseBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	PurchaseBhv_I.BaseBehavior.TableDbName = "Purchase"
	var purchase df.Behavior =PurchaseBhv_I
	PurchaseBhv_I.BaseBehavior.Behavior=&purchase
	PurchaseBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	PurchasePaymentBhv_I = new(PurchasePaymentBhv)
	PurchasePaymentBhv_I.BaseBehavior = new(df.BaseBehavior)
	PurchasePaymentBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	PurchasePaymentBhv_I.BaseBehavior.TableDbName = "PurchasePayment"
	var purchasePayment df.Behavior =PurchasePaymentBhv_I
	PurchasePaymentBhv_I.BaseBehavior.Behavior=&purchasePayment
	PurchasePaymentBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	RegionBhv_I = new(RegionBhv)
	RegionBhv_I.BaseBehavior = new(df.BaseBehavior)
	RegionBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	RegionBhv_I.BaseBehavior.TableDbName = "Region"
	var region df.Behavior =RegionBhv_I
	RegionBhv_I.BaseBehavior.Behavior=&region
	RegionBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	ServiceRankBhv_I = new(ServiceRankBhv)
	ServiceRankBhv_I.BaseBehavior = new(df.BaseBehavior)
	ServiceRankBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	ServiceRankBhv_I.BaseBehavior.TableDbName = "ServiceRank"
	var serviceRank df.Behavior =ServiceRankBhv_I
	ServiceRankBhv_I.BaseBehavior.Behavior=&serviceRank
	ServiceRankBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	SummaryProductBhv_I = new(SummaryProductBhv)
	SummaryProductBhv_I.BaseBehavior = new(df.BaseBehavior)
	SummaryProductBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	SummaryProductBhv_I.BaseBehavior.TableDbName = "SummaryProduct"
	var summaryProduct df.Behavior =SummaryProductBhv_I
	SummaryProductBhv_I.BaseBehavior.Behavior=&summaryProduct
	SummaryProductBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	VendorDateFkBhv_I = new(VendorDateFkBhv)
	VendorDateFkBhv_I.BaseBehavior = new(df.BaseBehavior)
	VendorDateFkBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	VendorDateFkBhv_I.BaseBehavior.TableDbName = "VendorDateFk"
	var vendorDateFk df.Behavior =VendorDateFkBhv_I
	VendorDateFkBhv_I.BaseBehavior.Behavior=&vendorDateFk
	VendorDateFkBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	VendorDatePkBhv_I = new(VendorDatePkBhv)
	VendorDatePkBhv_I.BaseBehavior = new(df.BaseBehavior)
	VendorDatePkBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	VendorDatePkBhv_I.BaseBehavior.TableDbName = "VendorDatePk"
	var vendorDatePk df.Behavior =VendorDatePkBhv_I
	VendorDatePkBhv_I.BaseBehavior.Behavior=&vendorDatePk
	VendorDatePkBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	VendorInheritInuBhv_I = new(VendorInheritInuBhv)
	VendorInheritInuBhv_I.BaseBehavior = new(df.BaseBehavior)
	VendorInheritInuBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	VendorInheritInuBhv_I.BaseBehavior.TableDbName = "VendorInheritInu"
	var vendorInheritInu df.Behavior =VendorInheritInuBhv_I
	VendorInheritInuBhv_I.BaseBehavior.Behavior=&vendorInheritInu
	VendorInheritInuBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	VendorInheritNekoBhv_I = new(VendorInheritNekoBhv)
	VendorInheritNekoBhv_I.BaseBehavior = new(df.BaseBehavior)
	VendorInheritNekoBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	VendorInheritNekoBhv_I.BaseBehavior.TableDbName = "VendorInheritNeko"
	var vendorInheritNeko df.Behavior =VendorInheritNekoBhv_I
	VendorInheritNekoBhv_I.BaseBehavior.Behavior=&vendorInheritNeko
	VendorInheritNekoBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	VendorLargeDataBhv_I = new(VendorLargeDataBhv)
	VendorLargeDataBhv_I.BaseBehavior = new(df.BaseBehavior)
	VendorLargeDataBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	VendorLargeDataBhv_I.BaseBehavior.TableDbName = "VendorLargeData"
	var vendorLargeData df.Behavior =VendorLargeDataBhv_I
	VendorLargeDataBhv_I.BaseBehavior.Behavior=&vendorLargeData
	VendorLargeDataBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	VendorLargeDataRefBhv_I = new(VendorLargeDataRefBhv)
	VendorLargeDataRefBhv_I.BaseBehavior = new(df.BaseBehavior)
	VendorLargeDataRefBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	VendorLargeDataRefBhv_I.BaseBehavior.TableDbName = "VendorLargeDataRef"
	var vendorLargeDataRef df.Behavior =VendorLargeDataRefBhv_I
	VendorLargeDataRefBhv_I.BaseBehavior.Behavior=&vendorLargeDataRef
	VendorLargeDataRefBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	VendorPartManBhv_I = new(VendorPartManBhv)
	VendorPartManBhv_I.BaseBehavior = new(df.BaseBehavior)
	VendorPartManBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	VendorPartManBhv_I.BaseBehavior.TableDbName = "VendorPartMan"
	var vendorPartMan df.Behavior =VendorPartManBhv_I
	VendorPartManBhv_I.BaseBehavior.Behavior=&vendorPartMan
	VendorPartManBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	VendorPartManHighBhv_I = new(VendorPartManHighBhv)
	VendorPartManHighBhv_I.BaseBehavior = new(df.BaseBehavior)
	VendorPartManHighBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	VendorPartManHighBhv_I.BaseBehavior.TableDbName = "VendorPartManHigh"
	var vendorPartManHigh df.Behavior =VendorPartManHighBhv_I
	VendorPartManHighBhv_I.BaseBehavior.Behavior=&vendorPartManHigh
	VendorPartManHighBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	WhiteCompoundPkBhv_I = new(WhiteCompoundPkBhv)
	WhiteCompoundPkBhv_I.BaseBehavior = new(df.BaseBehavior)
	WhiteCompoundPkBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	WhiteCompoundPkBhv_I.BaseBehavior.TableDbName = "WhiteCompoundPk"
	var whiteCompoundPk df.Behavior =WhiteCompoundPkBhv_I
	WhiteCompoundPkBhv_I.BaseBehavior.Behavior=&whiteCompoundPk
	WhiteCompoundPkBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	WhiteCompoundPkRefBhv_I = new(WhiteCompoundPkRefBhv)
	WhiteCompoundPkRefBhv_I.BaseBehavior = new(df.BaseBehavior)
	WhiteCompoundPkRefBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	WhiteCompoundPkRefBhv_I.BaseBehavior.TableDbName = "WhiteCompoundPkRef"
	var whiteCompoundPkRef df.Behavior =WhiteCompoundPkRefBhv_I
	WhiteCompoundPkRefBhv_I.BaseBehavior.Behavior=&whiteCompoundPkRef
	WhiteCompoundPkRefBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	WhiteCompoundPkWrongOrderBhv_I = new(WhiteCompoundPkWrongOrderBhv)
	WhiteCompoundPkWrongOrderBhv_I.BaseBehavior = new(df.BaseBehavior)
	WhiteCompoundPkWrongOrderBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	WhiteCompoundPkWrongOrderBhv_I.BaseBehavior.TableDbName = "WhiteCompoundPkWrongOrder"
	var whiteCompoundPkWrongOrder df.Behavior =WhiteCompoundPkWrongOrderBhv_I
	WhiteCompoundPkWrongOrderBhv_I.BaseBehavior.Behavior=&whiteCompoundPkWrongOrder
	WhiteCompoundPkWrongOrderBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	WhiteNotPkBhv_I = new(WhiteNotPkBhv)
	WhiteNotPkBhv_I.BaseBehavior = new(df.BaseBehavior)
	WhiteNotPkBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	WhiteNotPkBhv_I.BaseBehavior.TableDbName = "WhiteNotPk"
	var whiteNotPk df.Behavior =WhiteNotPkBhv_I
	WhiteNotPkBhv_I.BaseBehavior.Behavior=&whiteNotPk
	WhiteNotPkBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	WhiteSameNameBhv_I = new(WhiteSameNameBhv)
	WhiteSameNameBhv_I.BaseBehavior = new(df.BaseBehavior)
	WhiteSameNameBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	WhiteSameNameBhv_I.BaseBehavior.TableDbName = "WhiteSameName"
	var whiteSameName df.Behavior =WhiteSameNameBhv_I
	WhiteSameNameBhv_I.BaseBehavior.Behavior=&whiteSameName
	WhiteSameNameBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	WhiteSameNameRefBhv_I = new(WhiteSameNameRefBhv)
	WhiteSameNameRefBhv_I.BaseBehavior = new(df.BaseBehavior)
	WhiteSameNameRefBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	WhiteSameNameRefBhv_I.BaseBehavior.TableDbName = "WhiteSameNameRef"
	var whiteSameNameRef df.Behavior =WhiteSameNameRefBhv_I
	WhiteSameNameRefBhv_I.BaseBehavior.Behavior=&whiteSameNameRef
	WhiteSameNameRefBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	WhiteXlsManBhv_I = new(WhiteXlsManBhv)
	WhiteXlsManBhv_I.BaseBehavior = new(df.BaseBehavior)
	WhiteXlsManBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	WhiteXlsManBhv_I.BaseBehavior.TableDbName = "WhiteXlsMan"
	var whiteXlsMan df.Behavior =WhiteXlsManBhv_I
	WhiteXlsManBhv_I.BaseBehavior.Behavior=&whiteXlsMan
	WhiteXlsManBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	WithdrawalReasonBhv_I = new(WithdrawalReasonBhv)
	WithdrawalReasonBhv_I.BaseBehavior = new(df.BaseBehavior)
	WithdrawalReasonBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	WithdrawalReasonBhv_I.BaseBehavior.TableDbName = "WithdrawalReason"
	var withdrawalReason df.Behavior =WithdrawalReasonBhv_I
	WithdrawalReasonBhv_I.BaseBehavior.Behavior=&withdrawalReason
	WithdrawalReasonBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
}