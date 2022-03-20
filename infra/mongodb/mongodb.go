package mongodb

var (
	client *MongoDB
)

type MongoDB struct {
}

type Config struct {
}

func Init(cfg Config) *MongoDB {
	if client != nil {
		return client
	}

	client = &MongoDB{}

	return client
}
