package categorygrpc

import (
	"context"
	sandjmav1 "github.com/Sanchir01/protos_files_job/pkg/gen/golang/category"
	"google.golang.org/grpc"
	"log/slog"
)

type Categories interface {
	AllCategory(ctx context.Context) (*sandjmav1.GetAllCategoryResponse, error)
}

type categoryServerApi struct {
	sandjmav1.UnimplementedCategoriesServer
	categories Categories
}

func NewCategoryServerApi(gRPC *grpc.Server, categ Categories) {
	sandjmav1.RegisterCategoriesServer(gRPC, &categoryServerApi{categories: categ})
}

func (s *categoryServerApi) GetAllCategory(ctx context.Context, req *sandjmav1.Empty) (*sandjmav1.GetAllCategoryResponse, error) {
	slog.Info("swsawdcsfd", req)
	allCategories, err := s.categories.AllCategory(ctx)
	if err != nil {
		return nil, err
	}
	return &sandjmav1.GetAllCategoryResponse{Category: allCategories.Category}, nil
}
