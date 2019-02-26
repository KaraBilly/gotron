package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Unknwon/goconfig"
	_ "github.com/denisenkom/go-mssqldb"
	"gotron/src/apiCall"
	"gotron/src/sqlReader"
	"log"
)

const baseAdr = "d:\\GoWorkspace\\gotron\\src\\config\\"

func init() {
	//baseAdr = os.Getenv("GOPATH")
}

//const OneSecond = 1*time.Second + 10*time.Millisecond
func main() {
	/*c := cron.New()
	c.Start()
	defer c.Stop()*/
	//c.AddFunc("*/5 * * * * ?", func() { fmt.Println("Every hour on the half hour") })
	//c.AddFunc("@hourly",      func() { fmt.Println("Every hour") })
	//select {

	//}
	//base := os.Getenv("GOPATH")
	sqlReader.SPRun()
	point := apiCall.Get()
	data, err := json.MarshalIndent(point, "", "   ")
	if err != nil {
		log.Fatal("jsonMarshalIndent failed", err.Error())
	}
	fmt.Printf("%s\n", data)

	//nosqlReader.GetOne()
	//nosqlReader.GetMany()
	//test 5
	cfg, err := goconfig.LoadConfigFile(baseAdr + "conf.ini")
	if err != nil {
		panic("错误")
		//fmt.Println(e)
	}
	/*userName, err := cfg.GetValue("mysql", "username")
	passWord,err:=cfg.GetValue("mysql", "password")
	url,err:=cfg.GetValue("mysql", "url")
	dbName,err:=cfg.GetValue("mysql", "gocron")
	path := strings.Join([]string{userName, ":", passWord, "@tcp(",url,")/", dbName, "?charset=utf8"}, "")
	//valut, err := cfg.Int("must", int)*/
	path, err := cfg.GetValue("mssql", "basket")
	conn, err := sql.Open("mssql", path)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()
	//fmt.Println(value)

}
