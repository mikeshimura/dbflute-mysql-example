package com.mssoftech.dbflute.bsentity;

import java.util.List;
import java.util.ArrayList;

import org.dbflute.Entity;
import org.dbflute.dbmeta.DBMeta;
import org.dbflute.dbmeta.AbstractEntity;
import org.dbflute.dbmeta.accessory.DomainEntity;
import org.dbflute.optional.OptionalEntity;
import com.mssoftech.dbflute.allcommon.DBMetaInstanceHandler;
import com.mssoftech.dbflute.exentity.*;

/**
 * The entity of member_withdrawal as TABLE. <br>
 * 会員退会情報: 退会会員の退会に関する詳細な情報。<br>
 * 退会会員のみデータが存在する。（"1 : 0...1" のパターンの one-to-one）<br>
 * 共通カラムがあってバージョンNOがないパターン。基本的に更新が入ることはないが、登録ユーザを保持したいのと、重要なデータなので更新系の情報も。
 * <pre>
 * [primary-key]
 *     member_id
 * 
 * [column]
 *     member_id, withdrawal_reason_code, withdrawal_reason_input_text, withdrawal_datetime, register_datetime, register_process, register_user, update_datetime, update_process, update_user
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
 *     member, withdrawal_reason
 * 
 * [referrer table]
 *     
 * 
 * [foreign property]
 *     member, withdrawalReason
 * 
 * [referrer property]
 *     
 * 
 * [get/set template]
 * /= = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
 * Integer memberId = entity.getMemberId();
 * String withdrawalReasonCode = entity.getWithdrawalReasonCode();
 * String withdrawalReasonInputText = entity.getWithdrawalReasonInputText();
 * java.time.LocalDateTime withdrawalDatetime = entity.getWithdrawalDatetime();
 * java.time.LocalDateTime registerDatetime = entity.getRegisterDatetime();
 * String registerProcess = entity.getRegisterProcess();
 * String registerUser = entity.getRegisterUser();
 * java.time.LocalDateTime updateDatetime = entity.getUpdateDatetime();
 * String updateProcess = entity.getUpdateProcess();
 * String updateUser = entity.getUpdateUser();
 * entity.setMemberId(memberId);
 * entity.setWithdrawalReasonCode(withdrawalReasonCode);
 * entity.setWithdrawalReasonInputText(withdrawalReasonInputText);
 * entity.setWithdrawalDatetime(withdrawalDatetime);
 * entity.setRegisterDatetime(registerDatetime);
 * entity.setRegisterProcess(registerProcess);
 * entity.setRegisterUser(registerUser);
 * entity.setUpdateDatetime(updateDatetime);
 * entity.setUpdateProcess(updateProcess);
 * entity.setUpdateUser(updateUser);
 * = = = = = = = = = =/
 * </pre>
 * @author DBFlute(AutoGenerator)
 */
public abstract class BsMemberWithdrawal extends AbstractEntity implements DomainEntity {

    // ===================================================================================
    //                                                                          Definition
    //                                                                          ==========
    /** The serial version UID for object serialization. (Default) */
    private static final long serialVersionUID = 1L;

    // ===================================================================================
    //                                                                           Attribute
    //                                                                           =========
    /** member_id: {PK, NotNull, int4(10), FK to member} */
    protected Integer _memberId;

    /** withdrawal_reason_code: {bpchar(3), FK to withdrawal_reason} */
    protected String _withdrawalReasonCode;

    /** withdrawal_reason_input_text: {text(2147483647)} */
    protected String _withdrawalReasonInputText;

    /** withdrawal_datetime: {NotNull, timestamp(26, 3)} */
    protected java.time.LocalDateTime _withdrawalDatetime;

    /** register_datetime: {NotNull, timestamp(26, 3)} */
    protected java.time.LocalDateTime _registerDatetime;

    /** register_process: {NotNull, varchar(200)} */
    protected String _registerProcess;

    /** register_user: {NotNull, varchar(200)} */
    protected String _registerUser;

    /** update_datetime: {NotNull, timestamp(26, 3)} */
    protected java.time.LocalDateTime _updateDatetime;

    /** update_process: {NotNull, varchar(200)} */
    protected String _updateProcess;

    /** update_user: {NotNull, varchar(200)} */
    protected String _updateUser;

    // ===================================================================================
    //                                                                             DB Meta
    //                                                                             =======
    /** {@inheritDoc} */
    public DBMeta asDBMeta() {
        return DBMetaInstanceHandler.findDBMeta(asTableDbName());
    }

