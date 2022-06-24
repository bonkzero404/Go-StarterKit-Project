package driver

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SuiteMock struct {
	suite.Suite
	DB   *gorm.DB
	Mock sqlmock.Sqlmock
}

func ConnectorMock(s *SuiteMock) {
	var err error

	DBMock, s.Mock, err = sqlmock.New()
	assert.NoError(s.T(), err)

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 DBMock,
		PreferSimpleProtocol: true,
	})

	s.DB, err = gorm.Open(dialector, &gorm.Config{})

	assert.NoError(s.T(), err)
}
