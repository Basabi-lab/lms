package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/kokoax/music_lab/adapters/controllers"
	"github.com/kokoax/music_lab/domains/models"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	controllers.NewAlbumHandler(r, db)

	return r
}

func gormConnect() *gorm.DB {
	u, _ := url.Parse(os.Getenv("CLEARDB_DATABASE_URL"))
	CONNECT := fmt.Sprintf("%s@tcp(%s:3306)%s", u.User.String(), u.Host, u.Path)
	db, err := gorm.Open("mysql", CONNECT)

	if err != nil {
		panic(err.Error())
	}

	db.LogMode(true)
	db.Set("gorm:table_options", "ENGINE=InnoDB")

	db.AutoMigrate(&models.Music{}, &models.Album{})

	return db
}

func main() {
	db := gormConnect()
	r := setupRouter(db)
	r.Run(":" + os.Getenv("PORT"))
}
