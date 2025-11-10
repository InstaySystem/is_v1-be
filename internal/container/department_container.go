package container

import (
	"github.com/InstaySystem/is-be/internal/handler"
	repoImpl "github.com/InstaySystem/is-be/internal/repository/implement"
	svcImpl "github.com/InstaySystem/is-be/internal/service/implement"
	"github.com/InstaySystem/is-be/pkg/snowflake"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DepartmentContainer struct {
	Hdl *handler.DepartmentHandler
}

func NewDepartmentContainer(
	db *gorm.DB,
	sfGen snowflake.Generator,
	logger *zap.Logger,
) *DepartmentContainer {
	repo := repoImpl.NewDepartmentRepository(db)
	svc := svcImpl.NewDepartmentService(repo, sfGen, logger)
	hdl := handler.NewDepartmentHandler(svc)

	return &DepartmentContainer{hdl}
}
