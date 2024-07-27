package repository

import (
	"database/sql"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"kinolove/internal/entity/.gen/kinolove/public/model"
	. "kinolove/internal/entity/.gen/kinolove/public/table"
)

type RoleRepositoryImpl struct {
	db *sql.DB
}

func NewRoleRepository(db *sql.DB) *RoleRepositoryImpl {
	return &RoleRepositoryImpl{db: db}
}

func (r *RoleRepositoryImpl) GetUserRolesIds(usrId uuid.UUID) (*[]int64, error) {
	roles := make([]model.UsersRoles, 0)
	stmt := SELECT(UsersRoles.RoleID).FROM(UsersRoles).WHERE(UsersRoles.UserID.EQ(UUID(usrId)))
	if err := stmt.Query(r.db, &roles); err != nil {
		return nil, err
	}

	result := make([]int64, 0, len(roles))

	for _, role := range roles {
		result = append(result, role.RoleID)
	}

	return &result, nil
}
