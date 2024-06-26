package follow

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

	tests := []struct {
		name         string
		method       string
		url          string
		data         any
		expectedCode int
		authHeader   string
	}{
		{
			name:         "follow",
			method:       http.MethodPost,
			url:          "/api/v1/follow/",
			data:         FollowRequest{LeaderID: 1, FollowerID: 2},
			expectedCode: http.StatusOK,
		},
		{
			name:         "followers list",
			method:       http.MethodGet,
			url:          "/api/v1/follow/followers/1",
			expectedCode: http.StatusOK,
		},
		{
			name:         "leaders list",
			method:       http.MethodGet,
			url:          "/api/v1/follow/leaders/1",
			expectedCode: http.StatusOK,
		},
		{
			name:         "follow exists",
			method:       http.MethodPost,
			url:          "/api/v1/follow/",
			data:         FollowRequest{LeaderID: 1, FollowerID: 2},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "follow with validation error",
			method:       http.MethodPost,
			url:          "/api/v1/follow/",
			data:         FollowRequest{LeaderID: 1},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "UnFollow with validation error",
			method:       http.MethodDelete,
			url:          "/api/v1/follow/",
			data:         FollowRequest{LeaderID: 1},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "UnFollow with not exists",
			method:       http.MethodDelete,
			url:          "/api/v1/follow/",
			data:         FollowRequest{LeaderID: 1, FollowerID: 50},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "UnFollow ",
			method:       http.MethodDelete,
			url:          "/api/v1/follow/",
			data:         FollowRequest{LeaderID: 1, FollowerID: 2},
			expectedCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		w := httptest.NewRecorder()

		jsonValue, _ := json.Marshal(tt.data)
		req, _ := http.NewRequest(tt.method, tt.url, bytes.NewBuffer(jsonValue))
		req.Header.Add("Authorization", tt.authHeader)
		s.router.ServeHTTP(w, req)

		s.Equal(tt.expectedCode, w.Code, tt.name)
	}

}
