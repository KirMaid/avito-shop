package repositories

import "errors"

var ErrCacheMiss = errors.New("cache miss")
var ErrEmptyCacheData = errors.New("empty cache data")
