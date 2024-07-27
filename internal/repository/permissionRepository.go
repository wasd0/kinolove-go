package repository

import (
	"database/sql"
	"fmt"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"kinolove/internal/entity/.gen/kinolove/public/model"
	. "kinolove/internal/entity/.gen/kinolove/public/table"
)

type PermissionRepositoryImpl struct {
	db *sql.DB
}

func NewPermissionRepository(db *sql.DB) *PermissionRepositoryImpl {
	return &PermissionRepositoryImpl{db: db}
}

func (p *PermissionRepositoryImpl) FindUserPermissions(usrId uuid.UUID) (*[]model.UsersPermissions, error) {
	stmt := SELECT(UsersPermissions.AllColumns).
		FROM(UsersPermissions).WHERE(UsersPermissions.UserID.EQ(UUID(usrId)))

	permissions := make([]model.UsersPermissions, 0)

	if err := stmt.Query(p.db, &permissions); err != nil {
		return nil, fmt.Errorf("cannot find permissions, err: %v", err)
	}

	return &permissions, nil
}

func (p *PermissionRepositoryImpl) FindRolePermissions(roleIds *[]int64) (*[]model.RolesPermissions, error) {
	if len(*roleIds) == 0 {
		return nil, qrm.ErrNoRows
	}

	sqlIds := make([]Expression, 0, len(*roleIds))

	for _, id := range *roleIds {
		sqlIds = append(sqlIds, Int64(id))
	}

	stmt := SELECT(RolesPermissions.AllColumns).
		FROM(RolesPermissions).WHERE(RolesPermissions.RoleID.IN(sqlIds...))

	permissions := make([]model.RolesPermissions, 0)

	if err := stmt.Query(p.db, &permissions); err != nil {
		msg := fmt.Sprintf("cannot find permissions, err: %v", err)
		return nil, errors.New(msg)
	}

	return &permissions, nil
}
