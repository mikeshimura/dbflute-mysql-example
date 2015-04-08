package main

import (
	"database/sql"
	"dbflute/adf/pmb"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mikeshimura/dbflute/df"
	"github.com/mikeshimura/dbflute/log"
)

func main() {
	log.InternalDebugFlag = true
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
	if err1 != nil {
		log.ErrorConv("test", err1.Error())
		tx.Rollback()
		return
	}
	fmt.Println("result no:", l)
	tx.Commit()
}
