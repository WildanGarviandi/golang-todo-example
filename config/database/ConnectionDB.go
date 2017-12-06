package database

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kellinreaver/hafood-backend/models"
)

var db *gorm.DB

func ConnectDB()(*gorm.DB, error) {
	var err error

	db, err = gorm.Open("mysql", "root:rahasia@/hafood?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		return nil, err
	}

	var model = models.TodoModel{}

	db := db.AutoMigrate(&model)

	return db, err
}