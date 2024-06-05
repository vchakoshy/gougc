package app

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/vchakoshy/gougc/docs"

	"github.com/vchakoshy/gougc/service/auth"
	"github.com/vchakoshy/gougc/service/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type App struct {
	db     *gorm.DB
	router *gin.Engine
}

func NewApp() App {
	// docker run --name go-postgres --rm -p 5432:5432 -e POSTGRES_PASSWORD=123456 -d postgres
	dsn := "user=postgres password=123456 host=localhost dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tehran"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: false,
	}})
	if err != nil {
		log.Fatal("gorm.Open error ", err)
	}
	return App{
		db:     db,
		router: gin.Default(),
	}
}

func (a App) Run() {
	v1 := a.router.Group("/api/v1")
	{
		user.NewModule(a.db, v1)
		auth.NewModule(a.db, v1)
	}

	a.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	a.router.Run()

}
