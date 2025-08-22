package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/shikidy/golang_grpc_train_sso/internal/app/grpc"
)

type App struct {
	GRPCserv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	grpcApp := grpcapp.New(log, grpcPort)
	return &App{
		GRPCserv: grpcApp,
	}
}
