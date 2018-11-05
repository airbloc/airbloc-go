package warehouse

import "time"

type Config struct {
	DefaultStorage string

	Http struct {
		Timeout         time.Duration
		MaxConnsPerHost int
	}

	LocalStorage struct {
		SavePath string
		Endpoint string
	}

	S3 struct {
		Region     string
		AccessKey  string
		Bucket     string
		PathPrefix string
	}
}

func DefaultConfig() (config Config) {
	config.DefaultStorage = "local"
	config.Http.Timeout = 30 * time.Second
	config.Http.MaxConnsPerHost = 10

	config.LocalStorage.SavePath = "local/warehouse/"
	config.LocalStorage.Endpoint = "http://localhost:9125/"
	return
}
