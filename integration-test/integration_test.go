package integration_test

//import (
//	"avitoshop/config"
//	handlershttp "avitoshop/internal/app/handlers/http"
//	"avitoshop/internal/app/middleware"
//	"avitoshop/internal/app/repositories"
//	auth "avitoshop/internal/app/usecases/auth"
//	buygood "avitoshop/internal/app/usecases/buy_good"
//	sendcoins "avitoshop/internal/app/usecases/send_coins"
//	userinfo "avitoshop/internal/app/usecases/user_info"
//	"avitoshop/pkg/logger"
//	"avitoshop/pkg/postgres"
//	"avitoshop/pkg/redis"
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"github.com/stretchr/testify/assert"
//	"io"
//	"io/ioutil"
//	"net/http"
//	"net/http/httptest"
//	"os"
//	"strconv"
//	"testing"
//	"time"
//)
//
//func setupTestServer() *gin.Engine {
//
//	fmt.Println("Environment variables:")
//	for _, env := range os.Environ() {
//		fmt.Println(env)
//	}
//
//	pgURL := fmt.Sprintf(
//		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
//		os.Getenv("DATABASE_USER"),
//		os.Getenv("DATABASE_PASSWORD"),
//		os.Getenv("DATABASE_HOST"),
//		getEnvAsInt("DATABASE_PORT", 5432),
//		os.Getenv("DATABASE_NAME"),
//	)
//	cfg := &config.Config{
//		Logger: config.Logger{
//			LogLevel: "debug",
//		},
//		Postgres: config.Postgres{
//			URL:     pgURL,
//			PoolMax: 10,
//		},
//		Redis: config.Redis{
//			Host:     os.Getenv("REDIS_HOST"),
//			Port:     getEnvAsInt("REDIS_PORT", 6379),
//			Password: os.Getenv("REDIS_PASSWORD"),
//			DB:       getEnvAsInt("REDIS_DB", 0),
//		},
//		Auth: config.Auth{
//			HashSalt:   "salt",
//			SigningKey: "secret",
//			TokenTTL:   "1h",
//		},
//		HTTP: config.HTTP{
//			Port: "8080",
//		},
//	}
//
//	l := logger.New(cfg.Logger.LogLevel)
//
//	pg, err := postgres.New(cfg.Postgres.URL, postgres.MaxPoolSize(cfg.Postgres.PoolMax))
//	if err != nil {
//		panic(err)
//	}
//	userRepo := repositories.NewUserRepository(pg.Pool)
//	inventoryRepo := repositories.NewInventoryRepository(pg.Pool)
//	transactionRepo := repositories.NewTransactionRepository(pg.Pool)
//	goodRepo := repositories.NewGoodRepository(pg.Pool)
//
//	redisAddr := fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port)
//	redisClient, err := redis.New(
//		redisAddr,
//		cfg.Redis.Password,
//		cfg.Redis.DB,
//	)
//
//	redisUserRepo := repositories.NewRedisUserRepository(redisClient.Client, 0)
//	redisInventoryRepo := repositories.NewRedisInventoryRepository(redisClient.Client, 0)
//	redisGoodRepo := repositories.NewRedisGoodRepository(redisClient.Client, 0)
//	redisTransactionRepo := repositories.NewRedisTransactionRepository(redisClient.Client, 0)
//
//	expireDuration, _ := time.ParseDuration(cfg.Auth.TokenTTL)
//	authUseCase := auth.NewAuthUseCase(
//		userRepo,
//		redisUserRepo,
//		cfg.Auth.HashSalt,
//		[]byte(cfg.Auth.SigningKey),
//		expireDuration,
//	)
//	userInfoUseCase := userinfo.NewUserInfoUseCase(
//		userRepo,
//		inventoryRepo,
//		transactionRepo,
//		goodRepo,
//		redisUserRepo,
//		redisInventoryRepo,
//		redisGoodRepo,
//		redisTransactionRepo,
//	)
//	sendCoinsUseCase := sendcoins.NewSendCoinsUseCase(
//		pg.Pool,
//		userRepo,
//		transactionRepo,
//		redisUserRepo,
//		redisTransactionRepo,
//	)
//	buyGoodUseCase := buygood.NewBuyGoodUseCase(
//		pg.Pool,
//		userRepo,
//		goodRepo,
//		inventoryRepo,
//		redisUserRepo,
//		redisGoodRepo,
//		redisInventoryRepo,
//	)
//
//	authMiddleware := middleware.AuthMiddleware([]byte(cfg.Auth.SigningKey))
//
//	handler := gin.New()
//	handlershttp.NewRouter(handler, l, authMiddleware, *authUseCase, *userInfoUseCase, *sendCoinsUseCase, *buyGoodUseCase)
//
//	return handler
//}
//
//func TestAuthEndpoint(t *testing.T) {
//	router := setupTestServer()
//	server := httptest.NewServer(router)
//	defer server.Close()
//
//	authReq := map[string]string{
//		"username": "user1",
//		"password": "password1",
//	}
//	body, _ := json.Marshal(authReq)
//	req, _ := http.NewRequest("POST", server.URL+"/api/auth", bytes.NewBuffer(body))
//	req.Header.Set("Content-Type", "application/json; charset=utf-8")
//
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	assert.NoError(t, err)
//	defer resp.Body.Close()
//
//	assert.Equal(t, http.StatusOK, resp.StatusCode)
//
//	respBody, _ := io.ReadAll(resp.Body)
//	var authResp map[string]string
//	json.Unmarshal(respBody, &authResp)
//
//	assert.NotEmpty(t, authResp["token"])
//}
//
//func TestInfoEndpoint(t *testing.T) {
//	router := setupTestServer()
//	server := httptest.NewServer(router)
//	defer server.Close()
//
//	authReq := map[string]string{
//		"username": "user1",
//		"password": "password1",
//	}
//	body, _ := json.Marshal(authReq)
//	req, _ := http.NewRequest("POST", server.URL+"/api/auth", bytes.NewBuffer(body))
//	req.Header.Set("Content-Type", "application/json; charset=utf-8")
//
//	client := &http.Client{}
//	resp, _ := client.Do(req)
//	respBody, _ := ioutil.ReadAll(resp.Body)
//	var authResp map[string]string
//	json.Unmarshal(respBody, &authResp)
//
//	req, _ = http.NewRequest("GET", server.URL+"/api/info", nil)
//	req.Header.Set("Authorization", "Bearer "+authResp["token"])
//
//	resp, _ = client.Do(req)
//	defer resp.Body.Close()
//
//	assert.Equal(t, http.StatusOK, resp.StatusCode)
//
//	respBody, _ = ioutil.ReadAll(resp.Body)
//	var infoResp map[string]interface{}
//	json.Unmarshal(respBody, &infoResp)
//
//	assert.NotNil(t, infoResp["coins"])
//	assert.NotNil(t, infoResp["inventory"])
//	assert.NotNil(t, infoResp["coinHistory"])
//}
//
//func TestSendCoinEndpoint(t *testing.T) {
//	router := setupTestServer()
//	server := httptest.NewServer(router)
//	defer server.Close()
//
//	authReq := map[string]string{
//		"username": "user1",
//		"password": "password1",
//	}
//	body, _ := json.Marshal(authReq)
//	req, _ := http.NewRequest("POST", server.URL+"/api/auth", bytes.NewBuffer(body))
//	req.Header.Set("Content-Type", "application/json; charset=utf-8")
//
//	client := &http.Client{}
//	resp, _ := client.Do(req)
//	respBody, _ := io.ReadAll(resp.Body)
//	var authResp map[string]string
//	json.Unmarshal(respBody, &authResp)
//
//	sendReq := map[string]interface{}{
//		"toUser": "user2",
//		"amount": 10,
//	}
//	body, _ = json.Marshal(sendReq)
//	req, _ = http.NewRequest("POST", server.URL+"/api/sendCoin", bytes.NewBuffer(body))
//	req.Header.Set("Authorization", "Bearer "+authResp["token"])
//	req.Header.Set("Content-Type", "application/json; charset=utf-8")
//
//	resp, _ = client.Do(req)
//	defer resp.Body.Close()
//
//	assert.Equal(t, http.StatusOK, resp.StatusCode)
//}
//
//func TestBuyItemEndpoint(t *testing.T) {
//	router := setupTestServer()
//	server := httptest.NewServer(router)
//	defer server.Close()
//
//	authReq := map[string]string{
//		"username": "user1",
//		"password": "password1",
//	}
//	body, _ := json.Marshal(authReq)
//	req, _ := http.NewRequest("POST", server.URL+"/api/auth", bytes.NewBuffer(body))
//	req.Header.Set("Content-Type", "application/json; charset=utf-8")
//
//	client := &http.Client{}
//	resp, _ := client.Do(req)
//	respBody, _ := io.ReadAll(resp.Body)
//	var authResp map[string]string
//	json.Unmarshal(respBody, &authResp)
//
//	req, _ = http.NewRequest("GET", server.URL+"/api/buy/sword", nil)
//	req.Header.Set("Authorization", "Bearer "+authResp["token"])
//
//	resp, _ = client.Do(req)
//	defer resp.Body.Close()
//
//	assert.Equal(t, http.StatusOK, resp.StatusCode)
//}
//
//func getEnvAsInt(key string, defaultValue int) int {
//	value := os.Getenv(key)
//	if value == "" {
//		return defaultValue
//	}
//	result, err := strconv.Atoi(value)
//	if err != nil {
//		return defaultValue
//	}
//	return result
//}
