package com.mssoftech.dbflute.cbean.nss;

import com.mssoftech.dbflute.cbean.cq.ProductCategoryCQ;

/**
 * The nest select set-upper of product_category.
 * @author DBFlute(AutoGenerator)
 */
public class ProductCategoryNss {

    // ===================================================================================
    //                                                                           Attribute
    //                                                                           =========
    protected final ProductCategoryCQ _query;
    public ProductCategoryNss(ProductCategoryCQ query) { _query = query; }
    public boolean hasConditionQuery() { return _query != null; }

    // ===================================================================================
    //                                                                     Nested Relation
    //                                                                     ===============
    /**
     * With nested relation columns to select clause. <br>
     * product_category by my parent_category_code, named 'productCategorySelf'.
     * @return The set-upper of more nested relation. {...with[nested-relation].with[more-nested-relation]} (NotNull)
     */
    public ProductCategoryNss withProductCategorySelf() {
        _query.xdoNss(() -> _query.queryProductCategorySelf());
        return new ProductCategoryNss(_query.queryProductCategorySelf());
    }
}
