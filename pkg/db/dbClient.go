package db

import (
	"github.com/aksharau/GoGormExamples/pkg/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbClient struct {
	db *gorm.DB
}

func GetDBClient() DbClient {
	dbC := DbClient{}
	dbC.db = dbConn()

	return dbC

}

func dbConn() (db *gorm.DB) {
<<<<<<< HEAD
=======

>>>>>>> 9e05657d28cafec9212b2acb586f927fc7162ec6

	dsn := "root:*****@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func (dbC *DbClient) SaveWeather(rec model.CityWeather) {
	dbC.db.Create(&rec)
}

func (dbC *DbClient) GetAllRec() []model.CityWeather {
	inst := model.CityWeather{}
	rec := []model.CityWeather{}
	dbC.db.Model(&inst).Select(inst.GetAllFields()).Find(&rec)

	return rec
}
