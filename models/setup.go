package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDataBase() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	Dbdriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	DB, err = gorm.Open(Dbdriver, DBURL)

	if err != nil {
		log.Fatal("connection error:", err)
	} else {
		log.Println("We are connected to the database ", Dbdriver)
	}

	sqlDB := DB.DB()
	err = ConfigureConnectionPool(sqlDB)
	if err != nil {
		log.Fatal("DB connection pool configure error:", err)
	}

	DB.AutoMigrate(&User{})
}

func ConfigureConnectionPool(sqlDB *sql.DB) error {
	DbIdleConns, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
	if err != nil {
		log.Fatal("Cannot read DB_MAX_IDLE_CONNS!")
		return err
	}
	DbMaxConns, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
	if err != nil {
		log.Fatal("Cannot read DB_MAX_OPEN_CONNS!")
		return err
	}
	DbLifetime, err := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME"))
	if err != nil {
		log.Fatal("Cannot read DB_MAX_LIFETIME!")
		return err
	}
	sqlDB.SetMaxIdleConns(DbIdleConns)
	sqlDB.SetMaxOpenConns(DbMaxConns)
	sqlDB.SetConnMaxLifetime(time.Hour * time.Duration(DbLifetime))
	return nil
}
