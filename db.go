package main

import "fmt"

type MySQL struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

var (
	localCfg = MySQL{}
)

func getConnectionString(cfg MySQL) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
}
