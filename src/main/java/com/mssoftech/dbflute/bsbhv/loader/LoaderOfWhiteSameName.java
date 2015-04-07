package com.mssoftech.dbflute.bsbhv.loader;

import java.util.List;

import org.dbflute.bhv.*;
import org.dbflute.bhv.referrer.*;
import com.mssoftech.dbflute.exbhv.*;
import com.mssoftech.dbflute.exentity.*;
import com.mssoftech.dbflute.cbean.*;

/**
 * The referrer loader of white_same_name as TABLE. <br>
 * <pre>
 * [primary key]
 *     same_name_id
 *
 * [column]
 *     same_name_id, same_name_name, same_name_integer, next_schema_product_id
 *
 * [sequence]
 *     
 *
 * [identity]
 *     
 *
 * [version-no]
 *     
 *
 * [foreign table]
 *     
 *
 * [referrer table]
 *     white_same_name_ref
 *
 * [foreign property]
 *     
 *
 * [referrer property]
 *     whiteSameNameRefList
 * </pre>
 * @author DBFlute(AutoGenerator)
 */
public class LoaderOfWhiteSameName {

    // ===================================================================================
    //                                                                           Attribute
    //                                                                           =========
    protected List<WhiteSameName> _selectedList;
    protected BehaviorSelector _selector;
    protected WhiteSameNameBhv _myBhv; // lazy-loaded

    // ===================================================================================
    //                                                                   Ready for Loading
    //                                                                   =================
    public LoaderOfWhiteSameName ready(List<WhiteSameName> selectedList, BehaviorSelector selector)
    { _selectedList = selectedList; _selector = selector; return this; }

    protected WhiteSameNameBhv myBhv()
    { if (_myBhv != null) { return _myBhv; } else { _myBhv = _selector.select(WhiteSameNameBhv.class); return _myBhv; } }

    // ===================================================================================
    //                                                                       Load Referrer
    //                                                                       =============
    protected List<WhiteSameNameRef> _referrerWhiteSameNameRef;

    /**
     * Load referrer of whiteSameNameRefList by the set-upper of referrer. <br>
     * white_same_name_ref by same_name_id, named 'whiteSameNameRefList'.
     * <pre>
     * <span style="color: #0000C0">whiteSameNameBhv</span>.<span style="color: #994747">load</span>(<span style="color: #553000">whiteSameNameList</span>, <span style="color: #553000">nameLoader</span> <span style="color: #90226C; font-weight: bold"><span style="font-size: 120%">-</span>&gt;</span> {
     *     <span style="color: #553000">nameLoader</span>.<span style="color: #CC4747">loadWhiteSameNameRef</span>(<span style="color: #553000">refCB</span> <span style="color: #90226C; font-weight: bold"><span style="font-size: 120%">-</span>&gt;</span> {
     *         <span style="color: #553000">refCB</span>.setupSelect...
     *         <span style="color: #553000">refCB</span>.query().set...
     *         <span style="color: #553000">refCB</span>.query().addOrderBy...
     *     }); <span style="color: #3F7E5E">// you can load nested referrer from here</span>
     *     <span style="color: #3F7E5E">//}).withNestedReferrer(<span style="color: #553000">refLoader</span> -&gt; {</span>
     *     <span style="color: #3F7E5E">//    refLoader.load...</span>
     *     <span style="color: #3F7E5E">//});</span>
     * });
     * for (WhiteSameName whiteSameName : <span style="color: #553000">whiteSameNameList</span>) {
     *     ... = whiteSameName.<span style="color: #CC4747">getWhiteSameNameRefList()</span>;
     * }
     * </pre>
     * About internal policy, the value of primary key (and others too) is treated as case-insensitive. <br>
     * The condition-bean, which the set-upper provides, has settings before callback as follows:
     * <pre>
     * cb.query().setSameNameId_InScope(pkList);
     * cb.query().addOrderBy_SameNameId_Asc();
     * </pre>
     * @param refCBLambda The callback to set up referrer condition-bean for loading referrer. (NotNull)
     * @return The callback interface which you can load nested referrer by calling withNestedReferrer(). (NotNull)
     */
    public NestedReferrerLoaderGateway<LoaderOfWhiteSameNameRef> loadWhiteSameNameRef(ReferrerConditionSetupper<WhiteSameNameRefCB> refCBLambda) {
        myBhv().loadWhiteSameNameRef(_selectedList, refCBLambda).withNestedReferrer(refLs -> _referrerWhiteSameNameRef = refLs);
        return hd -> hd.handle(new LoaderOfWhiteSameNameRef().ready(_referrerWhiteSameNameRef, _selector));
    }

    // ===================================================================================
    //                                                                    Pull out Foreign
    //                                                                    ================
    // ===================================================================================
    //                                                                            Accessor
    //                                                                            ========
    public List<WhiteSameName> getSelectedList() { return _selectedList; }
    public BehaviorSelector getSelector() { return _selector; }
}
