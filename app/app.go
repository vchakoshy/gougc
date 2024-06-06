package app

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/vchakoshy/gougc/docs"
	"github.com/vchakoshy/gougc/pkg"

	"github.com/vchakoshy/gougc/service/attachment"
	"github.com/vchakoshy/gougc/service/auth"
	"github.com/vchakoshy/gougc/service/follow"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type App struct {
	db     *gorm.DB
	router *gin.Engine

	AuthModule       *auth.Module
	FollowModule     *follow.Module
	AttachmentModule *attachment.Module
}

const defaultDbDSN = "user=postgres password=123456 host=localhost dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tehran"

func NewApp() App {
	dsn := pkg.GetEnv("DB_DSN", defaultDbDSN)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})
	if err != nil {
		log.Fatal("gorm.Open error ", err)
	}

	return App{
		db:     db,
		router: gin.Default(),
	}
}

func (a *App) Setup() {
	a.AuthModule = auth.NewModule(a.db)
	a.FollowModule = follow.NewModule(a.db)
	a.AttachmentModule = attachment.NewModule(a.db, "../media/partial")

	v1 := a.router.Group("/api/v1")

	a.AuthModule.SetupRoutes(v1)
	a.FollowModule.SetupRoutes(v1)
	a.AttachmentModule.SetupRoutes(v1)

	a.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}

func (a *App) Run() {
	a.router.Run()
}
