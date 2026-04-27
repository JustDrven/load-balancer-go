package manager

import (
	"fmt"
	"time"

	"dev.justdrven/loadbalancer/data"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	CONNECTION_STRING_DEFAULT string        = "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	TRANS_TIMEOUT             time.Duration = 5 * time.Second
)

func OrmInit(configData data.Config) *gorm.DB {
	fmt.Println("[ORM-MANAGER] Initializing..")

	connectinString := createConnectionString(configData)

	db, err := gorm.Open(mysql.Open(connectinString), &gorm.Config{
		DefaultTransactionTimeout: TRANS_TIMEOUT,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("[ORM-MANAGER] Starting configuration of DB")
	db.AutoMigrate(&data.Service{})

	fmt.Println("[ORM-MANAGER] Done")
	return db
}

func createConnectionString(cnf data.Config) string {
	addr := cnf.Address
	username := cnf.Username
	password := cnf.Password
	dbname := cnf.DbName

	return fmt.Sprintf(CONNECTION_STRING_DEFAULT, username, password, addr, dbname)
}
