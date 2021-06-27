package component

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type DBConnection struct {
	Username string
	Password string
	Name     string
	Host     string
	Port     string
}

func (db *DBConnection) GetURL() string {
	return db.Username + ":" + db.Password + "@tcp" + "(" + db.Host + ":" + db.Port + ")/" + db.Name + "?" + "parseTime=true&loc=Local"
}

func GetDBConnection() (*gorm.DB, error) {
	dba := DBConnection{
		Username: os.Getenv("dbUserName"),
		Password: os.Getenv("dbPassword"),
		Name:     os.Getenv("dbName"),
		Host:     os.Getenv("dbHost"),
		Port:     os.Getenv("dbPort"),
	}

	db, err := gorm.Open(mysql.Open(dba.GetURL()), &gorm.Config{})
	return db, err
}
