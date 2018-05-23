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
		ar := dbs.NewAlbumMysql(db)

		aac := usecases.NewAlbumAllUsecase(ar)
		aap := presenters.NewAlbumAllPresenter()
		agbiu := usecases.NewAlbumGetByIDUsecase(ar)
		agbip := presenters.NewAlbumGetByIDPresenter()
		acc := usecases.NewAlbumCreateUsecase(ar)
		acp := presenters.NewAlbumCreatePresenter()

		ac := controllers.NewAlbumController(
			aac,
			aap,
			agbiu,
			agbip,
			acc,
			acp,
			usecases.NewAlbumClearUsecase(ar),
			presenters.NewAlbumClearPresenter(),
		)

		api.GET("/album", ac.GetAll)
		api.GET("/album/:id", ac.GetById)
		api.POST("/album", ac.Post)
		api.PATCH("/album/clear", ac.Clear)

		sr := dbs.NewSongMysql(db)

		sac := usecases.NewSongAllUsecase(sr)
		sap := presenters.NewSongAllPresenter()
		sgbiu := usecases.NewSongGetByIDUsecase(sr)
		sgbip := presenters.NewSongGetByIDPresenter()
		scc := usecases.NewSongCreateUsecase(sr)
		scp := presenters.NewSongCreatePresenter()

		sc := controllers.NewSongController(
			sac,
			sap,
			sgbiu,
			sgbip,
			scc,
			scp,
			usecases.NewSongClearUsecase(sr),
			presenters.NewSongClearPresenter(),
		)

		api.GET("/song", sc.GetAll)
		api.GET("/song/:id", sc.GetById)
		api.POST("/song", sc.Post)
		api.PATCH("/song/clear", sc.Clear)

		artistRepo := dbs.NewArtistMysql(db)

		artistAllU := usecases.NewArtistAllUsecase(artistRepo)
		artistAllP := presenters.NewArtistAllPresenter()
		artistCreateU := usecases.NewArtistGetByIDUsecase(artistRepo)
		artistCreateP := presenters.NewArtistGetByIDPresenter()
		artistGetByIDU := usecases.NewArtistCreateUsecase(artistRepo)
		artistGetByIDP := presenters.NewArtistCreatePresenter()

		artistCtrl := controllers.NewArtistController(
			artistAllU,
			artistAllP,
			artistCreateU,
			artistCreateP,
			artistGetByIDU,
			artistGetByIDP,
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
