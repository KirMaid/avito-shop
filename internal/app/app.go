package app

import (
	"avitoshop/config"
	"avitoshop/internal/app/handlers/http"
	"avitoshop/internal/app/middleware"
	"avitoshop/internal/app/repositories"
	"avitoshop/internal/app/usecases"
	usecase "avitoshop/internal/app/usecases/user_info"
	"avitoshop/pkg/httpserver"
	"avitoshop/pkg/logger"
	"avitoshop/pkg/postgres"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Logger.LogLevel)
	pg, err := postgres.New(cfg.Postgres.URL, postgres.MaxPoolSize(cfg.Postgres.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}

	userRepo := repositories.NewUserRepository(pg.Pool)
	coinHistoryRepo := repositories.NewCoinHistoryRepository(pg.Pool)
	inventoryRepo := repositories.NewInventoryRepository(pg.Pool)
	transactionRepo := repositories.NewTransactionRepository(pg.Pool)

	expireDuration, _ := time.ParseDuration(cfg.Auth.TokenTTL)
	authUseCase := usecases.NewAuthUseCase(
		userRepo,
		cfg.Auth.HashSalt,
		[]byte(cfg.Auth.SigningKey),
		expireDuration,
	)
	userInfoUseCase := usecase.NewUserInfoUseCase(userRepo, coinHistoryRepo, inventoryRepo, transactionRepo)

	authMiddleware := middleware.AuthMiddleware([]byte(cfg.Auth.SigningKey))

	handler := gin.New()
	http.NewRouter(handler, l, authMiddleware, *authUseCase, *userInfoUseCase)

	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
