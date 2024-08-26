package categorygrpc

import (
	"context"
	"log/slog"

	featureCategory "github.com/Sanchir01/microservice_sandjma_category/internal/feature/category"
	"github.com/Sanchir01/microservice_sandjma_category/internal/modain/models"
	sandjmav1 "github.com/Sanchir01/protos_files_job/pkg/gen/golang/category"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Categories interface {
	AllCategory(ctx context.Context) ([]models.Category, error)
}

type categoryServerApi struct {
	sandjmav1.UnimplementedCategoriesServer
	categories Categories
}

func NewCategoryServerApi(gRPC *grpc.Server, categ Categories) {
	sandjmav1.RegisterCategoriesServer(gRPC, &categoryServerApi{categories: categ})
}

func (s *categoryServerApi) GetAllCategory(ctx context.Context, _ *emptypb.Empty) (*sandjmav1.GetAllCategoryResponse, error) {

	allCategories, err := s.categories.AllCategory(ctx)
	slog.Error("error", err)
	if err != nil {
		return nil, err
	}
	convertCategory, err := featureCategory.MapCategoryToGRPCModel(allCategories)
	if err != nil {
		return nil, err
	}
	return &sandjmav1.GetAllCategoryResponse{Category: convertCategory}, nil
}
