package grpcapp

import (
	"fmt"
	categorygrpc "github.com/Sanchir01/microservice_sandjma_category/internal/grpc/category"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type GrpcApp struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func NewGrpcApp(log *slog.Logger, port int, categoryService categorygrpc.Categories) *GrpcApp {
	gRPCServer := grpc.NewServer()

	categorygrpc.NewCategoryServerApi(gRPCServer, categoryService)
	return &GrpcApp{log, gRPCServer, port}
}

func (g *GrpcApp) MustRun() {
	if err := g.Run(); err != nil {
		g.log.Error("error starting server", err.Error())
		panic(err)
	}
}

func (g *GrpcApp) Run() error {
	const op = "grcpapp.GrpcApp.running"

	log := g.log.With(
		slog.String("op", op),
		slog.Int("port", g.port),
	)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", g.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("starting gRPC server")

	if err := g.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (g *GrpcApp) Stop() {
	const op = "grpcapp.GrpcApp.Stop"

	g.log.With(slog.String("op", op), slog.Int("port", g.port))

	g.gRPCServer.GracefulStop()
}
