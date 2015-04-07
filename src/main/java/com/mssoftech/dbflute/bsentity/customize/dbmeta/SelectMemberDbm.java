package com.mssoftech.dbflute.bsentity.customize.dbmeta;

import java.util.List;
import java.util.Map;

import org.dbflute.Entity;
import org.dbflute.dbmeta.AbstractDBMeta;
import org.dbflute.dbmeta.info.*;
import org.dbflute.dbmeta.name.*;
import org.dbflute.dbmeta.property.PropertyGateway;
import org.dbflute.dbway.DBDef;
import com.mssoftech.dbflute.allcommon.*;
import com.mssoftech.dbflute.exentity.customize.*;

/**
 * The DB meta of SelectMember. (Singleton)
 * @author DBFlute(AutoGenerator)
 */
public class SelectMemberDbm extends AbstractDBMeta {

    // ===================================================================================
    //                                                                           Singleton
    //                                                                           =========
    private static final SelectMemberDbm _instance = new SelectMemberDbm();
    private SelectMemberDbm() {}
    public static SelectMemberDbm getInstance() { return _instance; }

    // ===================================================================================
    //                                                                       Current DBDef
    //                                                                       =============
    public String getProjectName() { return DBCurrent.getInstance().projectName(); }
    public String getProjectPrefix() { return DBCurrent.getInstance().projectPrefix(); }
    public String getGenerationGapBasePrefix() { return DBCurrent.getInstance().generationGapBasePrefix(); }
    public DBDef getCurrentDBDef() { return DBCurrent.getInstance().currentDBDef(); }

    // ===================================================================================
    //                                                                    Property Gateway
    //                                                                    ================
    // -----------------------------------------------------
    //                                       Column Property
    //                                       ---------------
    protected final Map<String, PropertyGateway> _epgMap = newHashMap();
    { xsetupEpg(); }
    protected void xsetupEpg() {
        setupEpg(_epgMap, et -> ((SelectMember)et).getMemberId(), (et, vl) -> ((SelectMember)et).setMemberId(cti(vl)), "memberId");
        setupEpg(_epgMap, et -> ((SelectMember)et).getMemberName(), (et, vl) -> ((SelectMember)et).setMemberName((String)vl), "memberName");
        setupEpg(_epgMap, et -> ((SelectMember)et).getMemberAccount(), (et, vl) -> ((SelectMember)et).setMemberAccount((String)vl), "memberAccount");
        setupEpg(_epgMap, et -> ((SelectMember)et).getBirthdate(), (et, vl) -> ((SelectMember)et).setBirthdate(ctld(vl)), "birthdate");
        setupEpg(_epgMap, et -> ((SelectMember)et).getFormalizedDatetime(), (et, vl) -> ((SelectMember)et).setFormalizedDatetime(ctldt(vl)), "formalizedDatetime");
        setupEpg(_epgMap, et -> ((SelectMember)et).getMemberStatusCode(), (et, vl) -> ((SelectMember)et).setMemberStatusCode((String)vl), "memberStatusCode");
        setupEpg(_epgMap, et -> ((SelectMember)et).getMemberStatusName(), (et, vl) -> ((SelectMember)et).setMemberStatusName((String)vl), "memberStatusName");
        setupEpg(_epgMap, et -> ((SelectMember)et).getDescription(), (et, vl) -> ((SelectMember)et).setDescription((String)vl), "description");
    }
    public PropertyGateway findPropertyGateway(String prop)
    { return doFindEpg(_epgMap, prop); }

    // ===================================================================================
    //                                                                          Table Info
    //                                                                          ==========
    protected final String _tableDbName = "SelectMember";
    protected final String _tableDispName = "SelectMember";
    protected final String _tablePropertyName = "selectMember";
    protected final TableSqlName _tableSqlName = new TableSqlName("SelectMember", _tableDbName);
    { _tableSqlName.xacceptFilter(DBFluteConfig.getInstance().getTableSqlNameFilter()); }
    public String getTableDbName() { return _tableDbName; }
    public String getTableDispName() { return _tableDispName; }
    public String getTablePropertyName() { return _tablePropertyName; }
    public TableSqlName getTableSqlName() { return _tableSqlName; }

    // ===================================================================================
    //                                                                         Column Info
    //                                                                         ===========
    protected final ColumnInfo _columnMemberId = cci("member_id", "member_id", null, null, Integer.class, "memberId", null, false, false, false, "serial", 10, 0, null, false, null, null, null, null, null, false);
//"sql.NullInt64"
    protected final ColumnInfo _columnMemberName = cci("member_name", "member_name", null, null, String.class, "memberName", null, false, false, false, "varchar", 200, 0, null, false, null, null, null, null, null, false);
//"sql.NullString"
    protected final ColumnInfo _columnMemberAccount = cci("member_account", "member_account", null, null, String.class, "memberAccount", null, false, false, false, "varchar", 50, 0, null, false, null, null, null, null, null, false);
//"sql.NullString"
    protected final ColumnInfo _columnBirthdate = cci("birthdate", "birthdate", null, null, java.time.LocalDate.class, "birthdate", null, false, false, false, "date", 13, 0, null, false, null, null, null, null, null, false);
//"df.NullDate"
    protected final ColumnInfo _columnFormalizedDatetime = cci("formalized_datetime", "formalized_datetime", null, null, java.time.LocalDateTime.class, "formalizedDatetime", null, false, false, false, "timestamp", 26, 3, null, false, null, null, null, null, null, false);
//"df.NullTimestamp"
    protected final ColumnInfo _columnMemberStatusCode = cci("member_status_code", "member_status_code", null, null, String.class, "memberStatusCode", null, false, false, false, "bpchar", 3, 0, null, false, null, null, null, null, null, false);
//"sql.NullString"
    protected final ColumnInfo _columnMemberStatusName = cci("member_status_name", "member_status_name", null, null, String.class, "memberStatusName", null, false, false, false, "varchar", 50, 0, null, false, null, null, null, null, null, false);
//"sql.NullString"
    protected final ColumnInfo _columnDescription = cci("description", "description", null, null, String.class, "description", null, false, false, false, "varchar", 200, 0, null, false, null, null, null, null, null, false);
//"sql.NullString"