    /** {@inheritDoc} */
    public String asTableDbName() {
        return "member_withdrawal";
    }

    // ===================================================================================
    //                                                                        Key Handling
    //                                                                        ============
    /** {@inheritDoc} */
    public boolean hasPrimaryKeyValue() {
        if (_memberId == null) { return false; }
        return true;
    }

    // ===================================================================================
    //                                                                    Foreign Property
    //                                                                    ================
    /** member by my member_id, named 'member'. */
    protected OptionalEntity<Member> _member;

    /**
     * [get] member by my member_id, named 'member'. <br>
     * Optional: alwaysPresent(), ifPresent().orElse(), get(), ...
     * @return The entity of foreign property 'member'. (NotNull, EmptyAllowed: when e.g. null FK column, no setupSelect)
     */
    public OptionalEntity<Member> getMember() {
        if (_member == null) { _member = OptionalEntity.relationEmpty(this, "member"); }
        return _member;
    }

    /**
     * [set] member by my member_id, named 'member'.
     * @param member The entity of foreign property 'member'. (NullAllowed)
     */
    public void setMember(OptionalEntity<Member> member) {
        _member = member;
    }

    /** withdrawal_reason by my withdrawal_reason_code, named 'withdrawalReason'. */
    protected OptionalEntity<WithdrawalReason> _withdrawalReason;

    /**
     * [get] withdrawal_reason by my withdrawal_reason_code, named 'withdrawalReason'. <br>
     * Optional: alwaysPresent(), ifPresent().orElse(), get(), ...
     * @return The entity of foreign property 'withdrawalReason'. (NotNull, EmptyAllowed: when e.g. null FK column, no setupSelect)
     */
    public OptionalEntity<WithdrawalReason> getWithdrawalReason() {
        if (_withdrawalReason == null) { _withdrawalReason = OptionalEntity.relationEmpty(this, "withdrawalReason"); }
        return _withdrawalReason;
    }

    /**
     * [set] withdrawal_reason by my withdrawal_reason_code, named 'withdrawalReason'.
     * @param withdrawalReason The entity of foreign property 'withdrawalReason'. (NullAllowed)
     */
    public void setWithdrawalReason(OptionalEntity<WithdrawalReason> withdrawalReason) {
        _withdrawalReason = withdrawalReason;
    }

    // ===================================================================================
    //                                                                   Referrer Property
    //                                                                   =================
    protected <ELEMENT> List<ELEMENT> newReferrerList() {
        return new ArrayList<ELEMENT>();
    }

    // ===================================================================================
    //                                                                      Basic Override
    //                                                                      ==============
    @Override
    protected boolean doEquals(Object obj) {
        if (obj instanceof BsMemberWithdrawal) {
            BsMemberWithdrawal other = (BsMemberWithdrawal)obj;
            if (!xSV(_memberId, other._memberId)) { return false; }
            return true;
        } else {
            return false;
        }
    }

    @Override
    protected int doHashCode(int initial) {
        int hs = initial;
        hs = xCH(hs, asTableDbName());
        hs = xCH(hs, _memberId);
        return hs;
    }

    @Override
    protected String doBuildStringWithRelation(String li) {
        StringBuilder sb = new StringBuilder();
        if (_member != null && _member.isPresent())
        { sb.append(li).append(xbRDS(_member, "member")); }
        if (_withdrawalReason != null && _withdrawalReason.isPresent())
        { sb.append(li).append(xbRDS(_withdrawalReason, "withdrawalReason")); }
        return sb.toString();
    }
    protected <ET extends Entity> String xbRDS(org.dbflute.optional.OptionalEntity<ET> et, String name) { // buildRelationDisplayString()
        return et.get().buildDisplayString(name, true, true);
    }

    @Override
    protected String doBuildColumnString(String dm) {
        StringBuilder sb = new StringBuilder();
        sb.append(dm).append(xfND(_memberId));
        sb.append(dm).append(xfND(_withdrawalReasonCode));
        sb.append(dm).append(xfND(_withdrawalReasonInputText));
        sb.append(dm).append(xfND(_withdrawalDatetime));
        sb.append(dm).append(xfND(_registerDatetime));
        sb.append(dm).append(xfND(_registerProcess));
        sb.append(dm).append(xfND(_registerUser));
        sb.append(dm).append(xfND(_updateDatetime));
        sb.append(dm).append(xfND(_updateProcess));
        sb.append(dm).append(xfND(_updateUser));
        if (sb.length() > dm.length()) {
            sb.delete(0, dm.length());
        }
        sb.insert(0, "{").append("}");
        return sb.toString();
    }

