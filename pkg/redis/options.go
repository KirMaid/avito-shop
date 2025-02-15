package redis

import "time"

type Option func(*Redis)

func MaxRetries(retries int) Option {
	return func(r *Redis) {
		r.maxRetries = retries
	}
}

func PoolSize(size int) Option {
	return func(r *Redis) {
		r.poolSize = size
	}
}

func ConnTimeout(timeout time.Duration) Option {
	return func(r *Redis) {
		r.connTimeout = timeout
	}
}
