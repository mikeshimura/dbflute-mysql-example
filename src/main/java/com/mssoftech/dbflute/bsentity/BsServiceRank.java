package com.mssoftech.dbflute.bsentity;

import java.util.List;
import java.util.ArrayList;

import org.dbflute.dbmeta.DBMeta;
import org.dbflute.dbmeta.AbstractEntity;
import org.dbflute.dbmeta.accessory.DomainEntity;
import com.mssoftech.dbflute.allcommon.DBMetaInstanceHandler;
import com.mssoftech.dbflute.exentity.*;

/**
 * The entity of service_rank as TABLE. <br>
 * サービスランク: 会員のサービスレベルを表現するランク。<br>
 * （ゴールドとかプラチナとか）
 * <pre>
 * [primary-key]
 *     service_rank_code
 * 
 * [column]
 *     service_rank_code, service_rank_name, service_point_incidence, new_acceptable_flg, description, display_order
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
 *     member_service
 * 
 * [foreign property]
 *     
 * 
 * [referrer property]
 *     memberServiceList
 * 
 * [get/set template]
 * /= = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
 * String serviceRankCode = entity.getServiceRankCode();
 * String serviceRankName = entity.getServiceRankName();
 * java.math.BigDecimal servicePointIncidence = entity.getServicePointIncidence();
 * Integer newAcceptableFlg = entity.getNewAcceptableFlg();
 * String description = entity.getDescription();
 * Integer displayOrder = entity.getDisplayOrder();
 * entity.setServiceRankCode(serviceRankCode);
 * entity.setServiceRankName(serviceRankName);
 * entity.setServicePointIncidence(servicePointIncidence);
 * entity.setNewAcceptableFlg(newAcceptableFlg);
 * entity.setDescription(description);
 * entity.setDisplayOrder(displayOrder);
 * = = = = = = = = = =/
 * </pre>
 * @author DBFlute(AutoGenerator)
 */
public abstract class BsServiceRank extends AbstractEntity implements DomainEntity {

    // ===================================================================================
    //                                                                          Definition
    //                                                                          ==========
    /** The serial version UID for object serialization. (Default) */
    private static final long serialVersionUID = 1L;

    // ===================================================================================
    //                                                                           Attribute
    //                                                                           =========
    /** service_rank_code: {PK, NotNull, bpchar(3)} */
    protected String _serviceRankCode;

    /** service_rank_name: {NotNull, varchar(50)} */
    protected String _serviceRankName;

    /** service_point_incidence: {NotNull, numeric(5, 3)} */
    protected java.math.BigDecimal _servicePointIncidence;

    /** new_acceptable_flg: {NotNull, int4(10)} */
    protected Integer _newAcceptableFlg;

    /** description: {NotNull, varchar(200)} */
    protected String _description;

    /** display_order: {UQ, NotNull, int4(10)} */
    protected Integer _displayOrder;

    // ===================================================================================
    //                                                                             DB Meta
    //                                                                             =======
    /** {@inheritDoc} */
    public DBMeta asDBMeta() {
        return DBMetaInstanceHandler.findDBMeta(asTableDbName());
    }

    /** {@inheritDoc} */
    public String asTableDbName() {
        return "service_rank";
    }

    // ===================================================================================
    //                                                                        Key Handling
    //                                                                        ============
    /** {@inheritDoc} */
    public boolean hasPrimaryKeyValue() {
        if (_serviceRankCode == null) { return false; }
        return true;
    }

    /**
     * To be unique by the unique column. <br>
     * You can update the entity by the key when entity update (NOT batch update).
     * @param displayOrder : UQ, NotNull, int4(10). (NotNull)
     */
    public void uniqueBy(Integer displayOrder) {
        __uniqueDrivenProperties.clear();
        __uniqueDrivenProperties.addPropertyName("displayOrder");
        setDisplayOrder(displayOrder);
    }

    // ===================================================================================
    //                                                                    Foreign Property
    //                                                                    ================
    // ===================================================================================
    //                                                                   Referrer Property
    //                                                                   =================
    /** member_service by service_rank_code, named 'memberServiceList'. */
    protected List<MemberService> _memberServiceList;

