package db

import (
	"fmt"
	"github.com/erfanfs10/Az-Backend/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var DB *sqlx.DB

func ConnectDb() {

	dbHost := utils.GetEnv("DB_HOST")
	dbUser := utils.GetEnv("DB_USER")
	dbPassword := utils.GetEnv("DB_PASSWORD")
	dbName := utils.GetEnv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true",
		dbUser, dbPassword, dbHost, dbName)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("Can not connect to DB")
		log.Fatalln(err.Error())
	}
	fmt.Println("DB Connected")
	DB = db
}

func CreateTables() {
	for name, query := range schemas {
		_, err := DB.Exec(query.query)
		if err != nil {
			fmt.Println("Can not create table ", name)
			log.Fatalln(err)
		}
	}
	fmt.Println("All tables created")
}
