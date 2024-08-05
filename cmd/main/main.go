package main

import (
	"github.com/Sanchir01/microservice_sandjma_category/internal/config"
	"github.com/Sanchir01/microservice_sandjma_category/pkg/lib/logger/handlers/slogpretty"
	"log/slog"
	"os"
)

var (
	development = "development"
	production  = "production"
)

func main() {
	cfg := config.MustLoad()

	lg := setupLogger(cfg.Env)

	lg.Info("starting application", slog.Any("config", cfg))

	//TODO: start app

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
