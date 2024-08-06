package app

import (
	grpcapp "github.com/Sanchir01/microservice_sandjma_category/internal/app/grpc"
	"log/slog"
)

type App struct {
	GRPCSrv *grpcapp.GrpcApp
}

func NewApp(log *slog.Logger, grpcPort int) *App {
	grpcApp := grpcapp.NewGrpcApp(log, grpcPort)
	return &App{
		GRPCSrv: grpcApp,
	}
}