    /**
     * member_id: {serial(10), refers to member.member_id}
     * @return The information object of specified column. (NotNull)
     */
    public ColumnInfo columnMemberId() { return _columnMemberId; }
    /**
     * member_name: {varchar(200), refers to member.member_name}
     * @return The information object of specified column. (NotNull)
     */
    public ColumnInfo columnMemberName() { return _columnMemberName; }
    /**
     * member_account: {varchar(50), refers to member.member_account}
     * @return The information object of specified column. (NotNull)
     */
    public ColumnInfo columnMemberAccount() { return _columnMemberAccount; }
    /**
     * birthdate: {date(13), refers to member.birthdate}
     * @return The information object of specified column. (NotNull)
     */
    public ColumnInfo columnBirthdate() { return _columnBirthdate; }
    /**
     * formalized_datetime: {timestamp(26, 3), refers to member.formalized_datetime}
     * @return The information object of specified column. (NotNull)
     */
    public ColumnInfo columnFormalizedDatetime() { return _columnFormalizedDatetime; }
    /**
     * member_status_code: {bpchar(3), refers to member.member_status_code}
     * @return The information object of specified column. (NotNull)
     */
    public ColumnInfo columnMemberStatusCode() { return _columnMemberStatusCode; }
    /**
     * member_status_name: {varchar(50), refers to member_status.member_status_name}
     * @return The information object of specified column. (NotNull)
     */
    public ColumnInfo columnMemberStatusName() { return _columnMemberStatusName; }
    /**
     * description: {varchar(200), refers to member_status.description}
     * @return The information object of specified column. (NotNull)
     */
    public ColumnInfo columnDescription() { return _columnDescription; }

    protected List<ColumnInfo> ccil() {
        List<ColumnInfo> ls = newArrayList();
        ls.add(columnMemberId());
        ls.add(columnMemberName());
        ls.add(columnMemberAccount());
        ls.add(columnBirthdate());
        ls.add(columnFormalizedDatetime());
        ls.add(columnMemberStatusCode());
        ls.add(columnMemberStatusName());
        ls.add(columnDescription());
        return ls;
    }

    { initializeInformationResource(); }

    // ===================================================================================
    //                                                                         Unique Info
    //                                                                         ===========
    // -----------------------------------------------------
    //                                       Primary Element
    //                                       ---------------
    protected UniqueInfo cpui() {
        throw new UnsupportedOperationException("The table does not have primary key: " + getTableDbName());
    }
    public boolean hasPrimaryKey() { return false; }
    public boolean hasCompoundPrimaryKey() { return false; }

    // ===================================================================================
    //                                                                       Relation Info
    //                                                                       =============
    // cannot cache because it uses related DB meta instance while booting
    // (instead, cached by super's collection)
    // -----------------------------------------------------
    //                                      Foreign Property
    //                                      ----------------

    // -----------------------------------------------------
    //                                     Referrer Property
    //                                     -----------------

    // ===================================================================================
    //                                                                        Various Info
    //                                                                        ============

    // ===================================================================================
    //                                                                           Type Name
    //                                                                           =========
    public String getEntityTypeName() { return "com.mssoftech.dbflute.exentity.customize.SelectMember"; }
    public String getConditionBeanTypeName() { return null; }
    public String getBehaviorTypeName() { return null; }

    // ===================================================================================
    //                                                                         Object Type
    //                                                                         ===========
    public Class<SelectMember> getEntityType() { return SelectMember.class; }

    // ===================================================================================
    //                                                                     Object Instance
    //                                                                     ===============
    public SelectMember newEntity() { return new SelectMember(); }

    // ===================================================================================
    //                                                                   Map Communication
    //                                                                   =================
    public void acceptPrimaryKeyMap(Entity et, Map<String, ? extends Object> mp)
    { doAcceptPrimaryKeyMap((SelectMember)et, mp); }
    public void acceptAllColumnMap(Entity et, Map<String, ? extends Object> mp)
    { doAcceptAllColumnMap((SelectMember)et, mp); }
    public Map<String, Object> extractPrimaryKeyMap(Entity et) { return doExtractPrimaryKeyMap(et); }
    public Map<String, Object> extractAllColumnMap(Entity et) { return doExtractAllColumnMap(et); }
}
