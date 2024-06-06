package attachment

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ModuleTestSuite struct {
	suite.Suite
	db      *gorm.DB
	usecase *Usecase
	router  *gin.Engine
}

func (s *ModuleTestSuite) SetupSuite() {
	var err error
	s.db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		panic(err)
	}

	s.usecase = NewUsecase(s.db)
	s.router = gin.New()
	m := NewModule(s.db)
	m.SetupRoutes(s.router.Group("/api/v1"))
}

func (s *ModuleTestSuite) TearDownSuite() {
	err := os.Remove("test.db")
	s.Nil(err)
}

func TestModuleTestSuite(t *testing.T) {
	suite.Run(t, &ModuleTestSuite{})
}

func (s *ModuleTestSuite) TestAll() {

}

// TestSuccessfulUpload tests that tusd can perform a single upload
// from actual HTTP requests.
func TestSuccessfulUpload(t *testing.T) {
}
