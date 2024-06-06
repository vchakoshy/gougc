package attachment

import (
	"github.com/gin-gonic/gin"
	"github.com/tus/tusd/v2/pkg/filestore"
	tusd "github.com/tus/tusd/v2/pkg/handler"

	"github.com/vchakoshy/gougc/models"
	"gorm.io/gorm"
)

type Module struct {
	db         *gorm.DB
	Usecase    Usecase
	tusHandler *tusd.Handler
	tusStore   filestore.FileStore
}

func NewModule(db *gorm.DB, mediaPath string) *Module {
	db.AutoMigrate(&models.Attachment{})

	tusStore := filestore.FileStore{
		Path: mediaPath,
	}

	composer := tusd.NewStoreComposer()
	tusStore.UseIn(composer)
	th, _ := tusd.NewHandler(tusd.Config{
		BasePath:                "/api/v1/attachment/files/",
		StoreComposer:           composer,
		RespectForwardedHeaders: true,
	})

	return &Module{
		db:         db,
		tusHandler: th,
		tusStore:   tusStore,
	}
}

func (m *Module) SetupRoutes(router *gin.RouterGroup) {
	r := router.Group("/attachment")
	{
		files := r.Group("/files/")
		{
			files.OPTIONS("/", gin.WrapH(m.tusHandler))
			files.POST("/", gin.WrapF(m.tusHandler.PostFile))
			files.HEAD("/:id", gin.WrapF(m.tusHandler.HeadFile))
			files.PATCH("/:id", gin.WrapF(m.tusHandler.PatchFile))
			files.GET("/:id", gin.WrapF(m.tusHandler.GetFile))
		}

		r.POST("/save/", func(ctx *gin.Context) {
			// f, err := m.tusStore.GetUpload(ctx, ctx.PostForm("file_id"))

			// info, err := f.GetInfo(ctx)

			// src, err := f.GetReader(ctx)

			// srcByte, err := io.ReadAll(src)
		})
	}
}
