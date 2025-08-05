package config

import "time"

// Default configuration

const (
	EnvProduction  = "production"
	EnvDevelopment = "development"
	EnvTest        = "test"

	ServerListen       = "0.0.0.0"
	ServerPort         = "8000"
	ServerTimeoutRead  = "15s"
	ServerTimeoutWrite = "15s"
	ServerTimeoutIdle  = "60s"

	MariadbHost = "127.0.0.1"
	MariadbPort = "3306"

	RedisHost = "127.0.0.1"
	RedisPort = "6379"
	RedisPassword = "pass"
	RedisDb   = "0"
)

// Defined constant store

const (
	CachingExtremeShortDuration = time.Minute * 5
	CachingShortDuration        = time.Hour
	CachingMediumDuration       = time.Hour * 4
	CachingLongDuration         = time.Hour * 8

	StoreBlockingState    = "block:ip:%s" // ip
	StoreBlockingPattern  = "block:ip:*"
	StoreSMSProviderIndex = "config:sms:provider"
)

// Tracer Attribute

const (
	EventCacheHit  = "Cache HIT"
	EventCacheMiss = "Cache MISS"
)

// Authorized context key

const (
	B2BAuthorizeContext = "b2b_auth"
)
