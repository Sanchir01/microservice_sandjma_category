package categorygrpc

import (
	"context"
	sandjmav1 "github.com/Sanchir01/protos_files_job/pkg/gen/golang/category"
	"google.golang.org/grpc"
	"log/slog"
)

type categoryServerApi struct {
	sandjmav1.UnimplementedCategoriesServer
}

func NewCategoryServerApi(gRPC *grpc.Server) {
	sandjmav1.RegisterCategoriesServer(gRPC, &categoryServerApi{})
}

func (s *categoryServerApi) GetCategory(ctx context.Context, req *sandjmav1.Empty) (*sandjmav1.GetAllCategoryResponse, error) {
	slog.Info("swsawdcsfd", req)
	//return &sandjmav1.GetAllCategoryResponse{Category: nil}, nil
	panic("implement me")
}
