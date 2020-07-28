package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"

	"go-gin-test/context"
)

var DB *gorm.DB

//var DB = Init()

// Opening a database
func Init(conf *context.Config) {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	fmt.Println(555555555, conf.GetEnv())
	//fmt.Println(432342342, conf.ConfigFile())

	sec, err := conf.GetCfg().GetSection("database")

	if err != nil {
		log.Fatal("Fail to get ini section 'database': %v", err)
	}
	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err := gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Fatal("Fail open db : %v", err)
	}
	//CloseDB()
	db.SingularTable(true)

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	if err := db.DB().Ping(); err != nil {
		log.Fatalln(err)
	}

	DB = db
}

func CloseDB() {
	defer DB.Close()
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}
