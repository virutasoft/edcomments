package configuration

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	//no se necesit sino inicializarlo
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//configuration crea la estructura para la conexion
type configuration struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

//getConfiguration obtiene los datos para configurar la url
func getConfiguration() configuration {
	var c configuration
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

//GetConnection permite obtener una conexion a nuestra db
func GetConnection() *gorm.DB {
	c := getConfiguration()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", c.User, c.Password, c.Server, c.Port, c.Database)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db

}
