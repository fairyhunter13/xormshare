package main

type User struct {
	ID      uint64 `xorm:"'id' pk autoincr notnull BIGINT" json:"id"`
	Name    string `xorm:"'name' index notnull VARCHAR(50)" json:"name"`
	Address string `xorm:"'address' notnull VARCHAR(100)" json:"address"`
}
