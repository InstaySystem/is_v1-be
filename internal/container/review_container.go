package container

import (
	"github.com/InstaySystem/is_v1-be/internal/handler"
	"github.com/InstaySystem/is_v1-be/internal/repository"
	svcImpl "github.com/InstaySystem/is_v1-be/internal/service/implement"
	"github.com/InstaySystem/is_v1-be/pkg/snowflake"
	"go.uber.org/zap"
)

type ReviewContainer struct {
	Hdl *handler.ReviewHandler
}

func NewReviewContainer(
	reviewRepo repository.ReviewRepository,
	sfGen snowflake.Generator,
	logger *zap.Logger,
) *ReviewContainer {
	svc := svcImpl.NewReviewService(reviewRepo, sfGen, logger)
	hdl := handler.NewReviewHandler(svc)

	return &ReviewContainer{hdl}
}
