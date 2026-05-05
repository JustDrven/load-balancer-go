package orm

import (
	"fmt"

	"dev.justdrven/loadbalancer/internal/config"
	"dev.justdrven/loadbalancer/internal/service"
	"dev.justdrven/loadbalancer/pkg"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Initialize(configData config.Config) (*gorm.DB, error) {
	fmt.Println("[ORM-MANAGER] Initializing..")

	connectinString := createConnectionString(configData)

	db, err := gorm.Open(mysql.Open(connectinString), &gorm.Config{
		DefaultTransactionTimeout: pkg.TRANS_TIMEOUT,
	})
	if err != nil {
		return nil, err
	}

	fmt.Println("[ORM-MANAGER] Starting configuration of DB")
	db.AutoMigrate(&service.Service{})

	fmt.Println("[ORM-MANAGER] Done")
	return db, nil
}

func createConnectionString(cnf config.Config) string {
	addr := cnf.Address
	username := cnf.Username
	password := cnf.Password
	dbname := cnf.DbName

	return fmt.Sprintf(pkg.CONNECTION_STRING_DEFAULT, username, password, addr, dbname)
}
