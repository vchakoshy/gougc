package app

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
	app    *App
	db     *gorm.DB
	router *gin.Engine
}

func (s *ModuleTestSuite) SetupSuite() {
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

func (s *ModuleTestSuite) TearDownSuite() {
	err := os.Remove("test.db")
	s.Nil(err)
}

func TestModuleTestSuite(t *testing.T) {
	suite.Run(t, &ModuleTestSuite{})
}

func (s *ModuleTestSuite) TestUcs() {
	err := s.app.FollowModule.Usecase.Follow(1, 2)
	s.Nil(err)
}

func (s *ModuleTestSuite) TestAll() {

	tests := []struct {
		name         string
		method       string
		url          string
		data         any
		expectedCode int
		authHeader   string
	}{}

	for _, tt := range tests {
		w := httptest.NewRecorder()

		jsonValue, _ := json.Marshal(tt.data)
		req, _ := http.NewRequest(tt.method, tt.url, bytes.NewBuffer(jsonValue))
		req.Header.Add("Authorization", tt.authHeader)
		s.router.ServeHTTP(w, req)

		s.Equal(tt.expectedCode, w.Code, tt.name)
	}

}
