package categoryservice

import (
	"context"
	"github.com/Sanchir01/microservice_sandjma_category/internal/modain/models"
	"log/slog"
)

type Category struct {
	log             *slog.Logger
	serviceCategory ServiceCategory
	appProvider     AppProvider
}
type ServiceCategory interface {
	AllCategory(ctx context.Context) ([]models.Category, error)
}

type AppProvider interface {
	App(ctx context.Context, appID int) (models.App, error)
}

func New(log *slog.Logger, serviceCategory ServiceCategory, appProvider AppProvider) *Category {
	return &Category{
		log: log, serviceCategory: serviceCategory, appProvider: appProvider,
	}
}

func (s *Category) AllCategory(ctx context.Context) ([]models.Category, error) {
	const op = "service.Category.AllCategory"
	log := s.log.With(
		slog.String("op", op),
		slog.String("category", "AllCategory"))
	log.Info("get all categories service")
	panic("implement me")
}
