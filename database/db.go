package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectaComBancoDeDados() {

	stringConnect := "root:Pac4386@01@tcp(127.0.0.1:3306)/bd_nsc?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(stringConnect), &gorm.Config{})
	if err != nil {
		log.Panic("Erro aoconectar no banco de dados")
	}
}
