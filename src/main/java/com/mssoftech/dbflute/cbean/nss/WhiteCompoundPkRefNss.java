package com.mssoftech.dbflute.cbean.nss;

import com.mssoftech.dbflute.cbean.cq.WhiteCompoundPkRefCQ;

/**
 * The nest select set-upper of white_compound_pk_ref.
 * @author DBFlute(AutoGenerator)
 */
public class WhiteCompoundPkRefNss {

    // ===================================================================================
    //                                                                           Attribute
    //                                                                           =========
    protected final WhiteCompoundPkRefCQ _query;
    public WhiteCompoundPkRefNss(WhiteCompoundPkRefCQ query) { _query = query; }
    public boolean hasConditionQuery() { return _query != null; }

    // ===================================================================================
    //                                                                     Nested Relation
    //                                                                     ===============
    /**
     * With nested relation columns to select clause. <br>
     * white_compound_pk by my ref_first_id, ref_second_id, named 'whiteCompoundPk'.
     */
    public void withWhiteCompoundPk() {
        _query.xdoNss(() -> _query.queryWhiteCompoundPk());
    }
}