    @Override
    protected String doBuildRelationString(String dm) {
        StringBuilder sb = new StringBuilder();
        if (_member != null && _member.isPresent())
        { sb.append(dm).append("member"); }
        if (_withdrawalReason != null && _withdrawalReason.isPresent())
        { sb.append(dm).append("withdrawalReason"); }
        if (sb.length() > dm.length()) {
            sb.delete(0, dm.length()).insert(0, "(").append(")");
        }
        return sb.toString();
    }

    @Override
    public MemberWithdrawal clone() {
        return (MemberWithdrawal)super.clone();
    }

    // ===================================================================================
    //                                                                            Accessor
    //                                                                            ========
    /**
     * [get] member_id: {PK, NotNull, int4(10), FK to member} <br>
     * @return The value of the column 'member_id'. (basically NotNull if selected: for the constraint)
     */
    public Integer getMemberId() {
        checkSpecifiedProperty("memberId");
        return _memberId;
    }

    /**
     * [set] member_id: {PK, NotNull, int4(10), FK to member} <br>
     * @param memberId The value of the column 'member_id'. (basically NotNull if update: for the constraint)
     */
    public void setMemberId(Integer memberId) {
        registerModifiedProperty("memberId");
        _memberId = memberId;
    }

    /**
     * [get] withdrawal_reason_code: {bpchar(3), FK to withdrawal_reason} <br>
     * 退会理由コード: 退会した定型理由を参照するコード。<br>
     * 何も言わずに退会する会員もいるので必須項目ではない。
     * @return The value of the column 'withdrawal_reason_code'. (NullAllowed even if selected: for no constraint)
     */
    public String getWithdrawalReasonCode() {
        checkSpecifiedProperty("withdrawalReasonCode");
        return _withdrawalReasonCode;
    }

    /**
     * [set] withdrawal_reason_code: {bpchar(3), FK to withdrawal_reason} <br>
     * 退会理由コード: 退会した定型理由を参照するコード。<br>
     * 何も言わずに退会する会員もいるので必須項目ではない。
     * @param withdrawalReasonCode The value of the column 'withdrawal_reason_code'. (NullAllowed: null update allowed for no constraint)
     */
    public void setWithdrawalReasonCode(String withdrawalReasonCode) {
        registerModifiedProperty("withdrawalReasonCode");
        _withdrawalReasonCode = withdrawalReasonCode;
    }

    /**
     * [get] withdrawal_reason_input_text: {text(2147483647)} <br>
     * 退会理由入力テキスト: 会員がフリーテキストで入力できる退会理由。<br>
     * もう言いたいこと言ってもらう感じ。<br>
     * サイト運営側としては真摯に受け止めて改善していきたい。
     * @return The value of the column 'withdrawal_reason_input_text'. (NullAllowed even if selected: for no constraint)
     */
    public String getWithdrawalReasonInputText() {
        checkSpecifiedProperty("withdrawalReasonInputText");
        return _withdrawalReasonInputText;
    }

    /**
     * [set] withdrawal_reason_input_text: {text(2147483647)} <br>
     * 退会理由入力テキスト: 会員がフリーテキストで入力できる退会理由。<br>
     * もう言いたいこと言ってもらう感じ。<br>
     * サイト運営側としては真摯に受け止めて改善していきたい。
     * @param withdrawalReasonInputText The value of the column 'withdrawal_reason_input_text'. (NullAllowed: null update allowed for no constraint)
     */
    public void setWithdrawalReasonInputText(String withdrawalReasonInputText) {
        registerModifiedProperty("withdrawalReasonInputText");
        _withdrawalReasonInputText = withdrawalReasonInputText;
    }

    /**
     * [get] withdrawal_datetime: {NotNull, timestamp(26, 3)} <br>
     * 退会日時: 退会した瞬間の日時。<br>
     * 正式会員日時と違い、こっちはone-to-oneの別テーブルで。
     * @return The value of the column 'withdrawal_datetime'. (basically NotNull if selected: for the constraint)
     */
    public java.time.LocalDateTime getWithdrawalDatetime() {
        checkSpecifiedProperty("withdrawalDatetime");
        return _withdrawalDatetime;
    }

