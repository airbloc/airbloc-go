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
