package mysqlx

import (
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func InitMysql(dataSource string) *xorm.Engine {
	db, err := xorm.NewEngine("mysql", dataSource)

	if err != nil {
		// db.Close()
		panic(err)
	}
	//连接测试
	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}
