package auth

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/vchakoshy/gougc/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type UsecaseTestSuite struct {
	suite.Suite
	db      *gorm.DB
	usecase *Usecase
}

func (s *UsecaseTestSuite) SetupSuite() {
	var err error
	s.db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		panic(err)
	}
	s.db.AutoMigrate(&models.User{})

	s.usecase = NewUsecase(s.db)
}

func (s *UsecaseTestSuite) TearDownSuite() {
	err := os.Remove("test.db")
	s.Nil(err)
}

func TestNoteTestSuite(t *testing.T) {
	suite.Run(t, &UsecaseTestSuite{})
}

func (s *UsecaseTestSuite) TestRegister() {
	// register user
	o, err := s.usecase.Register(RegisterForm{Username: "vahid", Password: "123456"})
	s.Nil(err)
	s.Greater(o.ID, uint(0))
	s.NotEqual("123456", o.Password)

	pok := s.usecase.CheckPasswordHash(o.Password, "123456")
	s.True(pok)

	// log in by registered user
	o, err = s.usecase.Login(LoginForm{Username: "vahid", Password: "123456"})
	s.Nil(err)
	s.NotZero(o.ID)

	// log in by wrong password
	o, err = s.usecase.Login(LoginForm{Username: "vahid", Password: "12345"})
	s.Equal(ErrUserNotFound, err)
	s.Zero(o)

	// log in by wrong username
	o, err = s.usecase.Login(LoginForm{Username: "vahid_does_not_exists", Password: "12345"})
	s.Equal(gorm.ErrRecordNotFound, err)
	s.Zero(o)

	// register user with duplicate username
	o, err = s.usecase.Register(RegisterForm{Username: "vahid", Password: "123456"})
	s.NotNil(err)
	s.Zero(o.ID)

	// register user with bad password
	o, err = s.usecase.Register(RegisterForm{Username: "vahid", Password: strings.Repeat("1", 100)})
	s.NotNil(err)
	s.Zero(o.ID)

}
