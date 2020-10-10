package initialize

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"

	//"gorm.io/plugin/dbresolver"

	"go-gin-test/global"
)

//var DB = Init()

// Opening a database
func MysqlInit() {
	var (
		err                                       error
		dbName, user, password, host, tablePrefix string
	)
	var (
		err1                             error
		dbName1, user1, password1, host1 string
	)

	var (
		err2                             error
		dbName2, user2, password2, host2 string
	)

	masterSec, err := global.CTX_CONFIG.GetCfg().GetSection("database_master")
	slaveSec1, err1 := global.CTX_CONFIG.GetCfg().GetSection("database_slave1")
	slaveSec2, err2 := global.CTX_CONFIG.GetCfg().GetSection("database_slave2")

	if err != nil {
		log.Fatal("Fail to get ini section 'database_master': %v", err)
	}
	if err1 != nil {
		log.Fatal("Fail to get ini section 'database_slave1': %v", err)
	}
	if err2 != nil {
		log.Fatal("Fail to get ini section 'database_slave2': %v", err)
	}
	//dbType = sec.Key("TYPE").String()
	dbName = masterSec.Key("NAME").String()
	user = masterSec.Key("USER").String()
	password = masterSec.Key("PASSWORD").String()
	host = masterSec.Key("HOST").String()
	tablePrefix = masterSec.Key("TABLE_PREFIX").String()

	dbName1 = slaveSec1.Key("NAME").String()
	user1 = slaveSec1.Key("USER").String()
	password1 = slaveSec1.Key("PASSWORD").String()
	host1 = slaveSec1.Key("HOST").String()
	//tablePrefix1 = slaveSec1.Key("TABLE_PREFIX").String()

	dbName2 = slaveSec2.Key("NAME").String()
	user2 = slaveSec2.Key("USER").String()
	password2 = slaveSec2.Key("PASSWORD").String()
	host2 = slaveSec2.Key("HOST").String()
	//tablePrefix2 = slaveSec2.Key("TABLE_PREFIX").String()

	dnsMaster := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)

	dnsSlave1 := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user1,
		password1,
		host1,
		dbName1)
	dnsSlave2 := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user2,
		password2,
		host2,
		dbName2)

	//dns := "root:123456@tcp(127.0.0.1:3306)/gezi?charset=utf8&parseTime=True&loc=Local"
	//dns79 := "www:123456@tcp(10.15.4.79:5306)/gezi?charset=utf8&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gezi?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	//db.Use(dbresolver.Register(dbresolver.Config{
	//	// `db2` 作为 sources，`db3`、`db4` 作为 replicas
	//	Sources: []gorm.Dialector{mysql.Open("root:123456@tcp(127.0.0.1:3306)/gezi?charset=utf8&parseTime=True&loc=Local")},
	//	Replicas: []gorm.Dialector{
	//		mysql.Open("www:123456@tcp(10.15.4.79:5306)/gezi?charset=utf8&parseTime=True&loc=Local"),
	//		mysql.Open("www:123456@tcp(10.15.4.79:5306)/gezi?charset=utf8&parseTime=True&loc=Local")}),
	////	// sources/replicas 负载均衡策略
	//	Policy: dbresolver.RandomPolicy{},
	//}))
	db, err := gorm.Open(mysql.Open(dnsMaster), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix, // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,        // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})

	db.Use(dbresolver.Register(dbresolver.Config{
		// use `db2` as sources, `db3`, `db4` as replicas
		Sources: []gorm.Dialector{mysql.Open(dnsMaster)},
		Replicas: []gorm.Dialector{
			mysql.Open(dnsSlave1),
			mysql.Open(dnsSlave2)},
		// sources/replicas load balancing policy
		Policy: dbresolver.RandomPolicy{},
	}))

	//db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
	//	user,
	//	password,
	//	host,
	//	dbName)),&gorm.Config{})

	if err != nil {
		log.Fatal("Fail open db : %v", err)
	}

	db.Use(
		dbresolver.Register(dbresolver.Config{ /* xxx */ }).
			SetConnMaxIdleTime(time.Hour).
			SetConnMaxLifetime(24 * time.Hour).
			SetMaxIdleConns(100).
			SetMaxOpenConns(200))

	global.DB = db
}

func CloseDB() {
	//defer global.DB.Closed()
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return global.DB
}
