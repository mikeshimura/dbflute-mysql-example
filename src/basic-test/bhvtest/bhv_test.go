package bhvtest

import (
	"database/sql"
	"dbflute/adf/bhv"
	"dbflute/adf/cb"
	"dbflute/adf/centity"
	"dbflute/adf/entity"
	"dbflute/adf/pmb"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mikeshimura/dbflute/df"
	"github.com/mikeshimura/dbflute/log"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	var Db *sql.DB
	var err error
	Db, err = sql.Open("mysql", "exampledb:exampledb@/exampledb")
	if err != nil {
		log.ErrorConv("factory", err.Error())

		return
	}
	tx, _ := Db.Begin()
	cb1 := cb.CreateMemberCB()
	cb1.Query.SetMemberId_Equal(3)
	l, err1 := bhv.MemberBhv_I.SelectList(cb1, tx)
	if err1!=nil{
		log.ErrorConv("test",err1.Error())
				tx.Rollback()
		return
	}
	fmt.Println("result no:", l.AllRecordCount)
	tt := (l.List.Get(0)).(*entity.Member)
	fmt.Printf("MemberName %s:\n", tt.GetMemberName())
	tx.Commit()
}

func TestSelect2(t *testing.T) {
	var Db *sql.DB
	var err error
	Db, err = sql.Open("mysql", "exampledb:exampledb@/exampledb")
	if err != nil {
		log.ErrorConv("factory", err.Error())
		return
	}
	tx, _ := Db.Begin()
	cb1 := cb.CreateMemberCB()
	cb1.Query.SetMemberName_PrefixSearch("S")
	cb1.Query.AddOrderBy_MemberId_Asc()
	l, err1 := bhv.MemberBhv_I.SelectList(cb1, tx)
	if err1!=nil{
		log.ErrorConv("test",err1.Error())
				tx.Rollback()
		return
	}
	fmt.Println("result no:", l.AllRecordCount)
	tt := (l.List.Get(0)).(*entity.Member)
	fmt.Printf("MemberName %s:\n", tt.GetMemberName())
	tx.Commit()
}

func TestInsert(t *testing.T) {
	var Db *sql.DB
	var err error
	Db, err = sql.Open("mysql", "exampledb:exampledb@/exampledb")
	if err != nil {
		log.ErrorConv("factory", err.Error())
	}
	tx, _ := Db.Begin()
	member := entity.CreateMember()
	member.SetMemberName("testName")
	member.SetMemberAccount("testAccount")
	member.SetMemberStatusCode("FML")
	member.SetRegisterDatetime(df.CreateMysqlTimestamp(time.Now()))
	member.SetUpdateDatetime(df.CreateMysqlTimestamp(time.Now()))
	member.SetRegisterUser("TEST")
	member.SetUpdateUser("TEST")
	_, err1 := bhv.MemberBhv_I.Insert(member, tx)
	if err1!=nil{
		log.ErrorConv("test",err1.Error())
				tx.Rollback()
		return
	}
	fmt.Printf("MemberId %d:\n", member.GetMemberId())
	cb1 := cb.CreateMemberCB()
	cb1.Query.SetMemberName_Equal("testName")
	l, err2 := bhv.MemberBhv_I.SelectList(cb1, tx)
	if err2!=nil{
		log.ErrorConv("test",err2.Error())
				tx.Rollback()
		return
	}
	fmt.Println("result no:", l.AllRecordCount)
	tt := (l.List.Get(0)).(*entity.Member)
	fmt.Printf("MemberId %d:\n", tt.GetMemberId())
	df.TxRollback(tx)
}
func TestOutsideSelect(t *testing.T) {
	var Db *sql.DB
	var err error
	Db, err = sql.Open("mysql", "exampledb:exampledb@/exampledb")
	if err != nil {
		log.ErrorConv("factory", err.Error())
		return
	}
	tx, _ := Db.Begin()
	pmb := new(pmb.C_SelectMemberPmb)
	pmb.SetName(*new(df.NullString))
	l, err1 := pmb.SelectList(tx)
	if err1!=nil{
		log.ErrorConv("test",err1.Error())
		tx.Rollback()
		return
	}
	fmt.Println("result no:", l.AllRecordCount)
	tt := (l.List.Get(0)).(*centity.C_SelectMember)
	fmt.Printf("MemberName %s:\n", tt.GetMemberName().String)
	tx.Commit()
}

func TestOutsideUpdate(t *testing.T) {
	log.InternalDebugFlag=true
	var Db *sql.DB
	var err error
	Db, err = sql.Open("mysql", "exampledb:exampledb@/exampledb")
	if err != nil {
		log.ErrorConv("factory", err.Error())
		return
	}
	tx, _ := Db.Begin()
	pmb := new(pmb.C_UpdateMemberPmb)
	pmb.SetName(df.CreateNullString("Mihajlovic"))
	l, err1 := pmb.Execute(tx)
	if err1!=nil{
		log.ErrorConv("test",err1.Error())
		tx.Rollback()
		return
	}
	fmt.Println("result no:", l)
	tx.Commit()
}