package com.mssoftech.dbflute.bsentity.dbmeta;

import java.util.List;
import java.util.Map;

import org.dbflute.Entity;
import org.dbflute.dbmeta.AbstractDBMeta;
import org.dbflute.dbmeta.info.*;
import org.dbflute.dbmeta.name.*;
import org.dbflute.dbmeta.property.PropertyGateway;
import org.dbflute.dbway.DBDef;
import com.mssoftech.dbflute.allcommon.*;
import com.mssoftech.dbflute.exentity.*;

/**
 * The DB meta of white_same_name. (Singleton)
 * @author DBFlute(AutoGenerator)
 */
public class WhiteSameNameDbm extends AbstractDBMeta {

    // ===================================================================================
    //                                                                           Singleton
    //                                                                           =========
    private static final WhiteSameNameDbm _instance = new WhiteSameNameDbm();
    private WhiteSameNameDbm() {}
    public static WhiteSameNameDbm getInstance() { return _instance; }

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
        setupEpg(_epgMap, et -> ((WhiteSameName)et).getSameNameId(), (et, vl) -> ((WhiteSameName)et).setSameNameId(ctl(vl)), "sameNameId");
        setupEpg(_epgMap, et -> ((WhiteSameName)et).getSameNameName(), (et, vl) -> ((WhiteSameName)et).setSameNameName((String)vl), "sameNameName");
        setupEpg(_epgMap, et -> ((WhiteSameName)et).getSameNameInteger(), (et, vl) -> ((WhiteSameName)et).setSameNameInteger(cti(vl)), "sameNameInteger");
        setupEpg(_epgMap, et -> ((WhiteSameName)et).getNextSchemaProductId(), (et, vl) -> ((WhiteSameName)et).setNextSchemaProductId(cti(vl)), "nextSchemaProductId");
    }
    public PropertyGateway findPropertyGateway(String prop)
    { return doFindEpg(_epgMap, prop); }

    // ===================================================================================
    //                                                                          Table Info
    //                                                                          ==========
    protected final String _tableDbName = "white_same_name";
    protected final String _tableDispName = "white_same_name";
    protected final String _tablePropertyName = "whiteSameName";
    protected final TableSqlName _tableSqlName = new TableSqlName("white_same_name", _tableDbName);
    { _tableSqlName.xacceptFilter(DBFluteConfig.getInstance().getTableSqlNameFilter()); }
    public String getTableDbName() { return _tableDbName; }
    public String getTableDispName() { return _tableDispName; }
    public String getTablePropertyName() { return _tablePropertyName; }
    public TableSqlName getTableSqlName() { return _tableSqlName; }

    // ===================================================================================
    //                                                                         Column Info
    //                                                                         ===========
    protected final ColumnInfo _columnSameNameId = cci("same_name_id", "same_name_id", null, null, Long.class, "sameNameId", null, true, false, true, "int8", 19, 0, null, false, null, null, null, "whiteSameNameRefList", null, false);
//"int64"
    protected final ColumnInfo _columnSameNameName = cci("same_name_name", "same_name_name", null, null, String.class, "sameNameName", null, false, false, false, "varchar", 200, 0, null, false, null, null, null, null, null, false);
//"sql.NullString"
    protected final ColumnInfo _columnSameNameInteger = cci("same_name_integer", "same_name_integer", null, null, Integer.class, "sameNameInteger", null, false, false, false, "int4", 10, 0, null, false, null, null, null, null, null, false);
//"sql.NullInt64"
    protected final ColumnInfo _columnNextSchemaProductId = cci("next_schema_product_id", "next_schema_product_id", null, null, Integer.class, "nextSchemaProductId", null, false, false, false, "int4", 10, 0, null, false, null, null, null, null, null, false);
//"sql.NullInt64"

    /**
     * same_name_id: {PK, NotNull, int8(19)}
     * @return The information object of specified column. (NotNull)
     */
    public ColumnInfo columnSameNameId() { return _columnSameNameId; }
    /**
     * same_name_name: {varchar(200)}
     * @return The information object of specified column. (NotNull)
     */
    public ColumnInfo columnSameNameName() { return _columnSameNameName; }
    /**
     * same_name_integer: {int4(10)}
     * @return The information object of specified column. (NotNull)
     */
    public ColumnInfo columnSameNameInteger() { return _columnSameNameInteger; }
    /**
     * next_schema_product_id: {int4(10)}
     * @return The information object of specified column. (NotNull)
     */
    public ColumnInfo columnNextSchemaProductId() { return _columnNextSchemaProductId; }

    protected List<ColumnInfo> ccil() {
        List<ColumnInfo> ls = newArrayList();
        ls.add(columnSameNameId());
        ls.add(columnSameNameName());
        ls.add(columnSameNameInteger());
        ls.add(columnNextSchemaProductId());
        return ls;
    }

    { initializeInformationResource(); }

    // ===================================================================================
    //                                                                         Unique Info
    //                                                                         ===========
    // -----------------------------------------------------
    //                                       Primary Element
    //                                       ---------------
    protected UniqueInfo cpui() { return hpcpui(columnSameNameId()); }
    public boolean hasPrimaryKey() { return true; }
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
    /**
     * white_same_name_ref by same_name_id, named 'whiteSameNameRefList'.
     * @return The information object of referrer property. (NotNull)
     */
    public ReferrerInfo referrerWhiteSameNameRefList() {
        Map<ColumnInfo, ColumnInfo> mp = newLinkedHashMap(columnSameNameId(), WhiteSameNameRefDbm.getInstance().columnSameNameId());
        return cri("fk_white_same_name_ref", "whiteSameNameRefList", this, WhiteSameNameRefDbm.getInstance(), mp, false, "whiteSameName");
    }

    // ===================================================================================
    //                                                                        Various Info
    //                                                                        ============

    // ===================================================================================
    //                                                                           Type Name
    //                                                                           =========
    public String getEntityTypeName() { return "com.mssoftech.dbflute.exentity.WhiteSameName"; }
    public String getConditionBeanTypeName() { return "com.mssoftech.dbflute.cbean.WhiteSameNameCB"; }
    public String getBehaviorTypeName() { return "com.mssoftech.dbflute.exbhv.WhiteSameNameBhv"; }

    // ===================================================================================
    //                                                                         Object Type
    //                                                                         ===========
    public Class<WhiteSameName> getEntityType() { return WhiteSameName.class; }

    // ===================================================================================
    //                                                                     Object Instance
    //                                                                     ===============
    public WhiteSameName newEntity() { return new WhiteSameName(); }

    // ===================================================================================
    //                                                                   Map Communication
    //                                                                   =================
    public void acceptPrimaryKeyMap(Entity et, Map<String, ? extends Object> mp)
    { doAcceptPrimaryKeyMap((WhiteSameName)et, mp); }
    public void acceptAllColumnMap(Entity et, Map<String, ? extends Object> mp)
    { doAcceptAllColumnMap((WhiteSameName)et, mp); }
    public Map<String, Object> extractPrimaryKeyMap(Entity et) { return doExtractPrimaryKeyMap(et); }
    public Map<String, Object> extractAllColumnMap(Entity et) { return doExtractAllColumnMap(et); }
}