    /**
     * [get] member_service by service_rank_code, named 'memberServiceList'.
     * @return The entity list of referrer property 'memberServiceList'. (NotNull: even if no loading, returns empty list)
     */
    public List<MemberService> getMemberServiceList() {
        if (_memberServiceList == null) { _memberServiceList = newReferrerList(); }
        return _memberServiceList;
    }

    /**
     * [set] member_service by service_rank_code, named 'memberServiceList'.
     * @param memberServiceList The entity list of referrer property 'memberServiceList'. (NullAllowed)
     */
    public void setMemberServiceList(List<MemberService> memberServiceList) {
        _memberServiceList = memberServiceList;
    }

    protected <ELEMENT> List<ELEMENT> newReferrerList() {
        return new ArrayList<ELEMENT>();
    }

    // ===================================================================================
    //                                                                      Basic Override
    //                                                                      ==============
    @Override
    protected boolean doEquals(Object obj) {
        if (obj instanceof BsServiceRank) {
            BsServiceRank other = (BsServiceRank)obj;
            if (!xSV(_serviceRankCode, other._serviceRankCode)) { return false; }
            return true;
        } else {
            return false;
        }
    }

    @Override
    protected int doHashCode(int initial) {
        int hs = initial;
        hs = xCH(hs, asTableDbName());
        hs = xCH(hs, _serviceRankCode);
        return hs;
    }

    @Override
    protected String doBuildStringWithRelation(String li) {
        StringBuilder sb = new StringBuilder();
        if (_memberServiceList != null) { for (MemberService et : _memberServiceList)
        { if (et != null) { sb.append(li).append(xbRDS(et, "memberServiceList")); } } }
        return sb.toString();
    }

    @Override
    protected String doBuildColumnString(String dm) {
        StringBuilder sb = new StringBuilder();
        sb.append(dm).append(xfND(_serviceRankCode));
        sb.append(dm).append(xfND(_serviceRankName));
        sb.append(dm).append(xfND(_servicePointIncidence));
        sb.append(dm).append(xfND(_newAcceptableFlg));
        sb.append(dm).append(xfND(_description));
        sb.append(dm).append(xfND(_displayOrder));
        if (sb.length() > dm.length()) {
            sb.delete(0, dm.length());
        }
        sb.insert(0, "{").append("}");
        return sb.toString();
    }

    @Override
    protected String doBuildRelationString(String dm) {
        StringBuilder sb = new StringBuilder();
        if (_memberServiceList != null && !_memberServiceList.isEmpty())
        { sb.append(dm).append("memberServiceList"); }
        if (sb.length() > dm.length()) {
            sb.delete(0, dm.length()).insert(0, "(").append(")");
        }
        return sb.toString();
    }

    @Override
    public ServiceRank clone() {
        return (ServiceRank)super.clone();
    }

    // ===================================================================================
    //                                                                            Accessor
    //                                                                            ========
    /**
     * [get] service_rank_code: {PK, NotNull, bpchar(3)} <br>
     * サービスランクコード: サービスランクを識別するコード。
     * @return The value of the column 'service_rank_code'. (basically NotNull if selected: for the constraint)
     */
    public String getServiceRankCode() {
        checkSpecifiedProperty("serviceRankCode");
        return _serviceRankCode;
    }

    /**
     * [set] service_rank_code: {PK, NotNull, bpchar(3)} <br>
     * サービスランクコード: サービスランクを識別するコード。
     * @param serviceRankCode The value of the column 'service_rank_code'. (basically NotNull if update: for the constraint)
     */
    public void setServiceRankCode(String serviceRankCode) {
        registerModifiedProperty("serviceRankCode");
        _serviceRankCode = serviceRankCode;
    }

    /**
     * [get] service_rank_name: {NotNull, varchar(50)} <br>
     * サービスランク名称: サービスランクの名称。<br>
     * （ゴールドとかプラチナとか基本的には威厳のある名前）
     * @return The value of the column 'service_rank_name'. (basically NotNull if selected: for the constraint)
     */
    public String getServiceRankName() {
        checkSpecifiedProperty("serviceRankName");
        return _serviceRankName;
    }

