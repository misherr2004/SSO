package grpcapp

import (
	authgrpc "SSO/cmd/internal/grpc/auth"
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"strconv"
)

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

// создание нового gRPC сервера приложения
func New(
	log *slog.Logger,
	port int,
) *App {
	gRPCServer := grpc.NewServer()
	authgrpc.Register(gRPCServer)
	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"

	log := a.log.With(
		slog.String("op", op),
		slog.Int("port", a.port),
	)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server is running", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
}

// метод СТОП, чтобы остановить gRPC сервер
func (a *App) Stop() {
	const op = "grpcapp.Stop"

	a.log.With(slog.String("op", op)).
		Info("stopping gRPC server", slog.Int("port", a.port))

	a.gRPCServer.GracefulStop() //нельзя просто выкл приложение. некоторые опперации должны закончиться прежде
}
