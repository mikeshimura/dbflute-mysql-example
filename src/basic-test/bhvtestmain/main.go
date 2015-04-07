package main

import (
	"database/sql"
	"dbflute/adf/centity"
	"dbflute/adf/pmb"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/mikeshimura/dbflute/df"
	"github.com/mikeshimura/dbflute/log"
)

func main() {
	log.InternalDebugFlag= false
	var Db *sql.DB
	var err error
	Db, err = sql.Open("postgres", "user=exampledb password=exampledb dbname=exampledb sslmode=disable")

	if err != nil {
		log.ErrorConv("factory", err.Error())
	}
	tx, _ := Db.Begin()
	pmb := new(pmb.C_SelectMemberPmb)
	pmb.SetName(df.CreateNullString("S%"))
	l, err1 := pmb.SelectList(tx)
	if err1!=nil{
		log.ErrorConv("test",err1.Error())
		tx.Rollback()
		return
	}
	fmt.Println("result no:", l.AllRecordCount)
	tt := (l.List.Get(0)).(*centity.C_SelectMember)
	fmt.Printf("MemberName %s:\n", tt.GetMemberName())
	tx.Commit()
}