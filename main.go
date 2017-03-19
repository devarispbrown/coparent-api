package main

import (
	"fmt"
	"os"
	"strconv"

	_ "github.com/devarispbrown/coparent-api/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=postgres password=postgres dbname=coparent_dev sslmode=disable")
	name := "default"
	force := false
	verbose := true
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	portNumber, err := strconv.Atoi(port)
	if err != nil {
		panic("BLAH")
	}
	beego.BConfig.Listen.HTTPPort = portNumber
	beego.Run()
}
