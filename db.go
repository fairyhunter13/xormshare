package main

import (
	"fmt"

	"github.com/fairyhunter13/xorm"
)

type MySQL struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

var (
	localCfg = MySQL{
		Host:     "192.168.99.100",
		Port:     "3306",
		Username: "root",
		Password: "kitabisa",
		Database: "testing",
	}
	engine *xorm.Engine
)

func getConnectionString(cfg MySQL) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
}

func initEngine() {
	var err error
	engine, err = xorm.NewEngine("mysql", getConnectionString(localCfg))
	panicIfErr(err)

	initSync()
}

func initSync() {
	engine.Sync2(&User{})
}
