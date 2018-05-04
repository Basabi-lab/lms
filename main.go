package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/Basabi-lab/lms/adapters/controllers"
	"github.com/Basabi-lab/lms/adapters/dbs"
	"github.com/Basabi-lab/lms/adapters/presenters"
	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/usecases"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	ar := dbs.NewAlbumMysql(db)

	aac := usecases.NewAllAlbumUsecase(ar)
	aap := presenters.NewAllAlbumPresenter()

	ac := controllers.NewAlbumController(aac, aap)

	r.GET("/album", ac.GetAllAlbum)
	r.GET("/album/:id", ac.GetWithId)

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

	db.AutoMigrate(&models.Song{}, &models.Album{})

	return db
}

func main() {
	db := gormConnect()
	r := setupRouter(db)
	r.Run(":" + os.Getenv("PORT"))
}
