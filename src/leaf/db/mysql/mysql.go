package mysql

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MysqlConfig struct {
	UserName  string
	Password  string
	IpAddress string
	Port      int
	DbName    string
	Charset   string
}

var mySqlCfg MysqlConfig

func init() {
	data, err := ioutil.ReadFile("conf/mysql.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &mySqlCfg)
	if err != nil {
		log.Fatal("%v", err)
	}
}

func ConnectMysql() (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", mySqlCfg.UserName, mySqlCfg.Password, mySqlCfg.IpAddress, mySqlCfg.Port, mySqlCfg.DbName, mySqlCfg.Charset)
	return sqlx.Open("mysql", dsn)
}
