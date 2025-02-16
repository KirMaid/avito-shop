package app

import (
	"avitoshop/config"
	"avitoshop/internal/app/handlers/http"
	"avitoshop/internal/app/middleware"
	"avitoshop/internal/app/repositories"
	auth "avitoshop/internal/app/usecases/auth"
	buygood "avitoshop/internal/app/usecases/buy_good"
	sendcoins "avitoshop/internal/app/usecases/send_coins"
	userinfo "avitoshop/internal/app/usecases/user_info"
	"avitoshop/pkg/httpserver"
	"avitoshop/pkg/logger"
	"avitoshop/pkg/postgres"
	"avitoshop/pkg/redis"
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
	inventoryRepo := repositories.NewInventoryRepository(pg.Pool)
	transactionRepo := repositories.NewTransactionRepository(pg.Pool)
	goodRepo := repositories.NewGoodRepository(pg.Pool)

	//l.Fatal(fmt.Errorf("PasswordRedis: %s", cfg.Redis.Password))
	redisAddr := fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port)
	redisClient, err := redis.New(
		redisAddr,
		cfg.Redis.Password,
		cfg.Redis.DB,
	)

	redisUserRepo := repositories.NewRedisUserRepository(redisClient.Client, 0)
	redisInventoryRepo := repositories.NewRedisInventoryRepository(redisClient.Client, 0)
	redisGoodRepo := repositories.NewRedisGoodRepository(redisClient.Client, 0)
	//redisTransactionRepo := repositories.NewRedisTransactionRepository(redisClient.Client, 5000000)

	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - redis.New: %w", err))
	}

	expireDuration, _ := time.ParseDuration(cfg.Auth.TokenTTL)
	authUseCase := auth.NewAuthUseCase(
		userRepo,
		redisUserRepo,
		cfg.Auth.HashSalt,
		[]byte(cfg.Auth.SigningKey),
		expireDuration,
	)
	userInfoUseCase := userinfo.NewUserInfoUseCase(
		userRepo,
		inventoryRepo,
		transactionRepo,
		goodRepo,
		redisUserRepo,
		redisInventoryRepo,
		redisGoodRepo,
		//redisTransactionRepo,
	)
	sendCoinsUseCase := sendcoins.NewSendCoinsUseCase(
		pg.Pool,
		userRepo,
		transactionRepo,
		redisUserRepo,
	)
	buyGoodUseCase := buygood.NewBuyGoodUseCase(
		pg.Pool,
		userRepo,
		goodRepo,
		inventoryRepo,
		redisUserRepo,
		redisGoodRepo,
		redisInventoryRepo,
	)

	// RabbitMQ RPC Server
	//rmqRouter := amqprpc.NewRouter(translationUseCase)
	//
	//rmqServer, err := server.New(cfg.RMQ.URL, cfg.RMQ.ServerExchange, rmqRouter, l)
	//if err != nil {
	//	l.Fatal(fmt.Errorf("app - Run - rmqServer - server.New: %w", err))
	//}

	//TODО Вынести в переменные среды

	authMiddleware := middleware.AuthMiddleware([]byte(cfg.Auth.SigningKey))

	handler := gin.New()
	http.NewRouter(handler, l, authMiddleware, *authUseCase, *userInfoUseCase, *sendCoinsUseCase, *buyGoodUseCase)

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
