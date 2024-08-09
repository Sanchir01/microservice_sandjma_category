package app

import (
	grpcapp "github.com/Sanchir01/microservice_sandjma_category/internal/app/grpc"
	"github.com/Sanchir01/microservice_sandjma_category/internal/config"
	categoryservice "github.com/Sanchir01/microservice_sandjma_category/internal/service/category"
	storagePG "github.com/Sanchir01/microservice_sandjma_category/internal/storage/postgres"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log/slog"
)

type App struct {
	GRPCSrv *grpcapp.GrpcApp
}

func NewApp(log *slog.Logger, cfg *config.Config, db *sqlx.DB) *App {
	storage := storagePG.NewStorage(db)
	categoryService := categoryservice.New(log, storage)
	grpcApp := grpcapp.NewGrpcApp(log, cfg.GRPC.Port, categoryService)
	return &App{
		GRPCSrv: grpcApp,
	}
}
