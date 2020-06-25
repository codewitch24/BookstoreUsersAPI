package usersdb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
	"os"
)

var Config *viper.Viper

const (
	MysqlUsersUsername = "MYSQL_USERS_USERNAME"
	MysqlUsersPassword = "MYSQL_USERS_PASSWORD"
	MysqlUsersHost     = "MYSQL_USERS_HOST"
	MysqlUsersSchema   = "MYSQL_USERS_SCHEMA"
)

var (
	Client *sql.DB

	username = os.Getenv(MysqlUsersUsername)
	password = os.Getenv(MysqlUsersPassword)
	host     = os.Getenv(MysqlUsersHost)
	schema   = os.Getenv(MysqlUsersSchema)
)

func init() {
	Config = viper.New()
	Config.SetConfigName("config")
	Config.AddConfigPath(".")
	if err := Config.ReadInConfig(); err != nil {
		log.Panic(err)
	}
	datasource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)
	var e error
	Client, e = sql.Open("mysql", datasource)
	if e != nil {
		log.Panic(e)
	}
	if err := Client.Ping(); err != nil {
		log.Panic(err)
	}
	log.Println("database successfully configured")
}
