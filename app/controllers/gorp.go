package controllers

import (
	"database/sql"
	"fmt"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jeongseop/jsweb/app/models"
	r "github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var (
	Dbm *gorp.DbMap
)

func getConnectionString() string {
	c, err := ReadConfig("../../conf/db.conf")
	c2, err2 := ReadConfig("/home/gopher/go/src/github.com/jeongseop/jsweb/conf/db.conf")
	c3, err3 := ReadConfig("github.com/jeongseop/jsweb/conf/db.conf")
	c4, err4 := ReadConfig("/conf/db.conf")
	if err != nil {
		log.Println("t1")
	}
	if err2 != nil {
		log.Println("t2")
	}
	if err3 != nil {
		log.Println("t3")
	}
	if err4 != nil {
		log.Println("t4")
	}
	panic(err)
	host2 := getString(c2, "DEFAULT", "db.host")
	host3 := getString(c3, "DEFAULT", "db.host")
	host4 := getString(c4, "DEFAULT", "db.host")
	log.Printf("2[%s] 3[%s] 4[%s]\n", host2, host3, host4)

	host := getString(c, "DEFAULT", "db.host")
	port := getString(c, "DEFAULT", "db.port")
	user := getString(c, "DEFAULT", "db.user")
	password := getString(c, "DEFAULT", "db.password")
	protocol := getString(c, "DEFAULT", "db.protocol")
	name := getString(c, "DEFAULT", "db.name")

	/*r.Config.SetSection("dev")
	host := r.Config.StringDefault("db.host", "localhost")
	port := r.Config.StringDefault("db.port", "3306")
	user := r.Config.StringDefault("db.user", "")
	password := r.Config.StringDefault("db.password", "")
	protocol := r.Config.StringDefault("db.protocol", "tcp")
	name := r.Config.StringDefault("db.name", "")
	*/
	log.Printf("%s:%s@%s(%s:%s)/%s", user, password, protocol, host, port, name)

	return fmt.Sprintf("%s:%s@%s(%s:%s)/%s", user, password, protocol, host, port, name)
}

func InitDB() {
	dbConnectString := getConnectionString()
	if dbConnectString == "" {
		log.Fatal("getConnectionString failed!!!")
	}
	db, err := sql.Open("mysql", dbConnectString)
	if err != nil {
		panic(err)
	}
	Dbm = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
		for col, size := range colSizes {
			t.ColMap(col).MaxSize = size
		}
	}

	t := Dbm.AddTableWithName(models.Member{}, "member").SetKeys(false, "UserId")
	t.ColMap("Password").Transient = true
	setColumnSizes(t, map[string]int{})

	Dbm.TraceOn("[gorp]", r.INFO)
	Dbm.CreateTables()

	bcryptPassword, _ := bcrypt.GenerateFromPassword(
		[]byte("demo"), bcrypt.DefaultCost)
	demoMember := &models.Member{"jeongseop", "demo", "jeongsub3312@naver.com", bcryptPassword}
	if err := Dbm.Insert(demoMember); err != nil {
		panic(err)
	}
}

type GorpController struct {
	*r.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() r.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
