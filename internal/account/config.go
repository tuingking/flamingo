package account

import "time"

type Config struct {
	Service    ServiceConfig
	Repository RepositoryConfig
}

type ServiceConfig struct {
	Jwt JwtConfig
}

type JwtConfig struct {
	Secret           string
	AccessExpiresIn  time.Duration
	RefreshExpiresIn time.Duration
}

type RepositoryConfig struct {
}
