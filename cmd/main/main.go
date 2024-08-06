package main

import (
	"github.com/Sanchir01/microservice_sandjma_category/internal/app"
	"github.com/Sanchir01/microservice_sandjma_category/internal/config"
	"github.com/Sanchir01/microservice_sandjma_category/pkg/lib/logger/handlers/slogpretty"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

var (
	development = "development"
	production  = "production"
)

func main() {
	cfg := config.MustLoad()

	lg := setupLogger(cfg.Env)

	application := app.NewApp(lg, cfg.GRPC.Port)

	go func() {
		application.GRPCSrv.MustRun()
	}()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	lg.Info("stoppping application", slog.String("signal", sign.String()))

	application.GRPCSrv.Stop()

	// TODO: add graceful shutdown

	//TODO: init db
}

func setupLogger(env string) *slog.Logger {
	var lg *slog.Logger
	switch env {
	case development:
		lg = setupPrettySlog()
	case production:
		lg = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return lg
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}

func setupSlog() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
}
