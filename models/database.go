package models

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

const (
	HOST     = "localhost"
	USERNAME = "root"
	PASSWORD = "root"
	DBNAME   = "lara"
	PORTA    = 3306
)

var (
	Db  *gorm.DB
	err error
)

func DatabaseConnect() {
	connectString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", USERNAME, PASSWORD, HOST, PORTA, DBNAME)
	Db, err = gorm.Open(mysql.Open(connectString), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conectado ao banco de dados!")
}

func Migrate(tabelas ...interface{}) {
	err := Db.AutoMigrate(tabelas...)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteDatabase(tabelas ...interface{}) {
	err := Db.Migrator().DropTable(tabelas...)
	if err != nil {
		log.Fatal(err)
	}
}
