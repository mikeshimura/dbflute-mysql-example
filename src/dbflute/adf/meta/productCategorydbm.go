package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type ProductCategoryDbm_T struct {
	df.BaseDBMeta
	ColumnProductCategoryCode *df.ColumnInfo
	ColumnProductCategoryName *df.ColumnInfo
	ColumnParentCategoryCode *df.ColumnInfo
}

func (b *ProductCategoryDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *ProductCategoryDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var ProductCategoryDbm *ProductCategoryDbm_T

func Create_ProductCategoryDbm() {
	ProductCategoryDbm = new(ProductCategoryDbm_T)
	ProductCategoryDbm.TableDbName = "product_category"
	ProductCategoryDbm.TableDispName = "product_category"
	ProductCategoryDbm.TablePropertyName = "productCategory"
	ProductCategoryDbm.TableSqlName = new(df.TableSqlName)
	ProductCategoryDbm.TableSqlName.TableSqlName = "product_category"
	ProductCategoryDbm.TableSqlName.CorrespondingDbName = ProductCategoryDbm.TableDbName

	var productCategory df.DBMeta
	productCategory = ProductCategoryDbm
	ProductCategoryDbm.DBMeta=&productCategory
	productCategoryCodeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo product_category_code
	productCategoryCodeSqlName.ColumnSqlName = "product_category_code"
	productCategoryCodeSqlName.IrregularChar = false
	ProductCategoryDbm.ColumnProductCategoryCode = df.CCI(&productCategory, "product_category_code", productCategoryCodeSqlName, "", "", "String.class", "productCategoryCode", "", true, false,true, "bpchar", 3, 0, "",false,"","", "","productList,productCategorySelfList","",false,"string")
	productCategoryNameSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo product_category_name
	productCategoryNameSqlName.ColumnSqlName = "product_category_name"
	productCategoryNameSqlName.IrregularChar = false
	ProductCategoryDbm.ColumnProductCategoryName = df.CCI(&productCategory, "product_category_name", productCategoryNameSqlName, "", "", "String.class", "productCategoryName", "", false, false,true, "varchar", 50, 0, "",false,"","", "","","",false,"string")
	parentCategoryCodeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo parent_category_code
	parentCategoryCodeSqlName.ColumnSqlName = "parent_category_code"
	parentCategoryCodeSqlName.IrregularChar = false
	ProductCategoryDbm.ColumnParentCategoryCode = df.CCI(&productCategory, "parent_category_code", parentCategoryCodeSqlName, "", "", "String.class", "parentCategoryCode", "", false, false,false, "bpchar", 3, 0, "",false,"","", "productCategorySelf","","",false,"sql.NullString")

	ProductCategoryDbm.ColumnInfoList = new(df.List)
	ProductCategoryDbm.ColumnInfoList.Add(ProductCategoryDbm.ColumnProductCategoryCode)
	ProductCategoryDbm.ColumnInfoList.Add(ProductCategoryDbm.ColumnProductCategoryName)
	ProductCategoryDbm.ColumnInfoList.Add(ProductCategoryDbm.ColumnParentCategoryCode)


	ProductCategoryDbm.ColumnInfoMap=make(map[string]int)
	ProductCategoryDbm.ColumnInfoMap["productCategoryCode"]=0
		ProductCategoryDbm.ColumnInfoMap["productCategoryName"]=1
		ProductCategoryDbm.ColumnInfoMap["parentCategoryCode"]=2
	    ProductCategoryDbm.PrimaryKey = true
    ProductCategoryDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &productCategory
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(ProductCategoryDbm.ColumnProductCategoryCode)

	ProductCategoryDbm.PrimaryInfo = new(df.PrimaryInfo)
	ProductCategoryDbm.PrimaryInfo.UniqueInfo = ui
	
	var productCategoryMeta df.DBMeta = ProductCategoryDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["ProductCategory"] = &productCategoryMeta
}
