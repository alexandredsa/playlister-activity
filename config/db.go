package config

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const defaultLifeTime = time.Minute * 5
const defaultMaxIdleConns = 5
const defaultMaxOpenConns = 5

func GetDB(user string, pass string, host string, port int, dbName string) *gorm.DB {
	db, _ := gorm.Open("mysql", generateDsn(user, pass, host, port, dbName))
	return db
}

func generateDsn(user string, pass string, host string, port int, dbName string) string {
	return user + ":" + pass + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + dbName + "?charset=utf8&parseTime=True&loc=America%2FSao_Paulo"
}
