package controllers

import (
	"database/sql"
	"fmt"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jeongseop/jsweb/app/models"
	r "github.com/revel/revel"
	"log"
	"errors"
)

var (
	Dbm *gorp.DbMap
)

func getConnectionString() (string, error) {
	Config := NewCustomConfig()
	if err := Config.LoadConfig("db.conf"); err != nil {
		return "", err
	}
	if !Config.SetSection("DEFAULT") {
		return "", errors.New("Set Section Failed")
	}
	host := Config.GetStringDefault("db.host", "localhost")
	port := Config.GetStringDefault("db.port", "3306")
	user := Config.GetStringDefault("db.user", "user")
	password := Config.GetStringDefault("db.password", "password")
	protocol := Config.GetStringDefault("db.protocol", "tcp")
	name := Config.GetStringDefault("db.name", "default")

	log.Printf("%s:%s@%s(%s:%s)/%s", user, password, protocol, host, port, name)

	return fmt.Sprintf("%s:%s@%s(%s:%s)/%s", user, password, protocol, host, port, name), nil
}

func InitDB() {
	dbConnectString, err := getConnectionString()
	if err != nil {
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
	t.ColMap("Pwd").Transient = true
	//setColumnSizes(t, map[string]int{})

	t = Dbm.AddTableWithName(models.Project{}, "project").SetKeys(true, "ProjectId")
	setColumnSizes(t, map[string]int{
		"ProjectComment":2048,
		"LaunchUrl":1024,
	})

	Dbm.TraceOn("[gorp]", r.INFO)
	if err := Dbm.CreateTablesIfNotExists(); err != nil {
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
