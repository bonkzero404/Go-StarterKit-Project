package repositories

import (
	"go-starterkit-project/database/driver"
	"go-starterkit-project/domain/stores"
	"go-starterkit-project/modules/role/domain/interfaces"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	driver.SuiteMock
	repository interfaces.RoleRepositoryInterface
	role       *stores.Role
}

func (s *Suite) SetupSuite() {
	driver.ConnectorMock(&s.SuiteMock)
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_RoleRepository_CreateRole() {
	s.repository = NewRoleRepository(s.DB)

	s.role = &stores.Role{
		RoleName:        "Admin",
		RoleDescription: "Lorem Ipsum Dolor",
		IsActive:        true,
	}

	s.role.ID = uuid.New()
	s.role.CreatedAt = time.Now()
	s.role.UpdatedAt = time.Now()

	s.Mock.ExpectBegin()

	s.Mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "roles" ("id","created_at","updated_at","deleted_at","role_name","role_description","is_active")
		VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id"`)).
		WithArgs(s.role.ID, s.role.CreatedAt, s.role.UpdatedAt, nil, s.role.RoleName, s.role.RoleDescription, s.role.IsActive).
		WillReturnRows(sqlmock.NewRows([]string{"role_name"}).AddRow(s.role.ID))

	s.Mock.ExpectCommit()

	err := s.repository.CreateRole(s.role).Error
	assert.NoError(s.T(), err)

	err = s.Mock.ExpectationsWereMet()
	assert.NoError(s.T(), err)
}
