package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

const (
	_defaultMaxRetries  = 3
	_defaultPoolSize    = 10
	_defaultConnTimeout = 5 * time.Second
)

type Redis struct {
	maxRetries  int
	poolSize    int
	DB          int
	connTimeout time.Duration
	Client      *redis.Client
}

func New(addr string, password string, DB int, opts ...Option) (*Redis, error) {
	r := &Redis{
		maxRetries:  _defaultMaxRetries,
		poolSize:    _defaultPoolSize,
		connTimeout: _defaultConnTimeout,
	}

	for _, opt := range opts {
		opt(r)
	}

	client := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     "",
		DB:           DB,
		PoolSize:     r.poolSize,
		DialTimeout:  r.connTimeout,
		ReadTimeout:  r.connTimeout,
		WriteTimeout: r.connTimeout,
	})

	log.Printf("Address %s\n", addr)
	log.Printf("Password %s\n", password)

	var err error
	for i := 0; i < r.maxRetries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), r.connTimeout)
		defer cancel()

		err = client.Ping(ctx).Err()
		if err == nil {
			break
		}

		log.Printf("Redis is trying to connect, attempts left: %d", r.maxRetries-i-1)
		time.Sleep(r.connTimeout)
	}

	if err != nil {
		return nil, fmt.Errorf("redis - New - failed to connect after %d attempts: %w", r.maxRetries, err)
	}

	r.Client = client
	return r, nil
}

func (r *Redis) Close() error {
	if r.Client != nil {
		return r.Client.Close()
	}
	return nil
}
