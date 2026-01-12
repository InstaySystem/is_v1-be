package container

import (
	"github.com/InstaySystem/is_v1-be/internal/handler"
	"github.com/InstaySystem/is_v1-be/internal/repository"
	svcImpl "github.com/InstaySystem/is_v1-be/internal/service/implement"
	"go.uber.org/zap"
)

type BookingContainer struct {
	Hdl *handler.BookingHandler
}

func NewBookingContainer(
	bookingRepo repository.BookingRepository,
	logger *zap.Logger,
) *BookingContainer {
	svc := svcImpl.NewBookingService(bookingRepo, logger)
	hdl := handler.NewBookingHandler(svc)

	return &BookingContainer{hdl}
}
