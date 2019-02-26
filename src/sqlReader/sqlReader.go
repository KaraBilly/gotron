package sqlReader

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Unknwon/goconfig"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"net/url"
)

const (
	baseAdr = "d:\\GoWorkspace\\gotron\\src\\config\\"

//uspInsertOrUpdateCustomerPoint = "uspInsertOrUpdateCustomerPoint"
)

var conn string

func init() {
	cfg, err := goconfig.LoadConfigFile(baseAdr + "conf.ini")
	if err != nil {
		panic("错误")
		//fmt.Println(e)
	}
	conn, err = cfg.GetValue("mssql", "basket")
	if err != nil {
		log.Fatal("Cannot find value", err)
		//fmt.Println(e)
	}
}

func SPRun() {
	fmt.Printf(conn)
	//db, err := sql.Open("mssql", "Provider=SQLOLEDB;data source=WIN-BUUBCMHEJPP.ebz-chn-dev.mkaws.com,502;Integrated Security=SSPI;initial catalog=Basket")
	query := url.Values{}
	query.Add("app name", "Basket")
	//query.Add("Integrated Security","SSPI")

	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword("siteuser", "ImQvAHiuetq"),
		Host:   fmt.Sprintf("%s:%d", "WIN-BUUBCMHEJPP.ebz-chn-dev.mkaws.com", 502),
		// Path:  instance, // if connecting to an instance instead of a port
		//IntegratedSecurity:"SSPI",
		RawQuery: query.Encode(),
	}
	db, err := sql.Open("sqlserver", u.String())
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return
	}
	defer db.Close()
	//param := make(map[string]string)
	//m1
	//err = exec(db, cmd)
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//err = db.Ping();
	//if(err!=nil){
	//	log.Fatal(err)
	//}
	_, err2 := db.ExecContext(context.Background(), "uspInsertOrUpdateCustomerPoint", sql.Named("CustomerId", 3123123), sql.Named("Point", 1.00))
	if err2 != nil {
		log.Fatal(err)
	}
}
