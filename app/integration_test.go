package app

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type IntegrationSuite struct {
	suite.Suite
	app *App
	db  *gorm.DB
}

func (s *IntegrationSuite) SetupSuite() {
	var err error
	s.db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		panic(err)
	}

	a := &App{
		db:     s.db,
		router: gin.Default(),
	}
	a.Setup()

	s.app = a
}

func (s *IntegrationSuite) TearDownSuite() {
	err := os.Remove("test.db")
	s.Nil(err)
}

func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, &ModuleTestSuite{})
}

func (s *IntegrationSuite) TestUcs() {
	err := s.app.FollowModule.Usecase.Follow(1, 2)
	s.Nil(err)
}
