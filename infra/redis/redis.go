package redis

var (
	client *Redis
)

type Redis struct {
}

type Config struct {
}

func Init(cfg Config) *Redis {
	if client != nil {
		return client
	}

	client = &Redis{}

	return client
}
