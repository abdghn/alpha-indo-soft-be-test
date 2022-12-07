package database

import (
	"fmt"
	"github.com/abdghn/alpha-indo-soft-be-test/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
)

func ConnectDB() *gorm.DB {
	err := viper.ReadInConfig()
	if err != nil {
		return nil
	}

	dbUser := viper.Get("DB_USER").(string)
	dbPass := viper.Get("DB_PASSWORD").(string)
	dbHost := viper.Get("DB_HOST").(string)
	dbPort := viper.Get("DB_PORT").(string)
	dbName := viper.Get("DB_NAME").(string)

	consStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName)

	db, err := gorm.Open("mysql", consStr)
	if err != nil {
		log.Fatal("Error when connect db " + consStr + " : " + err.Error())
		return nil
	}

	db.Debug().AutoMigrate(
		models.Article{},
	)

	return db
}