    /**
     * [set] service_rank_name: {NotNull, varchar(50)} <br>
     * サービスランク名称: サービスランクの名称。<br>
     * （ゴールドとかプラチナとか基本的には威厳のある名前）
     * @param serviceRankName The value of the column 'service_rank_name'. (basically NotNull if update: for the constraint)
     */
    public void setServiceRankName(String serviceRankName) {
        registerModifiedProperty("serviceRankName");
        _serviceRankName = serviceRankName;
    }

    /**
     * [get] service_point_incidence: {NotNull, numeric(5, 3)} <br>
     * サービスポイント発生率: 購入ごとのサービスポイントの発生率。<br>
     * 購入価格にこの値をかけた数が発生ポイント。<br>
     * ExampleDBとして数少ない貴重な小数点ありのカラム。
     * @return The value of the column 'service_point_incidence'. (basically NotNull if selected: for the constraint)
     */
    public java.math.BigDecimal getServicePointIncidence() {
        checkSpecifiedProperty("servicePointIncidence");
        return _servicePointIncidence;
    }

    /**
     * [set] service_point_incidence: {NotNull, numeric(5, 3)} <br>
     * サービスポイント発生率: 購入ごとのサービスポイントの発生率。<br>
     * 購入価格にこの値をかけた数が発生ポイント。<br>
     * ExampleDBとして数少ない貴重な小数点ありのカラム。
     * @param servicePointIncidence The value of the column 'service_point_incidence'. (basically NotNull if update: for the constraint)
     */
    public void setServicePointIncidence(java.math.BigDecimal servicePointIncidence) {
        registerModifiedProperty("servicePointIncidence");
        _servicePointIncidence = servicePointIncidence;
    }

    /**
     * [get] new_acceptable_flg: {NotNull, int4(10)} <br>
     * 新規受け入れ可能フラグ: このランクへの新規受け入れができるかどうか。
     * @return The value of the column 'new_acceptable_flg'. (basically NotNull if selected: for the constraint)
     */
    public Integer getNewAcceptableFlg() {
        checkSpecifiedProperty("newAcceptableFlg");
        return _newAcceptableFlg;
    }

    /**
     * [set] new_acceptable_flg: {NotNull, int4(10)} <br>
     * 新規受け入れ可能フラグ: このランクへの新規受け入れができるかどうか。
     * @param newAcceptableFlg The value of the column 'new_acceptable_flg'. (basically NotNull if update: for the constraint)
     */
    public void setNewAcceptableFlg(Integer newAcceptableFlg) {
        registerModifiedProperty("newAcceptableFlg");
        _newAcceptableFlg = newAcceptableFlg;
    }

    /**
     * [get] description: {NotNull, varchar(200)} <br>
     * 説明: ランクに関する業務的な説明。
     * @return The value of the column 'description'. (basically NotNull if selected: for the constraint)
     */
    public String getDescription() {
        checkSpecifiedProperty("description");
        return _description;
    }

    /**
     * [set] description: {NotNull, varchar(200)} <br>
     * 説明: ランクに関する業務的な説明。
     * @param description The value of the column 'description'. (basically NotNull if update: for the constraint)
     */
    public void setDescription(String description) {
        registerModifiedProperty("description");
        _description = description;
    }

    /**
     * [get] display_order: {UQ, NotNull, int4(10)} <br>
     * 表示順: UI上の表示順を表現する番号。
     * @return The value of the column 'display_order'. (basically NotNull if selected: for the constraint)
     */
    public Integer getDisplayOrder() {
        checkSpecifiedProperty("displayOrder");
        return _displayOrder;
    }

    /**
     * [set] display_order: {UQ, NotNull, int4(10)} <br>
     * 表示順: UI上の表示順を表現する番号。
     * @param displayOrder The value of the column 'display_order'. (basically NotNull if update: for the constraint)
     */
    public void setDisplayOrder(Integer displayOrder) {
        registerModifiedProperty("displayOrder");
        _displayOrder = displayOrder;
    }
}
