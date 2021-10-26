package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySqlInstance : MySqlInstance Struct
type MySqlInstance struct {
	DB *gorm.DB
}

// MI : An instance of MySqlInstance Struct
var MI MySqlInstance

// ConnectDB - database connection
func ConnectDB() {

	var err error
	var user string = os.Getenv("APP_DB_USERNAME")
	var password string = os.Getenv("APP_DB_PASSWORD")
	var server string = os.Getenv("APP_DB_SERVER")
	var port string = os.Getenv("APP_DB_PORT")
	var dbname string = os.Getenv("APP_DB_NAME")

	connectionString :=
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, server, port, dbname)

	//connectionString := os.Getenv("database_url")

	MI = MySqlInstance{}

	MI.DB, err = gorm.Open(mysql.Open(connectionString))
	if err != nil {
		MI = MySqlInstance{}
		fmt.Println("Erro ao conectar no banco de dados", err)
		return
	}
	fmt.Println("Database connected!")

}

/*
func InitTables() {

	MI.DB.AutoMigrate(&models.User{})
}
*/
