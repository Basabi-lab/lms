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

	api := r.Group("/api")
	{
		albumRepo := dbs.NewAlbumMysql(db)

		albumCtrl := controllers.NewAlbumController(
			usecases.NewAlbumAllUsecase(albumRepo),
			presenters.NewAlbumAllPresenter(),
			usecases.NewAlbumGetByIDUsecase(albumRepo),
			presenters.NewAlbumGetByIDPresenter(),
			usecases.NewAlbumCreateUsecase(albumRepo),
			presenters.NewAlbumCreatePresenter(),
			usecases.NewAlbumClearUsecase(albumRepo),
			presenters.NewAlbumClearPresenter(),
		)

		api.GET("/album", albumCtrl.GetAll)
		api.GET("/album/:id", albumCtrl.GetById)
		api.POST("/album", albumCtrl.Post)
		api.PATCH("/album/clear", albumCtrl.Clear)

		songRepo := dbs.NewSongMysql(db)

		songCtrl := controllers.NewSongController(
			usecases.NewSongAllUsecase(songRepo),
			presenters.NewSongAllPresenter(),
			usecases.NewSongGetByIDUsecase(songRepo),
			presenters.NewSongGetByIDPresenter(),
			usecases.NewSongCreateUsecase(songRepo),
			presenters.NewSongCreatePresenter(),
			usecases.NewSongClearUsecase(songRepo),
			presenters.NewSongClearPresenter(),
		)

		api.GET("/song", songCtrl.GetAll)
		api.GET("/song/:id", songCtrl.GetById)
		api.POST("/song", songCtrl.Post)
		api.PATCH("/song/clear", songCtrl.Clear)

		artistRepo := dbs.NewArtistMysql(db)

		artistCtrl := controllers.NewArtistController(
			usecases.NewArtistAllUsecase(artistRepo),
			presenters.NewArtistAllPresenter(),
			usecases.NewArtistGetByIDUsecase(artistRepo),
			presenters.NewArtistGetByIDPresenter(),
			usecases.NewArtistCreateUsecase(artistRepo),
			presenters.NewArtistCreatePresenter(),
			usecases.NewArtistClearUsecase(artistRepo),
			presenters.NewArtistClearPresenter(),
		)

		api.GET("/artist", artistCtrl.GetAll)
		api.GET("/artist/:id", artistCtrl.GetById)
		api.POST("/artist", artistCtrl.Post)
		api.PATCH("/artist/clear", artistCtrl.Clear)
	}

	return r
}

func gormConnect() *gorm.DB {
	u, _ := url.Parse(os.Getenv("CLEARDB_DATABASE_URL"))
	CONNECT := fmt.Sprintf("%s@tcp(%s:3306)%s?parseTime=true", u.User.String(), u.Host, u.Path)
	db, err := gorm.Open("mysql", CONNECT)

	if err != nil {
		panic(err.Error())
	}

	db.LogMode(true)
	db.Set("gorm:table_options", "ENGINE=InnoDB")

	db.AutoMigrate(&models.Song{}, &models.Album{}, &models.Artist{})

	return db
}

func main() {
	db := gormConnect()
	r := setupRouter(db)
	r.Run(":" + os.Getenv("PORT"))
}
