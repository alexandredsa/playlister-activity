package main

import (
	"os"
	"strconv"

	"bitbucket.com/devplaylister/playlister-activity/config"
	"bitbucket.com/devplaylister/playlister-activity/routes"
	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload"
	"github.com/olivere/elastic"
)

var db *gorm.DB
var esClient *elastic.Client

func init() {
	dbport, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	db = config.GetDB(os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), dbport, os.Getenv("DB_NAME"))
	config.StartClient(os.Getenv("ELASTIC_SEARCH_URL"))
}

func main() {
	routes.SetUp()
	defer db.Close()
}