    /**
     * [set] withdrawal_datetime: {NotNull, timestamp(26, 3)} <br>
     * 退会日時: 退会した瞬間の日時。<br>
     * 正式会員日時と違い、こっちはone-to-oneの別テーブルで。
     * @param withdrawalDatetime The value of the column 'withdrawal_datetime'. (basically NotNull if update: for the constraint)
     */
    public void setWithdrawalDatetime(java.time.LocalDateTime withdrawalDatetime) {
        registerModifiedProperty("withdrawalDatetime");
        _withdrawalDatetime = withdrawalDatetime;
    }

    /**
     * [get] register_datetime: {NotNull, timestamp(26, 3)} <br>
     * @return The value of the column 'register_datetime'. (basically NotNull if selected: for the constraint)
     */
    public java.time.LocalDateTime getRegisterDatetime() {
        checkSpecifiedProperty("registerDatetime");
        return _registerDatetime;
    }

    /**
     * [set] register_datetime: {NotNull, timestamp(26, 3)} <br>
     * @param registerDatetime The value of the column 'register_datetime'. (basically NotNull if update: for the constraint)
     */
    public void setRegisterDatetime(java.time.LocalDateTime registerDatetime) {
        registerModifiedProperty("registerDatetime");
        _registerDatetime = registerDatetime;
    }

    /**
     * [get] register_process: {NotNull, varchar(200)} <br>
     * @return The value of the column 'register_process'. (basically NotNull if selected: for the constraint)
     */
    public String getRegisterProcess() {
        checkSpecifiedProperty("registerProcess");
        return _registerProcess;
    }

    /**
     * [set] register_process: {NotNull, varchar(200)} <br>
     * @param registerProcess The value of the column 'register_process'. (basically NotNull if update: for the constraint)
     */
    public void setRegisterProcess(String registerProcess) {
        registerModifiedProperty("registerProcess");
        _registerProcess = registerProcess;
    }

    /**
     * [get] register_user: {NotNull, varchar(200)} <br>
     * @return The value of the column 'register_user'. (basically NotNull if selected: for the constraint)
     */
    public String getRegisterUser() {
        checkSpecifiedProperty("registerUser");
        return _registerUser;
    }

    /**
     * [set] register_user: {NotNull, varchar(200)} <br>
     * @param registerUser The value of the column 'register_user'. (basically NotNull if update: for the constraint)
     */
    public void setRegisterUser(String registerUser) {
        registerModifiedProperty("registerUser");
        _registerUser = registerUser;
    }

    /**
     * [get] update_datetime: {NotNull, timestamp(26, 3)} <br>
     * @return The value of the column 'update_datetime'. (basically NotNull if selected: for the constraint)
     */
    public java.time.LocalDateTime getUpdateDatetime() {
        checkSpecifiedProperty("updateDatetime");
        return _updateDatetime;
    }

    /**
     * [set] update_datetime: {NotNull, timestamp(26, 3)} <br>
     * @param updateDatetime The value of the column 'update_datetime'. (basically NotNull if update: for the constraint)
     */
    public void setUpdateDatetime(java.time.LocalDateTime updateDatetime) {
        registerModifiedProperty("updateDatetime");
        _updateDatetime = updateDatetime;
    }

    /**
     * [get] update_process: {NotNull, varchar(200)} <br>
     * @return The value of the column 'update_process'. (basically NotNull if selected: for the constraint)
     */
    public String getUpdateProcess() {
        checkSpecifiedProperty("updateProcess");
        return _updateProcess;
    }

    /**
     * [set] update_process: {NotNull, varchar(200)} <br>
     * @param updateProcess The value of the column 'update_process'. (basically NotNull if update: for the constraint)
     */
    public void setUpdateProcess(String updateProcess) {
        registerModifiedProperty("updateProcess");
        _updateProcess = updateProcess;
    }

    /**
     * [get] update_user: {NotNull, varchar(200)} <br>
     * @return The value of the column 'update_user'. (basically NotNull if selected: for the constraint)
     */
    public String getUpdateUser() {
        checkSpecifiedProperty("updateUser");
        return _updateUser;
    }

    /**
     * [set] update_user: {NotNull, varchar(200)} <br>
     * @param updateUser The value of the column 'update_user'. (basically NotNull if update: for the constraint)
     */
    public void setUpdateUser(String updateUser) {
        registerModifiedProperty("updateUser");
        _updateUser = updateUser;
    }
}
